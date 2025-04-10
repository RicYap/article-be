package article

// Skeleton model
type User struct {
	UserID      int    `db:"UserID" json:"user_id"`
	Name        string `db:"Name" json:"name"`
	UserName    string `db:"UserName" json:"username"`
	RoleID      int    `db:"RoleID" json:"role_id"`
	RoleName    string `db:"RoleName" json:"role_name"`
	Password    string `db:"Name" json:"password"`
	PhoneNumber string `db:"PhoneNumber" json:"phone_number"`
}
