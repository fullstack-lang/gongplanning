// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Predecessors           map[*Predecessor]any
	Predecessors_mapString map[string]*Predecessor

	// insertion point for slice of pointers maps

	OnAfterPredecessorCreateCallback OnAfterCreateInterface[Predecessor]
	OnAfterPredecessorUpdateCallback OnAfterUpdateInterface[Predecessor]
	OnAfterPredecessorDeleteCallback OnAfterDeleteInterface[Predecessor]
	OnAfterPredecessorReadCallback   OnAfterReadInterface[Predecessor]

	Projects           map[*Project]any
	Projects_mapString map[string]*Project

	// insertion point for slice of pointers maps
	Project_Tasks_reverseMap map[*Task]*Project

	OnAfterProjectCreateCallback OnAfterCreateInterface[Project]
	OnAfterProjectUpdateCallback OnAfterUpdateInterface[Project]
	OnAfterProjectDeleteCallback OnAfterDeleteInterface[Project]
	OnAfterProjectReadCallback   OnAfterReadInterface[Project]

	Tasks           map[*Task]any
	Tasks_mapString map[string]*Task

	// insertion point for slice of pointers maps
	Task_Predecessors_reverseMap map[*Predecessor]*Task

	OnAfterTaskCreateCallback OnAfterCreateInterface[Task]
	OnAfterTaskUpdateCallback OnAfterUpdateInterface[Task]
	OnAfterTaskDeleteCallback OnAfterDeleteInterface[Task]
	OnAfterTaskReadCallback   OnAfterReadInterface[Task]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongplanning/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitPredecessor(predecessor *Predecessor)
	CheckoutPredecessor(predecessor *Predecessor)
	CommitProject(project *Project)
	CheckoutProject(project *Project)
	CommitTask(task *Task)
	CheckoutTask(task *Task)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Predecessors:           make(map[*Predecessor]any),
		Predecessors_mapString: make(map[string]*Predecessor),

		Projects:           make(map[*Project]any),
		Projects_mapString: make(map[string]*Project),

		Tasks:           make(map[*Task]any),
		Tasks_mapString: make(map[string]*Task),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Predecessor"] = len(stage.Predecessors)
	stage.Map_GongStructName_InstancesNb["Project"] = len(stage.Projects)
	stage.Map_GongStructName_InstancesNb["Task"] = len(stage.Tasks)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Predecessor"] = len(stage.Predecessors)
	stage.Map_GongStructName_InstancesNb["Project"] = len(stage.Projects)
	stage.Map_GongStructName_InstancesNb["Task"] = len(stage.Tasks)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts predecessor to the model stage
func (predecessor *Predecessor) Stage(stage *StageStruct) *Predecessor {
	stage.Predecessors[predecessor] = __member
	stage.Predecessors_mapString[predecessor.Name] = predecessor

	return predecessor
}

// Unstage removes predecessor off the model stage
func (predecessor *Predecessor) Unstage(stage *StageStruct) *Predecessor {
	delete(stage.Predecessors, predecessor)
	delete(stage.Predecessors_mapString, predecessor.Name)
	return predecessor
}

// UnstageVoid removes predecessor off the model stage
func (predecessor *Predecessor) UnstageVoid(stage *StageStruct) {
	delete(stage.Predecessors, predecessor)
	delete(stage.Predecessors_mapString, predecessor.Name)
}

// commit predecessor to the back repo (if it is already staged)
func (predecessor *Predecessor) Commit(stage *StageStruct) *Predecessor {
	if _, ok := stage.Predecessors[predecessor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPredecessor(predecessor)
		}
	}
	return predecessor
}

func (predecessor *Predecessor) CommitVoid(stage *StageStruct) {
	predecessor.Commit(stage)
}

// Checkout predecessor to the back repo (if it is already staged)
func (predecessor *Predecessor) Checkout(stage *StageStruct) *Predecessor {
	if _, ok := stage.Predecessors[predecessor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPredecessor(predecessor)
		}
	}
	return predecessor
}

// for satisfaction of GongStruct interface
func (predecessor *Predecessor) GetName() (res string) {
	return predecessor.Name
}

// Stage puts project to the model stage
func (project *Project) Stage(stage *StageStruct) *Project {
	stage.Projects[project] = __member
	stage.Projects_mapString[project.Name] = project

	return project
}

// Unstage removes project off the model stage
func (project *Project) Unstage(stage *StageStruct) *Project {
	delete(stage.Projects, project)
	delete(stage.Projects_mapString, project.Name)
	return project
}

// UnstageVoid removes project off the model stage
func (project *Project) UnstageVoid(stage *StageStruct) {
	delete(stage.Projects, project)
	delete(stage.Projects_mapString, project.Name)
}

// commit project to the back repo (if it is already staged)
func (project *Project) Commit(stage *StageStruct) *Project {
	if _, ok := stage.Projects[project]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProject(project)
		}
	}
	return project
}

func (project *Project) CommitVoid(stage *StageStruct) {
	project.Commit(stage)
}

// Checkout project to the back repo (if it is already staged)
func (project *Project) Checkout(stage *StageStruct) *Project {
	if _, ok := stage.Projects[project]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProject(project)
		}
	}
	return project
}

// for satisfaction of GongStruct interface
func (project *Project) GetName() (res string) {
	return project.Name
}

// Stage puts task to the model stage
func (task *Task) Stage(stage *StageStruct) *Task {
	stage.Tasks[task] = __member
	stage.Tasks_mapString[task.Name] = task

	return task
}

// Unstage removes task off the model stage
func (task *Task) Unstage(stage *StageStruct) *Task {
	delete(stage.Tasks, task)
	delete(stage.Tasks_mapString, task.Name)
	return task
}

// UnstageVoid removes task off the model stage
func (task *Task) UnstageVoid(stage *StageStruct) {
	delete(stage.Tasks, task)
	delete(stage.Tasks_mapString, task.Name)
}

// commit task to the back repo (if it is already staged)
func (task *Task) Commit(stage *StageStruct) *Task {
	if _, ok := stage.Tasks[task]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTask(task)
		}
	}
	return task
}

func (task *Task) CommitVoid(stage *StageStruct) {
	task.Commit(stage)
}

// Checkout task to the back repo (if it is already staged)
func (task *Task) Checkout(stage *StageStruct) *Task {
	if _, ok := stage.Tasks[task]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTask(task)
		}
	}
	return task
}

// for satisfaction of GongStruct interface
func (task *Task) GetName() (res string) {
	return task.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMPredecessor(Predecessor *Predecessor)
	CreateORMProject(Project *Project)
	CreateORMTask(Task *Task)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMPredecessor(Predecessor *Predecessor)
	DeleteORMProject(Project *Project)
	DeleteORMTask(Task *Task)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Predecessors = make(map[*Predecessor]any)
	stage.Predecessors_mapString = make(map[string]*Predecessor)

	stage.Projects = make(map[*Project]any)
	stage.Projects_mapString = make(map[string]*Project)

	stage.Tasks = make(map[*Task]any)
	stage.Tasks_mapString = make(map[string]*Task)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Predecessors = nil
	stage.Predecessors_mapString = nil

	stage.Projects = nil
	stage.Projects_mapString = nil

	stage.Tasks = nil
	stage.Tasks_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for predecessor := range stage.Predecessors {
		predecessor.Unstage(stage)
	}

	for project := range stage.Projects {
		project.Unstage(stage)
	}

	for task := range stage.Tasks {
		task.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
	// insertion point for generic types
	Predecessor | Project | Task
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	// insertion point for generic types
	*Predecessor | *Project | *Task
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any |
		// insertion point for generic types
		map[*Predecessor]any |
		map[*Project]any |
		map[*Task]any |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

type GongstructMapString interface {
	map[any]any |
		// insertion point for generic types
		map[string]*Predecessor |
		map[string]*Project |
		map[string]*Task |
		map[*any]any // because go does not support an extra "|" at the end of type specifications
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Predecessor]any:
		return any(&stage.Predecessors).(*Type)
	case map[*Project]any:
		return any(&stage.Projects).(*Type)
	case map[*Task]any:
		return any(&stage.Tasks).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Predecessor:
		return any(&stage.Predecessors_mapString).(*Type)
	case map[string]*Project:
		return any(&stage.Projects_mapString).(*Type)
	case map[string]*Task:
		return any(&stage.Tasks_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Predecessor:
		return any(&stage.Predecessors).(*map[*Type]any)
	case Project:
		return any(&stage.Projects).(*map[*Type]any)
	case Task:
		return any(&stage.Tasks).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Predecessor:
		return any(&stage.Predecessors).(*map[Type]any)
	case *Project:
		return any(&stage.Projects).(*map[Type]any)
	case *Task:
		return any(&stage.Tasks).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Predecessor:
		return any(&stage.Predecessors_mapString).(*map[string]*Type)
	case Project:
		return any(&stage.Projects_mapString).(*map[string]*Type)
	case Task:
		return any(&stage.Tasks_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Predecessor:
		return any(&Predecessor{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Task: &Task{Name: "Task"},
		}).(*Type)
	case Project:
		return any(&Project{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			Tasks: []*Task{{Name: "Tasks"}},
		}).(*Type)
	case Task:
		return any(&Task{
			// Initialisation of associations
			// field is initialized with an instance of Predecessor with the name of the field
			Predecessors: []*Predecessor{{Name: "Predecessors"}},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Predecessor
	case Predecessor:
		switch fieldname {
		// insertion point for per direct association field
		case "Task":
			res := make(map[*Task][]*Predecessor)
			for predecessor := range stage.Predecessors {
				if predecessor.Task != nil {
					task_ := predecessor.Task
					var predecessors []*Predecessor
					_, ok := res[task_]
					if ok {
						predecessors = res[task_]
					} else {
						predecessors = make([]*Predecessor, 0)
					}
					predecessors = append(predecessors, predecessor)
					res[task_] = predecessors
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Project
	case Project:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Predecessor
	case Predecessor:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Project
	case Project:
		switch fieldname {
		// insertion point for per direct association field
		case "Tasks":
			res := make(map[*Task]*Project)
			for project := range stage.Projects {
				for _, task_ := range project.Tasks {
					res[task_] = project
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
		case "Predecessors":
			res := make(map[*Predecessor]*Task)
			for task := range stage.Tasks {
				for _, predecessor_ := range task.Predecessors {
					res[predecessor_] = task
				}
			}
			return any(res).(map[*End]*Start)
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Predecessor:
		res = "Predecessor"
	case Project:
		res = "Project"
	case Task:
		res = "Task"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Predecessor:
		res = "Predecessor"
	case *Project:
		res = "Project"
	case *Task:
		res = "Task"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Predecessor:
		res = []string{"Name", "DependencyType", "Task"}
	case Project:
		res = []string{"Name", "Tasks"}
	case Task:
		res = []string{"Name", "Predecessors"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Predecessor:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Task"
		rf.Fieldname = "Predecessors"
		res = append(res, rf)
	case Project:
		var rf ReverseField
		_ = rf
	case Task:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Project"
		rf.Fieldname = "Tasks"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Predecessor:
		res = []string{"Name", "DependencyType", "Task"}
	case *Project:
		res = []string{"Name", "Tasks"}
	case *Task:
		res = []string{"Name", "Predecessors"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Predecessor:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DependencyType":
			enum := inferedInstance.DependencyType
			res = enum.ToCodeString()
		case "Task":
			if inferedInstance.Task != nil {
				res = inferedInstance.Task.Name
			}
		}
	case *Project:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Tasks":
			for idx, __instance__ := range inferedInstance.Tasks {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Task:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Predecessors":
			for idx, __instance__ := range inferedInstance.Predecessors {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Predecessor:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DependencyType":
			enum := inferedInstance.DependencyType
			res = enum.ToCodeString()
		case "Task":
			if inferedInstance.Task != nil {
				res = inferedInstance.Task.Name
			}
		}
	case Project:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Tasks":
			for idx, __instance__ := range inferedInstance.Tasks {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Task:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Predecessors":
			for idx, __instance__ := range inferedInstance.Predecessors {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
