package custom

/* @MT-TPL-IMPORT-START */
import (
	"os"
	"strconv"
	"time"

	model "github.com/0xunion/exercise_back/src/model"
	master_types "github.com/0xunion/exercise_back/src/types"
	"github.com/0xunion/exercise_back/src/util/auth"
	"github.com/0xunion/exercise_back/src/util/hash"
	"github.com/0xunion/exercise_back/src/util/num"
	"github.com/xuri/excelize/v2"
)

/* @MT-TPL-IMPORT-END */

/* @MT-TPL-SERVICE-START */
// /api/custom/admin/game/import/red_team Service 从文件导入红队信息
func ApiCustomAdminGameImportRedTeamService(
    user *master_types.User,
    GameId master_types.PrimaryId,
    RedTeamFileId master_types.PrimaryId,
) (*master_types.MasterResponse) {
    var apiCustomAdminGameImportRedTeamResponse struct {
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
	file_id := RedTeamFileId
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
	var standard_header = []string{"Phone", "Name", "Team"}
	for i, v := range header {
		if i >= len(standard_header) {
			return master_types.ErrorResponse(-500, "invalid header format, too many columns")
		}

		if v != standard_header[i] {
			return master_types.ErrorResponse(-500, "invalid header format, column "+strconv.Itoa(i)+" should be "+standard_header[i])
		}
	}

	type temp_user struct {
		Phone  string
		Name   string
		Team   string
		TeamId int
	}

	var users []*temp_user

	// read data
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) != len(header) {
			return master_types.ErrorResponse(-500, "invalid data format, row "+strconv.Itoa(i)+" has "+strconv.Itoa(len(row))+" columns, but header has "+strconv.Itoa(len(header))+" columns")
		}

		var user temp_user
		user.Phone = row[0]
		user.Name = row[1]
		user.Team = row[2]

		users = append(users, &user)
	}

	// check how many teams
	var team_names []string
	for _, user := range users {
		team_id := -1
		for i, team_name := range team_names {
			if user.Team == team_name {
				team_id = i
				break
			}
		}

		if team_id == -1 {
			user.TeamId = len(team_names)
			team_names = append(team_names, user.Team)
		} else {
			user.TeamId = team_id
		}
	}

	// create teams
	var groups []*master_types.Group
	for _, team_name := range team_names {
		group := &master_types.Group{
			Name:        team_name,
			Parent:      user.Id,
			CreateAt:    time.Now().Unix(),
			Description: "red team",
		}

		var id master_types.PrimaryId
		err = model.ModelInsert(group, &id)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}
		group.Id = id
		groups = append(groups, group)

		read_team := &master_types.ReadTeam{
			Name:   team_name,
			GameId: GameId,
			Gid:    group.Id,
			Score:  0,
		}

		err = model.ModelInsert(read_team, nil)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}
	}

	// create users
	// create excel file to store user info
	f.Close()
	f = excelize.NewFile()

	index, err := f.NewSheet("Sheet1")
	headers := []string{"Username", "Password", "Team", "Phone"}
	header_serial := []string{"A", "B", "C", "D"}

	for i, header := range headers {
		f.SetCellValue("Sheet1", header_serial[i]+"1", header)
	}

	admin := user
	for i, user := range users {
		phone_max_length := 4
		if len(user.Phone) < phone_max_length {
			phone_max_length = len(user.Phone)
		}
		user_model := &master_types.User{
			Username: team_names[user.TeamId] + "_" + user.Phone[:phone_max_length],
			Parent:   admin.Id,
		}

		clear_password := ""

		// check if phone already existss
		phone, err := model.ModelGet[master_types.Phone](
			model.NewMongoFilter(
				model.MongoKeyFilter("phone", user.Phone),
			),
		)

		if err != nil || phone == nil {
			// create password
			clear_password = user.Phone + strconv.Itoa(num.Random(1000, 9999))
			password := &master_types.Password{
				Uid:      user_model.Id,
				Password: auth.HashPassword(hash.Md5(clear_password)),
				CreateAt: time.Now().Unix(),
			}

			err = model.ModelInsert(password, nil)
			if err != nil {
				return master_types.ErrorResponse(-500, err.Error())
			}

			// create phone
			phone := &master_types.Phone{
				Uid:      user_model.Id,
				Phone:    user.Phone,
				CreateAt: time.Now().Unix(),
			}

			err = model.ModelInsert(phone, nil)
			if err != nil {
				return master_types.ErrorResponse(-500, err.Error())
			}

			// create user
			var uid master_types.PrimaryId
			if err = model.ModelInsert(user_model, &uid); err != nil {
				return master_types.ErrorResponse(-500, err.Error())
			}

			user_model.Id = uid
		} else {
			if phone == nil {
				return master_types.ErrorResponse(-500, "phone is nil")
			}
			clear_password = "用户已存在，密码需自行回忆，平台不会储存任何明文密码"
			user_model.Id = phone.Uid
		}

		// add user to group
		group := groups[user.TeamId]
		member := &master_types.GroupMember{
			Gid:      group.Id,
			Uid:      user_model.Id,
			Role:     master_types.GROUP_MEMBER_ROLE_USER,
			CreateAt: time.Now().Unix(),
		}

		err = model.ModelInsert(member, nil)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}

		// add user to game
		gamer := &master_types.Gamer{
			GameId:   GameId,
			GroupId:  group.Id,
			Owner:    admin.Id,
			Name:     user.Name,
			Phone:    user.Phone,
			Identity: master_types.GAMER_IDENTITY_ATTACKER,
			Score:    0,
			CreateAt: time.Now().Unix(),
		}

		err = model.ModelInsert(gamer, nil)
		if err != nil {
			return master_types.ErrorResponse(-500, err.Error())
		}

		// add user to excel file
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), user_model.Username)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), clear_password)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), team_names[user.TeamId])
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), user.Phone)
	}

	f.SetActiveSheet(index)
	// Save file
	random_hash := hash.Md5("rand-" + strconv.Itoa(num.Random(100000, 999999)) + "-" + strconv.FormatInt(time.Now().Unix(), 16))
	date := time.Now().Format("2006-01-02")
	file_path = "storage/storage/generate/" + date + "/" + random_hash + ".xlsx"
	os.MkdirAll("storage/storage/generate/"+date, 0777)

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

	apiCustomAdminGameImportRedTeamResponse.FileId = id
	/* @MT-TPL-SERVICE-RESP-START */

    return master_types.SuccessResponse(apiCustomAdminGameImportRedTeamResponse)
}

    /* @MT-TPL-SERVICE-RESP-END */
