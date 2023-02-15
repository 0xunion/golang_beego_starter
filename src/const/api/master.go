package api

import "strconv"

func ApplyJoinCluster(host string, port int) string {
	return "http://" + host + ":" + strconv.Itoa(port) + "/api/worker/join"
}
