package main

import (
	"sync"
	"fmt"
)

type DefaultEngine struct {
	user *DefaultUser
	game *Game
	gameChannel chan DefaultGameMsg
	inputChannel chan string
}

type DefaultUser struct {
	id string
}

type DefaultGameMsg struct {
	msgType string
	userWhite bool
	gameStatus string
	currMove string
	winner string
}

func (engine *DefaultEngine) Setup(gameChannel chan DefaultGameMsg,
								   inputChannel chan string) {

	engine.gameChannel = gameChannel
	engine.inputChannel = inputChannel

	engine.game = &Game{}
	engine.game.Setup()
}

func (engine *DefaultEngine) getUser() *DefaultUser {
	return engine.user
}

func (engine *DefaultEngine) getGameChannel() chan DefaultGameMsg {
	return engine.gameChannel
}

func (engine *DefaultEngine) run(wg *sync.WaitGroup) {
	defer wg.Done()

	printBoard(engine.game.board, engine.game.color)

	gameMsg := DefaultGameMsg{}
	gameMsg.msgType = "gameFull"

	for {
		printMoveList(engine.game.GetAllLegalMoves())
		engine.gameChannel <- gameMsg
		cmd := <- engine.inputChannel
		fromSqr, toSqr := engine.game.TranslateCommand(cmd)
		move, err := engine.game.ProcessMove(fromSqr, toSqr)
		if err != nil {
			fmt.Println(err)
		} else {
			engine.game.MakeMove(move)
		}

		printBoard(engine.game.board, engine.game.color)

		gameMsg := DefaultGameMsg{}
		gameMsg.msgType = "gameState"
		gameMsg.currMove = cmd
	}
}

func (gameMsg *DefaultGameMsg) getType() string {
	return gameMsg.msgType
}

func (gameMsg *DefaultGameMsg) isUserWhite() bool {
	return gameMsg.userWhite
}

func (gameMsg *DefaultGameMsg) getGameStatus() string {
	return gameMsg.gameStatus
}

func (gameMsg *DefaultGameMsg) getCurrMove() string {
	return gameMsg.currMove
}

func (gameMsg *DefaultGameMsg) getWinner() string {
	return gameMsg.winner
}