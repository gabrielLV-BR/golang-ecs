package main

import (
	"fmt"
	"gabriellv/main/ecs"
)

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

func main() {
	world := ecs.World{}.New()

	entity := world.NewEntity()
	world.AddComponent(entity, Velocity{10, 15, 20})
	world.AddComponent(entity, Name{"Charles"})

	entity2 := world.NewEntity()
	world.AddComponent(entity2, Name{"Daniel"})

	entity3 := world.NewEntity()
	world.AddComponent(entity3, Name{"Tupac"})
	world.AddComponent(entity3, Velocity{0, 100, 0})

	entity4 := world.NewEntity()
	world.AddComponent(entity4, Velocity{-1, -1, 0})

	query := ecs.Query{}.With(ecs.Id[Velocity]()).With(ecs.Id[Name]())

	for entity, components := range world.Run(query) {
		velocity, _ := ecs.Get[Velocity](query, components)
		name, _ := ecs.Get[Name](query, components)

		if velocity.Y > 50 {
			world.AddComponent(entity4, Name{"Dave"})
		}

		fmt.Printf(
			"Entidade %d: Velocity { %f, %f, %f }, Name { %s }\n",
			entity,
			velocity.X, velocity.Y, velocity.Z,
			name.Name)
	}
}
