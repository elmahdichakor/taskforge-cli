package dag

// CriticalPath calcule le chemin critique du DAG pour identifier la tâche prioritaire majeure
func (g *Graph) CriticalPath() ([]string, int, error) {
	sorted, err := g.TopologicalSort()
	if err != nil {
		return nil, 0, err
	}

	// Stocker la durée totale la plus longue pour atteindre chaque nœud (Earliest Finish)
	dist := make(map[string]int)
	predecessor := make(map[string]string)

	maxDuration := 0
	var criticalEndNode string

	for _, id := range sorted {
		node := g.Nodes[id]
		currentDist := dist[id] + node.Duration

		if currentDist > maxDuration {
			maxDuration = currentDist
			criticalEndNode = id
		}

		// Propager la durée aux tâches suivantes
		for _, otherNode := range g.Nodes {
			for _, dep := range otherNode.Dependencies {
				if dep == id {
					if dist[id]+node.Duration > dist[otherNode.ID] {
						dist[otherNode.ID] = dist[id] + node.Duration
						predecessor[otherNode.ID] = id
					}
				}
			}
		}
	}

	// Reconstruire le chemin critique à l'envers
	var path []string
	curr := criticalEndNode
	for curr != "" {
		path = append([]string{curr}, path...)
		curr = predecessor[curr]
	}

	return path, maxDuration, nil
}
