// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	PredecessorAPIs []*PredecessorAPI

	ProjectAPIs []*ProjectAPI

	TaskAPIs []*TaskAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, predecessorDB := range backRepo.BackRepoPredecessor.Map_PredecessorDBID_PredecessorDB {

		var predecessorAPI PredecessorAPI
		predecessorAPI.ID = predecessorDB.ID
		predecessorAPI.PredecessorPointersEncoding = predecessorDB.PredecessorPointersEncoding
		predecessorDB.CopyBasicFieldsToPredecessor_WOP(&predecessorAPI.Predecessor_WOP)

		backRepoData.PredecessorAPIs = append(backRepoData.PredecessorAPIs, &predecessorAPI)
	}

	for _, projectDB := range backRepo.BackRepoProject.Map_ProjectDBID_ProjectDB {

		var projectAPI ProjectAPI
		projectAPI.ID = projectDB.ID
		projectAPI.ProjectPointersEncoding = projectDB.ProjectPointersEncoding
		projectDB.CopyBasicFieldsToProject_WOP(&projectAPI.Project_WOP)

		backRepoData.ProjectAPIs = append(backRepoData.ProjectAPIs, &projectAPI)
	}

	for _, taskDB := range backRepo.BackRepoTask.Map_TaskDBID_TaskDB {

		var taskAPI TaskAPI
		taskAPI.ID = taskDB.ID
		taskAPI.TaskPointersEncoding = taskDB.TaskPointersEncoding
		taskDB.CopyBasicFieldsToTask_WOP(&taskAPI.Task_WOP)

		backRepoData.TaskAPIs = append(backRepoData.TaskAPIs, &taskAPI)
	}

}
