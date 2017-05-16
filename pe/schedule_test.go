package pe

import (
	"reflect"
	"testing"
	"time"
)

func TestPCBPool_TimeRotation(t *testing.T) {
	tests := []struct {
		name    string
		pcbPool *PCBPool
	}{
		{
			name:    "test1",
			pcbPool: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pcbPool.TimeRotation()
		})
	}
}

func TestPCB_detectAllResourceStatus(t *testing.T) {
	type fields struct {
		Name      string
		PID       int
		Status    string
		Priority  int
		CPUState  string
		Memory    string
		ReqResArr []RequestResource
		Parent    *PCB
		Children  *PCB
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "tests",
			fields: fields{ // 匿名struct可以作为map的value而且赋值方便，而匿名struct作为struct的成员则在初始化时 需要带上struct的名字才行
				Name:      "test1",
				PID:       1000,
				Status:    "ready",
				Priority:  0,
				CPUState:  "notused",
				Memory:    "notused",
				ReqResArr: []RequestResource{},
				Parent:    nil,
				Children:  nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PCB{
				Name:      tt.fields.Name,
				PID:       tt.fields.PID,
				Status:    tt.fields.Status,
				Priority:  tt.fields.Priority,
				CPUState:  tt.fields.CPUState,
				Memory:    tt.fields.Memory,
				ReqResArr: tt.fields.ReqResArr,
				Parent:    tt.fields.Parent,
				Children:  tt.fields.Children,
			}
			if got := p.detectAllResourceStatus(); got != tt.want {
				t.Errorf("PCB.detectAllResourceStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPCBPool_Schedule(t *testing.T) {

	pcb1 := new(PCB)
	pcb2 := new(PCB)
	pcb1ele := new(PCBEle)
	pcb2ele := new(PCBEle)

	pool := new(PCBPool)
	pcb1 = &PCB{
		Name:      "PCB1",
		PID:       1000,
		Status:    "ready",
		Priority:  2,
		CPUState:  "notused",
		Memory:    "notused",
		ReqResArr: []RequestResource{},
		Parent:    nil,
		Children:  pcb2,
	}
	pcb2 = &PCB{
		Name:      "PCB2",
		PID:       2000,
		Status:    "ready",
		Priority:  2,
		CPUState:  "notused",
		Memory:    "notused",
		ReqResArr: []RequestResource{},
		Parent:    pcb1,
		Children:  nil,
	}

	rpcb := pcb1
	rpcb.Status = "running"
	rpcb.CPUState = "using"
	rpcb.Memory = "using"

	// 在引用前不分配内存则不会引用的问题怎么解决？
	// new一个对象创建一个分配了内存的指针，这个对象中的值被初始化为0，
	// var 定义一个指针对象并不会为它分配一个内存空间，为nil（0x0）
	// 没有C++中的构造函数，对象的创建一般交给一个全局的创建函数来完成，返回的局部变量地址在函数返回后依然存在
	pcb1ele.Data = *pcb1
	pcb1ele.Next = pcb2ele

	pcb2ele.Data = *pcb2
	pcb2ele.Next = pcb1ele

	pool[2].Head = pcb1ele
	pool[2].Length = 2

	s := time.Now().Format("2006-01-02 15:04:05")
	start, _ := time.Parse("2006-01-02 15:04:05", s)

	tests := []struct {
		name    string
		pcbPool *PCBPool
		want    *Running
	}{
		{
			name:    "test1",
			pcbPool: pool,
			want: &Running{
				Process: *rpcb,
				Start:   start,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pcbPool.Schedule(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PCBPool.Schedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
