// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Predecessor_WOP struct {
	// insertion point
	Name string
	DependencyType DependencyType
}

func (from *Predecessor) CopyBasicFields(to *Predecessor) {
	// insertion point
	to.Name = from.Name
	to.DependencyType = from.DependencyType
}

type Project_WOP struct {
	// insertion point
	Name string
}

func (from *Project) CopyBasicFields(to *Project) {
	// insertion point
	to.Name = from.Name
}

type Task_WOP struct {
	// insertion point
	Name string
}

func (from *Task) CopyBasicFields(to *Task) {
	// insertion point
	to.Name = from.Name
}

