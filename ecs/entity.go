package ecs

// Identifies an entity in the ECS world.
type Entity uint32

// Holds the components an entity is attached to.
// TODO: Create interface for this.
type EntityComponents struct {
	ComponentIndices map[ComponentId]ComponentIndex
}

// Initializes the entity components structure.
func (components EntityComponents) New() EntityComponents {
	components.ComponentIndices = make(map[ComponentId]ComponentIndex)
	return components
}

// Adds a component to the entity binding.
func (components *EntityComponents) AddComponent(componentId ComponentId, componentIndex ComponentIndex) {
	components.ComponentIndices[componentId] = componentIndex
}

// Removes a component from the entity binding.
func (components *EntityComponents) RemoveComponent(componentId ComponentId) {
	delete(components.ComponentIndices, componentId)
}

// Tells whether the entity is bound to any components or not.
func (components *EntityComponents) Empty() bool {
	return len(components.ComponentIndices) == 0
}
