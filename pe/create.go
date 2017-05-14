package pe

import (
	"math/rand"
	"time"
)

// CreateRCB initialize resource
func CreateRCB(name string, t int) *RCB {
	return &RCB{
		Name:        name,
		Total:       t,
		Available:   t,
		BlockedList: Queue{},
	}
}

/*
CreatePCB means create a process
*/
func (p *PCB) CreatePCB(name string, level int) *PCB {
	const (
		status string = "ready"
		cpu    string = "notused"
		memory string = "notused"
	)

	pid := generateRandomPID()

	return &PCB{
		Name:      name,
		PID:       pid,
		Status:    status,
		Priority:  level,
		CPUState:  cpu,
		Memory:    memory,
		ReqResArr: []RequestResource{},
		parent:    p,
		children:  nil,
	}
}

// AppendPCBEle replace method "append"
func (pcbPool *PCBPool) AppendPCBEle(p *PCB) {

	var pcbele *PCBEle
	pcbele = new(PCBEle)
	pcbele.Data = *p

	if pcbPool[p.Priority].head == nil {
		pcbPool[p.Priority].head = pcbele
	} else {
		var h = pcbPool[p.Priority].head
		for h.next != nil {
			h = h.next
		}
		h.next = pcbele
	}
	pcbPool[p.Priority].length++

}

// RemovePCBEle remove PCB ele from PCBPool
func (pcbPool *PCBPool) RemovePCBEle(p *PCB) bool {
	//n为当前节点，h为前一个节点，初始状态为同一个节点
	var h, n = pcbPool[p.Priority].head, pcbPool[p.Priority].head

	if pcbPool[p.Priority].head == n && n.Data.Name == p.Name {
		pcbPool[p.Priority].head = pcbPool[p.Priority].head.next
		pcbPool[p.Priority].length--
		return true
	}
	for n != nil {
		if n.Data.Name == p.Name {
			//由于有垃圾回收，所以不用考虑释放内存
			h.next = n.next
			pcbPool[p.Priority].length--
			return true
		}
		h = n
		n = n.next
	}
	return false
}

// generateRandomPID gets the PID in the range of 2000~15000
func generateRandomPID() (pid int) {
	min := 2000
	max := 15000

	rand.Seed(time.Now().Unix())
	pid = rand.Intn(max-min) + min
	return
}
