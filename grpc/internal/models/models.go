package models

import "github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/date"

type Champion struct {
	ID                     string        `db:"id" json:"id"`
	TitleHolder            string        `db:"title_holder" json:"titleHolder"`
	TitleHolderNumber      int32         `db:"title_holder_number" json:"titleHolderNumber"`
	DateWon                date.NullTime `db:"date_won" json:"dateWon"`
	DateLost               date.NullTime `db:"date_lost" json:"dateLost"`
	Show                   string        `db:"show" json:"show"`
	CurrentChampion        bool          `db:"current_champion" json:"currentChampion"`
	TitleHolderOrderNumber int32         `db:"title_holder_order_number" json:"titleHolderOrderNumber"`
}

type Champions []*Champion

type ChampionList struct {
	Champions Champions
}

type CreateChampionInput struct {
	TitleHolder string
	DateWon     date.NullTime
	Show        string
}
