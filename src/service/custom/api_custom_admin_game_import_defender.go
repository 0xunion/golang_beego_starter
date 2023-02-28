package custom

/* @MT-TPL-IMPORT-START */
import (
	"strconv"
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	"github.com/xuri/excelize/v2"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/import/defender Service 
func ApiCustomAdminGameImportDefenderService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    DefenderFileId master_types.PrimaryId,
) (*master_types.MasterResponse) {
    var apiCustomAdminGameImportDefenderResponse struct {
        Success bool `json:"success"`
        FileId master_types.PrimaryId `json:"file_id"`
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
			return master_types.ErrorResponse(-500, "invalid row format, row "+strconv.Itoa(i)+" should have "+strconv.Itoa(len(header))+" columns")
		}

		// read data
		// check if defender exists
		for _, defender := range defenders {
			if defender.Name == row[1] {
				continue
			}
		}

		// do not exist, create new defender
		defender := master_types.Defender{
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
		var id master_types.PrimaryId
		err := model.ModelInsert(defender, &id)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}
		defender.Id = id
	}

	// insert asset
	for i, asset := range rows {
		if i == 0 {
			continue
		}
		// check if defender exists
		var defender *master_types.Defender
		for _, d := range defenders {
			if d.Name == asset[1] {
				defender = d
				break
			}
		}
		if defender == nil {
			return master_types.ErrorResponse(-500, "defender not found")
		}

		// create asset
		asset := master_types.Asset{
			Owner:      user.Id,
			DefenderId: defender.Id,
			Uri:        asset[0],
			Industry:   asset[2],
			CreateAt:   time.Now().Unix(),
			GameId:     GameId,
		}

		var id master_types.PrimaryId
		err := model.ModelInsert(&asset, &id)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}
	}

	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAdminGameImportDefenderResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
