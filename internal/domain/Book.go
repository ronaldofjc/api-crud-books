package domain

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func (book Book) IsEmpty() bool {
	return book == Book{}
}
