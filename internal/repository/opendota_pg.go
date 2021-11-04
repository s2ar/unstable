package repository

import (
	"context"
	"fmt"
)

type OpendotaPG struct {
	connection *DB
}

func (r *OpendotaPG) GetTopTeam() (*DBTeam, error) {
	team := &DBTeam{}
	_ = team.tableName
	err := r.connection.Model(team).OrderExpr("rating desc").Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (r *OpendotaPG) UpdateTeam(teams *DBTeam) error {
	ctx := context.Background()
	_, err := r.connection.
		WithContext(ctx).
		Model(teams).
		OnConflict("(id) DO UPDATE").
		Insert()
	if err != nil {
		return fmt.Errorf("SQL insert failed: %e", err)
	}
	return nil
}

func NewOpendotaPG(connection *DB) *OpendotaPG {
	repo := &OpendotaPG{
		connection: connection,
	}
	return repo
}
