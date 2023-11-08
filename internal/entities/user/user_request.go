package user

type UserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func (u UserRequest) Req2DB() UserDB {
	return UserDB(u)
}

func (u UserRequest) Req2Resp() UserResponse {
	return UserResponse{
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
	}
}
