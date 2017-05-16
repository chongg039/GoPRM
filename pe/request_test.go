package pe

import (
	"reflect"
	"testing"
)

func TestPCB_RequestResource(t *testing.T) {
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
	type args struct {
		rcbPool *RCBPool
		pcbPool *PCBPool
		rname   []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RCBPool
		want1  *PCBPool
	}{
	// TODO: Add test cases.
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
