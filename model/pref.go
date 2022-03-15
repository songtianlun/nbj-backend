package model

type UserPrefModel struct {
	BaseModel
	UserUID                   uint64 `json:"user_uid" gorm:"column:user_uid"`
	NotionToken               string `json:"notion_token" gorm:"column:notion_token"`
	NotionBillsID             string `json:"notion_bills_id" gorm:"column:notion_bills_id"`
	NotionAccountID           string `json:"notion_account_id" gorm:"column:notion_account_id"`
	NotionInvestmentID        string `json:"notion_investment_id" gorm:"column:notion_investment_id"`
	NotionInvestmentAccountID string `json:"notion_investment_account_id" gorm:"column:notion_investment_account_id"`
	NotionBudgetID            string `json:"notion_budget_id" gorm:"column:notion_budget_id"`
}

func (up *UserPrefModel) UpdateUserPref(uid uint64) error {
	up.UserUID = uid
	return DB.DB.Save(up).Error
}

func GetPrefByUID(uid uint64) (*UserPrefModel, error) {
	up := &UserPrefModel{}
	d := DB.DB.Where("user_uid = ?", uid).First(&up)

	return up, d.Error
}
