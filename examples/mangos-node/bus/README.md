# mangosNode bus

Example Code to start 3 Bus Nodes and connect all 3 for distributed messages.
###Bus 1
	
	package main
	
	import (
		"github.com/DanielRenne/mangosNode/bus"
		"log"
		"time"
	)
	
	const url1 = "tcp://127.0.0.1:600"
	const url2 = "tcp://127.0.0.1:601"
	const url3 = "tcp://127.0.0.1:602"
	
	func main() {
		var node bus.Node
	
		err := node.Listen(url1, handleBusMessage)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		err = node.Connect(url2)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		//Code a forever loop to stop main from exiting.
		for {
			time.Sleep(3 * time.Second)
			go node.Send([]byte("Bus1 Message"))
		}
	
	}
	
	func handleBusMessage(msg []byte) {
		log.Println(string(msg))
	}

	
### Bus 2

	package main
	
	import (
		"github.com/DanielRenne/mangosNode/bus"
		"log"
		"time"
	)
	
	const url1 = "tcp://127.0.0.1:600"
	const url2 = "tcp://127.0.0.1:601"
	const url3 = "tcp://127.0.0.1:602"
	
	func main() {
		var node bus.Node
	
		err := node.Listen(url2, handleBusMessage)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		err = node.Connect(url3)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		//Code a forever loop to stop main from exiting.
		for {
			time.Sleep(3 * time.Second)
			go node.Send([]byte("Bus2 Message"))
		}
	
	}
	
	func handleBusMessage(msg []byte) {
		log.Println(string(msg))
	}

### Bus 3
	
	package main
	
	import (
		"github.com/DanielRenne/mangosNode/bus"
		"log"
		"time"
	)
	
	const url1 = "tcp://127.0.0.1:600"
	const url2 = "tcp://127.0.0.1:601"
	const url3 = "tcp://127.0.0.1:602"
	
	func main() {
		var node bus.Node
	
		err := node.Listen(url3, handleBusMessage)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		err = node.Connect(url1)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		//Code a forever loop to stop main from exiting.
		for {
			time.Sleep(3 * time.Second)
			go node.Send([]byte("Bus3 Message"))
		}
	
	}
	
	func handleBusMessage(msg []byte) {
		log.Println(string(msg))
	}
