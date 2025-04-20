package ecs

// A generic function that runs based on a query.
type System func(entity Entity, components ...Component)

// A system coupled with the query that describes its arguments.
type SystemDetails struct {
	Query Query
	Run   System
}

// A storage system for the systems.
type SystemStorage struct {
	systems []SystemDetails
}

// Initializes the systems storage.
func (storage SystemStorage) New() SystemStorage {
	storage.systems = make([]SystemDetails, 0)
	return storage
}

// Adds a system to the storage.
func (storage *SystemStorage) AddSystem(system SystemDetails) {
	storage.systems = append(storage.systems, system)
}
