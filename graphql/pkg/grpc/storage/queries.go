package storage

import (
	"context"

	wwErrors "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/errors"
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *Service) GetAllChampions(ctx context.Context) (models.Champions, error) {
	champions := models.Champions{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		"date_won",
		"date_lost",
		"show",
		COALESCE("current_champion", false) "current_champion"
	FROM "champion"
	ORDER BY "date_lost" ASC
	;`

	if err := pgxscan.Select(ctx, s.Pool, &champions, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Champions{}, wwErrors.ErrNoChampionsReturned
		}

		return models.Champions{}, errors.Errorf("error in GetAllChampions(): %v", err)
	}

	return champions, nil
}

func (s *Service) GetPreviousChampion(ctx context.Context) (models.Champion, error) {
	champion := models.Champion{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		"date_won",
		"date_lost",
		"show",
		COALESCE("current_champion", false) "current_champion",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number"
	FROM "champion"
	WHERE  "current_champion" = true
	;`

	if err := pgxscan.Get(ctx, s.Pool, &champion, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Champion{}, wwErrors.ErrNoChampionsReturned
		}

		return models.Champion{}, errors.Errorf("error in GetPrevChamp(): %v", err)
	}

	return champion, nil
}

func (s *Service) GetChampionListByName(ctx context.Context, name string) (models.Champions, error) {
	champions := models.Champions{}
	query := `
	SELECT
		"id",
		"title_holder",
		COALESCE("title_holder_number", 0) "title_holder_number",
		"date_won",
		"date_lost",
		"show",
		COALESCE("current_champion", false) "current_champion",
		COALESCE("title_holder_order_number", 0) "title_holder_order_number"
	FROM "champion"
	WHERE  "title_holder" = $1
	;`

	if err := pgxscan.Select(ctx, s.Pool, &champions, query, name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Champions{}, wwErrors.ErrNoChampionsReturned
		}

		return models.Champions{}, errors.Errorf("error in GetChampionListByName(): %v", err)
	}

	return champions, nil
}
