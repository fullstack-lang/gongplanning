// generated code - do not edit

import { TaskAPI } from './task-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Predecessor } from './predecessor'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Task {

	static GONGSTRUCT_NAME = "Task"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Predecessors: Array<Predecessor> = []
}

export function CopyTaskToTaskAPI(task: Task, taskAPI: TaskAPI) {

	taskAPI.CreatedAt = task.CreatedAt
	taskAPI.DeletedAt = task.DeletedAt
	taskAPI.ID = task.ID

	// insertion point for basic fields copy operations
	taskAPI.Name = task.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	taskAPI.TaskPointersEncoding.Predecessors = []
	for (let _predecessor of task.Predecessors) {
		taskAPI.TaskPointersEncoding.Predecessors.push(_predecessor.ID)
	}

}

// CopyTaskAPIToTask update basic, pointers and slice of pointers fields of task
// from respectively the basic fields and encoded fields of pointers and slices of pointers of taskAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTaskAPIToTask(taskAPI: TaskAPI, task: Task, frontRepo: FrontRepo) {

	task.CreatedAt = taskAPI.CreatedAt
	task.DeletedAt = taskAPI.DeletedAt
	task.ID = taskAPI.ID

	// insertion point for basic fields copy operations
	task.Name = taskAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	task.Predecessors = new Array<Predecessor>()
	for (let _id of taskAPI.TaskPointersEncoding.Predecessors) {
		let _predecessor = frontRepo.map_ID_Predecessor.get(_id)
		if (_predecessor != undefined) {
			task.Predecessors.push(_predecessor!)
		}
	}
}
