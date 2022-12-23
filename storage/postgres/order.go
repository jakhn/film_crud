package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/jakhn/film_crud/models"
	"github.com/jakhn/film_crud/pkg/helper"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (f *orderRepo) Create(ctx context.Context, order *models.CreateOrder) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO orders(
			order_id,
			books_id, 
			users_id, 
			created_at,
			updated_at
		) VALUES ( $1, $2 , $3,now(), now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		order.BooksId,
		order.UsersId, 
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *orderRepo) GetByPKey(ctx context.Context, pkey *models.OrderPrimarKey) (*models.Order, error) {

	var (
		id        		sql.NullString
		booksId 			sql.NullString
		usersId 			sql.NullString 
		createdAt 		sql.NullString
		updatedAt 		sql.NullString
	)

	query := `
		SELECT
			order_id,
			books_id,
			users_id, 
			created_at,
			updated_at
		FROM
			orders
		WHERE order_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&booksId,
			&usersId, 
			&createdAt,
			&updatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &models.Order{
		Id: 		 id.String,
		BooksId:  	 booksId.String,
		UsersId: 	 usersId.String, 
		CreatedAt: 	 createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (f *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {

	var (
		resp   = models.GetListOrderResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			order_id,
			users.first_name || ' ' || users.last_name as fullname,
			books.price as price
			created_at,
			updated_at
		FROM
			orders
		JOIN users ON orders.users_id = users.user_id
		JOIN book ON book.book_id = orders.book_id
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id        			sql.NullString
			fullName 			sql.NullString
			price 				sql.NullInt64
			createdAt 			sql.NullString
			updatedAt 			sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id, 
			&fullName,
			&price,  
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Orders = append(resp.Orders, &models.OrderBook{
		Id: 		 		id.String,
		FullName:	     	fullName.String,
		Price: 		 		price.Int64, 
		CreatedAt: 	 		createdAt.String,
		UpdatedAt:   		updatedAt.String,
		})

	}

	return &resp, err
}

func (f *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			orders
		SET
			books_id = :books_id,
			users_id = :users_id, 
			updated_at = now()
		WHERE order_id = :order_id
	`

	params = map[string]interface{}{
		"order_id": 	req.Id,
		"books_id": 		req.BooksId,
		"users_id":  	req.UsersId, 
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *orderRepo) Delete(ctx context.Context, req *models.OrderPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM orders WHERE order_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
