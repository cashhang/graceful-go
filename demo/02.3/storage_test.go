package storage

import (
	"github.com/go-redis/redis"
	"github.com/golang/mock/gomock"
	"github.com/liyue201/graceful-go/demo/02.3/mock"
	"testing"
	"time"
)   
    
     
func Test_RedisStorage(t *testing.T)  {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	client := mock.NewMockUniversalClient(mockCtl)

	key := "aaaaa"
	val := "bbbbb" 

	expectSetResult := redis.NewStatusResult(val, nil)
	client.EXPECT().Set(key, val, time.Second).Return(expectSetResult)

	expectGetResult  := redis.NewStringResult(val, nil)
	client.EXPECT().Get(key).Return(expectGetResult)

	storage := NewRedisStorage(client)

	err := storage.Set(key, val, time.Second)
	if err != nil{
		t.Error(err)
		return
	}

	ret, err :=  storage.Get(key)
	if err != nil{
		t.Error(err)
		return
	}

	t.Logf("ret=%v", ret)

	if ret != val {
		t.Errorf("expect %v, but got %v",val, ret)
	}
}

