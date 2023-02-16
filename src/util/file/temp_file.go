package file

import (
	"io"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/0xunion/exercise_back/src/const/conf"
	"github.com/0xunion/exercise_back/src/routine"
)

var (
	root_temp_path = ""
)

func init() {
	rand.Seed(time.Now().UnixNano())
	root_temp_path = path.Clean(conf.WorkDir() + "/" + conf.TempPath())
	// stat the root temp path
	if _, err := os.Stat(root_temp_path); os.IsNotExist(err) {
		// create the root temp path
		os.MkdirAll(root_temp_path, os.ModePerm)
		routine.Info("[TempFile] Create temp path: " + root_temp_path)
	} else if err != nil {
		routine.Panic("[TempFile] Stat temp path failed: " + err.Error())
	}
}

func CreateTempFile() (string, io.WriteCloser, error) {
	file, err := os.Create(
		root_temp_path +
			"/tmp-" + strconv.FormatInt(time.Now().Unix(), 10) +
			"-" + strconv.Itoa(rand.Int()))
	if err != nil {
		return "", nil, err
	}

	return file.Name(), file, nil
}

func CreateTempDir() (string, func(), error) {
	dir := root_temp_path +
		"/tmp-" + strconv.FormatInt(time.Now().Unix(), 10) +
		"-" + strconv.Itoa(rand.Int())
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", nil, err
	}

	return dir, func() {
		// remove the temp dir
		os.RemoveAll(dir)
	}, nil
}
