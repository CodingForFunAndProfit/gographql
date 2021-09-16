package job

type Repository interface {
	GetJobs(employeeID, company string) ([]Job, error)
	GetJob(employeeID, jobid string) (Job, error)
	Update(Job) (Job, error)
}

type Job struct {
	ID         string `json:"id"`
	EmployeeID string `json:"employeeID"`
	Company    string `json:"company"`
	Title      string `json:"title"`
	Start      string `json:"start"`
	End        string `json:"end"`
}
