package main

import "github.com/qlfzn/etf-go/cmd/api"

func main(){
	app := api.Application{
		Addr: ":8080",
	}

	mux := app.NewServer()

	// run server
	app.Run(mux)
}