package types

import "time"

type Employee struct {
	ID    string
	Name  string
	Email string
}

type Team struct {
	ID   string
	Name string
}

type Project struct {
	ID   string
	Name string
}

type Story struct {
	ID          int64
	Name        string
	ProjectID   int64
	TeamID      string
	EpicID      string
	MilestoneID string
	Type        string // Feature/Chore/Bug
	Estimation  int64  // Story points
	Priority    *int64 // lower number - high priority
	Severity    *int64 // lower number - high severity
	RequesterID string // Who created story (acts as accepter)
	OwnerID     []string

	CteatedAt   time.Time
	RefinedAt   time.Time // When story was refined
	PlannedAt   time.Time // When story was planned first
	StartedAt   time.Time // When story was open last time before the complete
	CompletedAt time.Time

	BlockedTime time.Duration // duration of blocking by other stories

	// Done bool // is done or rejected
}

type WorkflowTransition struct {
	StoryID       int64
	OwnerID       int64
	FromState     string
	ToState       string
	CreatedAt     string
	Duration      time.Duration // Duration without weekends
	TotalDuration time.Duration
}
