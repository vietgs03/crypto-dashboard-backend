package repositories

import (
	"crypto-dashboard/gw-example/domain/entities"
	"crypto-dashboard/pkg/common"

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
