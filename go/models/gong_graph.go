// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Predecessor:
		ok = stage.IsStagedPredecessor(target)

	case *Project:
		ok = stage.IsStagedProject(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedPredecessor(predecessor *Predecessor) (ok bool) {

	_, ok = stage.Predecessors[predecessor]

	return
}

func (stage *StageStruct) IsStagedProject(project *Project) (ok bool) {

	_, ok = stage.Projects[project]

	return
}

func (stage *StageStruct) IsStagedTask(task *Task) (ok bool) {

	_, ok = stage.Tasks[task]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Predecessor:
		stage.StageBranchPredecessor(target)

	case *Project:
		stage.StageBranchProject(target)

	case *Task:
		stage.StageBranchTask(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchPredecessor(predecessor *Predecessor) {

	// check if instance is already staged
	if IsStaged(stage, predecessor) {
		return
	}

	predecessor.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if predecessor.Task != nil {
		StageBranch(stage, predecessor.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchProject(project *Project) {

	// check if instance is already staged
	if IsStaged(stage, project) {
		return
	}

	project.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range project.Tasks {
		StageBranch(stage, _task)
	}

}

func (stage *StageStruct) StageBranchTask(task *Task) {

	// check if instance is already staged
	if IsStaged(stage, task) {
		return
	}

	task.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _predecessor := range task.Predecessors {
		StageBranch(stage, _predecessor)
	}

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Predecessor:
		toT := CopyBranchPredecessor(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Project:
		toT := CopyBranchProject(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Task:
		toT := CopyBranchTask(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchPredecessor(mapOrigCopy map[any]any, predecessorFrom *Predecessor) (predecessorTo *Predecessor) {

	// predecessorFrom has already been copied
	if _predecessorTo, ok := mapOrigCopy[predecessorFrom]; ok {
		predecessorTo = _predecessorTo.(*Predecessor)
		return
	}

	predecessorTo = new(Predecessor)
	mapOrigCopy[predecessorFrom] = predecessorTo
	predecessorFrom.CopyBasicFields(predecessorTo)

	//insertion point for the staging of instances referenced by pointers
	if predecessorFrom.Task != nil {
		predecessorTo.Task = CopyBranchTask(mapOrigCopy, predecessorFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchProject(mapOrigCopy map[any]any, projectFrom *Project) (projectTo *Project) {

	// projectFrom has already been copied
	if _projectTo, ok := mapOrigCopy[projectFrom]; ok {
		projectTo = _projectTo.(*Project)
		return
	}

	projectTo = new(Project)
	mapOrigCopy[projectFrom] = projectTo
	projectFrom.CopyBasicFields(projectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range projectFrom.Tasks {
		projectTo.Tasks = append(projectTo.Tasks, CopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func CopyBranchTask(mapOrigCopy map[any]any, taskFrom *Task) (taskTo *Task) {

	// taskFrom has already been copied
	if _taskTo, ok := mapOrigCopy[taskFrom]; ok {
		taskTo = _taskTo.(*Task)
		return
	}

	taskTo = new(Task)
	mapOrigCopy[taskFrom] = taskTo
	taskFrom.CopyBasicFields(taskTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _predecessor := range taskFrom.Predecessors {
		taskTo.Predecessors = append(taskTo.Predecessors, CopyBranchPredecessor(mapOrigCopy, _predecessor))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Predecessor:
		stage.UnstageBranchPredecessor(target)

	case *Project:
		stage.UnstageBranchProject(target)

	case *Task:
		stage.UnstageBranchTask(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchPredecessor(predecessor *Predecessor) {

	// check if instance is already staged
	if !IsStaged(stage, predecessor) {
		return
	}

	predecessor.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if predecessor.Task != nil {
		UnstageBranch(stage, predecessor.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchProject(project *Project) {

	// check if instance is already staged
	if !IsStaged(stage, project) {
		return
	}

	project.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range project.Tasks {
		UnstageBranch(stage, _task)
	}

}

func (stage *StageStruct) UnstageBranchTask(task *Task) {

	// check if instance is already staged
	if !IsStaged(stage, task) {
		return
	}

	task.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _predecessor := range task.Predecessors {
		UnstageBranch(stage, _predecessor)
	}

}
