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

const Fieldcategame_home_redis = "LISTCATEGAME_BACKEND"
const Fieldcategame_home_client_redis = "LISTCATEGAME_FRONTEND"

func CateGamehome(c *fiber.Ctx) error {
	var obj entities.Model_categame
	var arraobj []entities.Model_categame
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcategame_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		categame_id, _ := jsonparser.GetString(value, "categame_id")
		categame_name, _ := jsonparser.GetString(value, "categame_name")
		categame_status, _ := jsonparser.GetString(value, "categame_status")
		categame_status_css, _ := jsonparser.GetString(value, "categame_status_css")
		categame_create, _ := jsonparser.GetString(value, "categame_create")
		categame_update, _ := jsonparser.GetString(value, "categame_update")

		obj.Categame_id = categame_id
		obj.Categame_name = categame_name
		obj.Categame_status = categame_status
		obj.Categame_status_css = categame_status_css
		obj.Categame_create = categame_create
		obj.Categame_update = categame_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_categameHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcategame_home_redis, result, 60*time.Minute)
		fmt.Println("CATEGAME MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("CATEGAME CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func CateGameSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_categamesave)
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

	// admin, idrecord, name, status, sData string
	result, err := models.Save_categame(
		client_admin,
		client.Categame_id, client.Categame_name, client.Categame_status, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_categame()
	return c.JSON(result)
}
func GameSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_gamesave)
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

	// admin, idrecord, idcategame, name, urlstaging, urlproduction, status, sData string, idprovider int) (helpers.Response, error
	result, err := models.Save_game(
		client_admin,
		client.Game_id, client.Game_idcategame, client.Game_name,
		client.Game_urlstaging, client.Game_urlproduction, client.Game_status,
		client.Sdata, client.Game_idprovider)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_categame()
	return c.JSON(result)
}
func _deleteredis_categame() {
	val_master := helpers.DeleteRedis(Fieldcategame_home_redis)
	fmt.Printf("Redis Delete BACKEND CATEGAME : %d", val_master)

	val_client := helpers.DeleteRedis(Fieldcategame_home_client_redis)
	fmt.Printf("Redis Delete CLIENT CATEGAME : %d", val_client)

}
