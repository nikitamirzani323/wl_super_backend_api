package controllers

import (
	"fmt"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/wl_super_backend_api/entities"
	"github.com/nikitamirzani323/wl_super_backend_api/helpers"
	"github.com/nikitamirzani323/wl_super_backend_api/models"
)

const Fieldmaster_home_redis = "LISTMASTER_BACKEND"
const Fieldmaster_home_client_redis = "LISTMASTER_FRONTEND"

func Masterhome(c *fiber.Ctx) error {
	var obj entities.Model_master
	var arraobj []entities.Model_master
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmaster_home_redis)
	jsonredis := []byte(resultredis)
	listcurr_RD, _, _, _ := jsonparser.Get(jsonredis, "listcurr")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		master_id, _ := jsonparser.GetString(value, "master_id")
		master_start, _ := jsonparser.GetString(value, "master_start")
		master_end, _ := jsonparser.GetString(value, "master_end")
		master_idcurr, _ := jsonparser.GetString(value, "master_idcurr")
		master_name, _ := jsonparser.GetString(value, "master_name")
		master_owner, _ := jsonparser.GetString(value, "master_owner")
		master_phone1, _ := jsonparser.GetString(value, "master_phone1")
		master_phone2, _ := jsonparser.GetString(value, "master_phone2")
		master_email, _ := jsonparser.GetString(value, "master_email")
		master_note, _ := jsonparser.GetString(value, "master_note")
		master_status, _ := jsonparser.GetString(value, "master_status")
		master_status_css, _ := jsonparser.GetString(value, "master_status_css")
		master_create, _ := jsonparser.GetString(value, "master_create")
		master_update, _ := jsonparser.GetString(value, "master_update")

		obj.Master_id = master_id
		obj.Master_start = master_start
		obj.Master_end = master_end
		obj.Master_idcurr = master_idcurr
		obj.Master_name = master_name
		obj.Master_owner = master_owner
		obj.Master_phone1 = master_phone1
		obj.Master_phone2 = master_phone2
		obj.Master_email = master_email
		obj.Master_note = master_note
		obj.Master_status = master_status
		obj.Master_status_css = master_status_css
		obj.Master_create = master_create
		obj.Master_update = master_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcurr_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		curr_id, _ := jsonparser.GetString(value, "curr_id")

		objcurr.Curr_id = curr_id
		arraobjcurr = append(arraobjcurr, objcurr)
	})

	if !flag {
		result, err := models.Fetch_masterHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmaster_home_redis, result, 60*time.Minute)
		fmt.Println("MASTER MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("MASTER CACHE")
		return c.JSON(fiber.Map{
			"status":   fiber.StatusOK,
			"message":  "Success",
			"record":   arraobj,
			"listcurr": arraobjcurr,
			"time":     time.Since(render_page).String(),
		})
	}
}
func MasterSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_mastersave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	result, err := models.Save_master(
		client_admin,
		client.Master_id, client.Master_idcurr, client.Master_name, client.Master_owner, client.Master_phone1, client.Master_phone2, client.Master_email, client.Master_note, client.Master_status, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_master()
	return c.JSON(result)
}

func _deleteredis_master() {
	val_master := helpers.DeleteRedis(Fieldmaster_home_redis)
	fmt.Printf("Redis Delete BACKEND CATEBANK : %d", val_master)

	val_client := helpers.DeleteRedis(Fieldcurr_home_client_redis)
	fmt.Printf("Redis Delete CLIENT CATEBANK : %d", val_client)

}
