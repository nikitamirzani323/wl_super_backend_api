package entities

type Model_master struct {
	Master_id         string      `json:"master_id"`
	Master_start      string      `json:"master_start"`
	Master_end        string      `json:"master_end"`
	Master_idcurr     string      `json:"master_idcurr"`
	Master_name       string      `json:"master_name"`
	Master_owner      string      `json:"master_owner"`
	Master_phone      string      `json:"master_phone"`
	Master_email      string      `json:"master_email"`
	Master_status     string      `json:"master_status"`
	Master_status_css string      `json:"master_status_css"`
	Master_list       interface{} `json:"master_list"`
	Master_create     string      `json:"master_create"`
	Master_update     string      `json:"master_update"`
}

type Controller_mastersave struct {
	Page          string `json:"page" validate:"required"`
	Sdata         string `json:"sdata" validate:"required"`
	Master_id     string `json:"master_id" validate:"required"`
	Master_idcurr string `json:"master_idcurr" validate:"required"`
	Master_name   string `json:"master_name" validate:"required"`
	Master_owner  string `json:"master_owner"`
	Master_phone  string `json:"master_phone"`
	Master_email  string `json:"master_email"`
	Master_status string `json:"master_status" validate:"required"`
}
