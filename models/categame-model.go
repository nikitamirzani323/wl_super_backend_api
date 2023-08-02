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

const database_categame_local = configs.DB_tbl_mst_cate_game

func Fetch_categameHome() (helpers.Response, error) {
	var obj entities.Model_categame
	var arraobj []entities.Model_categame
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idcategame , nmcategame, statuscategame,  
			createcategame, to_char(COALESCE(createdatecategame,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updatecategame, to_char(COALESCE(updatedatecategame,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_categame_local + `  
			ORDER BY createdatecategame DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcategame_db, nmcategame_db, statuscategame_db                                    string
			createcategame_db, createdatecategame_db, updatecategame_db, updatedatecategame_db string
		)

		err = row.Scan(&idcategame_db, &nmcategame_db, &statuscategame_db,
			&createcategame_db, &createdatecategame_db, &updatecategame_db, &updatedatecategame_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		status_css := configs.STATUS_CANCEL
		if createcategame_db != "" {
			create = createcategame_db + ", " + createdatecategame_db
		}
		if updatecategame_db != "" {
			update = updatecategame_db + ", " + updatedatecategame_db
		}
		if statuscategame_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Categame_id = idcategame_db
		obj.Categame_name = nmcategame_db
		obj.Categame_status = statuscategame_db
		obj.Categame_status_css = status_css
		obj.Categame_create = create
		obj.Categame_update = update
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
func Save_categame(admin, idrecord, name, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_categame_local, "idcategame", idrecord)
		if !flag {
			sql_insert := `
				insert into
				` + database_categame_local + ` (
					idcategame , nmcategame, statuscategame, 
					createcategame, createdatecategame  
				) values (
					$1, $2, $3,   
					$4, $5
				)
			`

			flag_insert, msg_insert := Exec_SQL(sql_insert, database_categame_local, "INSERT",
				idrecord, name, status,
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
				` + database_categame_local + `  
				SET nmcategame=$1, statuscategame=$2,   
				updatecategame=$3, updatedatecategame=$4    
				WHERE idcategame=$5   
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_categame_local, "UPDATE",
			name, status,
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
