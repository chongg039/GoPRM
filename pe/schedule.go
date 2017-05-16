package pe

import (
	"log"
	"time"
)

// 在每个优先级ready队列中使用时间片轮转

// TimeRotation is used in each priority's ready queue
func (pcbPool *PCBPool) TimeRotation() {
	if pcbPool == nil {
		return
	}
	for i := len(pcbPool) - 1; i >= 0; i-- {
		if pcbPool[i].Head == nil {
			continue
		}
		h := pcbPool[i].Head.Data

		pcbPool[i].Head = pcbPool[i].Head.Next
		pcbPool[i].Length--

		pcbPool.AppendPCBEle(&h)

	}
	return
}

// Schedule is based on priority 0, 1, 2
// 返回running运行时的指针地址，没有的话为nil
func (pcbPool *PCBPool) Schedule() *Running {
	var running *Running

	for priority := 2; priority >= 0; priority-- {
		if pcbPool[priority].Head == nil {
			break
		}
		h := pcbPool[priority].Head

		for {
			if h.Data.detectAllResourceStatus() == true {
				s := time.Now().Format("2006-01-02 15:04:05")
				start, _ := time.Parse("2006-01-02 15:04:05", s)
				h.Data.Status = "running"
				h.Data.CPUState = "using"
				h.Data.Memory = "using"
				running = &Running{h.Data, start}
				// 从原队列中移除
				judge := pcbPool.RemovePCBEle(&h.Data)
				if judge == true {
					break
				}
				h = h.Next
			}
			h = h.Next
		}

	}
	return running
}

// 检测进程的所有资源是否都已准备好
func (p *PCB) detectAllResourceStatus() bool {
	if len(p.ReqResArr) == 0 {
		return true
	}
	for i := 0; i < len(p.ReqResArr); i++ {
		if p.ReqResArr[i].OK == false {
			return false
		}
		continue
	}
	log.Printf("%s has requested all resources!", p.Name)
	return true
}

// 从队列中移除
// func (p *PCB) removeFromReadyQueue(queue Queue) Queue {
// 	for k, v := range queue {
// 		if p.Name == v.Name {
// 			r := queue[:k]
// 			l := queue[k+1:]
// 			queue = append(r, l...)
// 		}
// 	}
// 	return queue
// }
