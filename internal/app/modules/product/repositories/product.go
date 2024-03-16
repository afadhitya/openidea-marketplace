package productrepositories

import (
	"context"
	"log"

	productentities "github.com/widcha/openidea-marketplace/internal/app/modules/product"
)

type productRepo struct {
	repo *repo
}

func newProductRepo(repo *repo) IProduct {
	return &productRepo{
		repo: repo,
	}
}

func (p *productRepo) Create(ctx context.Context, product productentities.Product) error {
	query := `INSERT INTO products (id, name, price, image_url, stock, condition, tags, is_purchaseable, purchase_count, seller_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := p.repo.datasource.Postgre.ExecContext(ctx, query, product.Id, product.Name, product.Price, product.ImageURL, product.Stock, product.Condition, product.Tags, product.IsPurchasable, product.PurchaseCount, product.SellerID)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
