package repositories

import (
	"apiProducts/src/pb/products"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

type ProductRepository struct{}

const dataFile string = "products.txt"

func (p *ProductRepository) LoadData() (*products.ProductList, error) {
	var productList = products.ProductList{}
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file: %v", err)
	}
	err = proto.Unmarshal(data, &productList)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %v", err)
	}
	return &productList, nil
}

func (p *ProductRepository) SaveData(products *products.ProductList) error {
	data, err := proto.Marshal(products)
	if err != nil {
		return fmt.Errorf("failed to marshal product: %v", err)
	}
	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write data file: %v", err)
	}
	return nil
}

func (p *ProductRepository) FindAll() (*products.ProductList, error) {
	return p.LoadData()
}

func (p *ProductRepository) Create(product *products.Product) error {
	productsList, err := p.LoadData()
	if err != nil {
		return err
	}
	product.Id = fmt.Sprintf("%d", len(productsList.Products)+1)
	productsList.Products = append(productsList.Products, product)
	return p.SaveData(productsList)
}
