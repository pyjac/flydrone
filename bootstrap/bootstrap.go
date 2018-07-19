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

type droneFactory func(droneId string) Drone

var dronesFactories = make(map[string]droneFactory)

func Run(dcs []DroneConfig){
	for _, dc := range dcs {
		log.Printf("%+v\n", dc)
	}
}

func Register(droneType string, df droneFactory) {
	if _, exists := dronesFactories[droneType]; exists {
		log.Println(droneType, "Drone Factory already registered")
		return
	}

	log.Println("Register", droneType, "drone")
	dronesFactories[droneType] = df
}