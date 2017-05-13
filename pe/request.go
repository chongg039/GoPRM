package pe

import "log"

// 进程有一个写入blocked队列或执行（输出到running队列）的函数

// RequestResource lets process request resources
func (p *PCB) RequestResource(rcbPool *RCBPool, pcbPool *PCBPool, rname ...string) (*RCBPool, *PCBPool) {

	var rs RequestResource

	for i := 0; i < len(rname); i++ {
		rs.Name = rname[i]
		rs.OK = false
		p.ReqResArr = append(p.ReqResArr, rs)
	}

	for i := 0; i < len(p.ReqResArr); i++ {
		for j := 0; j < len(*rcbPool); j++ {

			if (*rcbPool)[j].Name == p.ReqResArr[i].Name && (*rcbPool)[j].Available > 0 {
				(*rcbPool)[j].Available--
				p.ReqResArr[i].OK = true
				log.Printf("%s request resource %s OK!", p.Name, p.ReqResArr[i].Name)

				// judge := p.detectAllResourceStatus()
				// if judge == true {
				// 	//模拟进程运行时间
				// 	p.running()
				// 	// 从就绪队列中删除
				// 	readyQueues[p.Priority] = p.removeFromReadyQueue(readyQueues[p.Priority])
				// 	// 释放资源
				// 	rp, finishedQueue = p.ReleaseResource(rp, finishedQueue)
				// 	log.Printf("finished queue : %+v\n", finishedQueue)

				// 	return rp, finishedQueue, readyQueues
				// }
				break
			} else if (*rcbPool)[j].Name == p.ReqResArr[i].Name && (*rcbPool)[j].Available == 0 {
				log.Printf("no available %s resource, add to waiting list", p.ReqResArr[i].Name)
				// 从就绪队列中删除
				pcbPool.RemovePCBEle(p)
				// 插入某资源的waitinglist
				p.Status = "blocked"
				(*rcbPool)[j].BlockedList = append((*rcbPool)[j].BlockedList, *p)
				// 从ready队列中移除
				return rcbPool, pcbPool
			} else {
				continue
			}
		}
		// 应该设置标志位判断请求是否合法
	}
	return rcbPool, pcbPool
}
