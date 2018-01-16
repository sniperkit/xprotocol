//Package push implements the pipeline pattern push node for mangos / nanomsg.
package push

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/push"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

//Node Structure used to send messages to other nodes.
type Node struct {
	url  string
	sock mangos.Socket
}

//Connect to a pull Server.
func (self *Node) Connect(url string) error {

	self.url = url

	var err error

	if self.sock, err = push.NewSocket(); err != nil {
		return err
	}
	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())
	if err = self.sock.Dial(url); err != nil {
		return err
	}

	return nil
}

//Push a message to a pull Server.
func (self *Node) Push(payload []byte) error {
	return self.sock.Send(payload)
}

//Closes the node connection.
func (self *Node) Close() {
	self.sock.Close()
}
