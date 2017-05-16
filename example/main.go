package main

import (
	"log"

	"github.com/chongg039/GoPRM/pe"
)

func main() {
	// var i *pe.PCB
	// log.Printf("%p", i)
	// log.Println(i == nil)

	// var j = new(pe.PCB)
	// log.Printf("%p", j)

	// InitPCB, PCBPool, RCBPool, finish, running := pe.Init()

	// log.Printf("%+v\n, %+v\n, %+v\n, %+v\n, %+v\n", InitPCB, PCBPool, RCBPool, finish, running)

	// var pcbpool = &pe.PCBPool{}
	// log.Printf("%p\n, %T", pcbpool, pcbpool)

	// running := pcbpool.Schedule()

	// log.Printf("%p\n, %T", running, running)
	pcb1 := new(pe.PCB)
	pcb2 := new(pe.PCB)
	pcb1ele := new(pe.PCBEle)
	pcb2ele := new(pe.PCBEle)

	pool := new(pe.PCBPool)
	pcb1 = &pe.PCB{
		Name:      "PCB1",
		PID:       1000,
		Status:    "running",
		Priority:  2,
		CPUState:  "notused",
		Memory:    "notused",
		ReqResArr: []pe.RequestResource{},
		Parent:    nil,
		Children:  pcb2,
	}
	pcb2 = &pe.PCB{
		Name:      "PCB2",
		PID:       2000,
		Status:    "ready",
		Priority:  2,
		CPUState:  "notused",
		Memory:    "notused",
		ReqResArr: []pe.RequestResource{},
		Parent:    pcb1,
		Children:  nil,
	}

	pcb1ele = &pe.PCBEle{
		Data: *pcb1,
		Next: pcb2ele,
	}
	pcb2ele = &pe.PCBEle{
		Data: *pcb2,
		Next: nil,
	}
	pool[2].Head = pcb1ele
	pool[2].Length = 2

	running := pool.Schedule()

	log.Printf("%p, %T, %+v", running, running, running)

}
