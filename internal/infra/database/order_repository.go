package database

import (
	"database/sql"

	"github.com/carlosgenuino/fieldservice/internal/entity"
)

type OrderRepositoryImpl struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		Db: db,
	}
}

func (o *OrderRepositoryImpl) Save(order *entity.Order) error {
	_, err := o.Db.Exec("INSERT INTO ORDERS (ID, PRICE, TAX, FINAL_PRICE) VALUES (?,?,?,?)", order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderRepositoryImpl) GetTotalTransactions() (int, error) {
	var total int
	err := o.Db.QueryRow("COUNT(*) FROM ORDERS").Scan(&total)

	if err != nil {
		return 0, err
	}

	return total, nil
}
