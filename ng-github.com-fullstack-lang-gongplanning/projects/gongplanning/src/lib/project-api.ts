// insertion point for imports
import { TaskAPI } from './task-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ProjectAPI {

	static GONGSTRUCT_NAME = "Project"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsExpanded: boolean = false

	// insertion point for other decls

	ProjectPointersEncoding: ProjectPointersEncoding = new ProjectPointersEncoding
}

export class ProjectPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Tasks: number[] = []
}
