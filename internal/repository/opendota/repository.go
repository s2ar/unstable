package opendota

import (
	"context"
	"fmt"

	"github.com/s2ar/unstable/internal/repository"
)

type DBTeam struct {
	tableName struct{} `pg:"teams"`
	TeamID    int32    `pg:"id,notnull,pk"`
	Rating    float32  `pg:"rating,notnull"`
	Name      string   `pg:"name"`
}

type repositoryImpl struct {
	connection *repository.DB
}

type Repository interface {
	UpdateTeam(team *DBTeam) error
	GetTopTeam() (*DBTeam, error)
}

func (r *repositoryImpl) GetTopTeam() (*DBTeam, error) {
	team := &DBTeam{}
	_ = team.tableName
	err := r.connection.Model(team).OrderExpr("rating desc").Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (r *repositoryImpl) UpdateTeam(teams *DBTeam) error {
	ctx := context.Background()
	_, err := r.connection.
		WithContext(ctx).
		Model(teams).
		OnConflict("(id) DO UPDATE").
		Insert()
	if err != nil {
		return fmt.Errorf("SQL insert failed: %w", err)
	}
	return nil
}

func New(connection *repository.DB) Repository {
	repo := &repositoryImpl{
		connection: connection,
	}
	return repo
}
