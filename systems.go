package main

import (
	"fmt"
	"gabriellv/main/ecs"
)

func PrintVelocityAndName(entity ecs.Entity, velocity Velocity, name Name) {
	fmt.Printf(
		"Entidade %d: Velocity { %f, %f, %f }, Name { %s }\n",
		entity,
		velocity.X, velocity.Y, velocity.Z,
		name.Name)
}
