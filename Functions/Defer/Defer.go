package  main

import (
	"fmt"
	"os"

	"io"
)

func main() {
	defer goodbye()
	fmt.Println("Hello")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	defer panic ("two")
	if err := fileCopy("readme.txt","readme-copy.txt"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func goodbye() {
	fmt.Println("GoodBye")
}

func fileCopy(sourceName string,destName string){
	src, err := os.Open(sourceName)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("closing",sourceName)
		src.Close()
	}()

	dst, err := os.Create(destName)
	if err != nil {
		return err
	}

	defer func() {
		fmt.Println("closing",destName)
		dst.Close()
	}()

	if _, err := io.Copy(dst,src); err != nil {
		return err
	}
	return nil

}