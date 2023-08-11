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

const database_adminrule_local = configs.DB_tbl_admingroup
const database_agen_local = configs.DB_tbl_mst_master_agen
const database_agenadminrule_local = configs.DB_tbl_mst_master_agen_admin_rule

func Fetch_adminruleHome() (helpers.Response, error) {
	var obj entities.Model_adminruleall
	var arraobj []entities.Model_adminruleall
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idadmin , ruleadmingroup 
			FROM ` + database_adminrule_local + ` 
			ORDER BY idadmin ASC  
		`

	row, err := con.QueryContext(ctx, sql_select)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idadmin_db, ruleadmingroup_db string
		)

		err = row.Scan(&idadmin_db, &ruleadmingroup_db)

		helpers.ErrorCheck(err)

		obj.Idadmin = idadmin_db
		obj.Ruleadmingroup = ruleadmingroup_db
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
func Fetch_agenadminruleHome() (helpers.Response, error) {
	var obj entities.Model_agenadminrule
	var arraobj []entities.Model_agenadminrule
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			A.idagenadminrule, B.nmagen, A.nmagenadminrule, A.ruleagenadminrule, 
			A.createagenadminrule, to_char(COALESCE(A.createagenadminruledate,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			A.updateagenadminrule, to_char(COALESCE(A.updatedateagenadminrule,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_agenadminrule_local + ` as A  
			JOIN ` + database_agen_local + ` as B ON B.idmasteragen = A.idmasteragen    
			ORDER BY A.idagenadminrule ASC  
	`

	row, err := con.QueryContext(ctx, sql_select)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idagenadminrule_db                                                                                     int
			nmagen_db, nmagenadminrule_db, ruleagenadminrule_db                                                    string
			createagenadminrule_db, createagenadminruledate_db, updateagenadminrule_db, updatedateagenadminrule_db string
		)

		err = row.Scan(&idagenadminrule_db, &nmagen_db, &nmagenadminrule_db, &ruleagenadminrule_db,
			&createagenadminrule_db, &createagenadminruledate_db, &updateagenadminrule_db, &updatedateagenadminrule_db)

		helpers.ErrorCheck(err)

		create := ""
		update := ""
		if createagenadminrule_db != "" {
			create = createagenadminrule_db + ", " + createagenadminruledate_db
		}
		if updateagenadminrule_db != "" {
			update = updateagenadminrule_db + ", " + updatedateagenadminrule_db
		}

		obj.Agenadminrule_id = idagenadminrule_db
		obj.Agenadminrule_nmagen = nmagen_db
		obj.Agenadminrule_name = nmagenadminrule_db
		obj.Agenadminrule_rule = ruleagenadminrule_db
		obj.Agenadminrule_create = create
		obj.Agenadminrule_update = update
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
func Save_adminrule(admin, idadmin, rule, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_adminrule_local, "idadmin ", idadmin)
		if !flag {
			sql_insert := `
				insert into
				` + database_adminrule_local + ` (
					idadmin 
				) values (
					$1
				) 
			`
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_adminrule_local, "INSERT", idadmin)

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
				` + database_adminrule_local + `   
				SET ruleadmingroup=$1 
				WHERE idadmin=$2 
			`
		flag_update, msg_update := Exec_SQL(sql_update, database_adminrule_local, "UPDATE", rule, idadmin)

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
func Save_agenadminrule(admin, idmasteragen, nmrule, rule, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	if sData == "New" {
		sql_insert := `
				insert into
				` + database_agenadminrule_local + ` (
					idagenadminrule, idmasteragen, nmagenadminrule,
					createagenadminrule, createagenadminruledate   
				) values (
					$1,$2,$3,
					$4,$5
				) 
			`
		field_column := database_agenadminrule_local + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, database_agenadminrule_local, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idmasteragen, nmrule,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_agenadminrule_local + `   
				SET nmagenadminrule=$1, ruleagenadminrule=$2, 
				updateagenadminrule=$3, updatedateagenadminrule=$4    
				WHERE idagenadminrule=$5 AND idmasteragen=$6   
			`
		flag_update, msg_update := Exec_SQL(sql_update, database_agenadminrule_local, "UPDATE",
			nmrule, rule,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord, idmasteragen)

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
