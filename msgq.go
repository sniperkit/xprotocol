// Copyright 2015 Isagani Mendoza (http://itjumpstart.wordpress.com)
// License: MIT
// Convenience TCP inter-process wrapper around mangos

package msgq

import (
	"errors"
	"fmt"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/bus"
	"github.com/go-mangos/mangos/protocol/pair"
	"github.com/go-mangos/mangos/protocol/pub"
	"github.com/go-mangos/mangos/protocol/pull"
	"github.com/go-mangos/mangos/protocol/push"
	"github.com/go-mangos/mangos/protocol/rep"
	"github.com/go-mangos/mangos/protocol/req"
	"github.com/go-mangos/mangos/protocol/respondent"
	"github.com/go-mangos/mangos/protocol/sub"
	"github.com/go-mangos/mangos/protocol/surveyor"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

/* Sender functions */

func Send(sock mangos.Socket, msg []byte) error {
	if err := sock.Send(msg); err != nil {
		s := fmt.Sprintf("socket.Send: %s", err)
		return errors.New(s)
	}
	return nil
}

func Dial(sock mangos.Socket, url string) error {
	if err := sock.Dial(url); err != nil {
		s := fmt.Sprintf("socket.Dial: %s", err)
		return errors.New(s)
	}
	return nil
}

/* Receiver functions */

func Listen(sock mangos.Socket, url string) error {
	if err := sock.Listen(url); err != nil {
		s := fmt.Sprintf("socket.Listen: %s", err)
		return errors.New(s)
	}
	return nil
}

func Receive(sock mangos.Socket) (msg []byte, err error) {
	msg, err = sock.Recv()
	if err != nil {
		s := fmt.Sprintf("socket.Recv: %s", err)
		return nil, errors.New(s)
	}
	return msg, nil
}

func newSocket(url, protocol string) (sock mangos.Socket, err error) {
	var str string

	switch protocol {
	case "bus":
		sock, err = bus.NewSocket()
		str = "bus.NewSocket: "
	case "pair":
		sock, err = pair.NewSocket()
		str = "pair.NewSocket: "
	case "push":
		sock, err = push.NewSocket()
		str = "push.NewSocket: "
	case "pull":
		sock, err = pull.NewSocket()
		str = "pull.NewSocket: "
	case "pub":
		sock, err = pub.NewSocket()
		str = "pub.NewSocket: "
	case "sub":
		sock, err = sub.NewSocket()
		str = "sub.NewSocket: "
	case "request":
		sock, err = req.NewSocket()
		str = "request.NewSocket: "
	case "reply":
		sock, err = rep.NewSocket()
		str = "reply.NewSocket: "
	case "surveyor":
		sock, err = surveyor.NewSocket()
		str = "surveyor.NewSocket: "
	case "respondent":
		sock, err = respondent.NewSocket()
		str = "respondent.NewSocket: "
	}

	if err != nil {
		s := fmt.Sprintf(str+" %s", err)
		return nil, errors.New(s)
	}

	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())

	return sock, nil
}

/* new instance functions */

func NewBus(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "bus")
}

func NewPair(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "pair")
}

func NewPushPipeline(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "push")
}

func NewPullPipeline(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "pull")
}

func NewPub(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "pub")
}

func NewSub(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "sub")
}

func NewRequest(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "request")
}

func NewReply(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "reply")
}

func NewSurveyor(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "surveyor")
}

func NewRespondent(url string) (sock mangos.Socket, err error) {
	return newSocket(url, "respondent")
}
