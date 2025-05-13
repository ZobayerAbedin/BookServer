package main

import "github.com/ZobayerAbedin/BookServer/internal"

func main() {
	var id int
	internal.Books, id = internal.InitBook()
	app := internal.App{}
	app.Initialise(internal.Books, id)
	app.Run("localhost:10000")
}
