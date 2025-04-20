package main

import "gabriellv/main/ecs"

type Velocity struct {
	X, Y, Z float32
}

func (vel Velocity) Id() ecs.ComponentId {
	return 0
}

type Name struct {
	Name string
}

func (name Name) Id() ecs.ComponentId {
	return 1
}
