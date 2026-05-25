package model

type AccountToPay struct {
	ID_USER             int     `json:"id_user"`
	DESCRIPTION         string  `json:"description"`
	DESCRIPTION_DETAILS string  `json:"description_details"`
	DATE_ACTION         string  `json:"date_action"`
	DATE_PREVIOUS       string  `json:"date_previous"`
	VALUE_PAG           float64 `json:"value_pag"`
	VALUE_ADD           float64 `json:"value_add"`
	VALUE_DISCOUNT      float64 `json:"value_discount"`
	NAME_PAG            string  `json:"name_pag"`
	PAID                bool    `json:"paid"`
}
