package internal

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DATABASE string = "finbot"
const TABLE string = "notes"

type Store struct {
	Driver   *mongo.Client
	Database *mongo.Database
	Options  *options.ClientOptions
}

func NewStore(opts string) *Store {
	return &Store{
		Options: getOptions(opts),
	}
}

func getOptions(opts string) *options.ClientOptions {
	return options.Client().ApplyURI(opts)
}

func (s *Store) Open() (context.Context, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, s.Options)
	if err != nil {
		return nil, err
	}
	s.Driver = client
	s.Database = s.Driver.Database(DATABASE)
	return ctx, nil
}

func (s *Store) Close(ctx context.Context) error {
	return s.Driver.Disconnect(ctx)
}
