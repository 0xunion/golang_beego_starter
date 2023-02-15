package routine

import (
	"context"
	"sync"
	"time"

	ants "github.com/panjf2000/ants/v2"
)

// RuntimeManage is a routine to manage runtime
type Routine func()

type RoutineNode struct {
	name string
	ctx  context.Context
	time int64
	fn   Routine
}

func (r *RoutineNode) GetTime() int64 {
	return r.time
}

func (r *RoutineNode) GetCtx() context.Context {
	return r.ctx
}

func (r *RoutineNode) GetFn() Routine {
	return r.fn
}

func (r *RoutineNode) GetName() string {
	return r.name
}

var routineMap sync.Map
var routinePool *ants.Pool
var routineMutex sync.Mutex

func init() {
	// init ants pool
	var err error
	Info("[RuntimeManage] Start init ants pool")
	routineMutex.Lock()
	routinePool, err = ants.NewPool(1000)
	if err != nil {
		Panic("[RuntimeManage] Init ants pool failed : %v", err)
	}
	routineMutex.Unlock()
}

// NewRuntimeManage returns a new RuntimeManage
// Remember : do not pass a function which will block forever
func Go(name string, fn Routine) {
	if routinePool == nil {
		// mutex lock until pool is not nil
		routineMutex.Lock()
		defer routineMutex.Unlock()
	}
	routineMap.Store(name, &RoutineNode{
		name: name,
		ctx:  context.Background(),
		fn:   fn,
		time: time.Now().Unix(),
	})
	routinePool.Submit(fn)
}

func GetAllRoutine() []RoutineNode {
	var routines []RoutineNode
	routineMap.Range(func(key, value interface{}) bool {
		routines = append(routines, *value.(*RoutineNode))
		return true
	})
	return routines
}

func GetRoutine(name string) *RoutineNode {
	if value, ok := routineMap.Load(name); ok {
		return value.(*RoutineNode)
	}
	return nil
}
