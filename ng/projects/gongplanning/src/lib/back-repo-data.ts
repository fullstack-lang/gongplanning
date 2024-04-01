// generated code - do not edit

//insertion point for imports
import { PredecessorAPI } from './predecessor-api'

import { ProjectAPI } from './project-api'

import { TaskAPI } from './task-api'


export class BackRepoData {
	// insertion point for declarations
	PredecessorAPIs = new Array<PredecessorAPI>()

	ProjectAPIs = new Array<ProjectAPI>()

	TaskAPIs = new Array<TaskAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.PredecessorAPIs = data?.PredecessorAPIs || [];

		this.ProjectAPIs = data?.ProjectAPIs || [];

		this.TaskAPIs = data?.TaskAPIs || [];

	}

}