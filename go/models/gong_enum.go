// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for DependencyType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (dependencytype DependencyType) ToString() (res string) {

	// migration of former implementation of enum
	switch dependencytype {
	// insertion code per enum code
	case FS:
		res = "FS"
	}
	return
}

func (dependencytype *DependencyType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FS":
		*dependencytype = FS
		return
	default:
		return errUnkownEnum
	}
}

func (dependencytype *DependencyType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FS":
		*dependencytype = FS
	default:
		return errUnkownEnum
	}
	return
}

func (dependencytype *DependencyType) ToCodeString() (res string) {

	switch *dependencytype {
	// insertion code per enum code
	case FS:
		res = "FS"
	}
	return
}

func (dependencytype DependencyType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FS")

	return
}

func (dependencytype DependencyType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FS")

	return
}

// Utility function for FormNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (formnames FormNames) ToString() (res string) {

	// migration of former implementation of enum
	switch formnames {
	// insertion code per enum code
	case ModelForm:
		res = "Model Form"
	case ShapeForm:
		res = "Shape Form"
	}
	return
}

func (formnames *FormNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Model Form":
		*formnames = ModelForm
		return
	case "Shape Form":
		*formnames = ShapeForm
		return
	default:
		return errUnkownEnum
	}
}

func (formnames *FormNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ModelForm":
		*formnames = ModelForm
	case "ShapeForm":
		*formnames = ShapeForm
	default:
		return errUnkownEnum
	}
	return
}

func (formnames *FormNames) ToCodeString() (res string) {

	switch *formnames {
	// insertion code per enum code
	case ModelForm:
		res = "ModelForm"
	case ShapeForm:
		res = "ShapeForm"
	}
	return
}

func (formnames FormNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ModelForm")
	res = append(res, "ShapeForm")

	return
}

func (formnames FormNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Model Form")
	res = append(res, "Shape Form")

	return
}

// Utility function for StackName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stackname StackName) ToString() (res string) {

	// migration of former implementation of enum
	switch stackname {
	// insertion code per enum code
	case StackNameDefault:
		res = "gongplanning"
	}
	return
}

func (stackname *StackName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "gongplanning":
		*stackname = StackNameDefault
		return
	default:
		return errUnkownEnum
	}
}

func (stackname *StackName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "StackNameDefault":
		*stackname = StackNameDefault
	default:
		return errUnkownEnum
	}
	return
}

func (stackname *StackName) ToCodeString() (res string) {

	switch *stackname {
	// insertion code per enum code
	case StackNameDefault:
		res = "StackNameDefault"
	}
	return
}

func (stackname StackName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "StackNameDefault")

	return
}

func (stackname StackName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "gongplanning")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
