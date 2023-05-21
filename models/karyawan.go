package models

type Karyawan struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"type:varchar(20)" json:"first_name"`
	LastName  string `gorm:"type:varchar(20)" json:"last_name"`
	Email     string `gorm:"type:varchar(40)" json:"email"`
	Address   string `gorm:"type:varchar(40)" json:"address"`
	InActive  string `gorm:"type:varchar(1)" json:"inactive"`
}
