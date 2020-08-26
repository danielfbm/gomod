package log

import "fmt"

func Log(exec string) error {
	fmt.Println("NEW log", exec)
	return nil
}
