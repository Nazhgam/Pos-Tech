package structures

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Task   string `json:"task"`
	Author string `json:"author"`
}
