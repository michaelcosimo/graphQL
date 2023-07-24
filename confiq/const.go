package confiq

var Users = map[string]string{
	"client_api_key": "client",
	"admin_api_key":  "admin",
}

// Define the roles
const (
	RoleClient = "client"
	RoleAdmin  = "admin"
)
