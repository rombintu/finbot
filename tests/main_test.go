package tests

import (
	"fmt"
	"testing"

	"github.com/rombintu/finbot/tools"
)

type Toster struct {
	Name   string
	Input  string
	Output bool
}

func TestNewNote(t *testing.T) {
	tosters := []Toster{
		{"Test1", "100 Food", true},
		{"Test2", "Food 100", true},
		{"Test3", "00 Food", false},
		{"Test4", "200 Food some comment", true},
		{"Test5", "fddd Food some comment", false},
		{"Test6", "fddd Food some commnt", false},
		{"Test7", "fddd food fsdgs 122121", true},
		{"Test8", "SomeCategory Food 211221112 some comment", true},
	}
	for _, toster := range tosters {
		note, isOk := tools.Filter(toster.Input)
		if toster.Output != isOk {
			fmt.Printf("Cost: %d, Cat: %s, Comm: %s\n", note.Cost, note.Category, note.Comment)
			t.Errorf("%s is FAILED [want: %t >> %t]", toster.Name, toster.Output, isOk)
		}
	}
}
