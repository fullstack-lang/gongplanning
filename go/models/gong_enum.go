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

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | DependencyType
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*DependencyType
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
