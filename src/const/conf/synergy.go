package conf

import (
	"strconv"

	"github.com/0xunion/exercise_back/src/routine"
)

/*
	this file is used to keep connection to master
	while init worker, we need to specify the master
	and the master will distribute token for worker
	after that, worker only accept request with token
*/

var worker_token string
var controller_key int

func init() {
	var err error
	controller_key, err = strconv.Atoi(getConf("controllerkey"))
	if err != nil {
		routine.Panic("[SYNERGY] An unexpected error occurred")
	}
}

func SetWorkerToken(token string) {
	worker_token = token
}

func GetWorkerToken() string {
	return worker_token
}

func GetControllerKey() int {
	return controller_key
}
