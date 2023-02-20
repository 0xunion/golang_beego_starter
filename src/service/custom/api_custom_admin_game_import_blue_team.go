package custom

/* @MT-TPL-IMPORT-START */
import (
	"strconv"
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/auth"
	"github.com/0xunion/exercise_back/src/util/hash"
	"github.com/0xunion/exercise_back/src/util/num"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/import/blue_team Service 从文件导入蓝队信息，文件格式参考蓝队模板
func ApiCustomAdminGameImportBlueTeamService(
	user *master_types.User,
	GameId master_types.PrimaryId,
	BlueTeamFileId master_types.PrimaryId,
) *master_types.MasterResponse {
	var apiCustomAdminGameImportBlueTeamResponse struct {
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
	file_id := BlueTeamFileId
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

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return master_types.ErrorResponse(-500, "failed to read excel file")
	}

	// read header
	header := rows[0]
	// check header format
	var standard_header = []string{"Phone", "Name", "DenfenderIndustry", "DenfenderName", "DefenderId"}
	for i, v := range header {
		if i >= len(standard_header) {
			return master_types.ErrorResponse(-500, "invalid header format, too many columns")
		}

		if v != standard_header[i] {
			return master_types.ErrorResponse(-500, "invalid header format, column "+strconv.Itoa(i)+" should be "+standard_header[i])
		}
	}

	// check if game exists
	game, err := model.ModelGet[master_types.Game](
		model.NewMongoFilter(
			model.IdFilter(GameId),
		),
	)
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	if game == nil {
		return master_types.ErrorResponse(-500, "game not found")
	}

	f.Close()

	f = excelize.NewFile()
	index, err := f.NewSheet("Sheet1")
	headers := []string{"Username", "Password", "Phone"}
	header_serial := []string{"A", "B", "C"}

	for i, v := range headers {
		f.SetCellValue("Sheet1", header_serial[i]+"1", v)
	}

	// create blue team members
	for i, row := range rows {
		if i == 0 {
			continue
		}

		phone := row[0]
		name := row[1]
		defender_id, err := primitive.ObjectIDFromHex(row[4])
		if err != nil {
			return master_types.ErrorResponse(-500, "invalid defender id")
		}

		var uid master_types.PrimaryId
		var clear_password string

		// check if user exists
		user, err := model.ModelGet[master_types.Phone](
			model.NewMongoFilter(
				model.MongoKeyFilter("phone", phone),
			),
		)

		if err != nil || user == nil {
			phone_max_length := 4
			if len(phone) < phone_max_length {
				phone_max_length = len(phone)
			}

			user_model := &master_types.User{
				Username: name,
				Parent:   user.Id,
			}

			var id master_types.PrimaryId
			if err := model.ModelInsert(user_model, &id); err != nil {
				return master_types.ErrorResponse(-500, "failed to create user")
			}

			uid = id

			// create phone
			phone_model := &master_types.Phone{
				Phone:    phone,
				Uid:      uid,
				CreateAt: time.Now().Unix(),
			}

			if err := model.ModelInsert(phone_model, nil); err != nil {
				return master_types.ErrorResponse(-500, "failed to create phone")
			}

			// create password
			clear_password = phone + strconv.Itoa(num.Random(1000, 9999))
			password_model := &master_types.Password{
				Uid:      uid,
				Password: auth.HashPassword(hash.Md5(clear_password)),
				CreateAt: time.Now().Unix(),
			}

			if err := model.ModelInsert(password_model, nil); err != nil {
				return master_types.ErrorResponse(-500, "failed to create password")
			}
		} else {
			if user == nil {
				return master_types.ErrorResponse(-500, "failed to get user")
			} else {
				uid = user.Id
				clear_password = "用户已存在，密码需自行回忆，平台不会储存任何明文密码"
			}
		}

		// create gamer
		gamer_model := &master_types.Gamer{
			Owner:    uid,
			Name:     name,
			Phone:    phone,
			CreateAt: time.Now().Unix(),
			Score:    0,
			GameId:   GameId,
			GroupId:  defender_id,
		}
		gamer_model.SetDefender()

		if err := model.ModelInsert(gamer_model, nil); err != nil {
			return master_types.ErrorResponse(-500, "failed to create gamer")
		}

		// write to excel
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), name)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), clear_password)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), phone)
	}

	f.SetActiveSheet(index)

	// save file
	random_hash := hash.Md5("rand-" + strconv.Itoa(num.Random(100000, 999999)) + "-" + strconv.FormatInt(time.Now().Unix(), 16))
	date := time.Now().Format("2006-01-02")
	file_path = "generate/" + date + "/" + random_hash

	err = f.SaveAs(file_path)
	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	out_file := &master_types.File{
		Owner:    user.Id,
		Path:     file_path,
		GameId:   GameId,
		CreateAt: time.Now().Unix(),
		Hash:     random_hash,
		Size:     0,
	}

	var id master_types.PrimaryId
	err = model.ModelInsert(out_file, &id)

	if err != nil {
		return master_types.ErrorResponse(-500, err.Error())
	}

	apiCustomAdminGameImportBlueTeamResponse.FileId = id
	/* @MT-TPL-SERVICE-RESP-START */

	return master_types.SuccessResponse(apiCustomAdminGameImportBlueTeamResponse)
}

/* @MT-TPL-SERVICE-RESP-END */
