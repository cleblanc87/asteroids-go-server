package main

import "bitbucket.org/cleblanc/roids-go-server/universe"
import "github.com/cleblanc87/pusher-http-go"
import "github.com/fzzy/radix/redis"
import "math/rand"
import "time"
import "fmt"
import "encoding/json"
import "net/http"

type test_struct struct {
    Test string
}

type UniverseMessage struct {
   Asteroids []string `json:"asteroids"`
   Timestamp int64 `json:"timestamp"`
}


func viewHandler(w http.ResponseWriter, r *http.Request){

  var t test_struct   
  if err := json.Unmarshal([]byte(r.PostFormValue("foo")), &t); err != nil {
  }
  r.ParseForm()
  fmt.Println(r.Form)

  fmt.Fprint(w, "thanks")
}


// connect to slanger
var pusher_client = pusher.Client{
  AppId: "$ID",
  Key: "$KEY",
  Secret: "$SECRET",
  Host: "roids.cloudapp.net:4567",
}

//connect to redis
var redis_client, err = redis.Dial("tcp", "localhost:6379")
var univ = &universe.Universe{}

func main(){
  //select db
  redis_client.Cmd("SELECT", "1")
  //our universe state
  univ.InitUniverse(redis_client)
  //seed rand
  rand.Seed(time.Now().UTC().UnixNano())

  //game loop
  last_time := time.Now().UTC().UnixNano()
  delta_time := float64(time.Now().UTC().UnixNano() - last_time)

  //game loop
  one_second := 1000000000
  ticker := 0.0

  //web
  http.HandleFunc("/registerPlayer/", viewHandler)
  go http.ListenAndServe(":8000", nil)

  for {
    delta_time = float64(time.Now().UTC().UnixNano() - last_time) / float64(one_second)
    ticker += delta_time
    last_time = time.Now().UTC().UnixNano()

    univ.Update(delta_time, redis_client)
    update(delta_time, int64(ticker))
    if(int64(ticker) > 1){
      ticker = 0
    }
  }

}


func update(delta_time float64, ticker int64){
 
  if(ticker > 1){
    //asteroids_json, _ := json.Marshal(universe_update_message)
    asteroids_json := univ.Dump()
    pusher_client.Trigger("presence-test", "universe_update", asteroids_json)
  }

}