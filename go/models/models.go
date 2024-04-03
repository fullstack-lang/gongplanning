package models

type Task struct {
	Name         string
	Predecessors []*Predecessor
}

type Project struct {
	Name  string
	Tasks []*Task

	IsExpanded bool
}

type Predecessor struct {
	Name           string
	DependencyType DependencyType
	Task           *Task
}

type DependencyType string

// Finish to Start (FS): A task must finish before the next task can start. This is the most common type of dependency.
// Start to Start (SS): A task must start before or at the same time as another task.
// Finish to Finish (FF): A task must finish at the same time as another task.
// Start to Finish (SF): A task must start before another task can finish (this is the least common type).
const (
	FS DependencyType = "FS"
	// SS DependencyType = "SS"
	// FF DependencyType = "FF"
	// SF DependencyType = "SF"
)
