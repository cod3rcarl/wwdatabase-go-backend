package client

import (
	"context"

	pb "github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/wwdatabase"
	"github.com/pkg/errors"
)

func (s *Service) DeleteChampion(ctx context.Context, id string) (
	*pb.DeleteChampionResponse, error) {
	i, err := s.store.DeleteChampion(ctx, id)
	if err != nil {
		return &pb.DeleteChampionResponse{}, errors.Wrap(err, "error mainContractor.GetMainContractors")
	}

	return &pb.DeleteChampionResponse{
		Id: i,
	}, nil
}
