package log

import "fmt"

func Log(exec string) error {
	fmt.Println("log", exec)
	return nil
}
