package bullet

import (
"bitbucket.org/cleblanc/roids-go-server/gameobject"
)

type Bullet struct{
  gameobject.GameObject
  Lifetime int `json:"lifetime"`
}

func (a *Bullet) Update(delta_time float64){
  //gameobject.update
  a.X = (float64(a.Velx) * delta_time) + a.X
  a.Y = (float64(a.Vely) * delta_time) + a.Y
}
