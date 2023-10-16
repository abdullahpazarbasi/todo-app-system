package domain_todo

type todoTagEntity struct {
	id  string
	key string
}

func (e *todoTagEntity) ID() string {
	return e.id
}

func (e *todoTagEntity) Key() string {
	return e.key
}
