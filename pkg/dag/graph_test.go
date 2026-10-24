package dag

import (
	"testing"
)

func TestGraphWorkflow(t *testing.T) {
	g := NewGraph()

	// Ajouter des tâches fictives avec dépendances
	g.AddTask("1", "Conception", 60, nil)
	g.AddTask("2", "Backend Dev", 120, []string{"1"})
	g.AddTask("3", "Frontend Dev", 90, []string{"1"})
	g.AddTask("4", "Tests & Déploiement", 30, []string{"2", "3"})

	// Tester le Tri Topologique
	sorted, err := g.TopologicalSort()
	if err != nil {
		t.Fatalf("Erreur inattendue lors du tri topologique : %v", err)
	}
	if len(sorted) != 4 {
		t.Errorf("Attendu 4 tâches triées, obtenu %d", len(sorted))
	}

	// Tester le Chemin Critique
	path, totalDuration, err := g.CriticalPath()
	if err != nil {
		t.Fatalf("Erreur inattendue lors du calcul du chemin critique : %v", err)
	}

	if totalDuration != 210 { // 60 (Conception) + 120 (Backend) + 30 (Tests) = 210
		t.Errorf("Durée critique attendue 210, obtenu %d", totalDuration)
	}

	if len(path) == 0 {
		t.Errorf("Le chemin critique ne doit pas être vide")
	}
}
