package client

import (
	"context"

	"github.com/pkg/errors"

	wwErrors "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/errors"
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/models"
	pb "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/wwdatabase"
)

func returnChampionList(champions models.Champions, err error) (*pb.ChampionsList, error) {
	if err != nil {
		if errors.Is(err, wwErrors.ErrNoChampionsReturned) {
			return nil, wwErrors.ErrNoChampionsReturned
		}

		return nil, errors.Wrap(err, "Error getting champions")
	}
	count := len(champions)

	if count == 0 {
		return nil, wwErrors.ErrNoChampionsReturned
	}
	pbChampions := make([]*pb.Champion, len(champions))
	for i := range champions {
		pbChampions[i] = models.ModelToPBChampion(champions[i])
	}

	return &pb.ChampionsList{TotalCount: int32(count), Champions: pbChampions}, nil
}

func (s *Service) GetChampionByOrderNumber(ctx context.Context, tn *pb.ChampionNumber) (*pb.ChampionResponse, error) {
	champion, err := s.store.GetChampionByOrderNumber(ctx, tn.TitleHolderOrderNumber)
	if err != nil {
		if errors.Is(err, wwErrors.ErrNoChampionsReturned) {
			return nil, wwErrors.ErrNoChampionsReturned
		}

		return nil, errors.Wrap(err, "Error getting champions")

	}

	return &pb.ChampionResponse{
		Champion: models.ModelToPBChampion(&champion),
	}, nil
}

func (s *Service) GetAllChampions(ctx context.Context, _ *pb.GetChampionsRequest) (*pb.ChampionsList, error) {
	champions, err := s.store.GetAllChampions(ctx)

	return returnChampionList(champions, err)
}

func (s *Service) GetChampionListByName(ctx context.Context, name string) (*pb.ChampionsList, error) {
	champions, err := s.store.GetChampionListByName(ctx, name)

	return returnChampionList(champions, err)
}
