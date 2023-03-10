package shortcut

import "time"

type Reference struct {
	ID         int    `json:"id"`
	EntityType string `json:"entity_type"`
	Name       string `json:"name"`
	AppURL     string `json:"app_url"`
}

type Action struct {
	ID         int    `json:"id"`
	EntityType string `json:"entity_type"`
	Action     string `json:"action"`
	Name       string `json:"name"`
	StoryType  string `json:"story_type"`
	AppURL     string `json:"app_url"`
	Changes    Change `json:"changes"`
}

type IterationChange struct {
	New int `json:"new"`
	Old int `json:"old"`
}

type PositionChange struct {
	New int64 `json:"new"`
	Old int64 `json:"old"`
}

type WorkflowStateChange struct {
	New int `json:"new"`
	Old int `json:"old"`
}

type StartedChange struct {
	New bool `json:"new"`
	Old bool `json:"old"`
}

type StartedAtChange struct {
	New time.Time `json:"new"`
	Old time.Time `json:"old"`
}

type CompletedAtChange struct {
	New time.Time `json:"new"`
}
type CompletedChange struct {
	New bool `json:"new"`
	Old bool `json:"old"`
}

type OwnerIdsChange struct {
	Adds []string `json:"adds"`
}

type Change struct {
	IterationChange     `json:"iteration_id"`
	PositionChange      `json:"position"`
	WorkflowStateChange `json:"workflow_state_id"`
	StartedChange       `json:"started"`
	StartedAtChange     `json:"started_at"`
	CompletedAtChange   `json:"completed_at"`
	CompletedChange     `json:"completed"`
	OwnerIdsChange      `json:"owner_ids"`
}

type Webhook struct {
	ID         string      `json:"id"`
	ChangedAt  time.Time   `json:"changed_at"`
	Version    string      `json:"version"`
	PrimaryID  int         `json:"primary_id"`
	MemberID   string      `json:"member_id"`
	Actions    []Action    `json:"actions"`
	References []Reference `json:"references"`
}
