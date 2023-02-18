package common

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/0xunion/exercise_back/src/model"
	"github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/hash"
	"github.com/0xunion/exercise_back/src/util/num"
)

/*
	this file is used to store the common code for all the service files
*/

func CreateFileService(user *types.User, file multipart.File, header *multipart.FileHeader) *types.MasterResponse {
	// check size
	if header.Size > 1024*1024*40 {
		return types.ErrorResponse(-400, "file size too large")
	}

	// save file
	random_hash := hash.Md5("rand-" + strconv.Itoa(num.Random(100000, 999999)) + "-" + strconv.FormatInt(time.Now().Unix(), 16))
	date := time.Now().Format("2006-01-02")
	file_path := "upload/" + date + "/" + random_hash

	file_obj := &types.File{
		Owner: user.Id,
		Hash: random_hash,
		Size: header.Size,
		Path: file_path,
		CreateAt: time.Now().Unix(),
	}

	var id types.PrimaryId
	err := model.ModelInsert(file_obj, &id)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	// copy to disk
	// open file
	file_disk, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}
	defer file_disk.Close()

	// copy
	_, err = io.Copy(file_disk, file)
	if err != nil {
		return types.ErrorResponse(-500, err.Error())
	}

	return types.SuccessResponse(id)
}

func GetFileService(user *types.User, fileId types.PrimaryId) (string, error) {
	file, err := model.ModelGet[types.File](
		model.NewMongoFilter(
			model.IdFilter(fileId),
		),
	)

	if err != nil {
		return "", err
	}

	if file.Owner != user.Id {
		return "", errors.New("permission denied")
	}

	return file.Path, nil
}