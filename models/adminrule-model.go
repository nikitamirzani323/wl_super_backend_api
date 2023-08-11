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

const database_adminrule_local = configs.DB_tbl_admingroup
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
			idagenadminrule, COALESCE(ruleagenadminrule,'')  
			FROM ` + database_agenadminrule_local + ` 
			ORDER BY idagenadminrule ASC  
		`

	row, err := con.QueryContext(ctx, sql_select)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idagenadminrule_db, ruleagenadminrule_db string
		)

		err = row.Scan(&idagenadminrule_db, &ruleagenadminrule_db)

		helpers.ErrorCheck(err)

		obj.Agenadminrule_id = idagenadminrule_db
		obj.Agenadminrule_rule = ruleagenadminrule_db
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
func Save_agenadminrule(admin, idadmin, rule, sData string) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_agenadminrule_local, "idagenadminrule", idadmin)
		if !flag {
			sql_insert := `
				insert into
				` + database_agenadminrule_local + ` (
					idagenadminrule, ruleagenadminrule, 
					createagenadminrule, createagenadminruledate   
				) values (
					$1,$2,
					$3,$4
				) 
			`
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_agenadminrule_local, "INSERT",
				idadmin, "", admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

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
				` + database_agenadminrule_local + `   
				SET ruleagenadminrule=$1, 
				updateagenadminrule=$2, updatedateagenadminrule=$3   
				WHERE idagenadminrule=$4  
			`
		flag_update, msg_update := Exec_SQL(sql_update, database_agenadminrule_local, "UPDATE",
			rule, admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idadmin)

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
