package client

type Employee struct {
	Id         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	Department Department `json:"department"`
	Role       Role       `json:"role"`
}

type Role string

const (
	RoleManager  = "manager"
	RoleEmployee = "employee"
)

type Department string

const (
	DepartmentSales       = "sales"
	DepartmentEngineering = "engineering"
	DepartmentMarketing   = "marketing"
)
