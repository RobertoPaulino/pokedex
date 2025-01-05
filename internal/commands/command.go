package commands 

type cliCommand struct {
  Name        string
  description string
  Callback    func(*config) error
  Config *config
}

type config struct {
  next string
  previous string
}
