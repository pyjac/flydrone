package drone

import (
	"flydrone/bootstrap"
	"reflect"
	"testing"
)

func Test_newSlowDrone(t *testing.T) {
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
			args: args{ID: "SLOW_ID"},
			want: &slowDrone{drone{id: "SLOW_ID", minSpeed: 50, maxSpeed: 60 }},
		},
		{
			name: "Testing Fast Drone",
			args: args{ID: "SLOW_ID_2"},
			want: &slowDrone{drone{id: "SLOW_ID_2", minSpeed: 50, maxSpeed: 60 }},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSlowDrone(tt.args.ID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSlowDrone() = %v, want %v", got, tt.want)
			}
		})
	}
}
