package utils

type User struct {
	UserId      string `json:"user_id" gorm:"primaryKey;column:user_id"`
	UserName    string `json:"userName" gorm:"column:userName"`
	UserRole    string `json:"userRole" gorm:"column:userRole"`
	UserEmail   string `json:"userEmail" gorm:"column:userEmail"`
	CreatedDate string `json:"createdDate" gorm:"column:createdDate"`
}
