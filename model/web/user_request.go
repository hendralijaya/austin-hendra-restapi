package web

type UserCreateRequest struct {
	Username string `json:"username" binding:"required,min=3,max=255"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

type UserUpdateRequest struct {
	Id       uint64
	Username string `json:"username" binding:"required,min=3,max=255"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}