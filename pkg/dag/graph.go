package dag

import (
	"fmt"
)

// TaskNode représente une tâche dans le DAG avec ses dépendances
type TaskNode struct {
	ID           string   // Identifiant unique de la tâche
	Title        string   // Titre ou description courte
	Duration     int      // Durée estimée en minutes
	Dependencies []string // Liste des IDs des tâches parentes requises
	Children     []string // Liste des IDs des tâches enfants dépendantes de celle-ci
}

// Graph représente le Graphe Orienté Acyclique (DAG)
type Graph struct {
	Nodes map[string]*TaskNode
}

// NewGraph initialise un nouveau graphe vide
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*TaskNode),
	}
}

// AddTask ajoute une tâche au graphe
func (g *Graph) AddTask(id, title string, duration int, dependencies []string) error {
	if _, exists := g.Nodes[id]; exists {
		return fmt.Errorf("la tâche avec l'ID %s existe déjà", id)
	}

	g.Nodes[id] = &TaskNode{
		ID:           id,
		Title:        title,
		Duration:     duration,
		Dependencies: dependencies,
		Children:     []string{},
	}
	return nil
}
