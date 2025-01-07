package pokemon

type Pokemon struct {
  Height int
  Weight int
  Stats struct {
    Hp int
    Attack int
    Defense int
    SpecialAttack int
    SpecialDefense int
    Speed int
  }
  Name string
  Types []string
}
