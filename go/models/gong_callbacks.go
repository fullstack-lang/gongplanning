// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Predecessor:
		if stage.OnAfterPredecessorCreateCallback != nil {
			stage.OnAfterPredecessorCreateCallback.OnAfterCreate(stage, target)
		}
	case *Project:
		if stage.OnAfterProjectCreateCallback != nil {
			stage.OnAfterProjectCreateCallback.OnAfterCreate(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskCreateCallback != nil {
			stage.OnAfterTaskCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Predecessor:
		newTarget := any(new).(*Predecessor)
		if stage.OnAfterPredecessorUpdateCallback != nil {
			stage.OnAfterPredecessorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Project:
		newTarget := any(new).(*Project)
		if stage.OnAfterProjectUpdateCallback != nil {
			stage.OnAfterProjectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Task:
		newTarget := any(new).(*Task)
		if stage.OnAfterTaskUpdateCallback != nil {
			stage.OnAfterTaskUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Predecessor:
		if stage.OnAfterPredecessorDeleteCallback != nil {
			staged := any(staged).(*Predecessor)
			stage.OnAfterPredecessorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Project:
		if stage.OnAfterProjectDeleteCallback != nil {
			staged := any(staged).(*Project)
			stage.OnAfterProjectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Task:
		if stage.OnAfterTaskDeleteCallback != nil {
			staged := any(staged).(*Task)
			stage.OnAfterTaskDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Predecessor:
		if stage.OnAfterPredecessorReadCallback != nil {
			stage.OnAfterPredecessorReadCallback.OnAfterRead(stage, target)
		}
	case *Project:
		if stage.OnAfterProjectReadCallback != nil {
			stage.OnAfterProjectReadCallback.OnAfterRead(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskReadCallback != nil {
			stage.OnAfterTaskReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Predecessor:
		stage.OnAfterPredecessorUpdateCallback = any(callback).(OnAfterUpdateInterface[Predecessor])
	
	case *Project:
		stage.OnAfterProjectUpdateCallback = any(callback).(OnAfterUpdateInterface[Project])
	
	case *Task:
		stage.OnAfterTaskUpdateCallback = any(callback).(OnAfterUpdateInterface[Task])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Predecessor:
		stage.OnAfterPredecessorCreateCallback = any(callback).(OnAfterCreateInterface[Predecessor])
	
	case *Project:
		stage.OnAfterProjectCreateCallback = any(callback).(OnAfterCreateInterface[Project])
	
	case *Task:
		stage.OnAfterTaskCreateCallback = any(callback).(OnAfterCreateInterface[Task])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Predecessor:
		stage.OnAfterPredecessorDeleteCallback = any(callback).(OnAfterDeleteInterface[Predecessor])
	
	case *Project:
		stage.OnAfterProjectDeleteCallback = any(callback).(OnAfterDeleteInterface[Project])
	
	case *Task:
		stage.OnAfterTaskDeleteCallback = any(callback).(OnAfterDeleteInterface[Task])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Predecessor:
		stage.OnAfterPredecessorReadCallback = any(callback).(OnAfterReadInterface[Predecessor])
	
	case *Project:
		stage.OnAfterProjectReadCallback = any(callback).(OnAfterReadInterface[Project])
	
	case *Task:
		stage.OnAfterTaskReadCallback = any(callback).(OnAfterReadInterface[Task])
	
	}
}
