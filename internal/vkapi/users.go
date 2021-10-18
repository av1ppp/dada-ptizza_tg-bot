package vkapi

// ==================
// UsersGet
// ==================

// UsersGetParams параметры метода UsersGet.
type UsersGetParams struct {
	UserIds  string
	Fields   string
	NameCase string
}

type UserInfo struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Deactivated     string `json:"deactivated"`
	IsClosed        bool   `json:"is_closed"`
	CanAccessClosed bool   `json:"can_access_closed"`
	Sex             int    `json:"sex"`

	Photo400Orig string `json:"photo_400_orig"`
}

// UsersGet позволяет получить расширенную информацию о пользователях.
func (api *API) UsersGet(p UsersGetParams) (*[]UserInfo, error) {
	resp, err := api.Request("users.get", p, &[]UserInfo{})
	if err != nil {
		return nil, err
	}
	return resp.(*[]UserInfo), nil
}
