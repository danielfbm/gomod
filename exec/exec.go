package exec

import "fmt"

func Exec(exec string) error {
	fmt.Println("NEW VERSION", exec)
	return nil
}
