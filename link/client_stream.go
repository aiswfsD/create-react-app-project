package link

import (
	"context"
	"fmt"
	"sync"

	"github.com/baetyl/baetyl-go/log"
	"github.com/baetyl/baetyl-go/utils"
)

type stream struct {
	cli  *Client
	conn Link_TalkClient
	tomb utils.Tomb
	once sync.Once
}

func (c *Client) connect() (*stream, error) {
	cs, err := c.cli.Talk(context.Background())
	if err != nil {
		return nil, err
	}
	s := &stream{
		cli:  c,
		conn: cs,
	}
	s.tomb.Go(s.receiving)
	return s, nil
}

func (s *stream) send(msg *Message) error {
	err := s.conn.Send(msg)
	if err != nil {
		s.die("failed to send message", err)
		return err
	}

	if ent := s.cli.log.Check(log.DebugLevel, "client sent a message"); ent != nil {
		ent.Write(log.Any("msg", fmt.Sprintf("%v", msg)))
	}

	return nil
}

func (s *stream) sending(curr *Message) *Message {
	s.cli.log.Info("client starts to send messages")
	defer s.cli.log.Info("client has stopped sending messages")

	var err error
	if curr != nil {
		err = s.send(curr)
		if err != nil {
			return curr
		}
	}
	for {
		select {
		case msg := <-s.cli.cache:
			err = s.send(msg)
			if err != nil {
				return msg
			}
		case <-s.cli.tomb.Dying():
			return nil
		case <-s.tomb.Dying():
			return nil
		}
	}
}

func (s *stream) receiving() error {
	s.cli.log.Info("client starts to receive messages")
	defer s.cli.log.Info("client has stopped receiving messages")

	var err error
	var msg *Message
	for {
		msg, err = s.conn.Recv()
		if err != nil {
			s.die("failed to receive message", err)
			return err
		}

		if ent := s.cli.log.Check(log.DebugLevel, "client received a message"); ent != nil {
			ent.Write(log.Any("msg", fmt.Sprintf("%v", msg)))
		}

		if msg.Ack() {
			err = s.cli.onAck(msg)
		} else {
			err = s.cli.onMsg(msg)
		}
		if err != nil {
			s.die("failed to handle message", err)
			return err
		}
	}
}

func (s *stream) die(msg string, err error) {
	s.once.Do(func() {
		s.tomb.Kill(err)
		s.cli.onErr(msg, err)
	})
}

// ! called in the same goroutine with sending
func (s *stream) close() error {
	s.die("", nil)
	s.conn.CloseSend()
	return s.tomb.Wait()
}
