package domain

type Book struct {
	ID          int    `json:"-" db:"id"`
	Title       string `json:"title" db:"title"`
	Body        string `json:"body" db:"body"`
	Description string `json:"description" db:"description"`
}

type Products struct {
	ID    int     `db:"id"`
	Name  string  `db:"name"`
	Type  string  `db:"type"`
	Count int     `db:"count"`
	Price float64 `db:"price"`
}

func NewBook(dados interface{}) Book {
	book := dados.(Book)
	return Book{
		Title:       book.Title,
		Body:        book.Body,
		Description: book.Description,
	}
}

type RepositoryBook interface {
	GetAll() ([]Book, error)
	GetOne(id int) (Book, error)
	Store(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(id int) error
}

type RepositoryProducts interface {
	GetAll() ([]Products, error)
	GetOne(id int) (Products, error)
	Store(book Products) (Products, error)
	Update(book Products) (Products, error)
	Delete(id int) error
}
