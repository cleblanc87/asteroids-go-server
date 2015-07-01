package gameobject

type GameObject struct {
  Id int `json:"id"`
  X float64 `json:"x"`
  Y float64 `json:"y"`
  Velx float64 `json:"velx"`
  Vely float64 `json:"vely"`
}