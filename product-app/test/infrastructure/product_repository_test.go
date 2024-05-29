package infrastructure

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"product-app/common/postgresql"
	"product-app/persistence"
	"testing"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})
}
