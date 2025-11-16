# Number Guessing Game

A simple command-line number guessing game written in Go. The player tries to guess a randomly generated number within a specified range and difficulty level.

## Features
- Multiple difficulty levels
- Input validation
- Replay option
- Modular code structure

## Project Structure
```
go.mod
cmd/
  main.go
internal/
  game/
    difficulty.go
    game.go
  utils/
    input.go
```

## Getting Started

### Prerequisites
- Go 1.18 or higher

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/bhaskar2610/number-guessing-game.git
   cd number-guessing-game
   ```
2. Build the project:
   ```sh
   go build -o guessing-game ./cmd/main.go
   ```

### Running the Game
```sh
./guessing-game
```

## How to Play
1. Select a difficulty level.
2. Try to guess the randomly generated number within the allowed attempts.
3. The game will provide feedback after each guess.
4. Play again or exit after finishing a round.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is licensed under the MIT License.
