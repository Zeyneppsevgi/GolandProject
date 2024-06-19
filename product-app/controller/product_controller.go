package controller

import "product-app/service"

type ProductController struct {
	productService service.IProductService
}
