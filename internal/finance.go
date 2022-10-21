package internal

import "github.com/rombintu/finbot/tools"

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
