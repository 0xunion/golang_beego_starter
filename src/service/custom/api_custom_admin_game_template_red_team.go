package custom

import (
	"os"
	"strconv"
	"time"

	"github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/hash"
	"github.com/0xunion/exercise_back/src/util/num"
	"github.com/xuri/excelize/v2"
	/* @MT-TPL-IMPORT-TIME-START */ /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/template/red_team Service 生成红队模板文件
func ApiCustomAdminGameTemplateRedTeamService(
	user *master_types.User,
	GameId master_types.PrimaryId,
) *master_types.MasterResponse {
	var apiCustomAdminGameTemplateRedTeamResponse struct {
		Success bool                   `json:"success"`
		FileId  master_types.PrimaryId `json:"file_id"`
	}

	access_controll := false
	if !access_controll && user.IsAdmin() {
		access_controll = true
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Permission denied")
	}
	/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do
	f := excelize.NewFile()
	defer f.Close()

	// Create a new sheet.
	index, error := f.NewSheet("Sheet1")
	if error != nil {
		return master_types.ErrorResponse(-500, error.Error())
	}

	header := []string{"Phone", "Name", "Team"}
	header_serial := []string{"A", "B", "C"}
	// write header
	for i, v := range header {
		f.SetCellValue("Sheet1", header_serial[i]+"1", v)
	}

	// write example
	f.SetCellValue("Sheet1", "A2", "12345678901")
	f.SetCellValue("Sheet1", "B2", "张三")
	f.SetCellValue("Sheet1", "C2", "红队")

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save file
	random_hash := hash.Md5("rand-" + strconv.Itoa(num.Random(100000, 999999)) + "-" + strconv.FormatInt(time.Now().Unix(), 16))
	date := time.Now().Format("2006-01-02")
	file_path := "storage/generate/" + date + "/" + random_hash + ".xlsx"
	// create dir
	os.MkdirAll("storage/generate/"+date, os.ModePerm)

	error = f.SaveAs(file_path)
	if error != nil {
		return master_types.ErrorResponse(-500, error.Error())
	}

	file := &master_types.File{
		Owner:    user.Id,
		Path:     file_path,
		GameId:   GameId,
		CreateAt: time.Now().Unix(),
		Hash:     random_hash,
		Size:     0,
	}
	file.SetPublicAccess()

	var id master_types.PrimaryId
	if err := model.ModelInsert(file, &id); err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	apiCustomAdminGameTemplateRedTeamResponse.FileId = id

	/* @MT-TPL-SERVICE-RESP-START */

	return master_types.SuccessResponse(apiCustomAdminGameTemplateRedTeamResponse)
}

/* @MT-TPL-SERVICE-RESP-END */
