# Snakeopoly

Snakeopoly offers a cynical and educational perspective on the classic snake game, casting players into a NOKIA-inspired digital arena. This Golang-developed game delves into Google's journey towards becoming a towering monopoly, challenging players to navigate an evocative "G"-shaped snake eager to amass data points and *special data points*. These points are not mere collectibles; they symbolize Google's strategic acquisitions and key milestones. Each *special data point* captured freezes the game, displaying a tongue-in-cheek quote from "Evil Google".

## Educational Narrative
Snakeopoly's narrative and gameplay are deeply inspired by Shoshana Zuboff's seminal work, "The Age of Surveillance Capitalism." The game serves as a critique and exploration of the mechanisms through which Google extends its influence across society and economy. By engaging with the game, players traverse a storyline that illuminates the transformation of personal data into a commodity and the societal implications of surveillance capitalism.

## Gameplay Progression
As players accumulate special data points, their avatar grows in level, evolving from a "Search Mogul" into a "Privacy Predator," then advancing to a "Household Invader," and ultimately becoming a "Surveillance Supremacist." This progression system not only enriches the gameplay experience but also mirrors the escalating stages of influence and control exhibited by Google in the real world.

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

To get started with Snakeopoly, clone this repository and ensure you have Golang 1.21.5.

Explore the game/ directory to understand the game's core logic and components.