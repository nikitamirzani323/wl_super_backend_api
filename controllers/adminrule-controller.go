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

const Fieldadminrule_home_redis = "LISTADMINRULE_BACKEND_ISBPANEL"
const Fieldagenadminrule_home_redis = "LISTAGENADMINRULE_BACKEND_ISBPANEL"

func Adminrulehome(c *fiber.Ctx) error {

	var obj entities.Responseredis_adminruleall
	var arraobj []entities.Responseredis_adminruleall
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldadminrule_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		Adminrule_idadmin, _ := jsonparser.GetString(value, "adminrule_idadmin")
		Adminrule_rule, _ := jsonparser.GetString(value, "adminrule_rule")

		obj.Adminrule_idadmin = Adminrule_idadmin
		obj.Adminrule_rule = Adminrule_rule
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_adminruleHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldadminrule_home_redis, result, 60*time.Minute)
		fmt.Println("ADMIN RULE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("ADMIN RULE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Agenadminrulehome(c *fiber.Ctx) error {
	var obj entities.Model_agenadminrule
	var arraobj []entities.Model_agenadminrule
	var objagen entities.Model_masteragen_share
	var arraobjagen []entities.Model_masteragen_share
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldagenadminrule_home_redis)
	jsonredis := []byte(resultredis)
	listagen_RD, _, _, _ := jsonparser.Get(jsonredis, "listagen")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		agenadminrule_id, _ := jsonparser.GetInt(value, "agenadminrule_id")
		agenadminrule_idagen, _ := jsonparser.GetString(value, "agenadminrule_idagen")
		agenadminrule_nmagen, _ := jsonparser.GetString(value, "agenadminrule_nmagen")
		agenadminrule_name, _ := jsonparser.GetString(value, "agenadminrule_name")
		agenadminrule_rule, _ := jsonparser.GetString(value, "agenadminrule_rule")
		agenadminrule_create, _ := jsonparser.GetString(value, "agenadminrule_create")
		agenadminrule_update, _ := jsonparser.GetString(value, "agenadminrule_update")

		obj.Agenadminrule_id = int(agenadminrule_id)
		obj.Agenadminrule_idagen = agenadminrule_idagen
		obj.Agenadminrule_nmagen = agenadminrule_nmagen
		obj.Agenadminrule_name = agenadminrule_name
		obj.Agenadminrule_rule = agenadminrule_rule
		obj.Agenadminrule_create = agenadminrule_create
		obj.Agenadminrule_update = agenadminrule_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listagen_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		masteragen_id, _ := jsonparser.GetString(value, "masteragen_id")
		masteragen_nmagen, _ := jsonparser.GetString(value, "masteragen_nmagen")

		objagen.Masteragen_id = masteragen_id
		objagen.Masteragen_nmagen = masteragen_nmagen
		arraobjagen = append(arraobjagen, objagen)
	})
	if !flag {
		result, err := models.Fetch_agenadminruleHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldagenadminrule_home_redis, result, 60*time.Minute)
		fmt.Println("AGEN ADMIN RULE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("AGEN ADMIN RULE CACHE")
		return c.JSON(fiber.Map{
			"status":   fiber.StatusOK,
			"message":  "Success",
			"record":   arraobj,
			"listagen": arraobjagen,
			"time":     time.Since(render_page).String(),
		})
	}
}
func AdminruleSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_adminrulesave)
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

	result, err := models.Save_adminrule(client_admin, client.Idadmin, client.Rule, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_adminrule()
	return c.JSON(result)
}
func AgenadminruleSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_agenadminrulesave)
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

	//admin, idmasteragen, nmrule, rule, sData string, idrecord int
	result, err := models.Save_agenadminrule(client_admin,
		client.Agenadminrule_idagen, client.Agenadminrule_name, client.Agenadminrule_rule, client.Sdata, client.Agenadminrule_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_adminrule()
	return c.JSON(result)
}
func _deleteredis_adminrule() {
	val_master := helpers.DeleteRedis(Fieldadminrule_home_redis)
	fmt.Printf("Redis Delete BACKEND ADMIN RULE : %d", val_master)

	val_master2 := helpers.DeleteRedis(Fieldagenadminrule_home_redis)
	fmt.Printf("Redis Delete BACKEND AGEN ADMIN RULE : %d", val_master2)

}
