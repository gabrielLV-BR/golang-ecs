package ecs

import "slices"

// Uniquely identifies a component.
type ComponentId uint32

// Identifies where the component is stored.
type ComponentIndex uint32

// Describes the minimum contract for a component to implement.
type Component interface {
	Id() ComponentId
}

// Returns the Id of the component.
func Id[T Component]() ComponentId {
	var component T
	return component.Id()
}

// Stores the concrete components.
type ComponentStorage struct {
	components []Component
}

// Initializes the component storage.
func (storage ComponentStorage) New() ComponentStorage {
	storage.components = make([]Component, 0)
	return storage
}

// Adds a new component to the storage.
func (storage *ComponentStorage) Store(component Component) ComponentIndex {
	index := len(storage.components)

	storage.components = append(storage.components, component)

	return ComponentIndex(index)
}

// Retrieves the stored component from the storage.
func (storage *ComponentStorage) Retrieve(componentIndex ComponentIndex) Component {
	return storage.components[componentIndex]
}

// Removes the component from the storage.
func (storage *ComponentStorage) Remove(componentIndex ComponentIndex) {
	deleteStart := int(componentIndex)
	deleteEnd := deleteStart + 1
	storage.components = slices.Delete(storage.components, deleteStart, deleteEnd)
}

// Tell whether the storage is empty or not.
func (storage *ComponentStorage) Empty() bool {
	return len(storage.components) == 0
}
