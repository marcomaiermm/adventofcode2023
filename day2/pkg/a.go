package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	RegexGameIDPattern   = string(`Game\s+(\d+)`)
	RegexRedGemPattern   = string(`(\d+)\s+red`)
	RegexBlueGemPattern  = string(`(\d+)\s+blue`)
	RegexGreenGemPattern = string(`(\d+)\s+green`)
	GameRegex            *regexp.Regexp
	RedGemRegex          *regexp.Regexp
	GreenGemRegex        *regexp.Regexp
	BlueGemRegex         *regexp.Regexp
)

type Balls struct {
	Blue  int
	Green int
	Red   int
}

type Game struct {
	Balls    []Balls
	possible bool
	ballSack Balls
	ID       int
}

func (g *Game) getGameID(gameStr string) error {
	matches := GameRegex.FindStringSubmatch(gameStr)
	if len(matches) < 2 {
		return fmt.Errorf("too few matches. Expected to have 3 Matches but got %d", len(matches))
	}

	gameId, err := strconv.Atoi(matches[1])
	if err != nil {
		return fmt.Errorf("i have troubles converting string id to int id: %s", err)
	}

	g.ID = gameId

	return nil
}

func (g *Game) putBallsInBallSack(ballStr string) error {
	balls, err := GetBallsFromBallSack(ballStr)
	if err != nil {
		return err
	}

	g.Balls = append(g.Balls, Balls{
		Green: balls.Green,
		Blue:  balls.Blue,
		Red:   balls.Red,
	})

	return nil
}

func (g *Game) verifyGame() {
	for _, balls := range g.Balls {
		if balls.Blue > g.ballSack.Blue || balls.Red > g.ballSack.Red || balls.Green > g.ballSack.Green {
			g.possible = false
			return
		}
	}
	g.possible = true
}

func NewGame(input string, bag Balls) *Game {
	game := &Game{ballSack: bag, possible: false}
	game.getGameID(input)
	fmt.Println(input)

	bags := strings.Split(input, ";")
	for _, bagStr := range bags {
		game.putBallsInBallSack(bagStr)
	}
	game.verifyGame()

	fmt.Println(game)
	return game
}

func GetPossibleGames(input []string, bag Balls) int {
	var n int
	for _, gameStr := range input {
		game := NewGame(gameStr, bag)
		if game.possible {
			n += game.ID
		}
	}
	return n
}

func GetBallsFromBallSack(ballStr string) (*Balls, error) {
	Red := 0
	Green := 0
	Blue := 0
	var err error

	redMatches := RedGemRegex.FindStringSubmatch(ballStr)
	blueMatches := BlueGemRegex.FindStringSubmatch(ballStr)
	greenMatches := GreenGemRegex.FindStringSubmatch(ballStr)

	if len(redMatches) >= 1 {
		Red, err = strconv.Atoi(redMatches[1])
		if err != nil {
			return nil, fmt.Errorf("i have troubles converting string to int: %s", err)
		}
	}

	if len(blueMatches) >= 1 {
		Blue, err = strconv.Atoi(blueMatches[1])
		if err != nil {
			return nil, fmt.Errorf("i have troubles converting string to int: %s", err)
		}
	}

	if len(greenMatches) >= 1 {
		Green, err = strconv.Atoi(greenMatches[1])
		if err != nil {
			return nil, fmt.Errorf("i have troubles converting string to int: %s", err)
		}
	}

	return &Balls{
		Red,
		Green,
		Blue,
	}, nil
}

func init() {
	var err error
	GameRegex, err = regexp.Compile(RegexGameIDPattern)
	if err != nil {
		fmt.Printf("Error compiling game regex: %s", err)
	}
	RedGemRegex, err = regexp.Compile(RegexRedGemPattern)
	if err != nil {
		fmt.Printf("Error compiling red gem regex: %s", err)
	}
	GreenGemRegex, err = regexp.Compile(RegexBlueGemPattern)
	if err != nil {
		fmt.Printf("Error compiling blue gem regex: %s", err)
	}
	BlueGemRegex, err = regexp.Compile(RegexGreenGemPattern)
	if err != nil {
		fmt.Printf("Error compiling green gem regex: %s", err)
	}
}
