// generated code - do not edit

import { ProjectAPI } from './project-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Task } from './task'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Project {

	static GONGSTRUCT_NAME = "Project"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsExpanded: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Tasks: Array<Task> = []
}

export function CopyProjectToProjectAPI(project: Project, projectAPI: ProjectAPI) {

	projectAPI.CreatedAt = project.CreatedAt
	projectAPI.DeletedAt = project.DeletedAt
	projectAPI.ID = project.ID

	// insertion point for basic fields copy operations
	projectAPI.Name = project.Name
	projectAPI.IsExpanded = project.IsExpanded

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	projectAPI.ProjectPointersEncoding.Tasks = []
	for (let _task of project.Tasks) {
		projectAPI.ProjectPointersEncoding.Tasks.push(_task.ID)
	}

}

// CopyProjectAPIToProject update basic, pointers and slice of pointers fields of project
// from respectively the basic fields and encoded fields of pointers and slices of pointers of projectAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyProjectAPIToProject(projectAPI: ProjectAPI, project: Project, frontRepo: FrontRepo) {

	project.CreatedAt = projectAPI.CreatedAt
	project.DeletedAt = projectAPI.DeletedAt
	project.ID = projectAPI.ID

	// insertion point for basic fields copy operations
	project.Name = projectAPI.Name
	project.IsExpanded = projectAPI.IsExpanded

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	project.Tasks = new Array<Task>()
	for (let _id of projectAPI.ProjectPointersEncoding.Tasks) {
		let _task = frontRepo.map_ID_Task.get(_id)
		if (_task != undefined) {
			project.Tasks.push(_task!)
		}
	}
}
