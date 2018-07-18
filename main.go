package main

import (
	"fmt"
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
	fmt.Printf("%v+", buf)
}