package user

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Profile struct {
	Name           string    `json:"name"`
	NativeLanguage string    `json:"native_language"`
	TargetLanguage string    `json:"target_language"`
	Difficulty     string    `json:"difficulty"`
	TotalScore     int       `json:"total_score"`
	BestScore      int       `json:"best_score"`
	GamesPlayed    int       `json:"games_played"`
	LastPlayed     time.Time `json:"last_played"`
	CreatedAt      time.Time `json:"created_at"`
}

func CreateProfile(name, nativeLanguage, targetLanguage, difficulty string) error {
	profile := Profile{
		Name:           name,
		NativeLanguage: nativeLanguage,
		TargetLanguage: targetLanguage,
		Difficulty:     difficulty,
		TotalScore:     0,
		BestScore:      0,
		GamesPlayed:    0,
		LastPlayed:     time.Now(),
		CreatedAt:      time.Now(),
	}

	return saveProfile(profile)
}

func GetProfile() (*Profile, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("profile not found: %w", err)
	}
	defer file.Close()

	var profile Profile
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&profile)
	if err != nil {
		return nil, fmt.Errorf("error reading profile: %w", err)
	}

	return &profile, nil
}

func UpdateSingleScore() error {
	profile, err := GetProfile()
	if err != nil {
		return err
	}

	profile.TotalScore++
	profile.LastPlayed = time.Now()

	if profile.TotalScore > profile.BestScore {
		profile.BestScore = profile.TotalScore
	}

	err = saveProfile(*profile)
	if err != nil {
		return err
	}

	newAchievements := CheckAchievements(profile.TotalScore)
	DisplayNewAchievements(newAchievements)

	return nil
}

func UpdateGamePlayed() error {
	profile, err := GetProfile()
	if err != nil {
		return err
	}

	profile.GamesPlayed++
	profile.LastPlayed = time.Now()

	return saveProfile(*profile)
}

func UpdateScore(sessionScore int) error {
	profile, err := GetProfile()
	if err != nil {
		return err
	}

	fmt.Printf("Before update - Total: %d, Best: %d, Games: %d\n", profile.TotalScore, profile.BestScore, profile.GamesPlayed)

	profile.TotalScore += sessionScore
	profile.GamesPlayed++
	profile.LastPlayed = time.Now()

	if sessionScore > profile.BestScore {
		profile.BestScore = sessionScore
	}

	fmt.Printf("After update - Total: %d, Best: %d, Games: %d\n", profile.TotalScore, profile.BestScore, profile.GamesPlayed)

	err = saveProfile(*profile)
	if err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}

	return nil
}

func UpdateDifficulty(difficulty string) error {
	profile, err := GetProfile()
	if err != nil {
		return err
	}

	profile.Difficulty = difficulty
	profile.LastPlayed = time.Now()

	return saveProfile(*profile)
}

func ShowProfile() {
	profile, err := GetProfile()
	if err != nil {
		fmt.Println("No profile found. Please create one first.")
		return
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("🎓 User Profile")
	t.AppendHeader(table.Row{"Field", "Value"})
	
	t.AppendRow(table.Row{"Name", profile.Name})
	t.AppendRow(table.Row{"Native Language", profile.NativeLanguage})
	t.AppendRow(table.Row{"Learning", profile.TargetLanguage})
	t.AppendRow(table.Row{"Difficulty", profile.Difficulty})
	t.AppendRow(table.Row{"Total Score", profile.TotalScore})
	t.AppendRow(table.Row{"Best Score", profile.BestScore})
	t.AppendRow(table.Row{"Games Played", profile.GamesPlayed})
	
	if profile.GamesPlayed > 0 {
		avgScore := float64(profile.TotalScore) / float64(profile.GamesPlayed)
		t.AppendRow(table.Row{"Average Score", fmt.Sprintf("%.1f", avgScore)})
	}
	
	t.AppendRow(table.Row{"Member Since", profile.CreatedAt.Format("2006-01-02")})
	t.AppendRow(table.Row{"Last Played", profile.LastPlayed.Format("2006-01-02 15:04")})
	
	t.Render()

	ShowAchievements(profile.TotalScore)
}

func ProfileExists() bool {
	_, err := os.Stat("config.json")
	return err == nil
}

func saveProfile(profile Profile) error {
	file, err := os.OpenFile("config.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening profile file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(profile)
	if err != nil {
		return fmt.Errorf("error writing profile: %w", err)
	}

	return nil
}