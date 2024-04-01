package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongplanning/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default models.StageStruct
var ___dummy__Time_Default time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.DependencyType": ref_models.DependencyType(""),

	"ref_models.FS": ref_models.FS,

	"ref_models.Predecessor": &(ref_models.Predecessor{}),

	"ref_models.Predecessor.DependencyType": (ref_models.Predecessor{}).DependencyType,

	"ref_models.Predecessor.Name": (ref_models.Predecessor{}).Name,

	"ref_models.Predecessor.Task": (ref_models.Predecessor{}).Task,

	"ref_models.Project": &(ref_models.Project{}),

	"ref_models.Project.Name": (ref_models.Project{}).Name,

	"ref_models.Project.Tasks": (ref_models.Project{}).Tasks,

	"ref_models.Task": &(ref_models.Task{}),

	"ref_models.Task.Name": (ref_models.Task{}).Name,

	"ref_models.Task.Predecessors": (ref_models.Task{}).Predecessors,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_DependencyType := (&models.Field{Name: `DependencyType`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000002_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_Default_DependencyType := (&models.GongEnumShape{Name: `Default-DependencyType`}).Stage(stage)

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_FS := (&models.GongEnumValueEntry{Name: `FS`}).Stage(stage)

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Predecessor := (&models.GongStructShape{Name: `Default-Predecessor`}).Stage(stage)
	__GongStructShape__000001_Default_Project := (&models.GongStructShape{Name: `Default-Project`}).Stage(stage)
	__GongStructShape__000002_Default_Task := (&models.GongStructShape{Name: `Default-Task`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Predecessors := (&models.Link{Name: `Predecessors`}).Stage(stage)
	__Link__000001_Task := (&models.Link{Name: `Task`}).Stage(stage)
	__Link__000002_Tasks := (&models.Link{Name: `Tasks`}).Stage(stage)

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_DependencyType := (&models.Position{Name: `Pos-Default-DependencyType`}).Stage(stage)
	__Position__000001_Pos_Default_Predecessor := (&models.Position{Name: `Pos-Default-Predecessor`}).Stage(stage)
	__Position__000002_Pos_Default_Project := (&models.Position{Name: `Pos-Default-Project`}).Stage(stage)
	__Position__000003_Pos_Default_Task := (&models.Position{Name: `Pos-Default-Task`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Predecessor_and_Default_Task := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Predecessor and Default-Task`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Project_and_Default_Task := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Project and Default-Task`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Task_and_Default_Predecessor := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Task and Default-Predecessor`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	// Field values setup
	__Field__000000_DependencyType.Name = `DependencyType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Predecessor.DependencyType]
	__Field__000000_DependencyType.Identifier = `ref_models.Predecessor.DependencyType`
	__Field__000000_DependencyType.FieldTypeAsString = ``
	__Field__000000_DependencyType.Structname = `Predecessor`
	__Field__000000_DependencyType.Fieldtypename = `DependencyType`

	// Field values setup
	__Field__000001_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Project.Name]
	__Field__000001_Name.Identifier = `ref_models.Project.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Project`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Predecessor.Name]
	__Field__000002_Name.Identifier = `ref_models.Predecessor.Name`
	__Field__000002_Name.FieldTypeAsString = ``
	__Field__000002_Name.Structname = `Predecessor`
	__Field__000002_Name.Fieldtypename = `string`

	// GongEnumShape values setup
	__GongEnumShape__000000_Default_DependencyType.Name = `Default-DependencyType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DependencyType]
	__GongEnumShape__000000_Default_DependencyType.Identifier = `ref_models.DependencyType`
	__GongEnumShape__000000_Default_DependencyType.Width = 240.000000
	__GongEnumShape__000000_Default_DependencyType.Height = 78.000000

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000000_FS.Name = `FS`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.DependencyType.FS]
	__GongEnumValueEntry__000000_FS.Identifier = `ref_models.DependencyType.FS`

	// GongStructShape values setup
	__GongStructShape__000000_Default_Predecessor.Name = `Default-Predecessor`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Predecessor]
	__GongStructShape__000000_Default_Predecessor.Identifier = `ref_models.Predecessor`
	__GongStructShape__000000_Default_Predecessor.ShowNbInstances = true
	__GongStructShape__000000_Default_Predecessor.NbInstances = 0
	__GongStructShape__000000_Default_Predecessor.Width = 340.000000
	__GongStructShape__000000_Default_Predecessor.Height = 93.000000
	__GongStructShape__000000_Default_Predecessor.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_Project.Name = `Default-Project`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Project]
	__GongStructShape__000001_Default_Project.Identifier = `ref_models.Project`
	__GongStructShape__000001_Default_Project.ShowNbInstances = true
	__GongStructShape__000001_Default_Project.NbInstances = 0
	__GongStructShape__000001_Default_Project.Width = 240.000000
	__GongStructShape__000001_Default_Project.Height = 78.000000
	__GongStructShape__000001_Default_Project.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Default_Task.Name = `Default-Task`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Task]
	__GongStructShape__000002_Default_Task.Identifier = `ref_models.Task`
	__GongStructShape__000002_Default_Task.ShowNbInstances = true
	__GongStructShape__000002_Default_Task.NbInstances = 0
	__GongStructShape__000002_Default_Task.Width = 240.000000
	__GongStructShape__000002_Default_Task.Height = 63.000000
	__GongStructShape__000002_Default_Task.IsSelected = false

	// Link values setup
	__Link__000000_Predecessors.Name = `Predecessors`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Task.Predecessors]
	__Link__000000_Predecessors.Identifier = `ref_models.Task.Predecessors`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Predecessor]
	__Link__000000_Predecessors.Fieldtypename = `ref_models.Predecessor`
	__Link__000000_Predecessors.FieldOffsetX = -50.000000
	__Link__000000_Predecessors.FieldOffsetY = -16.000000
	__Link__000000_Predecessors.TargetMultiplicity = models.MANY
	__Link__000000_Predecessors.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_Predecessors.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_Predecessors.SourceMultiplicity = models.MANY
	__Link__000000_Predecessors.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_Predecessors.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_Predecessors.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_Predecessors.StartRatio = 0.324236
	__Link__000000_Predecessors.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Predecessors.EndRatio = 0.500000
	__Link__000000_Predecessors.CornerOffsetRatio = 1.015873

	// Link values setup
	__Link__000001_Task.Name = `Task`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Predecessor.Task]
	__Link__000001_Task.Identifier = `ref_models.Predecessor.Task`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Task]
	__Link__000001_Task.Fieldtypename = `ref_models.Task`
	__Link__000001_Task.FieldOffsetX = -50.000000
	__Link__000001_Task.FieldOffsetY = -16.000000
	__Link__000001_Task.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_Task.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_Task.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_Task.SourceMultiplicity = models.MANY
	__Link__000001_Task.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_Task.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_Task.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Task.StartRatio = 0.500000
	__Link__000001_Task.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Task.EndRatio = 0.500000
	__Link__000001_Task.CornerOffsetRatio = 1.380000

	// Link values setup
	__Link__000002_Tasks.Name = `Tasks`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Project.Tasks]
	__Link__000002_Tasks.Identifier = `ref_models.Project.Tasks`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Task]
	__Link__000002_Tasks.Fieldtypename = `ref_models.Task`
	__Link__000002_Tasks.FieldOffsetX = -50.000000
	__Link__000002_Tasks.FieldOffsetY = -16.000000
	__Link__000002_Tasks.TargetMultiplicity = models.MANY
	__Link__000002_Tasks.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_Tasks.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_Tasks.SourceMultiplicity = models.MANY
	__Link__000002_Tasks.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_Tasks.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_Tasks.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Tasks.StartRatio = 0.500000
	__Link__000002_Tasks.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Tasks.EndRatio = 0.500000
	__Link__000002_Tasks.CornerOffsetRatio = 1.380000

	// Position values setup
	__Position__000000_Pos_Default_DependencyType.X = 65.000000
	__Position__000000_Pos_Default_DependencyType.Y = 196.000000
	__Position__000000_Pos_Default_DependencyType.Name = `Pos-Default-DependencyType`

	// Position values setup
	__Position__000001_Pos_Default_Predecessor.X = 698.999969
	__Position__000001_Pos_Default_Predecessor.Y = 187.000000
	__Position__000001_Pos_Default_Predecessor.Name = `Pos-Default-Predecessor`

	// Position values setup
	__Position__000002_Pos_Default_Project.X = 62.000000
	__Position__000002_Pos_Default_Project.Y = 36.000000
	__Position__000002_Pos_Default_Project.Name = `Pos-Default-Project`

	// Position values setup
	__Position__000003_Pos_Default_Task.X = 411.999969
	__Position__000003_Pos_Default_Task.Y = 44.000000
	__Position__000003_Pos_Default_Task.Name = `Pos-Default-Task`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Predecessor_and_Default_Task.X = 1001.499969
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Predecessor_and_Default_Task.Y = 123.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Predecessor_and_Default_Task.Name = `Verticle in class diagram Default in middle between Default-Predecessor and Default-Task`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Project_and_Default_Task.X = 596.999984
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Project_and_Default_Task.Y = 79.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Project_and_Default_Task.Name = `Verticle in class diagram Default in middle between Default-Project and Default-Task`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Task_and_Default_Predecessor.X = 915.499969
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Task_and_Default_Predecessor.Y = 147.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Task_and_Default_Predecessor.Name = `Verticle in class diagram Default in middle between Default-Task and Default-Predecessor`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Project)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Task)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Predecessor)
	__Classdiagram__000000_Default.GongEnumShapes = append(__Classdiagram__000000_Default.GongEnumShapes, __GongEnumShape__000000_Default_DependencyType)
	__GongEnumShape__000000_Default_DependencyType.Position = __Position__000000_Pos_Default_DependencyType
	__GongEnumShape__000000_Default_DependencyType.GongEnumValueEntrys = append(__GongEnumShape__000000_Default_DependencyType.GongEnumValueEntrys, __GongEnumValueEntry__000000_FS)
	__GongStructShape__000000_Default_Predecessor.Position = __Position__000001_Pos_Default_Predecessor
	__GongStructShape__000000_Default_Predecessor.Fields = append(__GongStructShape__000000_Default_Predecessor.Fields, __Field__000002_Name)
	__GongStructShape__000000_Default_Predecessor.Fields = append(__GongStructShape__000000_Default_Predecessor.Fields, __Field__000000_DependencyType)
	__GongStructShape__000000_Default_Predecessor.Links = append(__GongStructShape__000000_Default_Predecessor.Links, __Link__000001_Task)
	__GongStructShape__000001_Default_Project.Position = __Position__000002_Pos_Default_Project
	__GongStructShape__000001_Default_Project.Fields = append(__GongStructShape__000001_Default_Project.Fields, __Field__000001_Name)
	__GongStructShape__000001_Default_Project.Links = append(__GongStructShape__000001_Default_Project.Links, __Link__000002_Tasks)
	__GongStructShape__000002_Default_Task.Position = __Position__000003_Pos_Default_Task
	__GongStructShape__000002_Default_Task.Links = append(__GongStructShape__000002_Default_Task.Links, __Link__000000_Predecessors)
	__Link__000000_Predecessors.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Task_and_Default_Predecessor
	__Link__000001_Task.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Predecessor_and_Default_Task
	__Link__000002_Tasks.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Project_and_Default_Task
}


