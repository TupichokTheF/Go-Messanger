package schemas

type UserInfoSchema struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	UserEmail string `json:"user_email"`
}
