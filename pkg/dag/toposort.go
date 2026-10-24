package dag

import (
	"errors"
)

// TopologicalSort effectue un tri topologique et détecte les cycles potentiels
func (g *Graph) TopologicalSort() ([]string, error) {
	inDegree := make(map[string]int)
	for id := range g.Nodes {
		inDegree[id] = 0
	}

	// Calculer le degré d'entrée (in-degree) pour chaque nœud
	for _, node := range g.Nodes {
		for _, dep := range node.Dependencies {
			if _, exists := g.Nodes[dep]; !exists {
				return nil, errors.New("dépendance manquante : la tâche parente n'existe pas dans le graphe")
			}
			inDegree[node.ID]++
		}
	}

	// Queue pour stocker les nœuds ayant un in-degree de 0
	var queue []string
	for id, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, id)
		}
	}

	var sortedOrder []string

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, current)

		// Parcourir les tâches et réduire le degré d'entrée des enfants
		for _, node := range g.Nodes {
			for _, dep := range node.Dependencies {
				if dep == current {
					inDegree[node.ID]--
					if inDegree[node.ID] == 0 {
						queue = append(queue, node.ID)
					}
				}
			}
		}
	}

	// Si le nombre d'éléments triés est différent du total, il y a une boucle (cycle)
	if len(sortedOrder) != len(g.Nodes) {
		return nil, errors.New("cycle détecté dans le graphe : dépendances circulaires invalides")
	}

	return sortedOrder, nil
}
