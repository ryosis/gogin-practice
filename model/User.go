package model

type TUser struct {
	Userid   string `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Name     string
}

type Tabler interface {
	TableName() string
}

func (TUser) TableName() string {
	return "t_user"
}
