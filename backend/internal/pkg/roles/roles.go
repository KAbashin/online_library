package roles

const (
	RoleUser       = "user"
	RoleNewUser    = "new-user"
	RoleAdmin      = "admin"
	RoleSuperAdmin = "superadmin"
)

func IsAdmin(role string) bool {
	return role == RoleAdmin || role == RoleSuperAdmin
}

func IsSuperAdmin(role string) bool {
	return role == RoleSuperAdmin
}
