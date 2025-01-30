package repositories

import (
	"sync"

	"crypto-dashboard-backend/pkg/common"
	"crypto-dashboard-backend/templates/domain/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	EntityRepository struct {
		common.Repository[entities.Entity]
	}
)

var (
	entityPicRepo *EntityRepository
	repoOnce      sync.Once
)

func ProvideEntityRepository(db *pgxpool.Pool) *EntityRepository {
	repoOnce.Do(func() {
		base := common.NewRepository[entities.Entity](db)
		entityPicRepo = &EntityRepository{*base}
	})

	return entityPicRepo
}

func test() {
}
