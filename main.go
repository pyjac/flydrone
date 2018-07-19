package main

import (
	_ "flydrone/drone"
	"flydrone/bootstrap"
	"io/ioutil"
	"os"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Drones []bootstrap.DroneConfig
}

func main() {
	var conf Config
	reader, _ := os.Open("config.yaml")
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &conf)
	bootstrap.Run(conf.Drones)
	select {} // block forever
}