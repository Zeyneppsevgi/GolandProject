package service

import (
	"os"
	"product-app/service"
	"testing"
)

var "product-app/service"

var productService  service.IProductService

func TestMain(m *testing.M){
	productService = service.NewProductService()
	exitCode := m.Run()
	os.Exit(exitCode)
}