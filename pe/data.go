package pe

import (
	"time"
)

/*
PCB struct introduction
*/
type PCB struct {
	// name
	Name string
	// PID
	PID int
	// status: ready, running, blocked
	Status string
	// priority: 0, 1, 2 (Init, User, System)
	Priority int
	// when status is "running" CPU state is "using", others are "not used"
	CPUState string
	// when status is "running" memory is "using", others are "not used"
	Memory string
	// ResourceArr means process needed resources
	ReqResArr []RequestResource
	// 父节点
	Parent *PCB
	// 子节点
	Children *PCB
}

// PCBEle means node of PCB in PCBLinkList
type PCBEle struct {
	Data PCB
	Next *PCBEle
}

// PCBLinkList is a linklist of PCBEle
type PCBLinkList struct {
	Length int
	Head   *PCBEle
}

// PCBPool consist of 0, 1, 2 Queue
type PCBPool [3]PCBLinkList

// Queue consist of finished PCB
type Queue []PCB

// RCB control block
type RCB struct {
	// source name: R1, R2, R3, R4
	Name string
	// initial number
	Total int
	// available number
	Available int
	// waiting(blocked) queue
	BlockedList Queue
}

// RCBPool consist of resources
type RCBPool []RCB

// RequestResource is one process needed resource
type RequestResource struct {
	Name string
	OK   bool
}

// RequestResArr is type array of RequestResource
// type RequestResArr []RequestResource

// Running is on behalf of a process which status is "running"
type Running struct {
	Process PCB
	Start   time.Time
}
