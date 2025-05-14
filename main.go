package main

import "github.com/ZobayerAbedin/BookServer/internal"

func main() {
	var id int
	internal.BookDB, id = internal.InitBook()
	app := internal.App{}
	app.Initialise(internal.BookDB, id)
	app.Run("localhost:10000")
}
