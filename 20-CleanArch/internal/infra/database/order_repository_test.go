package database

import (
	"database/sql"
	"testing"

	"github.com/devfullcycle/goexpert/20-CleanArch/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("dock")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenList_ThenShouldReturnOrders() {
	order1, err := entity.NewOrder("1", 19.9, 2.0)
	suite.NoError(err)
	suite.NoError(order1.CalculateFinalPrice())

	order2, err := entity.NewOrder("2", 29.9, 4.0)
	suite.NoError(err)
	suite.NoError(order2.CalculateFinalPrice())

	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order1)
	suite.NoError(err)

	err = repo.Save(order2)
	suite.NoError(err)

	orders, err := repo.List()
	suite.NoError(err)
	suite.Len(orders, 2)
}
