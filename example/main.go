package main

import (
	"log"

	"github.com/chongg039/GoPRM/pe"
)

func main() {
	InitPCB, PCBPool, RCBPool, finish, running := pe.Init()

	log.Printf("%+v\n, %+v\n, %+v\n, %+v\n, %+v\n", InitPCB, PCBPool, RCBPool, finish, running)
}
