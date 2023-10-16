package usecase_port

type Todo interface {
	ID() string
	UserID() string
	Value() string
	IsCompleted() bool
}
