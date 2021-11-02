//go:generate mockgen -destination=./mocks.go -source=./service.go -package=opendota

package opendota

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/s2ar/unstable/config"
	RepositoryOpendota "github.com/s2ar/unstable/internal/repository/opendota"
)

const timeoutSec = 10
const methodTeams = "teams"

type Team struct {
	TeamID int32   `json:"team_id"`
	Rating float32 `json:"rating"`
	Name   string  `json:"name"`
}

type serviceImpl struct {
	Repository RepositoryOpendota.Repository
	Config     *config.Configuration
}

type Opendota interface {
	AddTeams() error
	GetTopTeam() (*Team, error)
}

func (c *serviceImpl) AddTeams() error {
	var myClient = &http.Client{Timeout: timeoutSec * time.Second}
	r, err := myClient.Get(c.Config.DataService.URL + methodTeams + "/?api_key=" + c.Config.DataService.APIKey)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	teams := []*Team{}
	err = json.NewDecoder(r.Body).Decode(&teams)
	if err != nil {
		return err
	}
	for _, t := range teams {
		// save to db
		err = c.Repository.UpdateTeam(&RepositoryOpendota.DBTeam{
			TeamID: t.TeamID,
			Rating: t.Rating,
			Name:   t.Name,
		})

		if err != nil {
			return errors.Wrap(err, "cannot add teams [repository]")
		}
	}

	log.Printf("upload %d teams", len(teams))
	return nil
}

func (c *serviceImpl) GetTopTeam() (*Team, error) {
	t, err := c.Repository.GetTopTeam()
	if err != nil {
		return nil, err
	}

	return &Team{
		TeamID: t.TeamID,
		Rating: t.Rating,
		Name:   t.Name,
	}, nil

}

func New(repository RepositoryOpendota.Repository, cfg *config.Configuration) Opendota {
	return &serviceImpl{
		Repository: repository,
		Config:     cfg,
	}
}
