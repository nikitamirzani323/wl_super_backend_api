package models

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/wl_super_backend_api/configs"
	"github.com/nikitamirzani323/wl_super_backend_api/db"
	"github.com/nikitamirzani323/wl_super_backend_api/entities"
	"github.com/nikitamirzani323/wl_super_backend_api/helpers"
	"github.com/nleeper/goment"
)

const database_catebank_local = configs.DB_tbl_mst_cate_bank

func Fetch_catebankHome() (helpers.Response, error) {
	var obj entities.Model_catebank
	var arraobj []entities.Model_catebank
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idcatebank , nmcatebank, statuscatebank,  
			createcatebank, to_char(COALESCE(createdatecatebank,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updatecatebank, to_char(COALESCE(updatedatecatebank,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_catebank_local + `  
			ORDER BY createdatecatebank DESC   `

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcatebank_db                                                                      int
			nmcatebank_db, statuscatebank_db                                                   string
			createcatebank_db, createdatecatebank_db, updatecatebank_db, updatedatecatebank_db string
		)

		err = row.Scan(&idcatebank_db, &nmcatebank_db, &statuscatebank_db,
			&createcatebank_db, &createdatecatebank_db, &updatecatebank_db, &updatedatecatebank_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if createcatebank_db != "" {
			create = createcatebank_db + ", " + createdatecatebank_db
		}
		if updatecatebank_db != "" {
			update = updatecatebank_db + ", " + updatedatecatebank_db
		}

		obj.Catebank_id = idcatebank_db
		obj.Catebank_name = nmcatebank_db
		obj.Catebank_status = statuscatebank_db
		obj.Catebank_create = create
		obj.Catebank_update = update
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
func Save_catebank(admin, name, status, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	if sData == "New" {
		sql_insert := `
				insert into
				` + database_catebank_local + ` (
					idcatebank , nmcatebank, statuscatebank, 
					createcatebank, createdatecatebank  
				) values (
					$1, $2, $3,   
					$4, $5
				)
			`
		field_column := database_catebank_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_catebank_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), name, status,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_catebank_local + `  
				SET nmcatebank=$1, statuscatebank=$2,   
				updatecatebank=$3, updatedatecatebank=$4    
				WHERE idcatebank=$5   
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_catebank_local, "UPDATE",
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
