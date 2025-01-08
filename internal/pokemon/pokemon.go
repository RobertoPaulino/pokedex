package pokemon

type Pokemon struct {
  Height int
  Weight int
  Stats PokemonStats   
  Name string
  Types []string
}

type PokemonStats struct {
  Hp int
  Attack int
  Defense int
  SpecialAttack int
  SpecialDefense int
  Speed int
}
