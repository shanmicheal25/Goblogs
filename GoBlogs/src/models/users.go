package models

//Users structure
type Users struct {
	UserID   uint   `gorm:"primary_key;type:int"  json:"user_id"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(50)" json:"email"`
	Mobile   string `gorm:"type:varchar(30)" json:"mobile"`
	Password string `gorm:"type:varchar(45)" json:"password"`
	Address  string `gorm:"type:varchar(255)" json:"address"`
}

//TableName return name of database table
func (u *Users) TableName() string {
	return "users"
}
