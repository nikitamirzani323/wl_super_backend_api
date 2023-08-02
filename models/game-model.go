package models

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/wl_super_backend_api/configs"
	"github.com/nikitamirzani323/wl_super_backend_api/db"
	"github.com/nikitamirzani323/wl_super_backend_api/entities"
	"github.com/nikitamirzani323/wl_super_backend_api/helpers"
	"github.com/nleeper/goment"
)

const database_game_local = configs.DB_tbl_mst_game

func Fetch_gameHome() (helpers.Response, error) {
	var obj entities.Model_game
	var arraobj []entities.Model_game
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idgame , idcategame, idprovider, nmgame, urlstaging, urlproduction, statusgame  
			creategame, to_char(COALESCE(createdategame,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updategame, to_char(COALESCE(updatedategame,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_game_local + `  
			ORDER BY createdategame DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idprovider_db                                                      int
			idgame_db, idcategame_db                                           string
			nmgame_db, urlstaging_db, urlproduction_db, statusgame_db          string
			creategame_db, createdategame_db, updategame_db, updatedategame_db string
		)

		err = row.Scan(&idgame_db, &idcategame_db, &idprovider_db,
			&nmgame_db, &urlstaging_db, &urlproduction_db, &statusgame_db,
			&creategame_db, &createdategame_db, &updategame_db, &updatedategame_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if creategame_db != "" {
			create = creategame_db + ", " + createdategame_db
		}
		if updategame_db != "" {
			update = updategame_db + ", " + updatedategame_db
		}
		if statusgame_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Game_id = idcategame_db
		obj.Game_idcategame = idcategame_db
		obj.Game_idprovider = idprovider_db
		obj.Game_name = nmgame_db
		obj.Game_urlstaging = urlstaging_db
		obj.Game_urlproduction = urlproduction_db
		obj.Game_status = statusgame_db
		obj.Game_status_css = status_css
		obj.Game_create = create
		obj.Game_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Save_game(admin, idrecord, idcategame, name, urlstaging, urlproduction, status, sData string, idprovider int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_game_local, "idgame", idrecord)
		if !flag {
			sql_insert := `
				insert into
				` + database_game_local + ` (
					idgame, idcategame ,idprovider , 
					nmgame, urlstaging, urlproduction, statusgame, 
					creategame, createdategame  
				) values (
					$1, $2, $3,   
					$4, $5, $6, $7,   
					$8, $9 
				)
			`

			flag_insert, msg_insert := Exec_SQL(sql_insert, database_game_local, "INSERT",
				idrecord, idcategame, idprovider,
				name, urlstaging, urlproduction,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_game_local + `  
				SET idcategame=$1, idprovider=$2,   
				nmgame=$3, urlstaging=$4, urlproduction=$5, statusgame=$6,      
				updategame=$7, updatedategame=$8     
				WHERE idgame=$9    
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_categame_local, "UPDATE",
			idcategame, idprovider, name, urlstaging, urlproduction, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
