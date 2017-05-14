package main

import (
	"log"

	"github.com/chongg039/GoPRM/pe"
)

func main() {
	InitPCB, PCBPool, RCBPool, finish, running := pe.Init()

	log.Printf("%+v\n, %+v\n, %+v\n, %+v\n, %+v\n", InitPCB, PCBPool, RCBPool, finish, running)
	// var pcbpool = &pe.PCBPool{}
	// log.Printf("%p\n, %T", pcbpool, pcbpool)

	// running := pcbpool.Schedule()

	// log.Printf("%p\n, %T", running, running)
}
