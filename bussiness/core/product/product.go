package product

import (
	"context"
	"fmt"
	"time"

	"github.com/ZweWT/backend-go/bussiness/data/store/product"
	"github.com/ZweWT/backend-go/bussiness/sys/auth"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Core manages the set of API's for product access.
type Core struct {
	log     *zap.SugaredLogger
	product product.Store
}

// NewCore constructs a core for product api access.
func NewCore(log *zap.SugaredLogger, db *sqlx.DB) Core {
	return Core{
		log:     log,
		product: product.NewStore(log, db),
	}
}

// Create inserts a new product into the database.
func (c Core) Create(ctx context.Context, claims auth.Claims, np product.NewProduct, now time.Time) (product.Product, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	product, err := c.product.Create(ctx, claims, np, now)
	if err != nil {
		return product, fmt.Errorf("create: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return product, nil
}

// Update replaces a product document in the database.
func (c Core) Update(ctx context.Context, claims auth.Claims, productID string, up product.UpdateProduct, now time.Time) error {

	// PERFORM PRE BUSINESS OPERATIONS

	if err := c.product.Update(ctx, claims, productID, up, now); err != nil {
		return fmt.Errorf("update: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return nil
}

// Delete removes a product from the database.
func (c Core) Delete(ctx context.Context, claims auth.Claims, productID string) error {

	// PERFORM PRE BUSINESS OPERATIONS

	if err := c.product.Delete(ctx, claims, productID); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return nil
}

// Query retrieves a list of existing products from the database.
func (c Core) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]product.Product, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	products, err := c.product.Query(ctx, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return products, nil
}

// QueryByID gets the specified product from the database.
func (c Core) QueryByID(ctx context.Context, claims auth.Claims, productID string) (product.Product, error) {

	// PERFORM PRE BUSINESS OPERATIONS

	product, err := c.product.QueryByID(ctx, claims, productID)
	if err != nil {
		return product, fmt.Errorf("query: %w", err)
	}

	// PERFORM POST BUSINESS OPERATIONS

	return product, nil
}
