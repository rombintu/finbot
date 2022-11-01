package internal

import (
	"time"

	"github.com/rombintu/finbot/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// func (s *Store) FUNC() error {
// 	ctx, err := s.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer s.Close(ctx)
// 	return nil
// }

func (s *Store) PutNote(note tools.Note) error {
	ctx, err := s.Open()
	if err != nil {
		return err
	}
	defer s.Close(ctx)
	if _, err := s.Database.Collection(TABLE).InsertOne(ctx, note); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCategories() ([]string, error) {
	ctx, err := s.Open()
	if err != nil {
		return []string{}, err
	}
	defer s.Close(ctx)
	cur, err := s.Database.Collection(TABLE).Find(ctx, bson.M{})
	if err != nil {
		return []string{}, err
	}
	var notes []tools.Note
	if err := cur.All(ctx, &notes); err != nil {
		return []string{}, err
	}
	var categories []string
	caser := cases.Title(language.Russian)
	for _, n := range notes {
		categories = append(categories, caser.String(n.Category))
	}
	return tools.Unique(categories), nil
}

func (s *Store) GetNotesByMonth() (map[string]int, error) {
	ctx, err := s.Open()
	if err != nil {
		return map[string]int{}, err
	}
	defer s.Close(ctx)
	cur, err := s.Database.Collection(TABLE).Find(ctx, bson.M{"timestamp": bson.M{
		"$gte": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -30)),
	}})
	if err != nil {
		return map[string]int{}, err
	}
	var notes []tools.Note
	if err := cur.All(ctx, &notes); err != nil {
		return map[string]int{}, err
	}
	var categories []string

	for _, n := range notes {
		categories = append(categories, n.Category)
	}
	payload := make(map[string]int)
	for _, c := range categories {
		for _, n := range notes {
			if n.Category == c {
				payload[c] += n.Cost
			}
		}
	}
	return payload, nil
}
