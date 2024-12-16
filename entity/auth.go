package entity

type User struct {
	ID       uint64 `gorm:"column:user_id;primaryKey;autoIncrement;not null;<-create"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}
