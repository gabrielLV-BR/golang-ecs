package ecs

// Holds the components that should be queried.
type Query struct {
	componentsToInclude []ComponentId
}

// Returns a new query with the added component Id.
func (query Query) With(componentId ComponentId) Query {
	if query.componentsToInclude == nil {
		query.componentsToInclude = make([]ComponentId, 0)
	}

	query.componentsToInclude = append(query.componentsToInclude, componentId)

	return query
}

// Gets the returned component from the given query.
func Get[T Component](query Query, components []Component) (T, bool) {
	componentId := Id[T]()

	for index, queriedId := range query.componentsToInclude {
		if queriedId == componentId {
			return components[index].(T), true
		}
	}

	panic("Queried value not found")
}

// Checks to see if the given component list matches the query.
func (query Query) Matches(componentIds []ComponentId) ([]ComponentId, bool) {

	presentComponents := make(map[ComponentId]any, len(componentIds))

	for _, component := range componentIds {
		presentComponents[component] = true
	}

	for _, requiredComponents := range query.componentsToInclude {
		if _, ok := presentComponents[requiredComponents]; !ok {
			return nil, false
		}
	}

	return query.componentsToInclude, true
}
