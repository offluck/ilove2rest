package user

type UserDB struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func (u UserDB) DB2Req() UserRequest {
	return UserRequest(u)
}

func (u UserDB) DB2Resp() UserResponse {
	return UserResponse{
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}
}
