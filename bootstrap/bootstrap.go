package bootstrap

import (
	"log"
	"time"
	"crypto/tls"
	"fmt"
	"bytes"
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)


type DroneConfig struct {
	Websocket string
	Drones []DroneConfigData
}

type DroneConfigData struct {
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

type addDroneMessage struct {
	DroneID string `json:"id"`
}

type droneMoveMessage struct {
	DroneID string `json:"id"`
	X uint32 `json:"x"`
	Y uint32 `json:"y"`
	Speed uint32 `json:"s"`
}


type droneFactory func(droneId string) Drone

var dronesFactories = make(map[string]droneFactory)

func Run(dcs DroneConfig){
	for _, dc := range dcs.Drones {
		droneFactory, exists := dronesFactories[dc.Type]
		if !exists {
			log.Printf("%s does not exit", dc.Type)
			continue
		}
		drone := droneFactory(dc.Id)
		droneWebSocketClient, err := newDroneMqttWebSocketClient(dcs.Websocket, drone.Id())
		if err != nil {
			log.Printf("Cannot create socket for Drone")
			continue
		}
		go func(d Drone, client MQTT.Client){
			ticker := time.NewTicker(1 * time.Second)
			for ; true; <-ticker.C {
				d.Move()

				droneMovePayload := droneMoveMessage{
					DroneID: d.Id(),
					X: d.X(),
					Y: d.Y(),
					Speed: d.GetSpeed(),
				}
				reqBodyBytes := new(bytes.Buffer)
				json.NewEncoder(reqBodyBytes).Encode(droneMovePayload)
				client.Publish("DRONE_MOVE", byte(0), false, reqBodyBytes.Bytes())
			}
		}(drone, droneWebSocketClient)
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

func newDroneMqttWebSocketClient(websocketPath string, droneID string) (MQTT.Client, error) {
	connOpts := MQTT.NewClientOptions().AddBroker(websocketPath).SetClientID(droneID).SetCleanSession(true)
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)
 
	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return nil, token.Error()
	}
	fmt.Printf("Connected to %s\n", websocketPath)

	testStruct := addDroneMessage{droneID}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(testStruct)

	client.Publish("ADD_DRONE", byte(0), false, reqBodyBytes.Bytes())
	return client, nil
}