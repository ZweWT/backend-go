package product

import (
	"context"
	"fmt"
	"time"

	"github.com/ZweWT/backend-go/bussiness/sys/auth"
	"github.com/ZweWT/backend-go/bussiness/sys/database"
	"github.com/ZweWT/backend-go/bussiness/sys/validate"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of API's for user access.
type Store struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

// NewStore constructs a user store for api access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) Store {
	return Store{
		log: log,
		db:  db,
	}
}

// Create inserts a new product into the database.
func (s Store) Create(ctx context.Context, claims auth.Claims, np NewProduct, now time.Time) (Product, error) {
	if err := validate.Check(np); err != nil {
		return Product{}, fmt.Errorf("validating data: %w", err)
	}

	if !claims.Authorized(auth.RoleAdmin) {
		return Product{}, database.ErrForbidden
	}

	product := Product{
		ID:          validate.GenerateID(),
		Name:        np.Name,
		CategoryID:  np.CategoryID,
		Cost:        np.Cost,
		DateCreated: now,
		DateUpdated: now,
	}

	const q = `
	INSERT INTO products
		(product_id, name, category_id, cost, date_created, date_updated)
	VALUES
		(:product_id, :name, :category_id, :cost, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, product); err != nil {
		return Product{}, fmt.Errorf("inserting user: %w", err)
	}

	return product, nil
}

// Update replaces a product document in the database.
func (s Store) Update(ctx context.Context, claims auth.Claims, productID string, up UpdateProduct, now time.Time) error {
	if err := validate.CheckID(productID); err != nil {
		return database.ErrInvalidID
	}

	if !claims.Authorized(auth.RoleAdmin) {
		return database.ErrForbidden
	}

	if err := validate.Check(up); err != nil {
		return fmt.Errorf("validating data: %w", err)
	}

	product, err := s.QueryByID(ctx, claims, productID)
	if err != nil {
		return fmt.Errorf("updating user userID[%s]: %w", productID, err)
	}

	if up.Name != nil {
		product.Name = *up.Name
	}
	if up.CategoryID != nil {
		product.CategoryID = *up.CategoryID
	}
	if up.Cost != nil {
		product.Cost = *up.Cost
	}
	product.DateUpdated = now

	const q = `
	UPDATE
		users
	SET 
		"name" = :name,
		"email" = :email,
		"roles" = :roles,
		"password_hash" = :password_hash,
		"date_updated" = :date_updated
	WHERE
		user_id = :user_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, product); err != nil {
		return fmt.Errorf("updating userID[%s]: %w", productID, err)
	}

	return nil
}

// Delete removes a product from the database.
func (s Store) Delete(ctx context.Context, claims auth.Claims, productID string) error {
	if err := validate.CheckID(productID); err != nil {
		return database.ErrInvalidID
	}

	if !claims.Authorized(auth.RoleAdmin) {
		return database.ErrForbidden
	}

	data := struct {
		productID string `db:"product_id"`
	}{
		productID: productID,
	}

	const q = `
	DELETE FROM
		products
	WHERE
		product_id = :product_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("deleting productID[%s]: %w", productID, err)
	}

	return nil
}

// QueryByID gets the specified product from the database.
func (s Store) QueryByID(ctx context.Context, claims auth.Claims, productID string) (Product, error) {
	if err := validate.CheckID(productID); err != nil {
		return Product{}, database.ErrInvalidID
	}

	data := struct {
		productID string `db:"product_id"`
	}{
		productID: productID,
	}

	const q = `
	SELECT
		*
	FROM
		products
	WHERE 
		product_id = :product_id`

	var product Product
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &product); err != nil {
		if err == database.ErrNotFound {
			return Product{}, database.ErrNotFound
		}
		return Product{}, fmt.Errorf("selecting productID[%q]: %w", productID, err)
	}

	return product, nil
}

// Query retrieves a list of existing products from the database.
func (s Store) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]Product, error) {
	data := struct {
		Offset      int `db:"offset"`
		RowsPerPage int `db:"rows_per_page"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}

	const q = `
	SELECT
		*
	FROM
		products
	ORDER BY
		product_id
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var products []Product
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &products); err != nil {
		if err == database.ErrNotFound {
			return nil, database.ErrNotFound
		}
		return nil, fmt.Errorf("selecting users: %w", err)
	}

	return products, nil
}
