// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongplanning/go/models"
	"github.com/fullstack-lang/gongplanning/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Task__dummysDeclaration__ models.Task
var __Task_time__dummyDeclaration time.Duration

var mutexTask sync.Mutex

// An TaskID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTask updateTask deleteTask
type TaskID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TaskInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTask updateTask
type TaskInput struct {
	// The Task to submit or modify
	// in: body
	Task *orm.TaskAPI
}

// GetTasks
//
// swagger:route GET /tasks tasks getTasks
//
// # Get all tasks
//
// Responses:
// default: genericError
//
//	200: taskDBResponse
func (controller *Controller) GetTasks(c *gin.Context) {

	// source slice
	var taskDBs []orm.TaskDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTasks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTask.GetDB()

	query := db.Find(&taskDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	taskAPIs := make([]orm.TaskAPI, 0)

	// for each task, update fields from the database nullable fields
	for idx := range taskDBs {
		taskDB := &taskDBs[idx]
		_ = taskDB
		var taskAPI orm.TaskAPI

		// insertion point for updating fields
		taskAPI.ID = taskDB.ID
		taskDB.CopyBasicFieldsToTask_WOP(&taskAPI.Task_WOP)
		taskAPI.TaskPointersEncoding = taskDB.TaskPointersEncoding
		taskAPIs = append(taskAPIs, taskAPI)
	}

	c.JSON(http.StatusOK, taskAPIs)
}

// PostTask
//
// swagger:route POST /tasks tasks postTask
//
// Creates a task
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTask(c *gin.Context) {

	mutexTask.Lock()
	defer mutexTask.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTasks", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTask.GetDB()

	// Validate input
	var input orm.TaskAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create task
	taskDB := orm.TaskDB{}
	taskDB.TaskPointersEncoding = input.TaskPointersEncoding
	taskDB.CopyBasicFieldsFromTask_WOP(&input.Task_WOP)

	query := db.Create(&taskDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTask.CheckoutPhaseOneInstance(&taskDB)
	task := backRepo.BackRepoTask.Map_TaskDBID_TaskPtr[taskDB.ID]

	if task != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), task)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, taskDB)
}

// GetTask
//
// swagger:route GET /tasks/{ID} tasks getTask
//
// Gets the details for a task.
//
// Responses:
// default: genericError
//
//	200: taskDBResponse
func (controller *Controller) GetTask(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTask", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTask.GetDB()

	// Get taskDB in DB
	var taskDB orm.TaskDB
	if err := db.First(&taskDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var taskAPI orm.TaskAPI
	taskAPI.ID = taskDB.ID
	taskAPI.TaskPointersEncoding = taskDB.TaskPointersEncoding
	taskDB.CopyBasicFieldsToTask_WOP(&taskAPI.Task_WOP)

	c.JSON(http.StatusOK, taskAPI)
}

// UpdateTask
//
// swagger:route PATCH /tasks/{ID} tasks updateTask
//
// # Update a task
//
// Responses:
// default: genericError
//
//	200: taskDBResponse
func (controller *Controller) UpdateTask(c *gin.Context) {

	mutexTask.Lock()
	defer mutexTask.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTask", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTask.GetDB()

	// Validate input
	var input orm.TaskAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var taskDB orm.TaskDB

	// fetch the task
	query := db.First(&taskDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	taskDB.CopyBasicFieldsFromTask_WOP(&input.Task_WOP)
	taskDB.TaskPointersEncoding = input.TaskPointersEncoding

	query = db.Model(&taskDB).Updates(taskDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	taskNew := new(models.Task)
	taskDB.CopyBasicFieldsToTask(taskNew)

	// redeem pointers
	taskDB.DecodePointers(backRepo, taskNew)

	// get stage instance from DB instance, and call callback function
	taskOld := backRepo.BackRepoTask.Map_TaskDBID_TaskPtr[taskDB.ID]
	if taskOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), taskOld, taskNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the taskDB
	c.JSON(http.StatusOK, taskDB)
}

// DeleteTask
//
// swagger:route DELETE /tasks/{ID} tasks deleteTask
//
// # Delete a task
//
// default: genericError
//
//	200: taskDBResponse
func (controller *Controller) DeleteTask(c *gin.Context) {

	mutexTask.Lock()
	defer mutexTask.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTask", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTask.GetDB()

	// Get model if exist
	var taskDB orm.TaskDB
	if err := db.First(&taskDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&taskDB)

	// get an instance (not staged) from DB instance, and call callback function
	taskDeleted := new(models.Task)
	taskDB.CopyBasicFieldsToTask(taskDeleted)

	// get stage instance from DB instance, and call callback function
	taskStaged := backRepo.BackRepoTask.Map_TaskDBID_TaskPtr[taskDB.ID]
	if taskStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), taskStaged, taskDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
