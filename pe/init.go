package pe

// InitReadyQueuesArr create queues array consist of queue which process's status is "ready"
// and priority is of 0, 1, 2
func (p *Process) InitReadyQueuesArr() (readyQueuesArr QueuesArr) {
	return
}

// InitResource initialize resource
func (r *Resource) InitResource(name string, t int) Resource {
	r.Name = name
	r.Total = t
	r.Available = t

	return *r
}

// InitResourcePool consist of resources
func (rp *ResourcePool) InitResourcePool(rs ...Resource) ResourcePool {
	*rp = rs
	return *rp
}

// 建立 running 输出队列

// blocked 队列在分别对应的 resource 中

// InitProcess should be used when start, and it creates a "Init" process
func (p *Process) InitProcess(name string) (readyQueuesArr QueuesArr) {
	const level int = 0

	readyQueuesArr = p.InitReadyQueuesArr()

	proc := p.CreateProcess(name, level, "R1")

	readyQueuesArr[proc.Priority] = append(readyQueuesArr[p.Priority], proc)

	return
}
