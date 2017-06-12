package pe

import "log"

// 进程有一个写入blocked队列或执行（输出到running队列）的函数

// 检测请求资源的PCB是否是某一资源block队列中的，是的话移除该blocked中的第一个元素
func (p *PCB) detectBlocked(blocked Queue) Queue {
	if len(blocked) != 0 && p == blocked[0] {
		blocked = blocked[1:]
		return blocked
	}
	return blocked
}

// RequestResource lets process request resources
func (p *PCB) RequestResource(rcbPool *RCBPool, pcbPool *PCBPool, rname ...string) (*RCBPool, *PCBPool) {

	if len(rname) == 0 {
		// 若不需要任何资源，检查所有资源是否准备就绪
		if len(p.ReqResArr) == 0 || p.detectAllResourceStatus() == true {
			log.Printf("OK, %s doesn't need any resource", p.Name)
			return rcbPool, pcbPool
		}
	}

	var rs RequestResource

	if len(p.ReqResArr) == 0 {
		for i := 0; i < len(rname); i++ {
			rs.Name = rname[i]
			rs.OK = false
			p.ReqResArr = append(p.ReqResArr, rs)
		}
	} else {
	Check:
		for i := 0; i < len(rname); i++ {
			for j := 0; j < len(p.ReqResArr); j++ {
				if rname[i] == p.ReqResArr[j].Name && p.ReqResArr[j].OK == true {
					log.Printf("already have requested resource %s", rname[i])
					continue Check
				}
			}
			rs.Name = rname[i]
			rs.OK = false
			p.ReqResArr = append(p.ReqResArr, rs)
		}
	}

	for i := 0; i < len(p.ReqResArr); i++ {
		for j := 0; j < len(*rcbPool); j++ {

			if (*rcbPool)[j].Name == p.ReqResArr[i].Name && p.ReqResArr[i].OK == false && (*rcbPool)[j].Available > 0 {
				(*rcbPool)[j].Available--
				p.ReqResArr[i].OK = true
				// 检测是不是blocked队列中的元素
				(*rcbPool)[j].BlockedList = p.detectBlocked((*rcbPool)[j].BlockedList)
				log.Printf("%s request resource %s OK!", p.Name, p.ReqResArr[i].Name)
				break
			} else if (*rcbPool)[j].Name == p.ReqResArr[i].Name && (*rcbPool)[j].Available == 0 && p.ReqResArr[i].OK == false {
				log.Printf("noreq available %s resource, add to waiting list", p.ReqResArr[i].Name)
				// 从就绪队列中删除
				pcbPool.RemovePCBEle(p)
				// 插入某资源的waitinglist
				p.Status = "blocked"
				(*rcbPool)[j].BlockedList = append((*rcbPool)[j].BlockedList, p)
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
