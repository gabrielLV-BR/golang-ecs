package ecs

// TODO: Generate these automatically. Involves bulking up the build system so will stay like this for a while.

// Adds a system which takes a single component to the system registry.
func AddSystem1[T1 Component](world *World, system func(entity Entity, component1 T1)) {
	query := Query{}.With(Id[T1]())
	callback := func(entity Entity, components ...Component) {
		c1, _ := Get[T1](query, components)
		system(entity, c1)
	}

	details := SystemDetails{query, callback}

	world.AddSystem(details)
}

// Adds a system which takes two components to the system registry.
func AddSystem2[T1 Component, T2 Component](world *World, system func(entity Entity, component1 T1, component2 T2)) {
	query := Query{}.With(Id[T1](), Id[T2]())
	callback := func(entity Entity, components ...Component) {
		c1, _ := Get[T1](query, components)
		c2, _ := Get[T2](query, components)
		system(entity, c1, c2)
	}

	details := SystemDetails{query, callback}

	world.AddSystem(details)
}

// Adds a system which takes three components to the system registry.
func AddSystem3[T1 Component, T2 Component, T3 Component](world *World, system func(entity Entity, component1 T1, component2 T2, component3 T3)) {
	query := Query{}.With(Id[T1](), Id[T2](), Id[T3]())
	callback := func(entity Entity, components ...Component) {
		c1, _ := Get[T1](query, components)
		c2, _ := Get[T2](query, components)
		c3, _ := Get[T3](query, components)
		system(entity, c1, c2, c3)
	}

	details := SystemDetails{query, callback}

	world.AddSystem(details)
}
