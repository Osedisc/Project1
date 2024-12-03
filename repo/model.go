package repo

type CreateRequest struct {
	Id          string `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	SurName     string `gorm:"column:sur_name"`
	DateOfBirth string `gorm:"column:date_of_birth"`
}

type CreateResponse struct {
}

type GetUserByIdRequest struct {
	Id string `gorm:"column:id"`
}

type GetUserByIdResponse struct {
	Id          string `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	SurName     string `gorm:"column:sur_name"`
	DateOfBirth string `gorm:"column:date_of_birth"`
}

func TableName() string {
	return "tbl_user"
}

type UpdateUserRequest struct {
	Id          string `gorm:"column:id"`
	Name        string `gorm:"column:name"`
	SurName     string `gorm:"column:sur_name"`
	DateOfBirth string `gorm:"column:date_of_birth"`
}

type UpdateUserResponse struct {
}
