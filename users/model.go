package users

type UserModel struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"column:email;uniqueIndex"`
	PasswordHash string `gorm:"column:password;not null"`
	FirstName    string `gorm:"column:first_name;not null"`
	LastName     string `gorm:"column:last_name;not null"`
}
