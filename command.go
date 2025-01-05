package main

type cliCommand struct {
  name        string
  description string
  callback    func() error
  config config
}

type config struct {
  next string
  previous string
}
