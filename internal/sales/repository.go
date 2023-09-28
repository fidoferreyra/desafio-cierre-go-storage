package sales

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

type Repository interface {
	Create(sales *domain.Sales) (int64, error)
	ReadAll() ([]*domain.Sales, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(sales *domain.Sales) (int64, error) {
	query := `INSERT INTO sales (product_id, invoices_id, quantity) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &sales.ProductId, &sales.InvoicesId, &sales.Quantity)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Sales, error) {
	query := `SELECT id, product_id, invoice_id, quantity FROM sales`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sales := make([]*domain.Sales, 0)
	for rows.Next() {
		sale := domain.Sales{}
		err := rows.Scan(&sale.Id, &sale.ProductId, &sale.InvoicesId, &sale.Quantity)
		if err != nil {
			return nil, err
		}
		sales = append(sales, &sale)
	}
	return sales, nil
}
