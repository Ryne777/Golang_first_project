package model

type Author struct {
	User
	ToDos []ToDo `json:"to_dos"`
}

func GetAuthor() *Author {
	var author Author
	return &author
}

func GetAuthors() *[]Author {
	var author []Author
	return &author
}
