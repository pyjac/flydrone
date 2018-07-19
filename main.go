package main

import (
	_ "flydrone/drone"
	"flydrone/bootstrap"
	"io/ioutil"
	"os"
	yaml "gopkg.in/yaml.v2"
)


func main() {
	var conf bootstrap.DroneConfig
	reader, _ := os.Open("config.yaml")
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &conf)
	bootstrap.Run(conf)
	select {} // block forever
}