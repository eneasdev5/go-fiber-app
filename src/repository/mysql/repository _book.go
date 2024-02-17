package mysql

import (
	"database/sql"

	"github.com/eneasdev5/go-fiber-app/src/domain"
)

type mysqlDBRepositoryBook struct {
	db *sql.DB
}

func NewMysqlDBRepositoryBook(db *sql.DB) domain.RepositoryBook {
	return &mysqlDBRepositoryBook{db: db}
}

func (sql *mysqlDBRepositoryBook) GetAll() ([]domain.Book, error) {
	var booksList []domain.Book = []domain.Book{}
	rows, err := sql.db.Query("SELECT id,title,body,description FROM u474998508_books.book;")
	if err != nil {
		return booksList, err
	}

	defer rows.Close()
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Body,
			&book.Description,
		)
		if err != nil {
			return booksList, err
		}
		booksList = append(booksList, book)
	}

	return booksList, nil
}
func (sql *mysqlDBRepositoryBook) GetOne(id int) (domain.Book, error) {
	return domain.Book{}, nil
}
func (sql *mysqlDBRepositoryBook) Store(book domain.Book) (domain.Book, error) {
	stmt, err := sql.db.Prepare("INSERT INTO u474998508_books.book VALUES(DEFAULT, ?,?,?);")
	if err != nil {
		return domain.Book{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		&book.Title,
		&book.Body,
		&book.Description,
	)
	if err != nil {
		return book, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return book, err
	}
	book.ID = int(lastID)

	return book, nil
}
func (sql *mysqlDBRepositoryBook) Update(book domain.Book) (domain.Book, error) {
	return domain.Book{}, nil
}
func (sql *mysqlDBRepositoryBook) Delete(id int) error {
	return nil
}
