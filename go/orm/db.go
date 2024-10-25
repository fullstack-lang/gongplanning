// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongplanning/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	predecessorDBs map[uint]*PredecessorDB

	nextIDPredecessorDB uint

	projectDBs map[uint]*ProjectDB

	nextIDProjectDB uint

	taskDBs map[uint]*TaskDB

	nextIDTaskDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		predecessorDBs: make(map[uint]*PredecessorDB),

		projectDBs: make(map[uint]*ProjectDB),

		taskDBs: make(map[uint]*TaskDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *PredecessorDB:
		db.nextIDPredecessorDB++
		v.ID = db.nextIDPredecessorDB
		db.predecessorDBs[v.ID] = v
	case *ProjectDB:
		db.nextIDProjectDB++
		v.ID = db.nextIDProjectDB
		db.projectDBs[v.ID] = v
	case *TaskDB:
		db.nextIDTaskDB++
		v.ID = db.nextIDTaskDB
		db.taskDBs[v.ID] = v
	default:
		return nil, errors.New("unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *PredecessorDB:
		delete(db.predecessorDBs, v.ID)
	case *ProjectDB:
		delete(db.projectDBs, v.ID)
	case *TaskDB:
		delete(db.taskDBs, v.ID)
	default:
		return nil, errors.New("unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *PredecessorDB:
		db.predecessorDBs[v.ID] = v
		return db, nil
	case *ProjectDB:
		db.projectDBs[v.ID] = v
		return db, nil
	case *TaskDB:
		db.taskDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *PredecessorDB:
		if existing, ok := db.predecessorDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *ProjectDB:
		if existing, ok := db.projectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	case *TaskDB:
		if existing, ok := db.taskDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("record not found")
		}
	default:
		return nil, errors.New("unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]PredecessorDB:
        *ptr = make([]PredecessorDB, 0, len(db.predecessorDBs))
        for _, v := range db.predecessorDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]ProjectDB:
        *ptr = make([]ProjectDB, 0, len(db.projectDBs))
        for _, v := range db.projectDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
	case *[]TaskDB:
        *ptr = make([]TaskDB, 0, len(db.taskDBs))
        for _, v := range db.taskDBs {
            *ptr = append(*ptr, *v)
        }
        return db, nil
    default:
        return nil, errors.New("Find: unsupported type")
    }
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("Do not process when conds is not a single parameter")
	}

	str, ok := conds[0].(string)

	if !ok {
		return nil, errors.New("conds[0] is not a string")
	}

	i, err := strconv.ParseUint(str, 10, 32) // Base 10, 32-bit unsigned int
	if err != nil {
		return nil, errors.New("conds[0] is not a string number")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *PredecessorDB:
		tmp, ok := db.predecessorDBs[uint(i)]

		predecessorDB, _ := instanceDB.(*PredecessorDB)
		*predecessorDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *ProjectDB:
		tmp, ok := db.projectDBs[uint(i)]

		projectDB, _ := instanceDB.(*ProjectDB)
		*projectDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	case *TaskDB:
		tmp, ok := db.taskDBs[uint(i)]

		taskDB, _ := instanceDB.(*TaskDB)
		*taskDB = *tmp
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unkown entry %d", i))
		}
	default:
		return nil, errors.New("Unkown type")
	}
	
	return db, nil
}

