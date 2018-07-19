package bootstrap

import (
	"log"
)

type DroneConfig struct {
	Id     string
	Type    string
}

type Drone interface {
	Move()
	X() uint32
	Y() uint32
	GetSpeed() uint32
	Id() string
}

func Run(dcs []DroneConfig){
	for _, dc := range dcs {
		log.Printf("%+v\n", dc)
	}
} 