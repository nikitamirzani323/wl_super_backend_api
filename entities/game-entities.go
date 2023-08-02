package entities

type Model_game struct {
	Game_id            string `json:"game_id"`
	Game_idcategame    string `json:"game_idcategame"`
	Game_idprovider    int    `json:"game_idprovider"`
	Game_name          string `json:"game_name"`
	Game_urlstaging    string `json:"game_urlstaging"`
	Game_urlproduction string `json:"game_urlproduction"`
	Game_status        string `json:"game_status"`
	Game_status_css    string `json:"game_status_css"`
	Game_create        string `json:"game_create"`
	Game_update        string `json:"game_update"`
}

type Controller_gamesave struct {
	Page               string `json:"page" validate:"required"`
	Sdata              string `json:"sdata" validate:"required"`
	Game_id            string `json:"game_id"`
	Game_idcategame    string `json:"game_idcategame" validate:"required"`
	Game_idprovider    int    `json:"game_idprovider" validate:"required"`
	Game_name          string `json:"game_name" validate:"required"`
	Game_urlstaging    string `json:"game_urlstaging" validate:"required"`
	Game_urlproduction string `json:"game_urlproduction" validate:"required"`
	Game_status        string `json:"game_status" validate:"required"`
}
