package custom

/* @MT-TPL-IMPORT-START */
import (
	"strconv"
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/hash"
	"github.com/0xunion/exercise_back/src/util/num"
	"github.com/xuri/excelize/v2"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/template/defender Service 生成防守方人员模板文件
func ApiCustomAdminGameTemplateDefenderService(
	user *master_types.User,
	GameId master_types.PrimaryId,
) *master_types.MasterResponse {
	var apiCustomAdminGameTemplateDefenderResponse struct {
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

	header := []string{"URI", "Name", "Industry"}
	header_serial := []string{"A", "B", "C"}
	// write header
	for i, v := range header {
		f.SetCellValue("Sheet1", header_serial[i]+"1", v)
	}

	// write example
	f.SetCellValue("Sheet1", "A2", "https://www.baidu.com")
	f.SetCellValue("Sheet1", "B2", "百度")
	f.SetCellValue("Sheet1", "C2", "搜索")

	f.SetCellValue("Sheet1", "A3", "tcp://111.111.111.111:8080")
	f.SetCellValue("Sheet1", "B3", "防守方1")
	f.SetCellValue("Sheet1", "C3", "医疗")

	f.SetCellValue("Sheet1", "A4", "mongodb://114.114.114.114:27017")
	f.SetCellValue("Sheet1", "B4", "防守方2")
	f.SetCellValue("Sheet1", "C4", "金融")

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save file
	random_hash := hash.Md5("rand-" + strconv.Itoa(num.Random(100000, 999999)) + "-" + strconv.FormatInt(time.Now().Unix(), 16))
	date := time.Now().Format("2006-01-02")
	file_path := "generate/" + date + "/" + random_hash

	err := f.SaveAs(file_path)
	if err != nil {
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

	apiCustomAdminGameTemplateDefenderResponse.FileId = id

	/* @MT-TPL-SERVICE-RESP-START */

	return master_types.SuccessResponse(apiCustomAdminGameTemplateDefenderResponse)
}

/* @MT-TPL-SERVICE-RESP-END */
