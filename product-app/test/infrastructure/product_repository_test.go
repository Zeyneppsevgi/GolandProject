package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"os"
	"product-app/common/postgresql"
	"product-app/persistence"
	"testing"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background() // Global ctx değişkeni atanıyor
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})
	productRepository = persistence.NewProductRepository(dbPool)
	fmt.Println("Before all tests")
	exitCode := m.Run()
	fmt.Println("After all tests")
	os.Exit(exitCode)
}

func setup() {
	TestDataInitialize(ctx, dbPool) // Global ctx ve dbPool kullanılıyor
}

func clear() {
	TruncateTestData(ctx, dbPool) // Global ctx ve dbPool kullanılıyor
}

func TestGetAllProducts(t *testing.T) {
	setup() // Global ctx ve dbPool kullanılıyor
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
	})
	clear() // Global ctx ve dbPool kullanılıyor
}
