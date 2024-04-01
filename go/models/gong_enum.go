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
	default:
		return errUnkownEnum
	}
	return
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

// Utility function for StackName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stackname StackName) ToString() (res string) {

	// migration of former implementation of enum
	switch stackname {
	// insertion code per enum code
	case StackNameDefault:
		res = "gongplanning-default"
	case StackNameTreeAtTheLeft:
		res = "gongplanning-tree-left"
	case StackNameSVGAtTheCenter:
		res = "gongplanning-svg-center"
	case StackNameFormAtTheRight:
		res = "gongplanning-form-right"
	}
	return
}

func (stackname *StackName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "gongplanning-default":
		*stackname = StackNameDefault
	case "gongplanning-tree-left":
		*stackname = StackNameTreeAtTheLeft
	case "gongplanning-svg-center":
		*stackname = StackNameSVGAtTheCenter
	case "gongplanning-form-right":
		*stackname = StackNameFormAtTheRight
	default:
		return errUnkownEnum
	}
	return
}

func (stackname *StackName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "StackNameDefault":
		*stackname = StackNameDefault
	case "StackNameTreeAtTheLeft":
		*stackname = StackNameTreeAtTheLeft
	case "StackNameSVGAtTheCenter":
		*stackname = StackNameSVGAtTheCenter
	case "StackNameFormAtTheRight":
		*stackname = StackNameFormAtTheRight
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
	case StackNameTreeAtTheLeft:
		res = "StackNameTreeAtTheLeft"
	case StackNameSVGAtTheCenter:
		res = "StackNameSVGAtTheCenter"
	case StackNameFormAtTheRight:
		res = "StackNameFormAtTheRight"
	}
	return
}

func (stackname StackName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "StackNameDefault")
	res = append(res, "StackNameTreeAtTheLeft")
	res = append(res, "StackNameSVGAtTheCenter")
	res = append(res, "StackNameFormAtTheRight")

	return
}

func (stackname StackName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "gongplanning-default")
	res = append(res, "gongplanning-tree-left")
	res = append(res, "gongplanning-svg-center")
	res = append(res, "gongplanning-form-right")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | DependencyType | StackName
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*DependencyType | *StackName
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
