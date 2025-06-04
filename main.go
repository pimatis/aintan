package main

import (
	"aintan/src"
	"aintan/src/ai"
	"aintan/src/user"
	"flag"
	"fmt"
)

func main() {
	var (
		setKey       = flag.String("set-key", "", "Set the API key for the application")
		profile      = flag.Bool("profile", false, "Show user profile")
		achievements = flag.Bool("achievements", false, "Show achievements")
		difficulty   = flag.String("difficulty", "", "Change difficulty level (easy/normal/hard/extra hard)")
		help         = flag.Bool("help", false, "Show help information")
	)
	
	flag.Parse()

	if *help {
		showHelp()
		return
	}

	if *setKey != "" {
		ai.SetConfig(*setKey)
		return
	}

	if *profile {
		user.ShowProfile()
		return
	}

	if *achievements {
		profile, err := user.GetProfile()
		if err != nil {
			fmt.Println("No profile found. Please create one first.")
			return
		}
		user.ShowAchievements(profile.TotalScore)
		return
	}

	if *difficulty != "" {
		validDifficulties := map[string]bool{
			"easy":       true,
			"normal":     true,
			"hard":       true,
			"extra hard": true,
		}

		if !validDifficulties[*difficulty] {
			fmt.Println("Invalid difficulty level. Valid options: easy, normal, hard, extra hard")
			return
		}

		err := user.UpdateDifficulty(*difficulty)
		if err != nil {
			fmt.Println("Error updating difficulty:", err)
			return
		}

		fmt.Printf("Difficulty level updated to: %s\n", *difficulty)
		return
	}

	_, err := ai.GetConfig()
	if err != nil {
		fmt.Println("Error: API key not found. Please configure your API key first.")
		fmt.Println("Use: go run main.go -set-key YOUR_API_KEY")
		fmt.Println("Or use: go run main.go -help for more information")
		return
	}

	fmt.Println("Welcome to Aintan - AI Language Learning Platform")
	fmt.Println("========================================")
	
	src.Create()
}

func showHelp() {
	fmt.Println("Aintan - AI Language Learning Platform")
	fmt.Println("=============================")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  go run main.go [OPTIONS]")
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("  -set-key <API_KEY>        Set your Google AI API key")
	fmt.Println("  -profile                  Show user profile and statistics")
	fmt.Println("  -achievements             Show achievements and progress")
	fmt.Println("  -difficulty <LEVEL>       Change difficulty level")
	fmt.Println("  -help                     Show this help message")
	fmt.Println()
	fmt.Println("DIFFICULTY LEVELS:")
	fmt.Println("  easy          Very simple vocabulary and basic phrases")
	fmt.Println("  normal        Intermediate vocabulary and common expressions")
	fmt.Println("  hard          Advanced vocabulary and complex grammar")
	fmt.Println("  extra hard    Expert-level vocabulary and linguistic concepts")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  go run main.go -set-key your-actual-api-key-here")
	fmt.Println("  go run main.go -profile")
	fmt.Println("  go run main.go -achievements")
	fmt.Println("  go run main.go -difficulty hard")
	fmt.Println("  go run main.go -help")
	fmt.Println("  go run main.go")
	fmt.Println()
	fmt.Println("DESCRIPTION:")
	fmt.Println("  This is an AI-powered language learning platform that generates")
	fmt.Println("  questions and validates answers for various languages.")
	fmt.Println("  When run without arguments, it starts the interactive learning mode.")
}