package game

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
)

type GameHandler struct {
	Game *Game
}

func (gh *GameHandler) handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GameHandler.handlerRoot !")

	// show README.md
}

func (gh *GameHandler) handlerTower(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.RawQuery) < 2 {
		errMsg := "Error reading tower coordinates"
		w.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(w, errMsg)
		return
	}

	y := r.URL.RawQuery[:1]
	x := r.URL.RawQuery[1:]

	Y := -1
	for i, c := range AlphabetY {
		if string(c) == y {
			Y = i + 1
			break
		}
	}
	if Y == -1 {
		errMsg := "Error reading Y coord: value not found in alphabet"
		//fmt.Println(errMsg)
		w.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(w, errMsg)
		return
	}

	X, err := strconv.Atoi(x)
	if err != nil {
		errMsg := fmt.Sprint("Error reading Y coord:", err)
		//fmt.Println(errMsg)
		w.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(w, errMsg)
		return
	}

	// check if the coordinates of the tower match with a constructible place (i.e. a wall in the ground)
	cell := gh.Game.StageGround[X][Y]
	if !cell.IsWall() {
		errMsg := fmt.Sprint("Error: Unconstructible cell, content:", cell.Content)
		//fmt.Println(errMsg)
		w.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(w, errMsg)
		return
	}

	// check if the coordinates match and existing tower in the stage
	cell = gh.Game.Stage[X][Y]
	if cell.IsTower() {
		done := gh.Game.Player.SpendMoney(Config.Game.MoneyPerTowerUpgrade)
		if !done {
			errMsg := "Not enough money to upgrade"
			//fmt.Println(errMsg)
			w.WriteHeader(http.StatusNotAcceptable)
			io.WriteString(w, errMsg)
			return
		}
		for _, tower := range gh.Game.Towers {
			if tower.X == X && tower.Y == Y {
				tower.Strength += 1
				break
			}
		}
	} else {
		done := gh.Game.Player.SpendMoney(Config.Game.MoneyPerTowerBuild)
		if !done {
			errMsg := "Not enough money to build"
			//fmt.Println(errMsg)
			w.WriteHeader(http.StatusNotAcceptable)
			io.WriteString(w, errMsg)
			return
		}
		tower := NewTower(X, Y)
		gh.Game.Towers = append(gh.Game.Towers, tower)
	}
}

type GameServer struct {
	Game *Game
}

func (gs *GameServer) Start() context.Context {
	const serverAddr = "localhost"

	gameHandler := GameHandler{Game: gs.Game}

	mux := http.NewServeMux()
	mux.HandleFunc("/", gameHandler.handlerRoot)
	mux.HandleFunc("/tower/", gameHandler.handlerTower)

	ctx, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, serverAddr, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server: %s\n", err)
		}
		cancelCtx()
	}()
	return ctx

}

func (gs *GameServer) End(ctx context.Context) {
	<-ctx.Done()
}
