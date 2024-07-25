package services

type Todo struct {
	Name     string
	Id       int
	Complete bool
}

var (
	Count int = 0
	Todos []Todo
)
