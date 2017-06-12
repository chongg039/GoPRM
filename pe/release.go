package pe

import (
	"time"
)

// time sleep 2 second to simulate process running
func (p *PCB) running() {
	time.Sleep(time.Second * 2)
}

// 判定资源池中是否有某进程仍在某资源的blocked队列中
func (rcbPool *RCBPool) detectBlockedQueue(p *PCB) bool {
	if len(p.ReqResArr) == 0 {
		return false
	}
	for i := 0; i < len(p.ReqResArr); i++ {
		if len(*rcbPool) == 0 {
			return false
		}
		for j := 0; j < len(*rcbPool); j++ {
			if len((*rcbPool)[j].BlockedList) == 0 {
				continue
			}
			for k := 0; k < len((*rcbPool)[j].BlockedList); k++ {
				if p.Name == (*rcbPool)[j].BlockedList[k].Name {
					return true
				}
			}
		}
	}
	return false
}

// func (blockedQueue *Queue) removeFromBlockedQueue(p *PCB) {
// 	for k, v := range *blockedQueue {
// 		if p.Name == v.Name {
// 			r := (*blockedQueue)[:k]
// 			l := (*blockedQueue)[k+1:]
// 			*blockedQueue = append(r, l...)
// 		}
// 	}
// }

// ReleaseResource should be used when process is already running
// 每次释放资源resource检测自身是否可用（>0），并从blocked队列中取出放入running
func (p *PCB) ReleaseResource(rcbPool *RCBPool, pcbPool *PCBPool) (*RCBPool, *PCBPool) {
	for i := 0; i < len(p.ReqResArr); i++ {
		if p.ReqResArr[i].OK == true {
			p.ReqResArr[i].OK = false
			for j := 0; j < len(*rcbPool); j++ {
				if (*rcbPool)[j].Name == p.ReqResArr[i].Name {
					(*rcbPool)[j].Available++

					// 检测blocked队列，让第一个执行request
					if len((*rcbPool)[j].BlockedList) != 0 {
						h := (*rcbPool)[j].BlockedList[0]
						s := []string{}
						rcbPool, pcbPool = h.RequestResource(rcbPool, pcbPool, s...)
						break
					}
					break
				}
				continue
			}
		}
	}
	p.ReqResArr = []RequestResource{}
	return rcbPool, pcbPool
}
