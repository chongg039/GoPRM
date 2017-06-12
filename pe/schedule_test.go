package pe

import (
	"reflect"
	"testing"
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

	// 在引用前不分配内存则不会引用的问题怎么解决？
	// new一个对象创建一个分配了内存的指针，这个对象中的值被初始化为0，
	// var 定义一个指针对象并不会为它分配一个内存空间，为nil（0x0）
	// 没有C++中的构造函数，对象的创建一般交给一个全局的创建函数来完成，返回的局部变量地址在函数返回后依然存在
	pcb1ele.Data = pcb1
	pcb1ele.Next = pcb2ele

	pcb2ele.Data = pcb2
	pcb2ele.Next = pcb1ele

	pool[2].Head = pcb1ele
	pool[2].Length = 2

	type args struct {
		running *PCB
	}
	tests := []struct {
		name    string
		pcbPool *PCBPool
		args    args
		want    *PCB
	}{
		{
			name:    "test1",
			pcbPool: pool,
			args: args{
				running: nil,
			},
			want: pcb1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pcbPool.Schedule(tt.args.running); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PCBPool.Schedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPCBPool_TimeOut(t *testing.T) {
	pcb1 := &PCB{
		Name:     "PCB1",
		PID:      1000,
		Status:   "ready",
		Priority: 2,
		CPUState: "notused",
		Memory:   "notused",
		ReqResArr: []RequestResource{
			{
				Name: "R1",
				OK:   true,
			},
		},
		Parent:   nil,
		Children: nil,
	}
	pcb2 := &PCB{
		Name:      "PCB2",
		PID:       2000,
		Status:    "running",
		Priority:  1,
		CPUState:  "using",
		Memory:    "using",
		ReqResArr: []RequestResource{},
		Parent:    nil,
		Children:  nil,
	}
	pcb1ele := new(PCBEle)
	pool := new(PCBPool)
	pcb1ele.Data = pcb1
	pcb1ele.Next = nil
	pool[2].Head = pcb1ele
	pool[2].Length = 1

	finishlist := new(Queue)
	*finishlist = append(*finishlist, pcb2)

	type args struct {
		running *PCB
		finish  Queue
	}
	tests := []struct {
		name    string
		pcbPool *PCBPool
		args    args
		want    *PCB
		want1   Queue
	}{
		{
			name:    "test1",
			pcbPool: pool,
			args: args{
				running: pcb2,
				finish:  Queue{},
			},
			want:  pcb1,
			want1: *finishlist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.pcbPool.TimeOut(tt.args.running, tt.args.finish)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PCBPool.TimeOut() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PCBPool.TimeOut() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
