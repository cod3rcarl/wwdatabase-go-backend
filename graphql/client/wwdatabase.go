package client

import (
	"context"
	"time"

	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/internal/server/graph/model"
	pb "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/wwdatabase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Client) CreateChampion(ctx context.Context, input *model.CreateChampionInput) (*model.CreateChampionPayload, error) {
	dateWon, err := time.Parse("2006-01-02 15:04", *input.DateWon)
	if err != nil {
		return nil, ErrDateIncorrect
	}
	pbChamp, err := s.wwdatabaseGRPCClient.AddChampion(ctx, &pb.NewChampionData{
		TitleHolder: input.TitleHolder,
		DateWon:     &timestamppb.Timestamp{Seconds: dateWon.Unix()},
		Show:        input.Show,
	})
	if err != nil {
		return nil, s.handleErr(err)
	}

	return &model.CreateChampionPayload{
		Success:  pbChamp.Success,
		Champion: pbChampionToModel(pbChamp.Champion),
	}, nil
}

func (s *Client) DeleteChampion(ctx context.Context, input *model.DeleteChampionInput) (*model.DeleteChampionPayload, error) {
	pbChamp, err := s.wwdatabaseGRPCClient.DeleteChampion(ctx, &pb.DeleteChampionRequest{
		Id: input.ID,
	})
	if err != nil {
		return nil, s.handleErr(err)
	}

	return &model.DeleteChampionPayload{
		Success: true,
		ID:      pbChamp.Id,
	}, nil
}

func (s *Client) GetChampions(ctx context.Context, filter *string) (
	*model.ChampionPayload, error,
) {
	if filter != nil {
		pbChamps, err := s.wwdatabaseGRPCClient.GetChampionByName(ctx, &pb.GetChampionByNameRequest{
			Name: *filter,
		})
		if err != nil {
			return nil, s.handleErr(err)
		}

		return pbChampionsToModel(pbChamps), nil
	}
	pbChamps, err := s.wwdatabaseGRPCClient.GetChampions(ctx, &pb.GetChampionsRequest{})
	if err != nil {
		return nil, s.handleErr(err)
	}

	return pbChampionsToModel(pbChamps), nil
}

func (s *Client) GetChampionDetailsByName(ctx context.Context, name string) (
	*model.ChampionPayload, error,
) {
	pbChamps, err := s.wwdatabaseGRPCClient.GetChampionByName(ctx, &pb.GetChampionByNameRequest{
		Name: name,
	})
	if err != nil {
		return nil, s.handleErr(err)
	}

	return pbChampionsToModel(pbChamps), nil
}

func (s *Client) GetChampionByOrderNumber(ctx context.Context, tn int32) (
	*model.Champion, error,
) {
	pbChamp, err := s.wwdatabaseGRPCClient.GetChampionByOrderNumber(ctx, &pb.ChampionNumber{
		TitleHolderOrderNumber: tn,
	})
	if err != nil {
		return nil, s.handleErr(err)
	}

	return pbChampionToModel(pbChamp.Champion), nil
}

func pbChampionsToModel(pbC *pb.ChampionsList) *model.ChampionPayload {
	champions := []*model.Champion{}

	for _, champion := range pbC.Champions {
		titlehn := int(champion.TitleHolderNumber)
		titlehnon := int(champion.TitleHolderOrderNumber)
		dw := champion.DateWon.String()
		dl := champion.DateLost.String()
		c := &model.Champion{
			TitleHolder:            champion.TitleHolder,
			TitleHolderNumber:      &titlehn,
			TitleHolderOrderNumber: &titlehnon,
			DateWon:                dw,
			DateLost:               &dl,
			Show:                   champion.Show,
			CurrentChampion:        &champion.CurrentChampion,
		}
		champions = append(champions, c)
	}
	count := len(champions)

	return &model.ChampionPayload{
		Champions:  champions,
		TotalCount: &count,
	}
}

func pbChampionToModel(pbC *pb.Champion) *model.Champion {
	titlehn := int(pbC.TitleHolderNumber)
	titlehnon := int(pbC.TitleHolderOrderNumber)
	dw := pbC.DateWon.String()
	dl := pbC.DateLost.String()

	return &model.Champion{
		TitleHolder:            pbC.TitleHolder,
		TitleHolderNumber:      &titlehn,
		TitleHolderOrderNumber: &titlehnon,
		DateWon:                dw,
		DateLost:               &dl,
		Show:                   pbC.Show,
		CurrentChampion:        &pbC.CurrentChampion,
		PreviousChampion:       &pbC.PreviousChampion,
	}
}
