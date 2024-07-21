// generated code - do not edit

import { PredecessorAPI } from './predecessor-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Task } from './task'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Predecessor {

	static GONGSTRUCT_NAME = "Predecessor"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DependencyType: string = ""

	// insertion point for pointers and slices of pointers declarations
	Task?: Task

}

export function CopyPredecessorToPredecessorAPI(predecessor: Predecessor, predecessorAPI: PredecessorAPI) {

	predecessorAPI.CreatedAt = predecessor.CreatedAt
	predecessorAPI.DeletedAt = predecessor.DeletedAt
	predecessorAPI.ID = predecessor.ID

	// insertion point for basic fields copy operations
	predecessorAPI.Name = predecessor.Name
	predecessorAPI.DependencyType = predecessor.DependencyType

	// insertion point for pointer fields encoding
	predecessorAPI.PredecessorPointersEncoding.TaskID.Valid = true
	if (predecessor.Task != undefined) {
		predecessorAPI.PredecessorPointersEncoding.TaskID.Int64 = predecessor.Task.ID  
	} else {
		predecessorAPI.PredecessorPointersEncoding.TaskID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyPredecessorAPIToPredecessor update basic, pointers and slice of pointers fields of predecessor
// from respectively the basic fields and encoded fields of pointers and slices of pointers of predecessorAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyPredecessorAPIToPredecessor(predecessorAPI: PredecessorAPI, predecessor: Predecessor, frontRepo: FrontRepo) {

	predecessor.CreatedAt = predecessorAPI.CreatedAt
	predecessor.DeletedAt = predecessorAPI.DeletedAt
	predecessor.ID = predecessorAPI.ID

	// insertion point for basic fields copy operations
	predecessor.Name = predecessorAPI.Name
	predecessor.DependencyType = predecessorAPI.DependencyType

	// insertion point for pointer fields encoding
	predecessor.Task = frontRepo.map_ID_Task.get(predecessorAPI.PredecessorPointersEncoding.TaskID.Int64)

	// insertion point for slice of pointers fields encoding
}
