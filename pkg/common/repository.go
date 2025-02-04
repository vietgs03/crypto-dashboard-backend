package common

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"crypto-dashboard-backend/pkg/response"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	PaginationQuery struct {
		CountSQL    string
		QuerySQL    string
		Params      []interface{}
		Page        int
		Take        int
		ItemMappers func(*sql.Rows) interface{}
	}

	Where struct {
		Condition string
		Params    []interface{}
	}
	FindOption struct {
		Where
		Page           int
		Take           int
		Order          []string
		Select         []string
		ExcludeDeleted bool
		Joins          []string
	}

	BaseEnityIf interface {
		TableName() string
	}

	RepositoryQueryIf[T BaseEnityIf] interface {
		BaseEnityIf
		FindByID(ctx context.Context, id string) (*T, *response.AppError)
		FindOne(ctx context.Context, option *FindOption) (*T, *response.AppError)
		// FindMapStruct[E comparable](ctx context.Context, option *FindOption) ([]E, *response.AppError)
		FindMany(ctx context.Context, option *FindOption) ([]T, *response.AppError)
		PaginationQuery(ctx context.Context, option *FindOption) ([]T, int, *response.AppError)
	}

	RepositoryCommandIf[T BaseEnityIf] interface {
		BaseEnityIf
		CreateOne(ctx context.Context, entity *T) (*T, *response.AppError)
		CreateMany(ctx context.Context, entities ...*T) []*T
		CreateWithOnConflicting(ctx context.Context, entities ...*T) []T
		Update(ctx context.Context, entity *T) (int *response.AppError)
		UpdateReturning(ctx context.Context, entity *T) ([]T, *response.AppError)
		UpdateOne(ctx context.Context, entity *T) *response.AppError
		UpdateOneReturning(ctx context.Context, entity *T) (*T, *response.AppError)
		UpdateMany(ctx context.Context, entity *T) *response.AppError
		UpdateById(ctx context.Context, id string) *response.AppError
		UpdateByMap(ctx context.Context, values *map[string]interface{}) *response.AppError
		DeleteByID(ctx context.Context, id string) *response.AppError
		DeleteMany(ctx context.Context, id string) *response.AppError
		CountBy(ctx context.Context, cond *FindOption) int
	}

	Repository[T BaseEnityIf] struct {
		BaseEnityIf
		Fields           map[string]bool
		FieldInsertCount int
		db               *pgxpool.Pool
	}
)

func NewRepository[T BaseEnityIf](db *pgxpool.Pool, entity T) *Repository[T] {
	t := reflect.TypeOf(entity)
	fields := make(map[string]bool, t.NumField())
	fieldInsertCount := 0
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("db") == "-" {
			continue
		}
		fieldInsertCount++
		fields[field.Name] = true
	}

	return &Repository[T]{
		db:               db,
		Fields:           fields,
		FieldInsertCount: fieldInsertCount,
	}
}

func (r *Repository[T]) DB() *pgxpool.Pool {
	return r.db
}

func (r *Repository[T]) CreateOne(ctx context.Context, entity *T) (*T, *response.AppError) {
	tableName := r.BaseEnityIf.TableName()
	t := reflect.TypeOf(entity)
	v := reflect.ValueOf(entity)

	columns := make([]string, 0, r.FieldInsertCount)
	values := make([]interface{}, 0, r.FieldInsertCount)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("db") == "-" {
			continue
		}

		columns = append(columns, field.Tag.Get("db"))
		values = append(values, v.Field(i).Interface())
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING *",
		tableName, strings.Join(columns, ","), strings.Repeat("?, ", r.FieldInsertCount-1)+"?")

	row := r.db.QueryRow(ctx, query, values...)

	var res T
	err := row.Scan(&res)
	if err != nil {
		return nil, response.DatabaseError(err)
	}

	return &res, nil
}

func (r *Repository[T]) CreateMany(ctx context.Context, entities ...*T) ([]T, *response.AppError) {
	if len(entities) == 0 {
		return nil, response.DatabaseError(fmt.Errorf("entities is empty"))
	}
	tableName := r.BaseEnityIf.TableName()
	t := reflect.TypeOf(entities[0])
	numberRecords := len(entities) * r.FieldInsertCount

	columns := make([]string, 0, r.FieldInsertCount)
	values := make([]interface{}, 0, numberRecords)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("db") == "-" {
			continue
		}
		columns = append(columns, field.Tag.Get("db"))
	}

	var valueStrings strings.Builder

	for _, entity := range entities {
		v := reflect.ValueOf(entity)
		value := make([]interface{}, 0, r.FieldInsertCount)
		valueStrings.WriteString("(")
		for i := 0; i < t.NumField(); i++ {
			value = append(value, v.Field(i).Interface())
		}
		valueStrings.WriteString("),")
		values = append(values, value)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s RETURNING *",
		tableName, strings.Join(columns, ","), valueStrings.String())

	rows, err := r.db.Query(ctx, query, values...)
	if err != nil {
		return nil, response.DatabaseError(err)
	}

	res := make([]T, 0, numberRecords)

	for {
		var entity T
		err := rows.Scan(&entity)
		if err == sql.ErrNoRows {
			break
		} else if err != nil {
			return nil, response.DatabaseError(err)
		}

		res = append(res, entity)
	}

	return res, nil
}
