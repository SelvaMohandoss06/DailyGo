package main

import "fmt"

func main() {
	var valid bool
	var command string
	fmt.Println(valid)
	fmt.Println(command)
	autosave, backupext, interval := true, `bak`, "10s"
	fmt.Printf("autosave:%+v \n", autosave)
	fmt.Printf("backup:%+v \n", backupext)
	fmt.Printf("interval: %+v\n", interval)
	foo := 5
	fmt.Println(foo)
	bar := true
	fmt.Println(bar)
	bookId := 1
	propername := "rerer"
	bodyHtml := `<body>This is the body of the page.  Link to <a href="https://www.gopherguides.com">Gopher Guides</a></body>`

	fmt.Println(bodyHtml)
	fmt.Println(bookId)
	fmt.Println(propername)
}
