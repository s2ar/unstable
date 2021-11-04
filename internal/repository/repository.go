package repository

type DBTeam struct {
	tableName struct{} `pg:"teams"`
	TeamID    int32    `pg:"id,notnull,pk"`
	Rating    float32  `pg:"rating,notnull"`
	Name      string   `pg:"name"`
}

type Opendota interface {
	UpdateTeam(team *DBTeam) error
	GetTopTeam() (*DBTeam, error)
}

type Repository struct {
	Opendota
}

func New(connection *DB) *Repository {
	return &Repository{
		Opendota: NewOpendotaPG(connection),
	}
}
