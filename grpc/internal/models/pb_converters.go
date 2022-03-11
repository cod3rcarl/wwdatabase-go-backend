package models

import (
	pb "github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/wwdatabase"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ModelToPBChampion(c *Champion) *pb.Champion {
	champion := &pb.Champion{
		Id:                c.ID,
		TitleHolder:       c.TitleHolder,
		TitleHolderNumber: c.TitleHolderNumber,
		DateWon:           &timestamppb.Timestamp{Seconds: c.DateWon.Unix()},
		DateLost:          &timestamppb.Timestamp{Seconds: c.DateLost.Unix()},
		Show:              c.Show,
		CurrentChampion:   c.CurrentChampion,
	}

	return champion
}

// func MakePBChampionsResponse(c *Champion) *pb.ChampionResponse {
// 	pbChampion := &pb.ChampionResponse{
// 		Champion: ModelToPBChampion(c),
// 	}

// 	return pbChampion
// }

// func PBMainContractorsToModel(pbMc *pb.MainContractorsResponse) MainContractors {
// 	maincontractors := make(MainContractors, len(pbMc.GetRecords()))
// 	for i, mc := range pbMc.GetRecords() {
// 		maincontractors[i] = &MainContractor{
// 			ID:           mc.Id,
// 			Name:         mc.Name,
// 			ShortName:    mc.ShortName,
// 			Phone:        mc.Phone,
// 			Country:      mc.Country,
// 			VatNumber:    mc.VatNumber,
// 			ContactName:  mc.ContactName,
// 			ContactEmail: mc.ContactEmail,
// 			ContactPhone: mc.ContactPhone,
// 			SalesforceID: mc.SalesforceId,
// 			AddressID:    mc.AddressId,
// 			IconID:       mc.IconId,
// 			LogoID:       mc.LogoId,
// 			SignatureID:  mc.SignatureId,
// 		}

// 		if mc.Deleted.IsValid() {
// 			maincontractors[i].Deleted = date.TimestampToNullTime(mc.Deleted)
// 		}

// 		if mc.CreatedAt.IsValid() {
// 			maincontractors[i].CreatedAt = date.TimestampToNullTime(mc.CreatedAt)
// 		}

// 		if mc.UpdatedAt.IsValid() {
// 			maincontractors[i].UpdatedAt = date.TimestampToNullTime(mc.UpdatedAt)
// 		}
// 	}

// 	return maincontractors
// }
