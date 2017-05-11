package pe

import "time"

/*
Process struct introduction
*/
type Process struct {
	// name
	Name string `json:"ProcName"`
	// PID
	PID int `json:"PID"`
	// status: ready, running, blocked
	Status string `json:"Status"`
	// priority: 0, 1, 2 (Init, User, System)
	Priority int `json:"Priority"`
	// when status is "running" CPU state is "using", others are "not used"
	CPUState string `json:"CPUState"`
	// when status is "running" memory is "using", others are "not used"
	Memory string `json:"Memory"`
	// ResourceArr means process needed resources
	RequestResArr RequestResArr `json:"RequestResArr"`
}

// Queue of process
type Queue []Process

// QueuesArr consist of 0, 1, 2 Queue
type QueuesArr [3]Queue

// Resource control block
type Resource struct {
	// source name: R1, R2, R3, R4
	Name string
	// initial number
	Total int
	// available number
	Available int
	// waiting(blocked) queue
	BlockedList Queue
}

// ResourcePool consist of resources
type ResourcePool []Resource

// RequestResource is one process needed resource
type RequestResource struct {
	Name string
	OK   bool
}

// RequestResArr is type array of RequestResource
type RequestResArr []RequestResource

// Running is on behalf of a process which status is "running"
type Running struct {
	Process Process
	Start   time.Time
}
