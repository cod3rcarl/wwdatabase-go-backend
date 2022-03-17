package storage

import (
	"context"

	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (s *Service) AddChampion(ctx context.Context, input models.CreateChampionInput) (models.Champion, error) {
	var createdChamp models.Champion
	id := uuid.New()
	query := `
	INSERT INTO "champion" (
	"id",
	"title_holder",
	"date_won",
	"show",
	"current_champion"
)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id,title_holder, COALESCE(title_holder_number, 0) "title_holder_number", date_won, date_lost, show, current_champion`

	if err := pgxscan.Get(
		ctx,
		s.Pool,
		&createdChamp,
		query,
		id.String(),
		input.TitleHolder,
		input.DateWon,
		input.Show,
		true,
	); err != nil {
		s.logger.Info(err.Error())
		return models.Champion{}, errors.Wrap(err, "error inserting champion")
	}

	return createdChamp, nil
}

func (s *Service) DeleteChampion(ctx context.Context, id string) (string, error) {
	query := `
			DELETE FROM "champion"
			WHERE id = $1
			RETURNING id
		`

	_, err := s.Pool.Exec(ctx, query, id)
	if err != nil {
		return "", errors.Errorf("error deleting champion: %v", err)
	}

	return id, nil
}
