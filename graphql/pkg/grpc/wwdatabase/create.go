package client

import (
	"context"

	"github.com/pkg/errors"

	wwErrors "github.com/cod3rcarl/wwdatabase-go-backend/graphql/client"
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/date"
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/models"
	pb "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/wwdatabase"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) AddChampion(ctx context.Context, req *pb.NewChampionData) (*pb.CreateChampionPayload, error) {
	prevChamp, err := s.store.GetCurrentChampion(ctx, true)
	if err != nil {
		return nil, errors.Wrap(err, "Error inserting champion")
	}

	history, err := s.store.GetChampionListByName(ctx, req.TitleHolder)

	champion, err := s.store.AddChampion(ctx, models.CreateChampionInput{
		TitleHolder:            req.TitleHolder,
		Show:                   req.Show,
		DateWon:                date.TimestampToNullTime(req.DateWon),
		TitleHolderNumber:      prevChamp.TitleHolderNumber + 1,
		TitleHolderOrderNumber: prevChamp.TitleHolderOrderNumber + 1,
		WrestlerID:             history[0].WrestlerID,
	})
	if err != nil {
		if errors.Is(err, wwErrors.ErrNoChampionsReturned) {
			return nil, wwErrors.ErrNoChampionsReturned
		}

		return nil, errors.Wrap(err, "Error inserting champion")
	}

	prevChamp, err = s.store.UpdateChampion(ctx, models.UpdateChampionInput{
		ID:              prevChamp.ID,
		DateLost:        champion.DateWon,
		CurrentChampion: false,
	})

	if err != nil {
		return nil, errors.Wrap(err, "Error updating old champion")
	}

	return &pb.CreateChampionPayload{
		Champion: &pb.Champion{
			Id:                     champion.ID,
			TitleHolder:            champion.TitleHolder,
			TitleHolderNumber:      prevChamp.TitleHolderNumber + 1,
			Show:                   champion.Show,
			DateWon:                &timestamppb.Timestamp{Seconds: champion.DateWon.Unix()},
			DateLost:               &timestamppb.Timestamp{Seconds: champion.DateLost.Unix()},
			CurrentChampion:        champion.CurrentChampion,
			PreviousChampion:       prevChamp.TitleHolder,
			TitleHolderOrderNumber: prevChamp.TitleHolderOrderNumber + 1,
			WrestlerId:             champion.WrestlerID,
		}, Success: true,
	}, nil
}
