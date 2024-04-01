// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Predecessor:
		// insertion point per field

	case *Project:
		// insertion point per field
		if fieldName == "Tasks" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Project) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Project)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Tasks).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Tasks = _inferedTypeInstance.Tasks[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Tasks =
								append(_inferedTypeInstance.Tasks, any(fieldInstance).(*Task))
						}
					}
				}
			}
		}

	case *Task:
		// insertion point per field
		if fieldName == "Predecessors" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Task) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Task)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Predecessors).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Predecessors = _inferedTypeInstance.Predecessors[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Predecessors =
								append(_inferedTypeInstance.Predecessors, any(fieldInstance).(*Predecessor))
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Predecessor
	// insertion point per field

	// Compute reverse map for named struct Project
	// insertion point per field
	clear(stage.Project_Tasks_reverseMap)
	stage.Project_Tasks_reverseMap = make(map[*Task]*Project)
	for project := range stage.Projects {
		_ = project
		for _, _task := range project.Tasks {
			stage.Project_Tasks_reverseMap[_task] = project
		}
	}

	// Compute reverse map for named struct Task
	// insertion point per field
	clear(stage.Task_Predecessors_reverseMap)
	stage.Task_Predecessors_reverseMap = make(map[*Predecessor]*Task)
	for task := range stage.Tasks {
		_ = task
		for _, _predecessor := range task.Predecessors {
			stage.Task_Predecessors_reverseMap[_predecessor] = task
		}
	}

}
