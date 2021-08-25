package ToDoListApp

type ToDoList struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id int
	UserId int
	ListId int
}

type ToDoItem struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`
}

type ListsItem struct {
	Id int
	ListId int
	ItemId int
}
