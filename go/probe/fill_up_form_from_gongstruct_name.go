// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongplanning/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Predecessor":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Predecessor Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PredecessorFormCallback(
			nil,
			probe,
			formGroup,
		)
		predecessor := new(models.Predecessor)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(predecessor, formGroup, probe)
	case "Project":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Project Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ProjectFormCallback(
			nil,
			probe,
			formGroup,
		)
		project := new(models.Project)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(project, formGroup, probe)
	case "Task":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Task Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TaskFormCallback(
			nil,
			probe,
			formGroup,
		)
		task := new(models.Task)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(task, formGroup, probe)
	}
	formStage.Commit()
}
