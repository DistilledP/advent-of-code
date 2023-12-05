package lib

import (
	"log"
	"strconv"
	"strings"
)

const (
	RedCubes   = 12
	GreenCubes = 13
	BlueCubes  = 14
)

type Game struct {
	ID    int
	Name  string
	Red   int
	Green int
	Blue  int
	POW   int
}

// Part01 result: 1853
func Part01(content string) int {
	games := buildGames(content)

	var idSum int
	for i := 0; i < len(games); i++ {
		game := games[i]
		if game.Red <= RedCubes && game.Green <= GreenCubes && game.Blue <= BlueCubes {
			idSum += game.ID
		}
	}

	return idSum
}

// Part02 result: 72706
func Part02(content string) int {
	games := buildGames(content)

	var powSum int
	for i := 0; i < len(games); i++ {
		game := games[i]
		powSum += game.POW
	}

	return powSum
}

// Helpers
func buildGames(content string) []Game {
	gameHistory := strings.Split(content, "\n")

	var gameList []Game
	for i := 0; i < len(gameHistory); i++ {
		gameParts := strings.Split(gameHistory[i], ":")
		rounds := strings.Split(gameParts[1], ";")

		gameNameID := strings.Split(strings.TrimSpace(gameParts[0]), " ")
		var gameID int
		if id, err := strconv.Atoi(gameNameID[1]); err == nil {
			gameID = id
		}

		game := Game{
			ID:   gameID,
			Name: gameParts[0],
		}

		for j := 0; j < len(rounds); j++ {
			round := strings.Split(rounds[j], ",")

			for k := 0; k < len(round); k++ {
				countAndColour := strings.Split(strings.TrimSpace(round[k]), " ")
				var ballCount int
				if count, err := strconv.Atoi(countAndColour[0]); err == nil {
					ballCount = count

				}

				switch strings.ToLower(countAndColour[1]) {
				case "red":
					if ballCount > game.Red {
						game.Red = ballCount
					}
				case "blue":
					if ballCount > game.Blue {
						game.Blue = ballCount

					}
				case "green":
					if ballCount > game.Green {
						game.Green = ballCount
					}
				default:
					log.Printf("unexpected colour: %s", countAndColour[1])
				}
			}
		}
		game.POW = game.Red * game.Blue * game.Green
		gameList = append(gameList, game)
	}

	return gameList
}
