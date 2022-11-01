package tools

import (
	"strconv"
	"strings"
	"time"
)

type Note struct {
	UUID      int64
	Cost      int
	Category  string
	Comment   string
	Timestamp time.Time
}

func tryFindMore(target string) (int, bool, error) {
	r := []rune(target)
	if r[len(r)-1] == rune('к') {
		cost, err := strconv.Atoi(string(r[:len(r)-1]))
		if err != nil {
			return -1, false, err
		}
		return cost * 1000, true, nil
	}
	return -1, false, nil
}

// returned Cost, []Index, isFind
func findCost(target []string) (int, int, bool) {
	for i, s := range target {
		cost, err := strconv.Atoi(s)
		if err == nil && cost > 0 {
			return cost, i, true
		}
		cost, ok, err := tryFindMore(s)
		if ok && err == nil && cost > 0 {
			return cost, i, true
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
	category := strings.ToLower(sWithoutCost[0])
	if withComment {
		comment = strings.Join(sWithoutCost[1:], " ")
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

func Unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
