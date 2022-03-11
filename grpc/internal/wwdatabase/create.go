package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/date"
	wwErrors "github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/errors"
	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/models"
	pb "github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/wwdatabase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) AddChampion(ctx context.Context, req *pb.NewChampionData) (*pb.CreateChampionPayload, error) {
	fmt.Println(req)

	prevChamp, err := s.store.GetPreviousChampion(ctx)
	if err != nil {
		return nil, errors.New("Error inserting champion")
	}
	fmt.Println(prevChamp)
	champion, err := s.store.AddChampion(ctx, models.CreateChampionInput{
		TitleHolder: req.TitleHolder,
		Show:        req.Show,
		DateWon:     date.TimestampToNullTime(req.DateWon),
	})
	if err != nil {
		if errors.Is(err, wwErrors.ErrNoChampionsReturned) {
			return nil, wwErrors.ErrNoChampionsReturned
		}
		return nil, errors.New("Error inserting champion")
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
		}, Success: true,
	}, nil
}
