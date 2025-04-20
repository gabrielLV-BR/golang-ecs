package ecs

import (
	"iter"
)

// Holds together the different structures and methods needed for interaction with the ECS system.
type World struct {
	EntityStorage
	SystemStorage
}

// Initializes the ECS world.
func (world World) New() World {
	world.EntityStorage = world.EntityStorage.New()

	return world
}

// Runs the registered systems.
func (world *World) Run() {
	for _, system := range world.systems {
		for entity, components := range world.Execute(system.Query) {
			system.Run(entity, components...)
		}
	}
}

// Runs the supplied query and returns an iterator with the results.
// TODO: This should be refactored as it sucks right now.
// TODO: This could use some goroutines.
// TODO: This could use some caching.
func (world World) Execute(query Query) iter.Seq2[Entity, []Component] {
	return func(yield func(Entity, []Component) bool) {
		for entity, components := range world.entityColumns {
			presentComponentIds := make([]ComponentId, 0, len(components.ComponentIndices))
			for componentId := range components.ComponentIndices {
				presentComponentIds = append(presentComponentIds, componentId)
			}

			queriedComponentIds, matches := query.Matches(presentComponentIds)

			if !matches {
				continue
			}

			componentInstances := make([]Component, 0, len(queriedComponentIds))

			for _, componentId := range queriedComponentIds {
				componentIndex, ok := components.ComponentIndices[componentId]

				if !ok {
					panic("Component did not exist")
				}

				componentStorage, ok := world.componentColumns[componentId]

				if !ok {
					panic("Queried component Id was not registered")
				}

				component := componentStorage.Retrieve(componentIndex)

				componentInstances = append(componentInstances, component)
			}

			if !yield(entity, componentInstances) {
				return
			}
		}
	}
}
