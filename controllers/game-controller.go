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

const Fieldgame_home_redis = "LISTGAME_BACKEND"
const Fieldgame_home_client_redis = "LISTGAME_FRONTEND"

func Gamehome(c *fiber.Ctx) error {
	var obj entities.Model_game
	var arraobj []entities.Model_game
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldgame_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		game_id, _ := jsonparser.GetString(value, "game_id")
		game_idcategame, _ := jsonparser.GetString(value, "game_idcategame")
		game_idprovider, _ := jsonparser.GetInt(value, "game_idprovider")
		game_name, _ := jsonparser.GetString(value, "game_name")
		game_urlstaging, _ := jsonparser.GetString(value, "game_urlstaging")
		game_urlproduction, _ := jsonparser.GetString(value, "game_urlproduction")
		game_status, _ := jsonparser.GetString(value, "game_status")
		game_status_css, _ := jsonparser.GetString(value, "game_status_css")
		game_create, _ := jsonparser.GetString(value, "game_create")
		game_update, _ := jsonparser.GetString(value, "game_update")

		obj.Game_id = game_id
		obj.Game_idcategame = game_idcategame
		obj.Game_idprovider = int(game_idprovider)
		obj.Game_name = game_name
		obj.Game_urlstaging = game_urlstaging
		obj.Game_urlproduction = game_urlproduction
		obj.Game_status = game_status
		obj.Game_status_css = game_status_css
		obj.Game_create = game_create
		obj.Game_update = game_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_gameHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcategame_home_redis, result, 60*time.Minute)
		fmt.Println("GAME MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("GAME CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
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

	_deleteredis_game()
	return c.JSON(result)
}

func _deleteredis_game() {
	val_master := helpers.DeleteRedis(Fieldgame_home_redis)
	fmt.Printf("Redis Delete BACKEND GAME : %d", val_master)

	val_client := helpers.DeleteRedis(Fieldgame_home_client_redis)
	fmt.Printf("Redis Delete CLIENT GAME : %d", val_client)

}
