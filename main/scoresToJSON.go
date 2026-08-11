package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// constant app name used to build the fixed leaderboard path
const leaderboardDir = "typetest"
 
// constant file path: <user config dir>/typetest/leaderboard.json
func leaderboardPath() (string, error) {
    configDir, err := os.UserConfigDir()
    if err != nil {
        return "", err
    }
    dir := filepath.Join(configDir, leaderboardDir)
    if err := os.MkdirAll(dir, 0o755); err != nil {
        return "", err
    }
    return filepath.Join(dir, "leaderboard.json"), nil
}

// []LeaderboardEntry struct
type LeaderboardEntry struct {
	WPM float64 `json:"wpm"`
	Timestamp time.Time `json:"timestamp"`
}

// load leaderboard, initialized in func main()
func loadLeaderboard() ([]LeaderboardEntry, error) {
	path, err := leaderboardPath()
	if err != nil {
		return nil, err
	}
	// read the file
	data, err := os.ReadFile(path)

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
	path, err := leaderboardPath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
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
