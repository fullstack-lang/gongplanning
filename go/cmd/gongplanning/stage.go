package main

import (
	"time"

	"github.com/fullstack-lang/gongplanning/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage models.StageStruct
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the _ gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage"] = _
// }

// _ will stage objects of database "stage"
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Predecessor

	// Declarations of staged instances of Project
	__Project__000000_Test2 := (&models.Project{Name: `Test2`}).Stage(stage)

	// Declarations of staged instances of Task

	// Setup of values

	// Project values setup
	__Project__000000_Test2.Name = `Test2`
	__Project__000000_Test2.IsExpanded = true

	// Setup of pointers
}


