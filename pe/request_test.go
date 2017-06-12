package pe

import (
	"reflect"
	"testing"
)

func TestPCB_RequestResource(t *testing.T) {
	pcbele := new(PCBEle)
	pcbpool := new(PCBPool)

	rcbpool := new(RCBPool)
	r1 := &RCB{
		Name:        "R1",
		Total:       1,
		Available:   1,
		BlockedList: Queue{},
	}
	pcb := &PCB{
		Name:      "testPCB",
		PID:       1000,
		Status:    "ready",
		Priority:  1,
		CPUState:  "notused",
		Memory:    "notused",
		ReqResArr: []RequestResource{},
		Parent:    nil,
		Children:  nil,
	}
	pcbele.Data = pcb
	pcbpool[1].Head = pcbele
	*rcbpool = append(*rcbpool, *r1)

	type args struct {
		rcbPool *RCBPool
		pcbPool *PCBPool
		rname   []string
	}
	tests := []struct {
		name   string
		fields PCB
		args   args
		want   *RCBPool
		want1  *PCBPool
	}{
		{
			name:   "test",
			fields: *pcb,
			args: args{
				rcbPool: rcbpool,
				pcbPool: pcbpool,
				rname:   []string{"R1"},
			},
			want:  rcbpool,
			want1: pcbpool,
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
			got, got1 := p.RequestResource(tt.args.rcbPool, tt.args.pcbPool, tt.args.rname...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PCB.RequestResource() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PCB.RequestResource() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
