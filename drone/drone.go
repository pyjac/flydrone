package drone

import (
	"math/rand"
	"time"
)

type drone struct {
	id       string
	x        uint32
	y        uint32
	speed    uint32
	minSpeed uint32
	maxSpeed uint32
}

func (d drone) X() uint32 {
	return d.x
}

func (d drone) Id() string {
	return d.id
}

func (d drone) Y() uint32 {
	return d.y
}

func (d drone) GetSpeed() uint32 {
	return d.speed
}

func (d *drone) Move() {
	rand.Seed(time.Now().Unix())
	speed := uint32(rand.Intn(int(d.maxSpeed-d.minSpeed))) + d.minSpeed
	d.accelerate(speed)
}

func (d *drone) accelerate(speed uint32) {
	d.speed = speed
	d.x += speed
	d.y += speed
}
