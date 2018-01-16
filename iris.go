package msgq

import (
	"gopkg.in/project-iris/iris-go.v1"
	"runtime"
	"time"
)

type Iris struct {
	conn *iris.Connection
}

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
*/

func (i *Iris) Connect(port int) (err error) {
	if port != 55555 {
		i.conn, err = iris.Connect(port)
	} else {
		i.conn, err = iris.Connect(55555)
	}
	return err
}

func (i *Iris) Close() error {
	return i.Close()
}

func (i *Iris) Publish(topic string, msg []byte) error {
	return i.conn.Publish(topic, msg)
}

func (i *Iris) Subscribe(topic string, handler iris.TopicHandler) error {
	return i.conn.Subscribe(topic, handler, &iris.TopicLimits{
		EventThreads: runtime.NumCPU(),
		EventMemory:  64 * 1024 * 1024,
	})
}

func (i *Iris) Unsubscribe(topic string) error {
	return i.conn.Unsubscribe(topic)
}

func (i *Iris) Broadcast(cluster string, msg []byte) error {
	return i.conn.Broadcast(cluster, msg)
}

func (i *Iris) Request(cluster string, msg []byte, timeout time.Duration) ([]byte, error) {
	return i.conn.Request(cluster, msg, timeout)
}
