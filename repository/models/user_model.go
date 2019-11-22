package models

type UserModel struct {
	BaseModel `storm:"inline"`
	Username  string `json:"username" storm:"unique"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" storm:"unique"`
	Password  string `json:"password"`
}

type UserModelWithOutPassword struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

func (user *UserModel) GetModelWithOutPassword() UserModelWithOutPassword {
	model := UserModelWithOutPassword{
		Username: user.Username,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
	return model
}
