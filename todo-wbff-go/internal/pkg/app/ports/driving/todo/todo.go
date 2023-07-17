package driving_app_ports_todo

type Todo interface {
	UserID() string
	Label() string
	IsCompleted() bool
}
