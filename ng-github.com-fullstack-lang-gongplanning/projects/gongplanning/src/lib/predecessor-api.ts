// insertion point for imports
import { TaskAPI } from './task-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class PredecessorAPI {

	static GONGSTRUCT_NAME = "Predecessor"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DependencyType: string = ""

	// insertion point for other decls

	PredecessorPointersEncoding: PredecessorPointersEncoding = new PredecessorPointersEncoding
}

export class PredecessorPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	TaskID: NullInt64 = new NullInt64 // if pointer is null, Task.ID = 0

}
