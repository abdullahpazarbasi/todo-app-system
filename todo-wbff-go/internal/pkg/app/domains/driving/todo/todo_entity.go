package driving_app_domains_todo

type todoEntity struct {
	id        string
	userID    string
	label     string
	completed bool
}

func NewTodoEntity(id string, userID string, label string, completed bool) *todoEntity {
	return &todoEntity{
		id:        id,
		userID:    userID,
		label:     label,
		completed: completed,
	}
}

func (t *todoEntity) ID() string {
	return t.id
}

func (t *todoEntity) UserID() string {
	return t.userID
}

func (t *todoEntity) Label() string {
	return t.label
}

func (t *todoEntity) IsCompleted() bool {
	return t.completed
}
