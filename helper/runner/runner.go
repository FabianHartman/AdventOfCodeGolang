package runner

func Run(function func() error) {
	err := function()
	if err != nil {
		panic(err)
	}
}
