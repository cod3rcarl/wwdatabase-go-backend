package models

import (
	pb "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/wwdatabase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelToPBChampion(c *Champion) *pb.Champion {
	champion := &pb.Champion{
		Id:                     c.ID,
		TitleHolder:            c.TitleHolder,
		TitleHolderNumber:      c.TitleHolderNumber,
		TitleHolderOrderNumber: c.TitleHolderOrderNumber,
		DateWon:                &timestamppb.Timestamp{Seconds: c.DateWon.Unix()},
		DateLost:               &timestamppb.Timestamp{Seconds: c.DateLost.Unix()},
		Show:                   c.Show,
		CurrentChampion:        c.CurrentChampion,
		WrestlerId:             c.WrestlerID,
	}

	return champion
}
