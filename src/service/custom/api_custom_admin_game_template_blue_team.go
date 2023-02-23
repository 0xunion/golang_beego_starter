package custom

import (
	"os"
	"strconv"
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/hash"
	"github.com/0xunion/exercise_back/src/util/num"
	"github.com/xuri/excelize/v2"
	/* @MT-TPL-IMPORT-TIME-START */ /* @MT-TPL-IMPORT-TIME-END */)

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/template/blue_team Service 生成蓝队模板文件
func ApiCustomAdminGameTemplateBlueTeamService(
	user *master_types.User,
	GameId master_types.PrimaryId,
) *master_types.MasterResponse {
	var apiCustomAdminGameTemplateBlueTeamResponse struct {
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

	header := []string{"Phone", "Name", "DenfenderIndustry", "DenfenderName", "DefenderId"}
	header_serial := []string{"A", "B", "C", "D", "E"}
	// write header
	for i, v := range header {
		f.SetCellValue("Sheet1", header_serial[i]+"1", v)
	}

	// write example
	// Get All defenders
	defenders, err := model.ModelGetAll[master_types.Defender](
		model.NewMongoFilter(
			model.MongoKeyFilter("game_id", GameId),
		),
	)

	// write exisitng defenders
	for i, v := range defenders {
		f.SetCellValue("Sheet1", header_serial[0]+strconv.Itoa(i+2), "13800138000")
		f.SetCellValue("Sheet1", header_serial[1]+strconv.Itoa(i+2), "张三")
		f.SetCellValue("Sheet1", header_serial[2]+strconv.Itoa(i+2), v.Industry)
		f.SetCellValue("Sheet1", header_serial[3]+strconv.Itoa(i+2), v.Name)
		f.SetCellValue("Sheet1", header_serial[4]+strconv.Itoa(i+2), v.Id.Hex())
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save file
	random_hash := hash.Md5("rand-" + strconv.Itoa(num.Random(100000, 999999)) + "-" + strconv.FormatInt(time.Now().Unix(), 16))
	date := time.Now().Format("2006-01-02")
	file_path := "storage/generate/" + date + "/" + random_hash + ".xlsx"
	os.MkdirAll("storage/generate/"+date, 0777)

	err = f.SaveAs(file_path)
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
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

	apiCustomAdminGameTemplateBlueTeamResponse.FileId = id

	/* @MT-TPL-SERVICE-RESP-START */

	return master_types.SuccessResponse(apiCustomAdminGameTemplateBlueTeamResponse)
}

/* @MT-TPL-SERVICE-RESP-END */
