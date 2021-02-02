package hhmmss

import "testing"

func Test__HhMmSs_ToHhMmSsString(t *testing.T) {
	hms, _ := NewHhMmSs("21:25:20")
	tests := []struct {
		name string
		hms  _HhMmSs
		want string
	}{
		{name: "Test__HhMmSs_ToHhMmSsString", hms: hms, want: "21:25:20"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hms.ToHhMmSsString(); got != tt.want {
				t.Errorf("ToHhMmSsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
