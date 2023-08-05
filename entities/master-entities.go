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
	Master_listadmin  interface{} `json:"master_listadmin"`
	Master_create     string      `json:"master_create"`
	Master_update     string      `json:"master_update"`
}
type Model_masteradmin struct {
	Masteradmin_id         int    `json:"masteradmin_id"`
	Masteradmin_tipe       string `json:"masteradmin_tipe"`
	Masteradmin_username   string `json:"masteradmin_username"`
	Masteradmin_name       string `json:"masteradmin_name"`
	Masteradmin_phone1     string `json:"masteradmin_phone1"`
	Masteradmin_phone2     string `json:"masteradmin_phone2"`
	Masteradmin_status     string `json:"masteradmin_status"`
	Masteradmin_status_css string `json:"masteradmin_status_css"`
	Masteradmin_create     string `json:"masteradmin_create"`
	Masteradmin_update     string `json:"masteradmin_update"`
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
type Controller_masteradminsave struct {
	Page                 string `json:"page" validate:"required"`
	Sdata                string `json:"sdata" validate:"required"`
	Masteradmin_id       int    `json:"masteradmin_id" `
	Masteradmin_idmaster string `json:"masteradmin_idmaster" validate:"required"`
	Masteradmin_tipe     string `json:"masteradmin_tipe" validate:"required"`
	Masteradmin_username string `json:"masteradmin_username" validate:"required"`
	Masteradmin_password string `json:"masteradmin_password" `
	Masteradmin_name     string `json:"masteradmin_name" validate:"required"`
	Masteradmin_phone1   string `json:"masteradmin_phone1" validate:"required"`
	Masteradmin_phone2   string `json:"masteradmin_phone2" `
	Masteradmin_status   string `json:"masteradmin_status" validate:"required"`
}
type Controller_masteragensave struct {
	Page                  string `json:"page" validate:"required"`
	Sdata                 string `json:"sdata" validate:"required"`
	Masteragen_id         string `json:"masteragen_id"`
	Masteragen_idmaster   string `json:"masteragen_idmaster" validate:"required"`
	Masteragen_idcurr     string `json:"masteragen_idcurr" validate:"required"`
	Masteragen_name       string `json:"masteragen_name" validate:"required"`
	Masteragen_owner      string `json:"masteragen_owner" validate:"required"`
	Masteragen_phone1     string `json:"masteragen_phone1" validate:"required"`
	Masteragen_phone2     string `json:"masteragen_phone2"`
	Masteragen_email      string `json:"masteragen_email"`
	Masteragen_note       string `json:"masteragen_note"`
	Masteragen_bank_id    string `json:"masteragen_bank_id" validate:"required"`
	Masteragen_bank_norek string `json:"masteragen_bank_norek" validate:"required"`
	Masteragen_bank_name  string `json:"masteragen_bank_name" validate:"required"`
	Masteragen_status     string `json:"masteragen_status" validate:"required"`
}
