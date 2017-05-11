package pe

import (
	"log"
	"time"
)

// 在每个优先级ready队列中使用时间片轮转

// TimeRotation is used in each priority's ready queue
func (rqa *QueuesArr) TimeRotation() {
	for i := len(rqa) - 1; i >= 0; i-- {
		rqa[i] = append(rqa[i][1:], rqa[i][0])
	}
}

// Schedule is based on priority 0, 1, 2
// 返回running运行时的指针地址，没有的话为nil
func (rqa *QueuesArr) Schedule() *Running {
	var running Running
Check:
	for priority := 2; priority >= 0; priority-- {
		for i := 0; i < len(rqa[priority]); i++ {
			// 所有资源准备就绪
			if rqa[priority][i].detectAllResourceStatus() == true {
				start := time.Now()
				rqa[priority][i].Status = "running"
				running = Running{rqa[priority][i], start}
				// 从原队列中移除
				rqa[priority] = append(rqa[priority][:i], rqa[priority][i+1:]...)
				break Check
			}
		}
	}
	return &running
}

// 检测进程的所有资源是否都已准备好
func (p *Process) detectAllResourceStatus() bool {
	for i := 0; i < len(p.RequestResArr); i++ {
		if p.RequestResArr[i].OK == false {
			return false
		}
		continue
	}
	log.Printf("%s has requested all resources!", p.Name)
	return true
}

// 从就绪队列中移除
func (p *Process) removeFromReadyQueue(readyQueue Queue) Queue {
	for k, v := range readyQueue {
		if p.Name == v.Name {
			r := readyQueue[:k]
			l := readyQueue[k+1:]
			readyQueue = append(r, l...)
		}
	}
	return readyQueue
}
