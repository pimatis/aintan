package ai

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genai"
)

type LanguageLearner struct {
	client          *genai.Client
	language        string
	nativeLanguage  string
	difficulty      string
	score           int
	questionHistory []string
}

func NewLanguageLearner(apiKey, language, nativeLanguage, difficulty string) (*LanguageLearner, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}

	return &LanguageLearner{
		client:          client,
		language:        language,
		nativeLanguage:  nativeLanguage,
		difficulty:      difficulty,
		score:           0,
		questionHistory: make([]string, 0),
	}, nil
}

func (ll *LanguageLearner) GenerateQuestion(ctx context.Context) (string, error) {
	var historyText string
	if len(ll.questionHistory) > 0 {
		historyText = fmt.Sprintf("\n\nDo NOT repeat these previous questions:\n%s", strings.Join(ll.questionHistory, "\n"))
	}

	difficultyPrompts := map[string]string{
		"easy":       "very simple vocabulary and basic phrases suitable for absolute beginners",
		"normal":     "intermediate vocabulary and common expressions for regular learners",
		"hard":       "advanced vocabulary, complex grammar, and challenging expressions",
		"extra hard": "expert-level vocabulary, complex sentence structures, and advanced linguistic concepts",
	}

	questionTypes := []string{
		"vocabulary translation from " + ll.language + " to " + ll.nativeLanguage,
		"vocabulary translation from " + ll.nativeLanguage + " to " + ll.language,
		"basic grammar fill-in-the-blank",
		"common phrase translation",
		"verb conjugation",
		"plural forms",
		"basic sentence structure",
	}

	difficultyLevel := difficultyPrompts[ll.difficulty]
	if difficultyLevel == "" {
		difficultyLevel = difficultyPrompts["normal"]
	}

	prompt := fmt.Sprintf("Generate ONE unique %s language learning question for a native %s speaker. Focus on: %s. Difficulty level: %s - use %s. Use %s language for instructions and explanations. Only return the question text, no additional explanations.%s",
		ll.language,
		ll.nativeLanguage,
		questionTypes[len(ll.questionHistory)%len(questionTypes)],
		ll.difficulty,
		difficultyLevel,
		ll.nativeLanguage,
		historyText)

	result, err := ll.client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate question: %w", err)
	}

	question := strings.TrimSpace(result.Text())
	ll.questionHistory = append(ll.questionHistory, question)

	if len(ll.questionHistory) > 10 {
		ll.questionHistory = ll.questionHistory[1:]
	}

	return question, nil
}

func (ll *LanguageLearner) CheckAnswer(ctx context.Context, question, userAnswer string) (bool, string, error) {
	prompt := fmt.Sprintf("Question: %s\nUser's answer: %s\n\nIs this answer correct for a native %s speaker learning %s? Respond with 'CORRECT' or 'INCORRECT' followed by a brief explanation in %s. If incorrect, provide the correct answer.",
		question, userAnswer, ll.nativeLanguage, ll.language, ll.nativeLanguage)

	result, err := ll.client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return false, "", fmt.Errorf("failed to check answer: %w", err)
	}

	response := strings.TrimSpace(result.Text())
	isCorrect := strings.HasPrefix(strings.ToUpper(response), "CORRECT")

	if isCorrect {
		ll.score++
	}

	return isCorrect, response, nil
}

func (ll *LanguageLearner) GetScore() int {
	return ll.score
}
