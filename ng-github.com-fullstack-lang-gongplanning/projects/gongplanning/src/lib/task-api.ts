// insertion point for imports
import { PredecessorAPI } from './predecessor-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class TaskAPI {

	static GONGSTRUCT_NAME = "Task"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other decls

	TaskPointersEncoding: TaskPointersEncoding = new TaskPointersEncoding
}

export class TaskPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Predecessors: number[] = []
}
