package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
)

func init() {

	pool = &redis.Pool{}

	dataSourceName := fmt.Sprintf("%s:%d", "127.0.0.1", 6379)
	//connect
	pool = &redis.Pool{
		MaxIdle:     50,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			fmt.Println("redis connect " + dataSourceName)
			c, err := redis.Dial("tcp", dataSourceName)
			if err != nil {
				panic(err.Error())
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

}

func Set(rw http.ResponseWriter, req *http.Request) {
	//vars := mux.Vars(req)
	key := "test"
	conn := pool.Get()
	defer func() {
		conn.Close()
	}()
	if conn == nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	c, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("read err:", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = conn.Do("set", key, string(c))
	if err != nil {
		log.Println("set err:", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = rw.Write([]byte(""))
	if err != nil {
		log.Println("rw err:", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

}
