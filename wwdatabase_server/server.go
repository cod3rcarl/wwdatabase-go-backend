package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"

	pb "github.com/cod3rcarl/wwdatabase-go-backend/wwdatabase"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func NewWwdatabaseServer() *WwdatabaseServer {
return &WwdatabaseServer{}
}

type WwdatabaseServer struct {
	conn *pgx.Conn
	db *pgxpool.Pool
	pb.UnimplementedWwdatabaseServer
}

func (server *WwdatabaseServer) Run() error {
	log.Println("Starting server..")
	port:= goDotEnvVariable("PORT")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to listen on port %v: %v", port, err)
	}

	s := grpc.NewServer()
	pb.RegisterWwdatabaseServer(s, server)

return s.Serve(lis)

}

func (server *WwdatabaseServer) GetCurrentChampion(ctx context.Context, input *pb.GetCurrentChampionParams) (*pb.Champion, error) {
var current_champion *pb.Champion = &pb.Champion{}
	rows, err := server.conn.Query(context.Background(), `SELECT * FROM champions
	WHERE current_champion=true;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		champion:= pb.Champion{}
		err = rows.Scan(&champion.TitleHolder, &champion.TitleHolderNumber, &champion.DateWon,&champion.DateLost, &champion.Show,&champion.PreviousChampion, &champion.CurrentChampion)
		if err != nil {
			return nil, err
		}

	current_champion = &champion
	}

return current_champion, nil
}

var user = goDotEnvVariable("POSTGRES_USER")
var password = goDotEnvVariable("POSTGRES_PASSWORD")
var host = goDotEnvVariable("POSTGRES_HOST")
var port = goDotEnvVariable("POSTGRES_PORT")
var database = goDotEnvVariable("POSTGRES_DB")
var dbConfig = Config {User: user, Password:password, Host:host, Port:port, Database:database}


func buildConnectionURL(dbConfig Config) string {
	dbConfig.User = url.QueryEscape(dbConfig.User)
	dbConfig.Password = url.QueryEscape(dbConfig.Password)
	args := []interface{}{dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", args...)
}


func main() {
// 	poolConfig, err := pgxpool.ParseConfig(buildConnectionURL(dbConfig))
// 	if err != nil {
// 	log.Fatalf("Unable to establish connection %v", err)
// }

// dbPool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
// 	if err != nil {
// 		log.Fatalf("failed to connect to database config: %v", err)
// 	}

// var wwdatabase_server *WwdatabaseServer = NewWwdatabaseServer()
// wwdatabase_server.db = dbPool
// if err := wwdatabase_server.Run(); err !=nil {
// 	log.Fatalf("Failed to serve: %v", err)
// }

	database_url := goDotEnvVariable("DATABASE_URL")
conn, err := pgx.Connect(context.Background(),database_url)
if err != nil {
	log.Fatalf("Unable to establish connection %v", err)
}
defer conn.Close(context.Background())
var wwdatabase_server *WwdatabaseServer = NewWwdatabaseServer()
wwdatabase_server.conn = conn
if err := wwdatabase_server.Run(); err !=nil {
	log.Fatalf("Failed to serve: %v", err)
}

}
