package domain

type EffortRecord struct {
	EffortDate    string
	EffortExpense string
	JiraID        string
	Description   string
	Qualification []string
}

type Appointment struct{}
