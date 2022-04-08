package irc

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Options struct {
	Addr           string
	Nick           string
	Channels       []string
	Timeout        time.Duration
	ReconnectDelay time.Duration
	TLSConfig      tls.Config
	Logger         *log.Logger
}

type IRC struct {
	log       *log.Logger
	conn      net.Conn
	Opts      *Options
	Nick      string
	callbacks map[string][]func(*IRC, Message)
}

func New(o *Options) *IRC {
	if o.Logger == nil {
		o.Logger = log.Default()
	}
	if o.Timeout == 0 {
		o.Timeout = time.Minute * 5
	}
	if o.ReconnectDelay == 0 {
		o.ReconnectDelay = time.Second * 10
	}

	return &IRC{log: o.Logger, Opts: o, callbacks: map[string][]func(*IRC, Message){}}
}

func (i *IRC) Start() {
	for i.run() {
		time.Sleep(i.Opts.ReconnectDelay)
	}
}

func (i *IRC) Register(cmd string, fn func(*IRC, Message)) {
	i.callbacks[cmd] = append(i.callbacks[cmd], fn)
}

func (i *IRC) Say(target, message string) error {
	return i.Write(fmt.Sprintf("PRIVMSG %s :%s", target, message))
}

func (i *IRC) dial() error {
	if i.conn != nil {
		_ = i.conn.Close()
	}
	i.log.Println("irc: connecting to", i.Opts.Addr)
	c, err := tls.Dial("tcp", i.Opts.Addr, &i.Opts.TLSConfig)
	if err != nil {
		return err
	}
	i.conn = c
	i.Nick = i.Opts.Nick
	n := i.Opts.Nick
	if err := i.Write(fmt.Sprintf("NICK %s", n)); err != nil {
		return err
	}
	if err := i.Write(fmt.Sprintf("USER %s %s %s :%s", n, n, n, n)); err != nil {
		return err
	}

	return nil
}

func (i *IRC) run() (reconnect bool) {
	reconnect = true
	if err := i.dial(); err != nil {
		i.log.Println("irc: connection error:", err)
	}
	_ = i.conn.SetReadDeadline(time.Now().Add(i.Opts.Timeout))

	s := bufio.NewScanner(i.conn)
	for s.Scan() {
		if err := i.processLine(s.Text()); err != nil {
			i.log.Println("irc: process:", err)
		}
		_ = i.conn.SetReadDeadline(time.Now().Add(i.Opts.Timeout))
	}
	if s.Err() != nil {
		i.log.Println("irc: connection error:", s.Err())
	}
	return
}

func (i *IRC) processLine(line string) error {
	m := parseMessage(line)

	switch m.Command {
	case "PING":
		if err := i.Write("PONG " + m.RawArgs); err != nil {
			return err
		}
	case "001":
		for _, channel := range i.Opts.Channels {
			if err := i.Write(fmt.Sprintf("JOIN %s", channel)); err != nil {
				return err
			}
		}
	case "433":
		i.Nick += "_"
		if err := i.Write(fmt.Sprintf("NICK %s", i.Nick)); err != nil {
			return err
		}
	}

	for _, fn := range i.callbacks[m.Command] {
		go fn(i, m)
	}
	return nil
}

func (i *IRC) Write(data string) error {
	_, err := i.conn.Write([]byte(data + "\r\n"))
	return err
}

type Message struct {
	Prefix  string
	Command string
	User    struct {
		Nick string
		User string
		Host string
	}
	Target  string
	Args    []string
	RawArgs string
}

func parseMessage(line string) Message {
	var m Message
	if strings.HasPrefix(line, ":") {
		line = line[1:]
		x := strings.SplitN(line, " ", 4)
		if len(x) < 3 {
			return m
		}
		m.Prefix = x[0]
		m.Command = x[1]
		if strings.Contains(m.Prefix, "!") && strings.Contains(m.Prefix, "@") {
			hs := strings.SplitN(m.Prefix, "@", 2)
			if len(hs) == 2 {
				m.User.Host = hs[1]
			}
			nu := strings.SplitN(hs[0], "!", 2)
			if len(nu) == 2 {
				m.User.Nick = nu[0]
				m.User.User = nu[1]
			}
		} else {
			m.User.Host = m.Prefix
		}
		m.Target = x[2]
		if len(x) == 4 {
			if x[3][0] == ':' {
				m.RawArgs = x[3][1:]
			} else {
				m.RawArgs = x[3]
			}
			m.Args = strings.Split(m.RawArgs, " ")
		}
	} else {
		x := strings.SplitN(line, " ", 2)
		if len(x) < 2 {
			return m
		}
		m.Command = x[0]
		m.RawArgs = x[1]
	}
	return m
}
