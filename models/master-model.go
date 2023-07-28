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

const database_master_local = configs.DB_tbl_mst_master

func Fetch_masterHome() (helpers.Response, error) {
	var obj entities.Model_master
	var arraobj []entities.Model_master
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idmaster ,  
			to_char(COALESCE(startjoinmaster,now()), 'YYYY-MM-DD HH24:MI:SS'),
			to_char(COALESCE(endjoinmaster,now()), 'YYYY-MM-DD HH24:MI:SS'),
			idcurr , nmmaster, nmowner, phoneowner, emailowner, statusmaster, 
			createmaster, to_char(COALESCE(createdatemaster,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updatemaster, to_char(COALESCE(updatedatemaster,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_master_local + `  
			ORDER BY createdatemaster DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idmaster_db, idcurr_db, nmmaster_db, nmowner_db, phoneowner_db, emailowner_db, statusmaster_db string
			startjoinmaster_db, endjoinmaster_db                                                           string
			createmaster_db, createdatemaster_db, updatemaster_db, updatedatemaster_db                     string
		)

		err = row.Scan(&idmaster_db, &startjoinmaster_db, &endjoinmaster_db,
			&nmmaster_db, &nmowner_db, &phoneowner_db, &emailowner_db, &statusmaster_db,
			&createmaster_db, &createdatemaster_db, &updatemaster_db, &updatedatemaster_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if createmaster_db != "" {
			create = createmaster_db + ", " + createdatemaster_db
		}
		if updatemaster_db != "" {
			update = updatemaster_db + ", " + updatedatemaster_db
		}

		obj.Master_id = idmaster_db
		obj.Master_start = startjoinmaster_db
		obj.Master_end = endjoinmaster_db
		obj.Master_idcurr = idcurr_db
		obj.Master_name = nmmaster_db
		obj.Master_phone = phoneowner_db
		obj.Master_email = emailowner_db
		obj.Master_create = create
		obj.Master_update = update
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
func Save_master(admin, idrecord, idcurr, name, owner, phone, email, status, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_master_local, "idmaster", idrecord)
		if !flag {
			sql_insert := `
				insert into
				` + database_master_local + ` (
					idmaster , startjoinmaster, endjoinmaster, 
					idcurr , nmmaster, nmowner , phoneowner, emailowner, statusmaster, 
					createmaster, createdatemaster,  
				) values (
					$1, $2, $3, $4,   
					$4, $5, $6, $7, $8, $9,    
					$10, $11
				)
			`
			start := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_master_local, "INSERT",
				idrecord, start, start, idcurr, name, owner, email, status,
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
				` + database_master_local + `  
				SET idcurr=$1, nmmaster=$2, nmowner=$3,  phoneowner=$4,  emailowner=$5, statusmaster=$6,  
				updatemaster=$7, updatedatemaster=$8    
				WHERE idmaster=$9   
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_master_local, "UPDATE",
			idcurr, name, owner, phone, email, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

		if flag_update {
			flag = true
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
