package model

type UserDetail struct {
	Username  string
	Name      string
	Statusemp string
}

type TUser struct {
	Userid   string `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Name     string
	Statusid string `gorm:"not null"`
}

type Tabler interface {
	TableName() string
}

func (TUser) TableName() string {
	return "t_user"
}
