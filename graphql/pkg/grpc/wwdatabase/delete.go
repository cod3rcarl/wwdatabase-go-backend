package client

import (
	"context"
	"strings"

	pb "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/wwdatabase"
	"github.com/pkg/errors"
)

func (s *Service) DeleteChampion(ctx context.Context, id string) (
	*pb.DeleteChampionResponse, error,
) {
	champErr := s.store.GetChampionByID(ctx, id)

	if champErr != nil {
		if strings.Contains(champErr.Error(), "no rows") {
			return &pb.DeleteChampionResponse{}, errors.Errorf("no champion with id %v", id)
		}

		return &pb.DeleteChampionResponse{}, errors.New(champErr.Error())
	}
	i, err := s.store.DeleteChampion(ctx, id)
	if err != nil {
		return &pb.DeleteChampionResponse{}, errors.Wrap(err, "error deleting champion")
	}

	return &pb.DeleteChampionResponse{
		Id: i,
	}, nil
}
