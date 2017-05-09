package main

import (
	"log"
	"time"

	"github.com/chongg039/GoPRM/pe"
)

func main() {
	var (
		// initProc, testProc1, testProc2, testProc3, testProc4 pe.Process
		initProc, testProc1, testProc2 pe.Process
		R1, R2, R3, R4                 pe.Resource
		ResPool                        pe.ResourcePool
		finishQueue                    pe.Queue
		readyQueues                    pe.QueuesArr
	)

	R1.InitResource("R1", 1)
	R2.InitResource("R2", 2)
	R3.InitResource("R3", 3)
	R4.InitResource("R4", 4)

	ResPool.InitResourcePool(R1, R2, R3, R4)

	log.Printf("resource pool : %+v\n", ResPool)

	// 创建三个ready队列，优先级为0，1，2
	readyQueues = initProc.InitProcess("InitProc")
	log.Printf("Ready Queues are %+v\n", readyQueues)

	ResPool, finishQueue, readyQueues = initProc.RequestResource(ResPool, finishQueue, readyQueues)

	log.Printf("resource pool : %+v\n", ResPool)
	log.Printf("finish queue : %+v\n", finishQueue)

	// testProc1.CreateProcess("testProc1", 2, "R1")
	go func() {
		testProc1.CreateProcess("testProc1", 2, "R1")
		log.Printf("Ready Queues are %+v\n", readyQueues)
		readyQueues = testProc1.InsertSortedQueue(readyQueues)
		log.Printf("Ready Queues are %+v\n", readyQueues)

		ResPool, finishQueue, readyQueues = testProc1.RequestResource(ResPool, finishQueue, readyQueues)
		log.Printf("resource pool : %+v\n", ResPool)
		log.Printf("finish queue : %+v\n", finishQueue)
		log.Printf("Ready Queues are %+v\n", readyQueues)
	}()

	go func() {
		testProc2.CreateProcess("testProc2", 1, "R1")
		log.Printf("Ready Queues are %+v\n", readyQueues)
		readyQueues = testProc2.InsertSortedQueue(readyQueues)
		log.Printf("Ready Queues are %+v\n", readyQueues)

		ResPool, finishQueue, readyQueues = testProc2.RequestResource(ResPool, finishQueue, readyQueues)
		log.Printf("resource pool : %+v\n", ResPool)
		log.Printf("finish queue : %+v\n", finishQueue)
		log.Printf("Ready Queues are %+v\n", readyQueues)
	}()

	time.Sleep(time.Second * 20)
	// testProc2.CreateProcess("testProc2", 1, "R1")

	// log.Printf("Ready Queues are %+v\n", readyQueues)

	// readyQueues = testProc1.InsertSortedQueue(readyQueues)
	// readyQueues = testProc2.InsertSortedQueue(readyQueues)
	// log.Printf("Ready Queues are %+v\n", readyQueues)

	// ResPool, finishQueue, readyQueues = testProc1.RequestResource(ResPool, finishQueue, readyQueues)
	// log.Printf("resource pool : %+v\n", ResPool)
	// log.Printf("finish queue : %+v\n", finishQueue)
	// log.Printf("Ready Queues are %+v\n", readyQueues)

	// ResPool, finishQueue, readyQueues = testProc2.RequestResource(ResPool, finishQueue, readyQueues)
	// log.Printf("resource pool : %+v\n", ResPool)
	// log.Printf("finish queue : %+v\n", finishQueue)
	// log.Printf("Ready Queues are %+v\n", readyQueues)

	// testProc1.CreateProcess("TestProc1", 2, "R1", "R2")
	// readyQueues = testProc1.InsertSortedQueue(readyQueues)
	// log.Printf("%+v\n", readyQueues)

	// testProc2.CreateProcess("TestProc2", 1, []string{"R3", "R2"})
	// readyQueues = testProc2.InsertSortedQueue(readyQueues)
	// log.Printf("%+v\n", readyQueues)

	// time.Sleep(time.Second * 3)

	// testProc3.CreateProcess("TestProc3", 2, []string{"R4", "R1"})
	// readyQueues = testProc3.InsertSortedQueue(readyQueues)
	// log.Printf("%+v\n", readyQueues)

	// time.Sleep(time.Second * 3)

	// testProc4.CreateProcess("TestProc4", 1, []string{"R3", "R1", "R2", "R4"})
	// readyQueues = testProc4.InsertSortedQueue(readyQueues)
	// log.Printf("%+v\n", readyQueues)
	// readyQueue = append(readyQueue, testProc1, testProc2, testProc3)
	// log.Println(readyQueue)

}
