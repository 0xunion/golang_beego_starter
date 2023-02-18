package custom

/* @MT-TPL-IMPORT-START */
import (
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	"github.com/xuri/excelize/v2"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
func ApiCustomAdminGameImportDefenderService(
	user *master_types.User,
	GameId master_types.PrimaryId,
	DefenderFileId master_types.PrimaryId,
) *master_types.MasterResponse {
	var apiCustomAdminGameImportDefenderResponse struct {
		Success bool `json:"success"`
	}

	// cache map, use to store the data that has been queried
	service_cache := make(map[string]interface{})

	access_controll := false
	if !access_controll && user.IsAdmin() {
		access_controll = true
	}

	if !access_controll {
		return master_types.ErrorResponse(-403, "Permission denied")
	}
	/* @MT-TPL-SERVICE-END */

	// TODO: add service code here, do what you want to do
	file_id := DefenderFileId
	file, err := model.ModelGet[master_types.File](
		model.NewMongoFilter(
			model.IdFilter(file_id),
		),
	)
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	if file == nil {
		return master_types.ErrorResponse(-500, "file not found")
	}

	if file.Owner != user.Id {
		return master_types.ErrorResponse(-500, "file not owned by you")
	}

	file_path := file.Path
	f, err := excelize.OpenFile(file_path)
	if err != nil {
		return master_types.ErrorResponse(-500, "failed to open excel file")
	}
	defer f.Close()

	// read row by row
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return master_types.ErrorResponse(-500, "failed to read excel file")
	}

	// read header
	header := rows[0]
	// check header format like "URI,Name,Idustry"
	var standard_header = []string{"URI", "Name", "Industry"}
	for i, v := range header {
		if i >= len(standard_header) {
			return master_types.ErrorResponse(-500, "invalid header format, too many columns")
		}
		if v != standard_header[i] {
			return master_types.ErrorResponse(-500, "invalid header format, column "+v+" should be "+standard_header[i])
		}
	}

	defenders := make([]*master_types.Defender, 0)

	// read data
	for i, row := range rows {
		if i == 0 {
			continue
		}
		// check row format
		if len(row) != len(header) {
			return master_types.ErrorResponse(-500, "invalid row format, row "+string(i)+" should have "+string(len(header))+" columns")
		}

		// read data
		defender := master_types.Defender{
			Uri:      row[0],
			Name:     row[1],
			Industry: row[2],
			Score:    10000,
			Owner:    user.Id,
			GameId:   GameId,
			CreateAt: time.Now().Unix(),
		}

		defenders = append(defenders, &defender)
	}

	// insert data
	for _, defender := range defenders {
		err := model.ModelInsert(defender, nil)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}
	}

	/* @MT-TPL-SERVICE-RESP-START */

	return master_types.SuccessResponse(apiCustomAdminGameImportDefenderResponse)
}

/* @MT-TPL-SERVICE-RESP-END */
