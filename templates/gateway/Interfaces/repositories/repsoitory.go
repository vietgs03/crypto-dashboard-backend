package repositories

import (
	"crypto-dashboard-backend/pkg/common"
	"crypto-dashboard-backend/templates/gateway/domain/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	EntityRepository struct {
		common.Repository[entities.Entity]
	}
)

var entityPicRepo *EntityRepository

func ProvideEntityRepository(db *pgxpool.Pool) *EntityRepository {
	base := common.NewRepository(db, entities.Entity{})
	entityPicRepo = &EntityRepository{*base}
	return entityPicRepo
}
