package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chongg039/GoPRM/pe"
)

func main() {
	monk := true
	reader := bufio.NewReader(os.Stdin)
	help := `Process & Resource Manager
---------------------
init:	Init something
cr:	Create a process
req:	Process request resource
rel:	Process release resource
to:	Time out to schedule
stop: 	exit
---------------------`
	fmt.Println(help)
	InitPCB, pcbPool, rcbPool, finish, running := pe.Init()
	log.Printf("init pcb %p", InitPCB)
	log.Printf("pcbPool %+v", pcbPool[0])
	log.Printf("rcbPool %+v", *rcbPool)
	log.Printf("finish %+v", finish)
	log.Printf("running pcb %p", running)
	log.Println("---------------------")
	for monk {
		log.Println("Input a command:")
		data, _, _ := reader.ReadLine()
		command := string(data)
		filed := strings.Fields(command)

		switch {
		case len(filed) == 3 && filed[0] == "cr":
			// f := running.FindPCB(filed[0])
			if running == nil {
				log.Println("no process is running and cant create")
				break
			}
			level, err := strconv.Atoi(filed[2])
			if err != nil {
				log.Println("please input priority range of [0, 1, 2]")
				break
			}
			c := running.CreatePCB(filed[1], level)
			pcbPool.AppendPCBEle(c)
			log.Printf("pcb pool is %+v", pcbPool)
		case len(filed) >= 2 && filed[0] == "req":
			rcbPool, pcbPool = running.RequestResource(rcbPool, pcbPool, filed[1:]...)
			log.Printf("running is %+v", running)
			log.Printf("rcb pool is %+v", rcbPool)
		case filed[0] == "rel":
			rcbPool, pcbPool = running.ReleaseResource(rcbPool, pcbPool)
			log.Printf("running is %+v", running)
			log.Printf("rcb pool is %+v", rcbPool)
		case filed[0] == "to":
			running, finish = pcbPool.TimeOut(running, finish)
			log.Printf("running is %+v", running)
			log.Printf("pcb pool is %+v", pcbPool)
			log.Printf("finish is %+v", finish)
		case filed[0] == "de":
			log.Println(running)
			rcbPool, running, finish = pcbPool.DestoryPCB(rcbPool, running, finish)
			log.Printf("running is %+v", running)
			log.Printf("rcb pool is %+v", rcbPool)
			log.Printf("pcb pool is %+v", pcbPool)
			log.Printf("finish is %+v", finish)
		case filed[0] == "stop":
			monk = false
			log.Println("exit OK")
		default:
			log.Println("can't parse command")
		}
		log.Println("---------------------")
	}
}
