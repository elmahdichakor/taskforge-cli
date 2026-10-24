# TaskForge-CLI Engine 🚀

> Un moteur d'orchestration de tâches et de focus en ligne de commande (CLI/TUI) à haut niveau de concurrence, résilient et zéro-latence.

---

## 🛠️ Informations du Projet & Lead Architect

* **Nom du Projet :** TaskForge-CLI Engine (`taskforged` / `taskforge`)
* **Lead Systems Architect :** El Mahdi Chakor
* **Contact :** +212 694946773 | elmahdichakor0@gmail.com
* **Événement :** Hackathon Taskforge (Devpost)

---

## 🏗️ Architecture Système (High-Level)

TaskForge-CLI est conçu selon une architecture Client-Démon répartie en local avec communication IPC zéro-latence :
┌────────────────────────────────────────────────────────────────────────┐
│                      CLIENT TUI / CLI (taskforge)                     │
│  - Interface réactive en Terminal (Go / Bubbletea & Lipgloss)          │
│  - Render Engine non-bloquant & Raccourcis Async                      │
└───────────────────────────────────┬────────────────────────────────────┘
│ Communication Bi-directionnelle
│ (Unix Domain Sockets)
▼
┌────────────────────────────────────────────────────────────────────────┐
│                     DÉMON ARRIÈRE-PLAN (taskforged)                   │
│  - Engine Concurrent Go (Goroutines / Channels)                        │
│  - Algorithme DAG (Tri Topologique & Chemin Critique / CPM)            │
│  - Monitor System & Focus Enforcer                                     │
└───────────────────────────────────┬────────────────────────────────────┘
│ Direct I/O (Zero Latency)
▼
┌────────────────────────────────────────────────────────────────────────┐
│                      BASE DE DONNÉES LOCALES                           │
│  - SQLite3 en Mode WAL (Write-Ahead Logging)                           │
│  - Synchronisation Atomique & Zero-Lock                                │
└───────────────────────────────────┬────────────────────────────────────┘


---

## 🔥 Fonctionnalités Clés & Modules Avancés

1. **Architecture Client-Démon IPC (Go / Sockets Unix) :** Communication asynchrone bidirectionnelle via `/tmp/taskforge.sock`.
2. **Algorithme DAG & Critical Path Method (CPM) :** Modélisation des tâches sous forme de Graphe Orienté Acyclique avec détection automatique de cycles invalides et calcul de la tâche bloquante majeure.
3. **Moteur de Stockage Résilient (SQLite3 WAL) :** Base de données ultra-rapide garantissant des transactions atomiques sans verrouillage.
4. **Client TUI Réactif (Bubbletea / Lipgloss) :** Dashboard terminal interactif avec mise en page dynamique.

---

## 💻 Compilation & Installation

### Prérequis
* **Go 1.21+**

### Compilation Locale
```bash
git clone [https://github.com/elmahdichakor/taskforge-cli.git](https://github.com/elmahdichakor/taskforge-cli.git)
cd taskforge-cli

# Compilation des exécutables
go build -o taskforged ./cmd/taskforged
go build -o taskforge ./cmd/taskforge
Exécution
Bash
# 1. Lancer le démon en arrière-plan
./taskforged &

# 2. Lancer l'interface TUI interactive
./taskforge
📂 Multi-Plateforme (Binaires Pré-compilés)
Les exécutables binaires compilés sont disponibles dans le répertoire bin/ :

Linux (AMD64) : bin/linux/

macOS (ARM64 Apple Silicon) : bin/macos/

Windows (AMD64) : bin/windows/

Développé avec passion par El Mahdi Chakor pour le Hackathon Taskforge sur Devpost.
