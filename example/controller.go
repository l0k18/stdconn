package main

import (
	"github.com/l0k18/log"
	"github.com/l0k18/stdconn/example/hello/hello"
	"github.com/l0k18/stdconn/worker"
)

func main() {
	log.L.SetLevel("trace", true)
	log.INFO("starting up example controller")
	cmd := worker.
		Spawn("go", "run", "hello/worker.go")
	client := hello.NewClient(cmd.StdConn)
	log.INFO("calling Hello.Say with 'worker'")
	log.INFO("reply:", client.Say("worker"))
	log.INFO("calling Hello.Bye")
	log.INFO("reply:", client.Bye())
}
