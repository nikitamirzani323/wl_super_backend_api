package entities

type Model_master struct {
	Master_id         string      `json:"master_id"`
	Master_start      string      `json:"master_start"`
	Master_end        string      `json:"master_end"`
	Master_idcurr     string      `json:"master_idcurr"`
	Master_name       string      `json:"master_name"`
	Master_owner      string      `json:"master_owner"`
	Master_phone1     string      `json:"master_phone1"`
	Master_phone2     string      `json:"master_phone2"`
	Master_email      string      `json:"master_email"`
	Master_note       string      `json:"master_note"`
	Master_bank_id    string      `json:"master_bank_id"`
	Master_bank_norek string      `json:"master_bank_norek"`
	Master_bank_name  string      `json:"master_bank_name"`
	Master_status     string      `json:"master_status"`
	Master_status_css string      `json:"master_status_css"`
	Master_credit_in  int         `json:"master_credit_in"`
	Master_credit_out int         `json:"master_credit_out"`
	Master_list       interface{} `json:"master_list"`
	Master_create     string      `json:"master_create"`
	Master_update     string      `json:"master_update"`
}

type Controller_mastersave struct {
	Page              string `json:"page" validate:"required"`
	Sdata             string `json:"sdata" validate:"required"`
	Master_id         string `json:"master_id" validate:"required"`
	Master_idcurr     string `json:"master_idcurr" validate:"required"`
	Master_name       string `json:"master_name" validate:"required"`
	Master_owner      string `json:"master_owner" validate:"required"`
	Master_phone1     string `json:"master_phone1" validate:"required"`
	Master_phone2     string `json:"master_phone2"`
	Master_email      string `json:"master_email"`
	Master_note       string `json:"master_note"`
	Master_bank_id    string `json:"master_bank_id" validate:"required"`
	Master_bank_norek string `json:"master_bank_norek" validate:"required"`
	Master_bank_name  string `json:"master_bank_name" validate:"required"`
	Master_status     string `json:"master_status" validate:"required"`
}
