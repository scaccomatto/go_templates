package main

import (
	"paru.net/gosimpleapp/internal/app"
)

func main() {
	//I choose to implement all the injection logic in the internal/app folder to not expose it as public.
	app.Start()
}
