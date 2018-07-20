package drone

import (
	"flydrone/bootstrap"
	"reflect"
	"testing"
)

func Test_newFastDrone(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name string
		args args
		want bootstrap.Drone
	}{
		{
			name: "Testing Fast Drone",
			args: args{ID: "FAST_ID"},
			want: &fastDrone{drone{id: "FAST_ID", minSpeed: 200, maxSpeed: 400 }},
		},
		{
			name: "Testing Fast Drone",
			args: args{ID: "FAST_ID_2"},
			want: &fastDrone{drone{id: "FAST_ID_2", minSpeed: 200, maxSpeed: 400 }},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFastDrone(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFastDrone() = %v, want %v", got, tt.want)
			}
		})
	}
}
