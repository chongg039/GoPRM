package pe

import (
	"reflect"
	"testing"
)

// func TestRCBPool_detectBlockedQueue(t *testing.T) {
// 	rr := RequestResource{
// 		Name: "R1",
// 		OK:   false,
// 	}
// 	pcb1 := &PCB{
// 		Name:      "pcb1",
// 		PID:       1000,
// 		Status:    "blocked",
// 		Priority:  1,
// 		CPUState:  "notused",
// 		Memory:    "notused",
// 		ReqResArr: []RequestResource{rr},
// 		Parent:    nil,
// 		Children:  nil,
// 	}
// 	pcb2 := &PCB{
// 		Name:      "pcb2",
// 		PID:       2000,
// 		Status:    "blocked",
// 		Priority:  2,
// 		CPUState:  "notused",
// 		Memory:    "notused",
// 		ReqResArr: []RequestResource{rr},
// 		Parent:    nil,
// 		Children:  nil,
// 	}
// 	rcbpool := new(RCBPool)
// 	r1 := &RCB{
// 		Name:        "R1",
// 		Total:       1,
// 		Available:   0,
// 		BlockedList: []PCB{*pcb1},
// 	}
// 	*rcbpool = append(*rcbpool, *r1)

// 	type args struct {
// 		p *PCB
// 	}
// 	tests := []struct {
// 		name    string
// 		rcbPool *RCBPool
// 		args    args
// 		want    bool
// 	}{
// 		{
// 			name:    "test1",
// 			rcbPool: rcbpool,
// 			args: args{
// 				p: pcb1,
// 			},
// 			want: true,
// 		},
// 		{
// 			name:    "test2",
// 			rcbPool: rcbpool,
// 			args: args{
// 				p: pcb2,
// 			},
// 			want: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.rcbPool.detectBlockedQueue(tt.args.p); got != tt.want {
// 				t.Errorf("RCBPool.detectBlockedQueue() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestPCB_ReleaseResource(t *testing.T) {
	// type fields struct {
	// 	Name      string
	// 	PID       int
	// 	Status    string
	// 	Priority  int
	// 	CPUState  string
	// 	Memory    string
	// 	ReqResArr []RequestResource
	// 	Parent    *PCB
	// 	Children  *PCB
	// }
	type args struct {
		rcbPool *RCBPool
		pcbPool *PCBPool
	}
	pcb1 := &PCB{
		Name:     "pcb1",
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
		Name:     "pcb2",
		PID:      2000,
		Status:   "ready",
		Priority: 1,
		CPUState: "notused",
		Memory:   "notused",
		ReqResArr: []RequestResource{
			{
				Name: "R1",
				OK:   false,
			},
		},
		Parent:   nil,
		Children: nil,
	}
	r1 := &RCB{
		Name:        "R1",
		Total:       1,
		Available:   0,
		BlockedList: Queue{pcb2},
	}
	a := args{
		rcbPool: &RCBPool{*r1},
		pcbPool: &PCBPool{},
	}
	tests := []struct {
		name   string
		fields *PCB
		args   args
		want   *RCBPool
		want1  *PCBPool
	}{
		{
			name:   "test1",
			fields: pcb1,
			args:   a,
			want:   a.rcbPool,
			want1:  a.pcbPool,
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
			got, got1 := p.ReleaseResource(tt.args.rcbPool, tt.args.pcbPool)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PCB.ReleaseResource() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PCB.ReleaseResource() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
