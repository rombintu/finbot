package tools

import (
	"strconv"
	"strings"
	"time"
)

type Note struct {
	Cost      int
	Category  string
	Comment   string
	Timestamp time.Time
}

// returned Cost, []Index, isFind
func findCost(target []string) (int, int, bool) {
	for i, s := range target {
		cost, err := strconv.Atoi(s)
		if err != nil {
			continue
		} else {
			if cost > 0 {
				return cost, i, true
			}
		}
	}
	return -1, 0, false
}

// Note and OK
func Filter(target string) (Note, bool) {
	s := strings.Split(target, " ")
	withComment := false
	switch size := len(s); {
	case size > 2:
		withComment = true
	case size < 2:
		return Note{}, false
	}

	cost, idx, isFind := findCost(s)
	if !isFind {
		return Note{}, false
	}
	comment := ""
	sWithoutCost := removeByIndex(s, idx)
	category := sWithoutCost[0]
	if withComment {
		comment = strings.Join(sWithoutCost[0:], " ")
	}
	return Note{
		Cost:      cost,
		Category:  category,
		Comment:   comment,
		Timestamp: time.Now(),
	}, true
}

func removeByIndex(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)
}
