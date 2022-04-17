package dto

type ProjectCreateReq struct {
	Name         string
	State        string
	Progress     float64
	Owner        string
	Participants []string
}

type ProjectUpdateReq struct {
	ID       uint
	Name     string
	State    string
	Progress float64
}

type Project struct {
	ID           uint
	Name         string
	State        string
	Progress     float64
	Owner        *Employee
	Participants []*Employee
}

type Employee struct {
	Id         string
	FirstName  string
	LastName   string
	Department string `json:",omitempty"`
}
