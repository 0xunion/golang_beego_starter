package meta

import (
	"reflect"
	"strings"
	"sync"
)

var metaNameMap = make(map[string]string)
var metaNameMutex sync.RWMutex

// GetTypeName get type name of a variable, for example:
// var a int = 1
// GetTypeName(a) => "int"
// type A struct {}
// var b A
// GetTypeName(b) => "A" not "package.A"
func GetTypeName(v interface{}) string {
	name := reflect.TypeOf(v).Name()
	if name == "" {
		name = reflect.TypeOf(v).String()
	}
	// check if name has been cached
	metaNameMutex.RLock()
	if metaName, ok := metaNameMap[name]; ok {
		metaNameMutex.RUnlock()
		return metaName
	}
	metaNameMutex.RUnlock()
	// split name by dot
	nameList := strings.Split(name, ".")
	// get the last part of name
	metaName := nameList[len(nameList)-1]
	// cache the name
	metaNameMutex.Lock()
	metaNameMap[name] = metaName
	metaNameMutex.Unlock()
	return metaName
}