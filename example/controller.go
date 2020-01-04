package main

import (
	"github.com/bindchain/core/pkg/log"
	"github.com/bindchain/core/pkg/stdconn/example/hello/hello"
	"github.com/bindchain/core/pkg/stdconn/worker"
)

func main() {
	log.L.SetLevel("trace", true)
	log.INFO("starting up example controller")
	cmd := worker.Spawn("go", "run", "hello/worker.go")
	client := hello.NewClient(cmd.StdConn)
	log.INFO("calling Hello.Say with 'worker'")
	log.INFO("reply:", client.Say("worker"))
	log.INFO("calling Hello.Bye")
	log.INFO("reply:", client.Bye())
}
