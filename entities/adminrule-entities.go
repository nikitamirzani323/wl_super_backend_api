package entities

type Model_adminruleall struct {
	Idadmin        string `json:"adminrule_idadmin"`
	Ruleadmingroup string `json:"adminrule_rule"`
}
type Model_agenadminrule struct {
	Agenadminrule_id   string `json:"agenadminrule_id"`
	Agenadminrule_rule string `json:"agenadminrule_rule"`
}

type Controller_adminrulesave struct {
	Sdata   string `json:"sdata" validate:"required"`
	Page    string `json:"page" validate:"required"`
	Idadmin string `json:"adminrule_idadmin" validate:"required"`
	Rule    string `json:"adminrule_rule" `
}
type Controller_agenadminrulesave struct {
	Sdata              string `json:"sdata" validate:"required"`
	Page               string `json:"page" validate:"required"`
	Agenadminrule_id   string `json:"agenadminrule_id" validate:"required"`
	Agenadminrule_rule string `json:"agenadminrule_rule"`
}

type Responseredis_adminruleall struct {
	Adminrule_idadmin string `json:"adminrule_idadmin"`
	Adminrule_rule    string `json:"adminrule_rule"`
}
