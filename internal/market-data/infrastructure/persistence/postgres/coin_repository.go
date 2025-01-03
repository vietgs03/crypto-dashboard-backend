package postgres

import (
	"context"
	"crypto-dashboard-backend/internal/market-data/domain/entity"
	"database/sql"
	"fmt"
)

type coinRepository struct {
	db *sql.DB
}

func NewCoinRepository(db *sql.DB) *coinRepository {
	return &coinRepository{db: db}
}

func (r *coinRepository) SaveCoins(ctx context.Context, coins []*entity.Coin) error {
	query := `
		INSERT INTO coins (id, symbol, name, price, market_cap, volume_24h, price_change, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (id) DO UPDATE SET
			symbol = EXCLUDED.symbol,
			name = EXCLUDED.name,
			price = EXCLUDED.price,
			market_cap = EXCLUDED.market_cap,
			volume_24h = EXCLUDED.volume_24h,
			price_change = EXCLUDED.price_change,
			updated_at = EXCLUDED.updated_at
	`
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, coin := range coins {
		_, err = stmt.ExecContext(
			ctx,
			coin.ID,
			coin.Symbol,
			coin.Name,
			coin.Price,
			coin.MarketCap,
			coin.Volume24h,
			coin.PriceChange,
			coin.UpdatedAt,
		)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *coinRepository) GetCoins(ctx context.Context) ([]*entity.Coin, error) {
	query := `
        SELECT id, symbol, name, price, market_cap, volume_24h, price_change, updated_at 
        FROM coins 
        ORDER BY market_cap DESC
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query coins: %w", err)
	}
	defer rows.Close()

	var coins []*entity.Coin
	for rows.Next() {
		coin := &entity.Coin{}
		err := rows.Scan(
			&coin.ID,
			&coin.Symbol,
			&coin.Name,
			&coin.Price,
			&coin.MarketCap,
			&coin.Volume24h,
			&coin.PriceChange,
			&coin.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan coin row: %w", err)
		}
		coins = append(coins, coin)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return coins, nil
}
