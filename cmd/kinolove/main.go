package main

import "kinolove/internal/app"

func main() {
	callback := app.Startup()
	defer callback()
}
