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
var __Project__dummysDeclaration__ models.Project
var __Project_time__dummyDeclaration time.Duration

var mutexProject sync.Mutex

// An ProjectID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getProject updateProject deleteProject
type ProjectID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ProjectInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postProject updateProject
type ProjectInput struct {
	// The Project to submit or modify
	// in: body
	Project *orm.ProjectAPI
}

// GetProjects
//
// swagger:route GET /projects projects getProjects
//
// # Get all projects
//
// Responses:
// default: genericError
//
//	200: projectDBResponse
func (controller *Controller) GetProjects(c *gin.Context) {

	// source slice
	var projectDBs []orm.ProjectDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetProjects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoProject.GetDB()

	_, err := db.Find(&projectDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	projectAPIs := make([]orm.ProjectAPI, 0)

	// for each project, update fields from the database nullable fields
	for idx := range projectDBs {
		projectDB := &projectDBs[idx]
		_ = projectDB
		var projectAPI orm.ProjectAPI

		// insertion point for updating fields
		projectAPI.ID = projectDB.ID
		projectDB.CopyBasicFieldsToProject_WOP(&projectAPI.Project_WOP)
		projectAPI.ProjectPointersEncoding = projectDB.ProjectPointersEncoding
		projectAPIs = append(projectAPIs, projectAPI)
	}

	c.JSON(http.StatusOK, projectAPIs)
}

// PostProject
//
// swagger:route POST /projects projects postProject
//
// Creates a project
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostProject(c *gin.Context) {

	mutexProject.Lock()
	defer mutexProject.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostProjects", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoProject.GetDB()

	// Validate input
	var input orm.ProjectAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create project
	projectDB := orm.ProjectDB{}
	projectDB.ProjectPointersEncoding = input.ProjectPointersEncoding
	projectDB.CopyBasicFieldsFromProject_WOP(&input.Project_WOP)

	_, err = db.Create(&projectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoProject.CheckoutPhaseOneInstance(&projectDB)
	project := backRepo.BackRepoProject.Map_ProjectDBID_ProjectPtr[projectDB.ID]

	if project != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), project)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, projectDB)
}

// GetProject
//
// swagger:route GET /projects/{ID} projects getProject
//
// Gets the details for a project.
//
// Responses:
// default: genericError
//
//	200: projectDBResponse
func (controller *Controller) GetProject(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetProject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoProject.GetDB()

	// Get projectDB in DB
	var projectDB orm.ProjectDB
	if _, err := db.First(&projectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var projectAPI orm.ProjectAPI
	projectAPI.ID = projectDB.ID
	projectAPI.ProjectPointersEncoding = projectDB.ProjectPointersEncoding
	projectDB.CopyBasicFieldsToProject_WOP(&projectAPI.Project_WOP)

	c.JSON(http.StatusOK, projectAPI)
}

// UpdateProject
//
// swagger:route PATCH /projects/{ID} projects updateProject
//
// # Update a project
//
// Responses:
// default: genericError
//
//	200: projectDBResponse
func (controller *Controller) UpdateProject(c *gin.Context) {

	mutexProject.Lock()
	defer mutexProject.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateProject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoProject.GetDB()

	// Validate input
	var input orm.ProjectAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var projectDB orm.ProjectDB

	// fetch the project
	_, err := db.First(&projectDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	projectDB.CopyBasicFieldsFromProject_WOP(&input.Project_WOP)
	projectDB.ProjectPointersEncoding = input.ProjectPointersEncoding

	db, _ = db.Model(&projectDB)
	_, err = db.Updates(projectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	projectNew := new(models.Project)
	projectDB.CopyBasicFieldsToProject(projectNew)

	// redeem pointers
	projectDB.DecodePointers(backRepo, projectNew)

	// get stage instance from DB instance, and call callback function
	projectOld := backRepo.BackRepoProject.Map_ProjectDBID_ProjectPtr[projectDB.ID]
	if projectOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), projectOld, projectNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the projectDB
	c.JSON(http.StatusOK, projectDB)
}

// DeleteProject
//
// swagger:route DELETE /projects/{ID} projects deleteProject
//
// # Delete a project
//
// default: genericError
//
//	200: projectDBResponse
func (controller *Controller) DeleteProject(c *gin.Context) {

	mutexProject.Lock()
	defer mutexProject.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteProject", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoProject.GetDB()

	// Get model if exist
	var projectDB orm.ProjectDB
	if _, err := db.First(&projectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&projectDB)

	// get an instance (not staged) from DB instance, and call callback function
	projectDeleted := new(models.Project)
	projectDB.CopyBasicFieldsToProject(projectDeleted)

	// get stage instance from DB instance, and call callback function
	projectStaged := backRepo.BackRepoProject.Map_ProjectDBID_ProjectPtr[projectDB.ID]
	if projectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), projectStaged, projectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
