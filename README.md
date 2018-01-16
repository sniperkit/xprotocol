# msgq
Convenience wrapper around [mangos](http://github.com/gdamore/mangos) and [Iris](https://github.com/project-iris/iris) to avoid boilerplate code when generating 
new instances of SP (scalability protocols) patterns and some syntax sugar regarding sender/receiver functions. 
Does not abstract away mangos and iris-go libraries.


### Mangos New Instance Functions

```go
func NewBus(url string) (sock mangos.Socket, err error)
```

```go
func NewPair(url string) (sock mangos.Socket, err error)
```

```go
func NewPushPipeline(url string) (sock mangos.Socket, err error)
```

```go
func NewPullPipeline(url string) (sock mangos.Socket, err error)
```

```go
func NewPub(url string) (sock mangos.Socket, err error)
```

```go
func NewSub(url string) (sock mangos.Socket, err error)
```

```go
func NewRequest(url string) (sock mangos.Socket, err error)
```

```go
func NewReply(url string) (sock mangos.Socket, err error)
```

```go
func NewSurveyor(url string) (sock mangos.Socket, err error)
```

```go
func NewRespondent(url string) (sock mangos.Socket, err error)
```

### Mangos Sender Functions

**Dial first, then send**

```go
func Dial(sock mangos.Socket, url string) error
```

```go
func Send(sock mangos.Socket, msg []byte) error
```

## Mangos Receiver Functions

**Listen first, then receive**

```go
func Listen(sock mangos.Socket, url string) error
```

```go
func Receive(sock mangos.Socket) (msg []byte, err error)
```

### NSQ Struct and Functions

```go
type Iris struct {
	conn *iris.Connection
}
```

```go
func (i *Iris) Connect(port int) (err error) 
```
```go
func (i *Iris) Close() error 
```
```go
func (i *Iris) Publish(topic string, msg []byte) error 
```

```go
func (i *Iris) Subscribe(topic string, handler iris.TopicHandler) error 
```

```go
func (i *Iris) Unsubscribe(topic string) error 
```

```go
func (i *Iris) Broadcast(cluster string, msg []byte) error 
```

```go
func (i *Iris) Request(cluster string, msg []byte, timeout time.Duration) ([]byte, error) 
```

### Example

See [https://github.com/ibmendoza/go-examples/tree/master/msgq](https://github.com/ibmendoza/go-examples/tree/master/msgq)

### License

MIT
