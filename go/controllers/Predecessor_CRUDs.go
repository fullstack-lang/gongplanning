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
var __Predecessor__dummysDeclaration__ models.Predecessor
var __Predecessor_time__dummyDeclaration time.Duration

var mutexPredecessor sync.Mutex

// An PredecessorID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPredecessor updatePredecessor deletePredecessor
type PredecessorID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PredecessorInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPredecessor updatePredecessor
type PredecessorInput struct {
	// The Predecessor to submit or modify
	// in: body
	Predecessor *orm.PredecessorAPI
}

// GetPredecessors
//
// swagger:route GET /predecessors predecessors getPredecessors
//
// # Get all predecessors
//
// Responses:
// default: genericError
//
//	200: predecessorDBResponse
func (controller *Controller) GetPredecessors(c *gin.Context) {

	// source slice
	var predecessorDBs []orm.PredecessorDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPredecessors", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPredecessor.GetDB()

	query := db.Find(&predecessorDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	predecessorAPIs := make([]orm.PredecessorAPI, 0)

	// for each predecessor, update fields from the database nullable fields
	for idx := range predecessorDBs {
		predecessorDB := &predecessorDBs[idx]
		_ = predecessorDB
		var predecessorAPI orm.PredecessorAPI

		// insertion point for updating fields
		predecessorAPI.ID = predecessorDB.ID
		predecessorDB.CopyBasicFieldsToPredecessor_WOP(&predecessorAPI.Predecessor_WOP)
		predecessorAPI.PredecessorPointersEncoding = predecessorDB.PredecessorPointersEncoding
		predecessorAPIs = append(predecessorAPIs, predecessorAPI)
	}

	c.JSON(http.StatusOK, predecessorAPIs)
}

// PostPredecessor
//
// swagger:route POST /predecessors predecessors postPredecessor
//
// Creates a predecessor
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostPredecessor(c *gin.Context) {

	mutexPredecessor.Lock()
	defer mutexPredecessor.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostPredecessors", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPredecessor.GetDB()

	// Validate input
	var input orm.PredecessorAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create predecessor
	predecessorDB := orm.PredecessorDB{}
	predecessorDB.PredecessorPointersEncoding = input.PredecessorPointersEncoding
	predecessorDB.CopyBasicFieldsFromPredecessor_WOP(&input.Predecessor_WOP)

	query := db.Create(&predecessorDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoPredecessor.CheckoutPhaseOneInstance(&predecessorDB)
	predecessor := backRepo.BackRepoPredecessor.Map_PredecessorDBID_PredecessorPtr[predecessorDB.ID]

	if predecessor != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), predecessor)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, predecessorDB)
}

// GetPredecessor
//
// swagger:route GET /predecessors/{ID} predecessors getPredecessor
//
// Gets the details for a predecessor.
//
// Responses:
// default: genericError
//
//	200: predecessorDBResponse
func (controller *Controller) GetPredecessor(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetPredecessor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPredecessor.GetDB()

	// Get predecessorDB in DB
	var predecessorDB orm.PredecessorDB
	if err := db.First(&predecessorDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var predecessorAPI orm.PredecessorAPI
	predecessorAPI.ID = predecessorDB.ID
	predecessorAPI.PredecessorPointersEncoding = predecessorDB.PredecessorPointersEncoding
	predecessorDB.CopyBasicFieldsToPredecessor_WOP(&predecessorAPI.Predecessor_WOP)

	c.JSON(http.StatusOK, predecessorAPI)
}

// UpdatePredecessor
//
// swagger:route PATCH /predecessors/{ID} predecessors updatePredecessor
//
// # Update a predecessor
//
// Responses:
// default: genericError
//
//	200: predecessorDBResponse
func (controller *Controller) UpdatePredecessor(c *gin.Context) {

	mutexPredecessor.Lock()
	defer mutexPredecessor.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdatePredecessor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPredecessor.GetDB()

	// Validate input
	var input orm.PredecessorAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var predecessorDB orm.PredecessorDB

	// fetch the predecessor
	query := db.First(&predecessorDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	predecessorDB.CopyBasicFieldsFromPredecessor_WOP(&input.Predecessor_WOP)
	predecessorDB.PredecessorPointersEncoding = input.PredecessorPointersEncoding

	query = db.Model(&predecessorDB).Updates(predecessorDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	predecessorNew := new(models.Predecessor)
	predecessorDB.CopyBasicFieldsToPredecessor(predecessorNew)

	// redeem pointers
	predecessorDB.DecodePointers(backRepo, predecessorNew)

	// get stage instance from DB instance, and call callback function
	predecessorOld := backRepo.BackRepoPredecessor.Map_PredecessorDBID_PredecessorPtr[predecessorDB.ID]
	if predecessorOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), predecessorOld, predecessorNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the predecessorDB
	c.JSON(http.StatusOK, predecessorDB)
}

// DeletePredecessor
//
// swagger:route DELETE /predecessors/{ID} predecessors deletePredecessor
//
// # Delete a predecessor
//
// default: genericError
//
//	200: predecessorDBResponse
func (controller *Controller) DeletePredecessor(c *gin.Context) {

	mutexPredecessor.Lock()
	defer mutexPredecessor.Unlock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeletePredecessor", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongplanning/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoPredecessor.GetDB()

	// Get model if exist
	var predecessorDB orm.PredecessorDB
	if err := db.First(&predecessorDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&predecessorDB)

	// get an instance (not staged) from DB instance, and call callback function
	predecessorDeleted := new(models.Predecessor)
	predecessorDB.CopyBasicFieldsToPredecessor(predecessorDeleted)

	// get stage instance from DB instance, and call callback function
	predecessorStaged := backRepo.BackRepoPredecessor.Map_PredecessorDBID_PredecessorPtr[predecessorDB.ID]
	if predecessorStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), predecessorStaged, predecessorDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
