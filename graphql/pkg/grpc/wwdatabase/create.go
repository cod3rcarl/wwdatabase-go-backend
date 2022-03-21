package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/date"
	wwErrors "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/errors"
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/models"
	pb "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/wwdatabase"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) AddChampion(ctx context.Context, req *pb.NewChampionData) (*pb.CreateChampionPayload, error) {
	champion, err := s.store.AddChampion(ctx, models.CreateChampionInput{
		TitleHolder: req.TitleHolder,
		Show:        req.Show,
		DateWon:     date.TimestampToNullTime(req.DateWon),
	})
	if err != nil {
		if errors.Is(err, wwErrors.ErrNoChampionsReturned) {
			return nil, wwErrors.ErrNoChampionsReturned
		}

		return nil, errors.Wrap(err, "Error inserting champion")
	}

	prevChamp, err := s.store.GetChampionByOrderNumber(ctx, champion.TitleHolderOrderNumber-1)
	if err != nil {
		return nil, errors.Wrap(err, "Error inserting champion")
	}

	return &pb.CreateChampionPayload{
		Champion: &pb.Champion{
			Id:                champion.ID,
			TitleHolder:       champion.TitleHolder,
			TitleHolderNumber: champion.TitleHolderNumber,
			Show:              champion.Show,
			DateWon:           &timestamppb.Timestamp{Seconds: champion.DateWon.Unix()},
			DateLost:          &timestamppb.Timestamp{Seconds: champion.DateLost.Unix()},
			CurrentChampion:   champion.CurrentChampion,
			PreviousChampion:  prevChamp.TitleHolder,
		}, Success: true,
	}, nil
}
