package pe

import (
	"log"
)

// DestoryPCB can destory process no matter the status running
func (pcbPool *PCBPool) DestoryPCB(rcbPool *RCBPool, p *PCB, finish Queue) (*RCBPool, *PCB, Queue) {
	if p.Status != "running" {
		log.Println("only running PCB can be destoried! exit")
		return rcbPool, p, finish
	}
	// 归还资源
	rcbPool, pcbPool = p.ReleaseResource(rcbPool, pcbPool)
	// 暂存running
	temp := p
	p = nil

	// 原running和其子PCB压入finish
	finish = append(finish, temp)
	for i := 0; i < len(temp.Children); i++ {
		tempc := temp.Children[i]
		// 在PCBPool中删除
		jug := pcbPool.RemovePCBEle(tempc)
		// 删除失败
		if jug == false {
			log.Println("Del failed! exit")
			return rcbPool, p, finish
		}
		finish = append(finish, tempc)
		//递归调用删除
		if len(tempc.Children) != 0 {
			for j := 0; j < len(tempc.Children); j++ {
				rcbPool, p, finish = pcbPool.DestoryPCB(rcbPool, tempc.Children[j], finish)
			}
		}
	}
	// 调度等待队列进入running
	p = pcbPool.Schedule(p)
	log.Printf("OK, PCB %p and its children has been destoried!", temp)
	return rcbPool, p, finish
}
