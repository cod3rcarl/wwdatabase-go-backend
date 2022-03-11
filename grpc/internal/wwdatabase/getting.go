package client

import (
	"context"

	"github.com/pkg/errors"

	wwErrors "github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/errors"
	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/models"
	pb "github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/wwdatabase"
)

func (s *Service) GetAllChampions(ctx context.Context, req *pb.GetChampionsRequest) (*pb.ChampionsList, error) {
	champions, err := s.store.GetAllChampions(ctx)
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
