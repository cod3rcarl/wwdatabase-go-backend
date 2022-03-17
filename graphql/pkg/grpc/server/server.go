package server

import (
	"context"

	"github.com/pkg/errors"

	pb "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/wwdatabase"
)

func (s *Service) GetChampions(ctx context.Context, req *pb.GetChampionsRequest) (
	*pb.ChampionsList, error,
) {
	s.logger.Info("gRPC call: GetAllChampions")
	pbMc, err := s.wwdatabase.GetAllChampions(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting champ")
	}

	return pbMc, nil
}

func (s *Service) GetChampionByName(ctx context.Context, req *pb.GetChampionByNameRequest) (
	*pb.ChampionsList, error,
) {
	s.logger.Info("gRPC call: GetChampionsByName")

	name := req.GetName()
	if name == "" {
		return nil, errors.New("must provide a name")
	}

	pbMc, err := s.wwdatabase.GetChampionListByName(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting champ")
	}

	return pbMc, nil
}

func (s *Service) AddChampion(ctx context.Context, req *pb.NewChampionData) (
	*pb.CreateChampionPayload, error,
) {
	s.logger.Info("gRPC call: AddChampion")
	pbMc, err := s.wwdatabase.AddChampion(ctx, req)
	if err != nil {
		return nil, errors.New("Cannot add champs")
	}

	return pbMc, nil
}

func (s *Service) DeleteChampion(ctx context.Context, req *pb.DeleteChampionRequest) (
	*pb.DeleteChampionResponse, error,
) {
	s.logger.Info("gRPC call: DeleteChampion")
	i, err := s.wwdatabase.DeleteChampion(ctx, req.Id)
	if err != nil {
		return nil, errors.New("Cannot delete champs")
	}

	return i, nil
}
