package bootstrap

import (
	"log"
)

type DroneConfig struct {
	Id     string
	Type    string
}


func Run(dcs []DroneConfig){
	for _, dc := range dcs {
		log.Printf("%+v\n", dc)
	}
} 