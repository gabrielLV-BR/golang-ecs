package ecs

// Stores entities, components and the relation between them.
type EntityStorage struct {
	componentColumns map[ComponentId]ComponentStorage
	entityColumns    map[Entity]EntityComponents
}

// Initializes the entity storage struct.
func (storage EntityStorage) New() EntityStorage {
	storage.componentColumns = make(map[ComponentId]ComponentStorage)
	storage.entityColumns = make(map[Entity]EntityComponents)

	return storage
}

// Creates a new entity and returns its Id.
// TODO: This should probably employ generational Ids.
func (storage *EntityStorage) NewEntity() Entity {
	entityId := Entity(len(storage.entityColumns))

	storage.entityColumns[entityId] = EntityComponents{}.New()

	return entityId
}

// Deletes an entity and every component it had.
// TODO: This is terribly inneficient right now.
// TODO: The deleted entity Id should probably be kept somewhere for reuse.
func (storage *EntityStorage) DeleteEntity(entity Entity) bool {
	entityComponents, ok := storage.entityColumns[entity]

	if !ok {
		return false
	}

	for componentId := range entityComponents.ComponentIndices {
		storage.RemoveComponent(entity, componentId)
	}

	delete(storage.entityColumns, entity)

	return true
}

// Adds a component to the store and binds it to the entity.
func (storage *EntityStorage) AddComponent(entity Entity, component Component) {
	componentId := component.Id()

	componentStorage, ok := storage.componentColumns[componentId]

	if !ok {
		componentStorage = ComponentStorage{}.New()
	}

	componentIndex := componentStorage.Store(component)

	entityComponents, ok := storage.entityColumns[entity]

	if !ok {
		entityComponents = EntityComponents{}.New()
	}

	entityComponents.AddComponent(componentId, componentIndex)

	storage.entityColumns[entity] = entityComponents
	storage.componentColumns[componentId] = componentStorage
}

// Removes the component of an entity from the storage.
func (storage *EntityStorage) RemoveComponent(entity Entity, componentId ComponentId) bool {
	entityComponents, ok := storage.entityColumns[entity]

	if !ok {
		return false
	}

	componentIndex, ok := entityComponents.ComponentIndices[componentId]

	if !ok {
		return false
	}

	componentStorage, ok := storage.componentColumns[componentId]

	if !ok {
		panic("Registered component on entity was not stored")
	}

	entityComponents.RemoveComponent(componentId)
	componentStorage.Remove(componentIndex)

	// Clean up "header" structs if no other component is present.

	if entityComponents.Empty() {
		delete(storage.entityColumns, entity)
	} else {
		storage.entityColumns[entity] = entityComponents
	}

	if componentStorage.Empty() {
		delete(storage.componentColumns, componentId)
	} else {
		storage.componentColumns[componentId] = componentStorage
	}

	return true
}
