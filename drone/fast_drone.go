package drone

import (
	"flydrone/bootstrap"
)

type fastDrone struct {
	drone
}

func newFastDrone(ID string) bootstrap.Drone {
	return &fastDrone{drone{id: ID, minSpeed: 200, maxSpeed: 400 }}
}

func init() {
	bootstrap.Register("fast_drone", newFastDrone)
}