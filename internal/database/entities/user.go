package entities

type User struct {
	Id         uint64 `gorm:"primaryKey;column:id"`
	Email      string `gorm:"column:email"`
	Password   string `gorm:"column:password"`
	ApiKeyHash string `gorm:"column:api_key_hash"`
}

func (user User) TableName() string {
	return "users"
}
