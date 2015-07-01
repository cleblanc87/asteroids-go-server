package asteroid

import (
"bitbucket.org/cleblanc/roids-go-server/gameobject"
"encoding/json"
)

type Asteroid struct {
  gameobject.GameObject
  Lifetime int `json:"lifetime"`
}

func (a *Asteroid) Update(delta_time float64){
 //gameobject.update
    a.X = (float64(a.Velx) * delta_time) + a.X
    a.Y = (float64(a.Vely) * delta_time) + a.Y

    //bounded check
    if a.X > 6400.0 {
      a.X = 0.0
    } else if a.X < 0 {
      a.X = 6400.0
    }
    if a.Y > 6400.0 {
      a.Y = 0
    } else if a.Y < 0 {
      a.Y = 6400.0
    }
}

func (a *Asteroid) Dump() string{
  asteroid_json, _ := json.Marshal(a)
  return string(asteroid_json)
}
