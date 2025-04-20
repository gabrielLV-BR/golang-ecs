package main

import (
	"gabriellv/main/ecs"
)

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

	ecs.AddSystem2(&world, PrintVelocityAndName)

	world.Run()
}
