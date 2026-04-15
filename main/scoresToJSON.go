package main

import (
	"encoding/json"
	"os"
	"sort"
	"time"
)

// leaderboard.json creates automatically on local machine
const leaderboardFile = "leaderboard.json"

// []LeaderboardEntry struct
type LeaderboardEntry struct {
	WPM float64 `json:"wpm"`
	Timestamp time.Time `json:"timestamp"`
}

// load leaderboard, initialized in func main()
func loadLeaderboard() ([]LeaderboardEntry, error) {
	// read the file
	data, err := os.ReadFile(leaderboardFile)

	// if does not exist return nil
	if err != nil {
		if os.IsNotExist(err) {
			return []LeaderboardEntry{}, nil
		}
		return nil, err
	}
	// unmarshal
	var scores []LeaderboardEntry
		if err := json.Unmarshal(data, &scores); err != nil {
		return nil, err
	}
	return scores, nil
}

// save score function
func saveScore(wpm float64) error {
	scores, err := loadLeaderboard()
	if err != nil {
		return err
	}
	scores = append(scores, LeaderboardEntry{
		WPM: wpm,
		Timestamp: time.Now(),
	})

	// sort wpm from high to low
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].WPM > scores[j].WPM
	})
	if len(scores) > 10 {
		scores = scores[:10]
	}
	
	// marshal
	data, err := json.MarshalIndent(scores, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(leaderboardFile, data, 0644)
}

// return the best WPM 
func bestScore(scores []LeaderboardEntry) float64 {
	if len(scores) == 0 {
		return 0
	}
	best := 0.0
	for _, s := range scores {
		if s.WPM > best {
			best = s.WPM
		}
	}
	return best
}
