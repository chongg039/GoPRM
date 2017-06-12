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

// CreatePCB means create a process
func (p *PCB) CreatePCB(name string, level int) *PCB {
	const (
		status string = "ready"
		cpu    string = "notused"
		memory string = "notused"
	)

	pid := generateRandomPID()

	np := &PCB{
		Name:      name,
		PID:       pid,
		Status:    status,
		Priority:  level,
		CPUState:  cpu,
		Memory:    memory,
		ReqResArr: []RequestResource{},
		Parent:    p,
		Children:  nil,
	}
	p.Children = append(p.Children, np)
	return np
}

// AppendPCBEle append PCB ele to PCBPool
func (pcbPool *PCBPool) AppendPCBEle(p *PCB) {

	if pcbPool[p.Priority].Head == nil {
		pcbPool[p.Priority].Head = p
	} else {
		var h = pcbPool[p.Priority].Head
		for h.Next != nil {
			h = h.Next
		}
		h.Next = p
	}
	pcbPool[p.Priority].Length++

}

// RemovePCBEle remove PCB ele from PCBPool
func (pcbPool *PCBPool) RemovePCBEle(p *PCB) bool {
	if pcbPool[p.Priority].Head == nil {
		return true
	}
	// n := pcbPool[p.Priority].Head
	if pcbPool[p.Priority].Head.Next == nil && pcbPool[p.Priority].Head.Name == p.Name {
		pcbPool[p.Priority].Length--
		pcbPool[p.Priority].Head = nil
		return true
	}
	for pcbPool[p.Priority].Head.Next != nil {
		if pcbPool[p.Priority].Head.Name == p.Name {
			pcbPool[p.Priority].Head = pcbPool[p.Priority].Head.Next
			pcbPool[p.Priority].Length--
			return true
		}
		pcbPool[p.Priority].Head = pcbPool[p.Priority].Head.Next
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

// FindPCB 通过名字来查询对应的PCB地址
func (p *PCB) FindPCB(s string) *PCB {
	if p != nil && p.Name == s {
		return p
	}
	return nil
}
