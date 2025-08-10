package timer

import (
	"fmt"
	"time"
)

func TimeExecution(function func() error) error {
	begin := time.Now()

	err := function()

	fmt.Println("Took:", time.Since(begin))

	return err
}
