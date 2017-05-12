package pe

import (
	"log"
	"time"
)

// 在每个优先级ready队列中使用时间片轮转

// TimeRotation is used in each priority's ready queue
func (pcbPool *PCBPool) TimeRotation() {
	for i := len(pcbPool) - 1; i >= 0; i-- {
		pcbPool[i] = append(pcbPool[i][1:], pcbPool[i][0])
	}
}

// Schedule is based on priority 0, 1, 2
// 返回running运行时的指针地址，没有的话为nil
func (pcbPool *PCBPool) Schedule() *Running {
	var running Running
Check:
	for priority := 2; priority >= 0; priority-- {
		for i := 0; i < len(pcbPool[priority]); i++ {
			// 所有资源准备就绪
			if pcbPool[priority][i].detectAllResourceStatus() == true {
				start := time.Now()
				pcbPool[priority][i].Status = "running"
				running = Running{pcbPool[priority][i], start}
				// 从原队列中移除
				pcbPool[priority] = append(pcbPool[priority][:i], pcbPool[priority][i+1:]...)
				break Check
			}
		}
	}
	return &running
}

// 检测进程的所有资源是否都已准备好
func (p *PCB) detectAllResourceStatus() bool {
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
func (p *PCB) removeFromReadyQueue(queue Queue) Queue {
	for k, v := range queue {
		if p.Name == v.Name {
			r := queue[:k]
			l := queue[k+1:]
			queue = append(r, l...)
		}
	}
	return queue
}
