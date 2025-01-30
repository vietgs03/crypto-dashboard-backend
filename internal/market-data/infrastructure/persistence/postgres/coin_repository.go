package postgres

import (
	"context"
	"crypto-dashboard-backend/internal/market-data/domain/entity"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type coinRepository struct {
	db *pgxpool.Pool
}

func NewCoinRepository(db *pgxpool.Pool) *coinRepository {
	return &coinRepository{db: db}
}

func (r *coinRepository) SaveCoins(ctx context.Context, coins []*entity.Coin) error {
	query := `
        INSERT INTO coins (
            id, symbol, name, price, market_cap, volume_24h, price_change,
            circulating_supply, total_supply, ath, ath_date, updated_at
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
        )
        ON CONFLICT (id) DO UPDATE SET
            symbol = EXCLUDED.symbol,
            name = EXCLUDED.name,
            price = EXCLUDED.price,
            market_cap = EXCLUDED.market_cap,
            volume_24h = EXCLUDED.volume_24h,
            price_change = EXCLUDED.price_change,
            circulating_supply = EXCLUDED.circulating_supply,
            total_supply = EXCLUDED.total_supply,
            ath = EXCLUDED.ath,
            ath_date = EXCLUDED.ath_date,
            updated_at = EXCLUDED.updated_at
    `

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	for _, coin := range coins {
		_, err = tx.Exec(ctx, query,
			coin.ID, coin.Symbol, coin.Name, coin.Price,
			coin.MarketCap, coin.Volume24h, coin.PriceChange,
			coin.CirculatingSupply, coin.TotalSupply,
			coin.ATH, coin.ATHDate, coin.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (r *coinRepository) GetCoins(ctx context.Context) ([]*entity.Coin, error) {
	query := `
        SELECT id, symbol, name, price, market_cap, volume_24h, price_change, updated_at 
        FROM coins 
        ORDER BY market_cap DESC
    `

	rows, err := r.db.Query(ctx, query)
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

	return coins, nil
}
