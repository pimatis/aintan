# Aintan - AI Language Learning Platform

Aintan is an AI-powered command-line language learning platform that generates personalized questions and validates answers using Google's Gemini AI. Learn any language with adaptive difficulty levels and track your progress with achievements.

## Features

- **AI-Powered Questions**: Dynamic question generation using Google Gemini AI
- **Multi-Language Support**: Learn any language from your native language
- **Difficulty Levels**: Choose from Easy, Normal, Hard, or Extra Hard
- **Achievement System**: Unlock achievements as you progress
- **Progress Tracking**: Track scores, games played, and learning statistics
- **Personalized Learning**: Adaptive questions based on your skill level
- **Beautiful Tables**: Clean CLI interface with formatted tables

## Quick Start

### Prerequisites

- Go 1.19 or higher
- Google AI API key (Gemini)

### Installation

1. **Clone the repository**
    ```bash
    git clone https://github.com/pimatis/aintan.git
    cd aintan
    ```

2. **Install dependencies**
    ```bash
    go mod tidy
    ```

3. **Build the application** (optional - for faster execution)
    ```bash
    go build -o aintan
    ```

4. **Get your Google AI API key**
    - Visit [Google AI Studio](https://makersuite.google.com/app/apikey)
    - Create a new API key
    - Copy the key for the next step

5. **Set up your API key**
    ```bash
    # Using built binary
    ./aintan -set-key YOUR_API_KEY_HERE
    
    # Or using go run
    go run main.go -set-key YOUR_API_KEY_HERE
    ```

6. **Start learning!**
    ```bash
    # Using built binary
    ./aintan
    
    # Or using go run
    go run main.go
    ```

## Usage

### Basic Commands

```bash
# Start interactive learning mode
./aintan                    # Using built binary
go run main.go              # Using go run

# Show user profile and statistics
./aintan -profile           # Using built binary
go run main.go -profile     # Using go run

# View achievements and progress
./aintan -achievements      # Using built binary
go run main.go -achievements # Using go run

# Change difficulty level
./aintan -difficulty hard   # Using built binary
go run main.go -difficulty hard # Using go run

# Set or update API key
./aintan -set-key YOUR_NEW_API_KEY # Using built binary
go run main.go -set-key YOUR_NEW_API_KEY # Using go run

# Show help
./aintan -help              # Using built binary
go run main.go -help        # Using go run
```

### Difficulty Levels

| Level | Description |
|-------|-------------|
| `easy` | Very simple vocabulary and basic phrases for absolute beginners |
| `normal` | Intermediate vocabulary and common expressions for regular learners |
| `hard` | Advanced vocabulary, complex grammar, and challenging expressions |
| `extra hard` | Expert-level vocabulary, complex sentence structures, and advanced linguistic concepts |

### First Time Setup

When you run Aintan for the first time, you'll be prompted to create a profile:

1. Enter your name
2. Specify your native language
3. Choose the language you want to learn
4. Select difficulty level

## Achievement System

Unlock achievements as you progress:

- **First Steps** - Answer your first question correctly (1 point)
- **Quick Learner** - Reach 5 correct answers
- **Persistent Scholar** - Achieve 10 correct answers
- **Language Enthusiast** - Score 15 correct answers
- **Dedicated Student** - Reach 25 correct answers
- **Academic Excellence** - Achieve 50 correct answers
- **Language Master** - Score 75 correct answers
- **Polyglot Prodigy** - Reach 100 correct answers
- **Linguistic Virtuoso** - Achieve 150 correct answers
- **Grand Scholar** - Score 200 correct answers

## Configuration

### Profile Configuration
User profiles are stored in `config.json` with the following structure:
```json
{
  "name": "Your Name",
  "native_language": "english",
  "target_language": "spanish",
  "difficulty": "normal",
  "total_score": 25,
  "best_score": 8,
  "games_played": 5,
  "last_played": "2024-01-15T10:30:00Z",
  "created_at": "2024-01-10T09:00:00Z"
}
```

### API Key Management
The Google AI API key is securely stored in `key.txt`. You can update it anytime using:
```bash
./aintan -set-key NEW_API_KEY          # Using built binary
go run main.go -set-key NEW_API_KEY    # Using go run
```

## Contributing

We welcome contributions! Here's how you can help:

1. **Fork the repository**
2. **Create a feature branch**
    ```bash
    git checkout -b feature/amazing-feature
    ```
3. **Make your changes**
4. **Commit your changes**
    ```bash
    git commit -m 'Add some amazing feature'
    ```
5. **Push to the branch**
    ```bash
    git push origin feature/amazing-feature
    ```
6. **Open a Pull Request**

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

<div align="center" style="display: flex; align-items: center; justify-content: space-between;">
    <p style="margin-left: 25rem; margin-top: 1.2rem;">Created by <a href="https://github.com/pimatis">Pimatis Labs</a></p>
    <img src="https://www.upload.ee/image/17796243/logo.png" alt="PiContent Logo" width="30" style="opacity: 0.2; position: absolute;">
</div>