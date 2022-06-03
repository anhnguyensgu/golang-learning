package todo

type TodoStatus string

const (
	NEW         TodoStatus = "NEW"
	IN_PROGRESS TodoStatus = "IN_PROGRESS"
	COMPLETED   TodoStatus = "COMPLETED"
)

type TodoStatusUpdateRequest struct {
	Status TodoStatus `json:"status"`
}

type TodoResponse struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
