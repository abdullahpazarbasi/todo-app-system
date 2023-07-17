package driven_app_domains_todo

type TodoEntity struct {
	ID     string            `field:"id"`
	UserID string            `field:"user_id"`
	Label  string            `field:"label"`
	Tags   TodoTagCollection `field:"tags"`
}
