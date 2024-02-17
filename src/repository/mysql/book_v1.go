package mysql

type Book struct {
	Title     string
	TotalPage int
	Author    string
}

type IBook interface {
	GetAllBook() []Book
}

func NewBook() IBook {
	return &Book{}
}

func (b *Book) GetAllBook() []Book {
	return loadBooks()
}

func loadBooks() []Book {
	return []Book{
		{
			Title:     "Domine Seu Foco",
			TotalPage: 60,
			Author:    "I. C. ROBLEDO",
		},
		{
			Title:     "Mente Suprema",
			TotalPage: 28,
			Author:    "Desconhecido",
		},
		{
			Title:     "MindSet",
			TotalPage: 120,
			Author:    "Carol S. Dweck",
		},
	}
}
