package driving_adapter_api_views

type Todo struct {
	UserID string             `json:"user_id"`
	ID     string             `json:"id"`
	Label  string             `json:"label"`
	Tags   *TodoTagCollection `json:"tags"`
}
