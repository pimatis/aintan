package user

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Achievement struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Requirement int    `json:"requirement"`
	Unlocked    bool   `json:"unlocked"`
}

var achievements = []Achievement{
	{
		Name:        "First Steps",
		Description: "Answer your first question correctly",
		Icon:        "🌱",
		Requirement: 1,
		Unlocked:    false,
	},
	{
		Name:        "Quick Learner",
		Description: "Reach 5 correct answers",
		Icon:        "⚡",
		Requirement: 5,
		Unlocked:    false,
	},
	{
		Name:        "Persistent Scholar",
		Description: "Achieve 10 correct answers",
		Icon:        "📚",
		Requirement: 10,
		Unlocked:    false,
	},
	{
		Name:        "Language Enthusiast",
		Description: "Score 15 correct answers",
		Icon:        "🌟",
		Requirement: 15,
		Unlocked:    false,
	},
	{
		Name:        "Dedicated Student",
		Description: "Reach 25 correct answers",
		Icon:        "🎯",
		Requirement: 25,
		Unlocked:    false,
	},
	{
		Name:        "Academic Excellence",
		Description: "Achieve 50 correct answers",
		Icon:        "🏆",
		Requirement: 50,
		Unlocked:    false,
	},
	{
		Name:        "Language Master",
		Description: "Score 75 correct answers",
		Icon:        "👑",
		Requirement: 75,
		Unlocked:    false,
	},
	{
		Name:        "Polyglot Prodigy",
		Description: "Reach 100 correct answers",
		Icon:        "🎖️",
		Requirement: 100,
		Unlocked:    false,
	},
	{
		Name:        "Linguistic Virtuoso",
		Description: "Achieve 150 correct answers",
		Icon:        "💎",
		Requirement: 150,
		Unlocked:    false,
	},
	{
		Name:        "Grand Scholar",
		Description: "Score 200 correct answers",
		Icon:        "🌌",
		Requirement: 200,
		Unlocked:    false,
	},
}

func CheckAchievements(totalScore int) []Achievement {
	var newAchievements []Achievement

	for i := range achievements {
		if !achievements[i].Unlocked && totalScore >= achievements[i].Requirement {
			achievements[i].Unlocked = true
			newAchievements = append(newAchievements, achievements[i])
		}
	}

	return newAchievements
}

func ShowAchievements(totalScore int) {
	fmt.Println()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("🏅 Achievements")
	t.AppendHeader(table.Row{"Status", "Achievement", "Description", "Progress"})

	unlockedCount := 0
	for _, achievement := range achievements {
		var status string
		var progress string

		if totalScore >= achievement.Requirement {
			status = "✅ " + achievement.Icon
			progress = "UNLOCKED"
			unlockedCount++
		} else {
			status = "🔒"
			progress = fmt.Sprintf("%d/%d", totalScore, achievement.Requirement)
		}

		t.AppendRow(table.Row{
			status,
			achievement.Name,
			achievement.Description,
			progress,
		})
	}

	t.Render()

	summaryTable := table.NewWriter()
	summaryTable.SetOutputMirror(os.Stdout)
	summaryTable.SetTitle("📊 Achievement Summary")
	summaryTable.AppendHeader(table.Row{"Metric", "Value"})

	summaryTable.AppendRow(table.Row{"Unlocked", fmt.Sprintf("%d/%d", unlockedCount, len(achievements))})
	summaryTable.AppendRow(table.Row{"Progress", fmt.Sprintf("%.1f%%", float64(unlockedCount)/float64(len(achievements))*100)})

	if unlockedCount < len(achievements) {
		nextAchievement := getNextAchievement(totalScore)
		if nextAchievement != nil {
			remaining := nextAchievement.Requirement - totalScore
			summaryTable.AppendRow(table.Row{"Next Goal", fmt.Sprintf("%s %s", nextAchievement.Icon, nextAchievement.Name)})
			summaryTable.AppendRow(table.Row{"Points Needed", remaining})
		}
	}

	summaryTable.Render()
}

func getNextAchievement(totalScore int) *Achievement {
	for _, achievement := range achievements {
		if totalScore < achievement.Requirement {
			return &achievement
		}
	}
	return nil
}

func DisplayNewAchievements(newAchievements []Achievement) {
	if len(newAchievements) > 0 {
		fmt.Println()

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("🎉 NEW ACHIEVEMENT UNLOCKED!")
		t.AppendHeader(table.Row{"Icon", "Achievement", "Description"})

		for _, achievement := range newAchievements {
			t.AppendRow(table.Row{
				achievement.Icon,
				achievement.Name,
				achievement.Description,
			})
		}

		t.Render()
		fmt.Println()
	}
}