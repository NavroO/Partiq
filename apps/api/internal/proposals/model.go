package proposals

type Proposal struct {
	ID        int    `json:"id"`
	ProcessID int    `json:"process_id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
