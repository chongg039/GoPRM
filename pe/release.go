package pe

import (
	"log"
	"time"
)

// time sleep 2 second to simulate process running
func (p *PCB) running() {
	time.Sleep(time.Second * 2)
}

// 判定资源池中是否有某进程仍在某资源的blocked队列中
func (p *PCB) detectBlockedQueue(rcbPool RCBPool) bool {
	for i := 0; i < len(p.RequestResArr); i++ {
		for j := 0; j < len(rcbPool); j++ {
			for k := 0; k < len(rcbPool[j].BlockedList); k++ {
				if p.Name == rcbPool[j].BlockedList[k].Name {
					return true
				}
			}
		}
	}
	return false
}

func (p *PCB) removeFromBlockedQueue(blockedQueue Queue) Queue {
	for k, v := range blockedQueue {
		if p.Name == v.Name {
			r := blockedQueue[:k]
			l := blockedQueue[k+1:]
			blockedQueue = append(r, l...)
		}
	}
	return blockedQueue
}

// ReleaseResource should be used when process is already running
// 每次释放资源resource检测自身是否可用（>0），并从blocked队列中取出放入running
func (p *PCB) ReleaseResource(rcbPool RCBPool, finishedQueue Queue) (RCBPool, Queue) {
	for i := 0; i < len(p.RequestResArr); i++ {
		for j := 0; j < len(rcbPool); j++ {
			if p.RequestResArr[i].Name == rcbPool[j].Name {
				rcbPool[j].Available++
				log.Printf("%s Already release resource, %s", p.Name, rcbPool[j].Name)

				if i == len(p.RequestResArr)-1 {
					p.Status = "finished"
					finishedQueue = append(finishedQueue, *p)
					log.Printf("Process %s has already finished, put it to finished queue", p.Name)

					// 检测该blocked队列
					if rcbPool[j].Available > 0 && len(rcbPool[j].BlockedList) > 0 {
						rcbPool[j].Available--
						//　检测所有blocked队列
						judge := p.detectBlockedQueue(rcbPool)
						if judge == false {
							rcbPool[j].BlockedList[0].Status = "finished"
							finishedQueue = append(finishedQueue, rcbPool[j].BlockedList[0])
							rcbPool[j].BlockedList = rcbPool[j].BlockedList[0].removeFromBlockedQueue(rcbPool[j].BlockedList)
						}
					}
					return rcbPool, finishedQueue
				}
			}
		}
	}
	return rcbPool, finishedQueue
}
