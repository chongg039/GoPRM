package pe

import (
	"log"
)

// TimeOut 模拟时间片到或外部中断
func (pcbPool *PCBPool) TimeOut(running *PCB, finish Queue) (*PCB, Queue) {
	// running块已被占用
	if running != nil {
		// 暂存原running
		temp := running
		temp.Status = "ready"
		temp.CPUState = "notused"
		temp.Memory = "notused"

		// 重置
		running = nil

		running = pcbPool.Schedule(running)
		pcbPool.AppendPCBEle(temp)
		// 调度的结果返回到finish队列
		finish = append(finish, temp)
	} else {
		running = pcbPool.Schedule(running)
	}
	return running, finish
}

// TimeRotation is used in each priority's ready queue
// 在每个优先级ready队列中使用时间片轮转
func (pcbPool *PCBPool) TimeRotation() {
	if pcbPool == nil {
		return
	}
	for i := len(pcbPool) - 1; i >= 0; i-- {
		if pcbPool[i].Head == nil {
			continue
		}
		h := pcbPool[i].Head

		pcbPool[i].Head = pcbPool[i].Head.Next
		pcbPool[i].Length--

		pcbPool.AppendPCBEle(h)

	}
	return
}

// Schedule is based on priority 0, 1, 2
// 返回running运行时的指针地址，没有的话为nil
func (pcbPool *PCBPool) Schedule(running *PCB) *PCB {
	for priority := 2; priority >= 0; priority-- {
		h := pcbPool[priority].Head
		if h != nil {
			h.Status = "running"
			h.CPUState = "using"
			h.Memory = "using"
			running = h
			// 从原队列中移除
			judge := pcbPool.RemovePCBEle(h)
			if judge == true {
				return running
			}
			h = h.Next
			continue
		}

	}
	return running
}

// 检测进程的所有资源是否都已准备好
func (p *PCB) detectAllResourceStatus() bool {
	if len(p.ReqResArr) == 0 {
		log.Printf("OK, %s doesn't need any resource", p.Name)
		return true
	}
	for i := 0; i < len(p.ReqResArr); i++ {
		if p.ReqResArr[i].OK == false {
			return false
		}
		continue
	}
	log.Printf("OK, %s has requested all resources!", p.Name)
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
