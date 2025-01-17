// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongplanning/go/db"
	"github.com/fullstack-lang/gongplanning/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Task_sql sql.NullBool
var dummy_Task_time time.Duration
var dummy_Task_sort sort.Float64Slice

// TaskAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model taskAPI
type TaskAPI struct {
	gorm.Model

	models.Task_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	TaskPointersEncoding TaskPointersEncoding
}

// TaskPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type TaskPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Predecessors is a slice of pointers to another Struct (optional or 0..1)
	Predecessors IntSlice `gorm:"type:TEXT"`
}

// TaskDB describes a task in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model taskDB
type TaskDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field taskDB.Name
	Name_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	TaskPointersEncoding
}

// TaskDBs arrays taskDBs
// swagger:response taskDBsResponse
type TaskDBs []TaskDB

// TaskDBResponse provides response
// swagger:response taskDBResponse
type TaskDBResponse struct {
	TaskDB
}

// TaskWOP is a Task without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type TaskWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var Task_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoTaskStruct struct {
	// stores TaskDB according to their gorm ID
	Map_TaskDBID_TaskDB map[uint]*TaskDB

	// stores TaskDB ID according to Task address
	Map_TaskPtr_TaskDBID map[*models.Task]uint

	// stores Task according to their gorm ID
	Map_TaskDBID_TaskPtr map[uint]*models.Task

	db db.DBInterface

	stage *models.StageStruct
}

func (backRepoTask *BackRepoTaskStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoTask.stage
	return
}

func (backRepoTask *BackRepoTaskStruct) GetDB() db.DBInterface {
	return backRepoTask.db
}

// GetTaskDBFromTaskPtr is a handy function to access the back repo instance from the stage instance
func (backRepoTask *BackRepoTaskStruct) GetTaskDBFromTaskPtr(task *models.Task) (taskDB *TaskDB) {
	id := backRepoTask.Map_TaskPtr_TaskDBID[task]
	taskDB = backRepoTask.Map_TaskDBID_TaskDB[id]
	return
}

// BackRepoTask.CommitPhaseOne commits all staged instances of Task to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoTask *BackRepoTaskStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for task := range stage.Tasks {
		backRepoTask.CommitPhaseOneInstance(task)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, task := range backRepoTask.Map_TaskDBID_TaskPtr {
		if _, ok := stage.Tasks[task]; !ok {
			backRepoTask.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoTask.CommitDeleteInstance commits deletion of Task to the BackRepo
func (backRepoTask *BackRepoTaskStruct) CommitDeleteInstance(id uint) (Error error) {

	task := backRepoTask.Map_TaskDBID_TaskPtr[id]

	// task is not staged anymore, remove taskDB
	taskDB := backRepoTask.Map_TaskDBID_TaskDB[id]
	db, _ := backRepoTask.db.Unscoped()
	_, err := db.Delete(taskDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoTask.Map_TaskPtr_TaskDBID, task)
	delete(backRepoTask.Map_TaskDBID_TaskPtr, id)
	delete(backRepoTask.Map_TaskDBID_TaskDB, id)

	return
}

// BackRepoTask.CommitPhaseOneInstance commits task staged instances of Task to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoTask *BackRepoTaskStruct) CommitPhaseOneInstance(task *models.Task) (Error error) {

	// check if the task is not commited yet
	if _, ok := backRepoTask.Map_TaskPtr_TaskDBID[task]; ok {
		return
	}

	// initiate task
	var taskDB TaskDB
	taskDB.CopyBasicFieldsFromTask(task)

	_, err := backRepoTask.db.Create(&taskDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoTask.Map_TaskPtr_TaskDBID[task] = taskDB.ID
	backRepoTask.Map_TaskDBID_TaskPtr[taskDB.ID] = task
	backRepoTask.Map_TaskDBID_TaskDB[taskDB.ID] = &taskDB

	return
}

// BackRepoTask.CommitPhaseTwo commits all staged instances of Task to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTask *BackRepoTaskStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, task := range backRepoTask.Map_TaskDBID_TaskPtr {
		backRepoTask.CommitPhaseTwoInstance(backRepo, idx, task)
	}

	return
}

// BackRepoTask.CommitPhaseTwoInstance commits {{structname }} of models.Task to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTask *BackRepoTaskStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, task *models.Task) (Error error) {

	// fetch matching taskDB
	if taskDB, ok := backRepoTask.Map_TaskDBID_TaskDB[idx]; ok {

		taskDB.CopyBasicFieldsFromTask(task)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		taskDB.TaskPointersEncoding.Predecessors = make([]int, 0)
		// 2. encode
		for _, predecessorAssocEnd := range task.Predecessors {
			predecessorAssocEnd_DB :=
				backRepo.BackRepoPredecessor.GetPredecessorDBFromPredecessorPtr(predecessorAssocEnd)
			
			// the stage might be inconsistant, meaning that the predecessorAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if predecessorAssocEnd_DB == nil {
				continue
			}
			
			taskDB.TaskPointersEncoding.Predecessors =
				append(taskDB.TaskPointersEncoding.Predecessors, int(predecessorAssocEnd_DB.ID))
		}

		_, err := backRepoTask.db.Save(taskDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Task intance %s", task.Name))
		return err
	}

	return
}

// BackRepoTask.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoTask *BackRepoTaskStruct) CheckoutPhaseOne() (Error error) {

	taskDBArray := make([]TaskDB, 0)
	_, err := backRepoTask.db.Find(&taskDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	taskInstancesToBeRemovedFromTheStage := make(map[*models.Task]any)
	for key, value := range backRepoTask.stage.Tasks {
		taskInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, taskDB := range taskDBArray {
		backRepoTask.CheckoutPhaseOneInstance(&taskDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		task, ok := backRepoTask.Map_TaskDBID_TaskPtr[taskDB.ID]
		if ok {
			delete(taskInstancesToBeRemovedFromTheStage, task)
		}
	}

	// remove from stage and back repo's 3 maps all tasks that are not in the checkout
	for task := range taskInstancesToBeRemovedFromTheStage {
		task.Unstage(backRepoTask.GetStage())

		// remove instance from the back repo 3 maps
		taskID := backRepoTask.Map_TaskPtr_TaskDBID[task]
		delete(backRepoTask.Map_TaskPtr_TaskDBID, task)
		delete(backRepoTask.Map_TaskDBID_TaskDB, taskID)
		delete(backRepoTask.Map_TaskDBID_TaskPtr, taskID)
	}

	return
}

// CheckoutPhaseOneInstance takes a taskDB that has been found in the DB, updates the backRepo and stages the
// models version of the taskDB
func (backRepoTask *BackRepoTaskStruct) CheckoutPhaseOneInstance(taskDB *TaskDB) (Error error) {

	task, ok := backRepoTask.Map_TaskDBID_TaskPtr[taskDB.ID]
	if !ok {
		task = new(models.Task)

		backRepoTask.Map_TaskDBID_TaskPtr[taskDB.ID] = task
		backRepoTask.Map_TaskPtr_TaskDBID[task] = taskDB.ID

		// append model store with the new element
		task.Name = taskDB.Name_Data.String
		task.Stage(backRepoTask.GetStage())
	}
	taskDB.CopyBasicFieldsToTask(task)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	task.Stage(backRepoTask.GetStage())

	// preserve pointer to taskDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_TaskDBID_TaskDB)[taskDB hold variable pointers
	taskDB_Data := *taskDB
	preservedPtrToTask := &taskDB_Data
	backRepoTask.Map_TaskDBID_TaskDB[taskDB.ID] = preservedPtrToTask

	return
}

// BackRepoTask.CheckoutPhaseTwo Checkouts all staged instances of Task to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTask *BackRepoTaskStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, taskDB := range backRepoTask.Map_TaskDBID_TaskDB {
		backRepoTask.CheckoutPhaseTwoInstance(backRepo, taskDB)
	}
	return
}

// BackRepoTask.CheckoutPhaseTwoInstance Checkouts staged instances of Task to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTask *BackRepoTaskStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, taskDB *TaskDB) (Error error) {

	task := backRepoTask.Map_TaskDBID_TaskPtr[taskDB.ID]

	taskDB.DecodePointers(backRepo, task)

	return
}

func (taskDB *TaskDB) DecodePointers(backRepo *BackRepoStruct, task *models.Task) {

	// insertion point for checkout of pointer encoding
	// This loop redeem task.Predecessors in the stage from the encode in the back repo
	// It parses all PredecessorDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	task.Predecessors = task.Predecessors[:0]
	for _, _Predecessorid := range taskDB.TaskPointersEncoding.Predecessors {
		task.Predecessors = append(task.Predecessors, backRepo.BackRepoPredecessor.Map_PredecessorDBID_PredecessorPtr[uint(_Predecessorid)])
	}

	return
}

// CommitTask allows commit of a single task (if already staged)
func (backRepo *BackRepoStruct) CommitTask(task *models.Task) {
	backRepo.BackRepoTask.CommitPhaseOneInstance(task)
	if id, ok := backRepo.BackRepoTask.Map_TaskPtr_TaskDBID[task]; ok {
		backRepo.BackRepoTask.CommitPhaseTwoInstance(backRepo, id, task)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitTask allows checkout of a single task (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutTask(task *models.Task) {
	// check if the task is staged
	if _, ok := backRepo.BackRepoTask.Map_TaskPtr_TaskDBID[task]; ok {

		if id, ok := backRepo.BackRepoTask.Map_TaskPtr_TaskDBID[task]; ok {
			var taskDB TaskDB
			taskDB.ID = id

			if _, err := backRepo.BackRepoTask.db.First(&taskDB, id); err != nil {
				log.Fatalln("CheckoutTask : Problem with getting object with id:", id)
			}
			backRepo.BackRepoTask.CheckoutPhaseOneInstance(&taskDB)
			backRepo.BackRepoTask.CheckoutPhaseTwoInstance(backRepo, &taskDB)
		}
	}
}

// CopyBasicFieldsFromTask
func (taskDB *TaskDB) CopyBasicFieldsFromTask(task *models.Task) {
	// insertion point for fields commit

	taskDB.Name_Data.String = task.Name
	taskDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromTask_WOP
func (taskDB *TaskDB) CopyBasicFieldsFromTask_WOP(task *models.Task_WOP) {
	// insertion point for fields commit

	taskDB.Name_Data.String = task.Name
	taskDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromTaskWOP
func (taskDB *TaskDB) CopyBasicFieldsFromTaskWOP(task *TaskWOP) {
	// insertion point for fields commit

	taskDB.Name_Data.String = task.Name
	taskDB.Name_Data.Valid = true
}

// CopyBasicFieldsToTask
func (taskDB *TaskDB) CopyBasicFieldsToTask(task *models.Task) {
	// insertion point for checkout of basic fields (back repo to stage)
	task.Name = taskDB.Name_Data.String
}

// CopyBasicFieldsToTask_WOP
func (taskDB *TaskDB) CopyBasicFieldsToTask_WOP(task *models.Task_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	task.Name = taskDB.Name_Data.String
}

// CopyBasicFieldsToTaskWOP
func (taskDB *TaskDB) CopyBasicFieldsToTaskWOP(task *TaskWOP) {
	task.ID = int(taskDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	task.Name = taskDB.Name_Data.String
}

// Backup generates a json file from a slice of all TaskDB instances in the backrepo
func (backRepoTask *BackRepoTaskStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "TaskDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*TaskDB, 0)
	for _, taskDB := range backRepoTask.Map_TaskDBID_TaskDB {
		forBackup = append(forBackup, taskDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Task ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Task file", err.Error())
	}
}

// Backup generates a json file from a slice of all TaskDB instances in the backrepo
func (backRepoTask *BackRepoTaskStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*TaskDB, 0)
	for _, taskDB := range backRepoTask.Map_TaskDBID_TaskDB {
		forBackup = append(forBackup, taskDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Task")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Task_Fields, -1)
	for _, taskDB := range forBackup {

		var taskWOP TaskWOP
		taskDB.CopyBasicFieldsToTaskWOP(&taskWOP)

		row := sh.AddRow()
		row.WriteStruct(&taskWOP, -1)
	}
}

// RestoreXL from the "Task" sheet all TaskDB instances
func (backRepoTask *BackRepoTaskStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoTaskid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Task"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoTask.rowVisitorTask)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoTask *BackRepoTaskStruct) rowVisitorTask(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var taskWOP TaskWOP
		row.ReadStruct(&taskWOP)

		// add the unmarshalled struct to the stage
		taskDB := new(TaskDB)
		taskDB.CopyBasicFieldsFromTaskWOP(&taskWOP)

		taskDB_ID_atBackupTime := taskDB.ID
		taskDB.ID = 0
		_, err := backRepoTask.db.Create(taskDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoTask.Map_TaskDBID_TaskDB[taskDB.ID] = taskDB
		BackRepoTaskid_atBckpTime_newID[taskDB_ID_atBackupTime] = taskDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "TaskDB.json" in dirPath that stores an array
// of TaskDB and stores it in the database
// the map BackRepoTaskid_atBckpTime_newID is updated accordingly
func (backRepoTask *BackRepoTaskStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoTaskid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "TaskDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Task file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*TaskDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_TaskDBID_TaskDB
	for _, taskDB := range forRestore {

		taskDB_ID_atBackupTime := taskDB.ID
		taskDB.ID = 0
		_, err := backRepoTask.db.Create(taskDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoTask.Map_TaskDBID_TaskDB[taskDB.ID] = taskDB
		BackRepoTaskid_atBckpTime_newID[taskDB_ID_atBackupTime] = taskDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Task file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Task>id_atBckpTime_newID
// to compute new index
func (backRepoTask *BackRepoTaskStruct) RestorePhaseTwo() {

	for _, taskDB := range backRepoTask.Map_TaskDBID_TaskDB {

		// next line of code is to avert unused variable compilation error
		_ = taskDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoTask.db.Model(taskDB)
		_, err := db.Updates(*taskDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoTask.ResetReversePointers commits all staged instances of Task to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoTask *BackRepoTaskStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, task := range backRepoTask.Map_TaskDBID_TaskPtr {
		backRepoTask.ResetReversePointersInstance(backRepo, idx, task)
	}

	return
}

func (backRepoTask *BackRepoTaskStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, task *models.Task) (Error error) {

	// fetch matching taskDB
	if taskDB, ok := backRepoTask.Map_TaskDBID_TaskDB[idx]; ok {
		_ = taskDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoTaskid_atBckpTime_newID map[uint]uint
