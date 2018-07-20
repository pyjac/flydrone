package bootstrap

import (
	"fmt"
	"reflect"
	"testing"
	"flydrone/bootstrap/mocks"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func Test_newDroneMqttWebSocketClient(t *testing.T) {
	type args struct {
		websocketPath string
		droneID       string
	}
	tests := []struct {
		name    string
		args    args
		want    MQTT.Client
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newDroneMqttWebSocketClient(tt.args.websocketPath, tt.args.droneID)
			if (err != nil) != tt.wantErr {
				t.Errorf("newDroneMqttWebSocketClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDroneMqttWebSocketClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	mockDroneFactory := func (ID string) Drone {
		return &mocks.Drone{}
	}
	dronesFactories["super_drone"] = mockDroneFactory
	type args struct {
		droneType string
		df        droneFactory
	}
	tests := []struct {
		name string
		args args
		err error
	}{
		{
			name: "Testing Registering new Drone",
			args: args{
				droneType: "ultra_drone",
				df: mockDroneFactory,
			},
			err: nil,
		},
		{
			name: "Testing Registering new Drone",
			args: args{
				droneType: "ultra_drone",
				df: mockDroneFactory,
			},
			err: fmt.Errorf("ultra_drone Drone Factory already registered"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Register(tt.args.droneType, tt.args.df)
			if err != tt.err && err.Error() != tt.err.Error() {
				t.Errorf("Expected %v, got %v", err, tt.err)
			}
		})
	}
}
