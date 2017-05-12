package pe

// InitPCB create queues array consist of queue which process's status is "ready"
// and priority is of 0, 1, 2
func (p *PCB) InitPCB() (pcbPool PCBPool) {
	return
}

// InitResource initialize resource
func (r *RCB) InitResource(name string, t int) RCB {
	r.Name = name
	r.Total = t
	r.Available = t

	return *r
}

// InitRCB consist of resources
func (rcbPool *RCBPool) InitRCB(rs ...RCB) RCBPool {
	*rcbPool = rs
	return *rcbPool
}

// 建立 running 输出队列

// blocked 队列在分别对应的 resource 中

// InitProcess should be used when start, and it creates a "Init" process
func (p *PCB) InitProcess(name string) (pcbPool PCBPool) {
	const level int = 0

	pcbPool = p.InitPCB()

	proc := p.CreateProcess(name, level, "R1")

	pcbPool[proc.Priority] = append(pcbPool[p.Priority], proc)

	return
}
