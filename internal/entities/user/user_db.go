package user

type UserDB struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

func (u UserDB) DB2Req() UserRequest {
	return UserRequest{
		Username:  &u.Username,
		Password:  &u.Password,
		FirstName: &u.FirstName,
		LastName:  &u.LastName,
		Email:     &u.Email,
		Phone:     &u.Phone,
	}
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
