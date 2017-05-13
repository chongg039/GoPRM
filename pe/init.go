package pe

// InitPCBPool create queues array consist of queue which process's status is "ready"
// and priority is of 0, 1, 2
func (p *PCB) InitPCBPool() *PCBPool {
	var pcbPool PCBPool
	return &pcbPool
}

// InitRCBPool consist of resources
func InitRCBPool(rs ...RCB) *RCBPool {
	var rcbPool RCBPool = rs
	return &rcbPool
}

// 建立 running 输出队列

// blocked 队列在分别对应的 resource 中

// Init 创建初始进程和R1，R2，R3，R4四种资源，并返回PCB池和RCB池的地址
func Init() (*PCB, *PCBPool, *RCBPool, *Queue, *Running) {
	const (
		name   string = "InitPCB"
		pid    int    = 1024
		status        = "ready"
		level  int    = 0
		cpu    string = "notused"
		memory string = "notused"
	)

	// 初始化Init进程
	initPCB := &PCB{
		Name:      name,
		PID:       1024,
		Status:    status,
		Priority:  level,
		CPUState:  cpu,
		Memory:    memory,
		ReqResArr: []RequestResource{},
		parent:    nil,
		children:  make([]interface{}, 5),
	}

	// 初始化PCB池
	pcbPool := initPCB.InitPCBPool()

	// 将Init进程放入进程池
	pcbPool.AppendPCBEle(initPCB)

	// 建立四个资源R1，R2，R3，R4
	R1 := CreateRCB("R1", 1)
	R2 := CreateRCB("R2", 2)
	R3 := CreateRCB("R3", 3)
	R4 := CreateRCB("R4", 4)

	// 初始化RCB池
	rcbPool := InitRCBPool(*R1, *R2, *R3, *R4)

	// 初始化finish队列
	finish := &Queue{}

	// 调度schedule，将Init进程移到running块中
	running := pcbPool.Schedule()

	return initPCB, pcbPool, rcbPool, finish, running
}
