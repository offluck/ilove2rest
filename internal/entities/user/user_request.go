package user

type UserRequest struct {
	Username  *string `json:"username"`
	Password  *string `json:"password"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

func (u UserRequest) Req2DB() UserDB {
	return UserDB{
		Username:  *u.Username,
		Password:  *u.Password,
		FirstName: *u.FirstName,
		LastName:  *u.LastName,
		Email:     *u.Email,
		Phone:     *u.Phone,
	}
}

func (u UserRequest) Req2Resp() UserResponse {
	return UserResponse{
		Username:  *u.Username,
		FirstName: *u.FirstName,
		LastName:  *u.LastName,
		Email:     *u.Email,
		Phone:     *u.Phone,
	}
}

func (u UserRequest) IsValid() bool {
	if u.Username == nil || *u.Username == "" {
		return false
	}

	if u.Password == nil || *u.Password == "" {
		return false
	}

	if u.FirstName == nil || *u.FirstName == "" {
		return false
	}

	if u.LastName == nil || *u.LastName == "" {
		return false
	}

	if u.Email == nil || *u.Email == "" {
		return false
	}

	if u.Phone == nil || *u.Phone == "" {
		return false
	}

	return true
}
