package src

import (
	"aintan/src/ai"
	"aintan/src/user"
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

func Create() {
	scanner := bufio.NewScanner(os.Stdin)

	if !user.ProfileExists() {
		fmt.Print("Welcome! What's your name? ")
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())

		fmt.Print("What is your native language? (e.g., Turkish, English, Spanish): ")
		scanner.Scan()
		nativeLanguage := strings.TrimSpace(scanner.Text())

		fmt.Print("Which language would you like to learn? (e.g., French, Spanish, German): ")
		scanner.Scan()
		language := strings.TrimSpace(scanner.Text())

		fmt.Print("Select difficulty level (easy/normal/hard/extra hard): ")
		scanner.Scan()
		difficulty := strings.ToLower(strings.TrimSpace(scanner.Text()))

		validDifficulties := map[string]bool{
			"easy":       true,
			"normal":     true,
			"hard":       true,
			"extra hard": true,
		}

		if !validDifficulties[difficulty] {
			fmt.Println("Invalid difficulty. Setting to 'normal'.")
			difficulty = "normal"
		}

		if name == "" || nativeLanguage == "" || language == "" {
			fmt.Println("Please provide all required information.")
			return
		}

		err := user.CreateProfile(name, nativeLanguage, language, difficulty)
		if err != nil {
			fmt.Println("Error creating profile:", err)
			return
		}

		fmt.Printf("Profile created successfully! Welcome %s!\n", name)
	}

	profile, err := user.GetProfile()
	if err != nil {
		fmt.Println("Error loading profile:", err)
		return
	}

	user.ShowProfile()
	fmt.Println()

	apiKey, err := ai.GetConfig()
	if err != nil {
		fmt.Println("Error getting API key:", err)
		return
	}

	learner, err := ai.NewLanguageLearner(apiKey, profile.TargetLanguage, profile.NativeLanguage, profile.Difficulty)
	if err != nil {
		fmt.Println("Error creating language learner:", err)
		return
	}

	fmt.Printf("Let's continue learning %s! (Difficulty: %s)\n", profile.TargetLanguage, profile.Difficulty)
	fmt.Println("Type 'quit' at any time to stop the session.")
	fmt.Println("----------------------------------------")

	ctx := context.Background()
	sessionActive := true
	sessionScore := 0

	for sessionActive {
		question, err := learner.GenerateQuestion(ctx)
		if err != nil {
			fmt.Println("Error generating question:", err)
			break
		}

		fmt.Printf("\nQuestion: %s\n", question)
		fmt.Print("Your answer: ")

		scanner.Scan()
		userAnswer := strings.TrimSpace(scanner.Text())

		if strings.ToLower(userAnswer) == "quit" {
			sessionActive = false
			break
		}

		if userAnswer == "" {
			fmt.Println("Please provide an answer or type 'quit' to exit.")
			continue
		}

		isCorrect, feedback, err := learner.CheckAnswer(ctx, question, userAnswer)
		if err != nil {
			fmt.Println("Error checking answer:", err)
			break
		}

		fmt.Printf("\n%s\n", feedback)

		if isCorrect {
			sessionScore++
			err = user.UpdateSingleScore()
			if err != nil {
				fmt.Println("Error updating profile:", err)
			}
			fmt.Printf("🎉 Correct! Session Score: %d\n", sessionScore)
			fmt.Println("Let's try another one!")
		} else {
			err = user.UpdateGamePlayed()
			if err != nil {
				fmt.Println("Error updating profile:", err)
			}
			fmt.Printf("❌ Game Over! Session Score: %d\n", sessionScore)
			fmt.Println("Don't worry, keep practicing!")
			sessionActive = false
		}

		fmt.Println("----------------------------------------")
	}

	fmt.Printf("\nSession completed! You scored %d points.\n", sessionScore)
	user.ShowProfile()
}