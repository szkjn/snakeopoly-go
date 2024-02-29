# Snakeopoly

Snakeopoly is a cynical and educational take on the classic snake game, viewed through the lens of Google's ascent to an ever-growing monopoly. Developed in Golang, this game immerses players in a NOKIA-inspired digital playground where they navigate a snake, symbolizing Google, on its quest to collect *data points* and *special data points* —each representing Google's strategic acquisitions and milestones.

## Game Architecture

The Snakeopoly codebase is structured as follows:
- assets/: Contains game assets like fonts and images.
    - fonts/: Font files for UI rendering.
    - images/30x30/: Image files for game characters and elements.
    - assets.go: Manages asset loading and processing.
    - competitors.csv: Stores competitors data (Name, slug, year, text, and level).
- game/: Main game logic and components.
    - cfg.go: Configuration constants (screen dimensions, colors, fonts).
    - datapoint.go: DataPoint logic (game objectives).
    - game.go: Core game structure and game state management.
    - snake.go: Snake entity logic.
    - ui.go: UI rendering and management.
    - shapes.go: Contains 2D array shapes to draw pixelated shapes on screen.
- .gitignore
- go.mod, go.sum: Go module files for managing dependencies.
- main.go: Entry point of the game.


## Key Features

- **State Management** : Implements a snake game with a welcome state, play state, and game over state.
- **User Inputs** : Handles user inputs for game interactions.
- **Asset Management** : Manages game assets like images and fonts efficiently.
- **Blink Theme Feature** : Introduces a "Blink Theme" feature that toggles between DayTheme and NightTheme, ensuring the theme resets to the player's chosen theme after completion.
- **Performance Optimization** : Focuses on addressing performance issues and optimizing response time as the project grows.

## Getting Started

To get started with Snakeopoly, clone this repository and ensure you have Golang 1.21.5 installed on your Windows machine. Open the project with Visual Studio Code and run main.go to launch the game. Explore the game/ directory to understand the game's core logic and components.