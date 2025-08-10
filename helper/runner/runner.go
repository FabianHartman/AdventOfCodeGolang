package runner

import "adventOfCode/helper/timer"

func Run(function func() error) {
	err := timer.TimeExecution(function)
	if err != nil {
		panic(err)
	}
}
