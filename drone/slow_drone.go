package drone

import (
	"flydrone/bootstrap"
)

type slowDrone struct {
	drone
}

func newSlowDrone(ID string) bootstrap.Drone {
	return &slowDrone{drone{id: ID, minSpeed: 50, maxSpeed: 60 }}
}

func init() {
	bootstrap.Register("slow_drone", newSlowDrone)
}