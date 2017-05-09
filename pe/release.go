package pe

import (
	"log"
	"time"
)

// time sleep 2 second to simulate process running
func (p *Process) running() {
	time.Sleep(time.Second * 2)
}

// 判定资源池中是否有某进程仍在某资源的blocked队列中
func (rp ResourcePool) detectBlockedQueue(p Process) bool {
	for i := 0; i < len(p.RequestResArr); i++ {
		for j := 0; j < len(rp); j++ {
			for k := 0; k < len(rp[j].BlockedList); k++ {
				if p.Name == rp[j].BlockedList[k].Name {
					return true
				}
			}
		}
	}
	return false
}

func (p *Process) removeFromBlockedQueue(blockedQueue Queue) Queue {
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
func (p *Process) ReleaseResource(rp ResourcePool, finishedQueue Queue) (ResourcePool, Queue) {
	for i := 0; i < len(p.RequestResArr); i++ {
		for j := 0; j < len(rp); j++ {
			if p.RequestResArr[i].Name == rp[j].Name {
				rp[j].Available++
				log.Printf("%s Already release resource, %s", p.Name, rp[j].Name)

				if i == len(p.RequestResArr)-1 {
					p.Status = "finished"
					finishedQueue = append(finishedQueue, *p)
					log.Printf("Process %s has already finished, put it to finished queue", p.Name)

					// 检测该blocked队列
					if rp[j].Available > 0 && len(rp[j].BlockedList) > 0 {
						rp[j].Available--
						//　检测所有blocked队列
						judge := rp.detectBlockedQueue(*p)
						if judge == false {
							rp[j].BlockedList[0].Status = "finished"
							finishedQueue = append(finishedQueue, rp[j].BlockedList[0])
							rp[j].BlockedList = rp[j].BlockedList[0].removeFromBlockedQueue(rp[j].BlockedList)
						}
					}
					return rp, finishedQueue
				}
			}
		}
	}
	return rp, finishedQueue
}
