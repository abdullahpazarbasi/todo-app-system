package driving_app_domains_todo

type todoCandidate struct {
	userID    string
	label     string
	completed bool
}

func NewTodoCandidate(userID string, label string, completed bool) *todoCandidate {
	return &todoCandidate{
		userID:    userID,
		label:     label,
		completed: completed,
	}
}

func (t *todoCandidate) UserID() string {
	return t.userID
}

func (t *todoCandidate) Label() string {
	return t.label
}

func (t *todoCandidate) IsCompleted() bool {
	return t.completed
}
