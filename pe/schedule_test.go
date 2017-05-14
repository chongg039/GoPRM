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
		parent    *PCB
		children  *PCB
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
				parent:    nil,
				children:  nil,
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
				parent:    tt.fields.parent,
				children:  tt.fields.children,
			}
			if got := p.detectAllResourceStatus(); got != tt.want {
				t.Errorf("PCB.detectAllResourceStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPCBPool_Schedule(t *testing.T) {
	var (
		pcb1, pcb2       *PCB
		pcb1ele, pcb2ele *PCBEle
	)
	pool := new(PCBPool)
	pcb1 = &PCB{
		Name:      "PCB1",
		PID:       1000,
		Status:    "running",
		Priority:  2,
		CPUState:  "notused",
		Memory:    "notused",
		ReqResArr: []RequestResource{},
		parent:    nil,
		children:  pcb2,
	}
	pcb2 = &PCB{
		Name:      "PCB2",
		PID:       2000,
		Status:    "ready",
		Priority:  2,
		CPUState:  "notused",
		Memory:    "notused",
		ReqResArr: []RequestResource{},
		parent:    pcb1,
		children:  nil,
	}

	pcb1ele = &PCBEle{
		Data: *pcb1,
		next: pcb2ele,
	}
	pcb2ele = &PCBEle{
		Data: *pcb2,
		next: nil,
	}
	pool[2].head = pcb1ele
	pool[2].length = 2

	tests := []struct {
		name    string
		pcbPool *PCBPool
		want    *Running
	}{
		{
			name:    "test1",
			pcbPool: pool,
			want: &Running{
				Process: *pcb1,
				Start:   time.Now(),
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
