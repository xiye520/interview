package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

type Lock struct {
	resource string
	token    string
	conn     redis.Conn
	timeout  int
}

func (lock *Lock) tryLock() (ok bool, err error) {
	_, err = redis.String(lock.conn.Do("SET", lock.key(), lock.token, "EX", int(lock.timeout), "NX"))
	if err == redis.ErrNil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
func (lock *Lock) Unlock() (err error) {
	_, err = lock.conn.Do("del", lock.key())
	return
}

func (lock *Lock) key() string {
	return fmt.Sprintf("redislock:%s", lock.resource)
}

func (lock *Lock) AddTimeout(ex_time int64) (ok bool, err error) {
	ttl_time, err := redis.Int64(lock.conn.Do("TTL", lock.key()))
	fmt.Println(ttl_time)
	if err != nil {
		log.Fatal(err)
	}
	if ttl_time > 0 {
		fmt.Println(11)
		_, err := redis.String(lock.conn.Do("SET", lock.key(), lock.token, "EX", int(ttl_time+ex_time)))
		if err == redis.ErrNil {
			return false, nil
		}
		if err != nil {
			return false, err
		}
	}
	return false, nil
}

func TryLock(conn redis.Conn, resource string, token string, DefaultTimeout int) (lock *Lock, ok bool, err error) {
	return TryLockWithTimeout(conn, resource, token, DefaultTimeout)
}

func TryLockWithTimeout(conn redis.Conn, resource string, token string, timeout int) (lock *Lock, ok bool, err error) {
	lock = &Lock{resource: resource, token: token, conn: conn, timeout: timeout}
	ok, err = lock.tryLock()
	if !ok || err != nil {
		lock = nil
	}
	return
}

func main() {
	fmt.Println("start")
	DefaultTimeout := 10
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	lock, ok, err := TryLock(conn, "anakin.sun", "token", int(DefaultTimeout))
	if err != nil {
		log.Fatal("error lock")
	}
	if !ok {
		log.Fatal("lock fail")
	}
	lock.AddTimeout(100)
	time.Sleep(time.Duration(DefaultTimeout) * time.Second)
	fmt.Println("end")
	defer lock.Unlock()
}
