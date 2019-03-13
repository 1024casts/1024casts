package redis

import (
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	client, err := NewRedisClient()
	if err != nil {
		t.Errorf("new redis client err: %v", err)
		return
	}

	var setGetKey = "test-set"
	var setGetValue = "test-content"
	client.Set(setGetKey, setGetValue, time.Second*100)

	expectValue := client.Get(setGetKey).Val()
	if setGetValue != expectValue {
		t.Fail()
		return
	}

	t.Log("redis set test pass")
}
