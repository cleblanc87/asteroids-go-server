package universe

import (
  "bitbucket.org/cleblanc/roids-go-server/asteroid"
  "bitbucket.org/cleblanc/roids-go-server/bullet"
  "bitbucket.org/cleblanc/roids-go-server/gameobject"
  "encoding/json"
  "fmt"
  "github.com/fzzy/radix/redis"
  "math/rand"

)

type Universe struct{
  Asteroids []*asteroid.Asteroid `json:"asteroids"`
  Bullets []*bullet.Bullet `json:"bullets"`
}

func (u *Universe) InitUniverse(redis_client *redis.Client){
  live_asteroids := redis_client.Cmd("KEYS", "asteroid-*")

  //load asteroids from db
  u.Asteroids = make([]*asteroid.Asteroid,0,len(live_asteroids.Elems))
  for key := range live_asteroids.Elems {
    asteroid := &asteroid.Asteroid{}
    temp_asteroid := redis_client.Cmd("GET", live_asteroids.Elems[key])
    if err := json.Unmarshal([]byte(temp_asteroid.String()), &asteroid); err != nil {
    }
    u.Asteroids = append(u.Asteroids,asteroid)
  }
}

func (u *Universe) Update(delta_time float64,redis_client *redis.Client){
  //here we run the game simulation

  //check redis for api spawned things
  //bullets


  //should we spawn more asteroids?
  u.SpawnAsteroid(redis_client)
    
  //asteroid updates a.Update()
  for _, a := range u.Asteroids{
    a.Update(delta_time)
    //write to redis
    redis_client.Cmd("SET", fmt.Sprintf("asteroid-%v", a.Id), a.Dump())
  }
}

func (u *Universe) SpawnAsteroid(redis_client *redis.Client){
  if rand.Float32() < 0.1 {
    live_asteroids := redis_client.Cmd("KEYS", "asteroid-*")
    num_asteroids := 1000
    if len(live_asteroids.Elems) <  num_asteroids {

      temp_asteroid := &asteroid.Asteroid{gameobject.GameObject{
                          Id: rand.Intn(500000),
                          X: rand.Float64(),
                          Y: rand.Float64(),
                          Velx: rand.Float64() * 100,
                          Vely: rand.Float64() * 100},
                          100}

      u.Asteroids = append(u.Asteroids,temp_asteroid)
      asteroid_json, _ := json.Marshal(temp_asteroid)

      redis_client.Cmd("SET", fmt.Sprintf("asteroid-%v", temp_asteroid.Id), asteroid_json)
    }

    // temp_bullet := &bullet.Bullet{gameobject.GameObject{
    //                     Id: rand.Intn(500000),
    //                     X: rand.Float64(),
    //                     Y: rand.Float64(),
    //                     Velx: rand.Float64() * 100,
    //                     Vely: rand.Float64() * 100},
    //                     100}
    
    // u.Bullets = append(u.Bullets,temp_bullet)

    // redis_client.Cmd("SET", fmt.Sprintf("bullet-%v", temp_bullet.Id), temp_bullet)

  }
}

func (u *Universe) Dump() []byte{
    universe_json, _ := json.Marshal(u)
    return universe_json
}