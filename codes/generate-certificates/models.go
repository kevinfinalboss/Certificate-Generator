package main

type CertificateData struct {
	UUID            string `json:"uuid"`
	ParticipantName string `json:"participant_name"`
	CompanyName     string `json:"company_name"`
	CourseName      string `json:"course_name"`
	TotalHours      string `json:"total_hours"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	DirectorName    string `json:"director_name"`
	CoordinatorName string `json:"coordinator_name"`
	CertificateID   string `json:"certificate_id"`
	IssueDate       string `json:"issue_date"`
}
