package mapcache

import (
	"context"
	"errors"
	"sync"
	"time"
)

var GlobalMapMutex = &sync.RWMutex{}
var GlobalMap = make(map[interface{}]interface{})
var IsTtlInitialized = make(map[string]bool)

func Load(ctx context.Context, pkg string, attribute interface{}) (interface{}, bool) {
	GlobalMapMutex.RLock()
	childMapItf, ok := GlobalMap[pkg]
	if !ok {
		return nil, false
	}
	GlobalMapMutex.RUnlock()

	childMap, ok := childMapItf.(sync.Map)
	if !ok {
		return nil, false
	}

	val, ok := childMap.Load(attribute)
	if !ok {
		return nil, false
	}

	return val, true
}

func Save(ctx context.Context, pkg string, attribute interface{}, vrb interface{}) {
	GlobalMapMutex.Lock()
	x, ok := GlobalMap[pkg]
	if !ok {
		GlobalMap[pkg] = make(map[interface{}]interface{})
	}

	y, ok := x.(sync.Map)
	if !ok {
		y = sync.Map{}
	}

	y.Store(attribute, vrb)

	GlobalMap[pkg] = y

	GlobalMapMutex.Unlock()
}

func Delete(ctx context.Context, pkg string) {
	GlobalMapMutex.Lock()
	delete(GlobalMap, pkg)
	GlobalMapMutex.Unlock()
}

func InitTTL(ctx context.Context, pkg string, ttl int) error {
	if IsTtlInitialized[pkg] {
		return errors.New("TTL already initialized")
	}

	go func() {
		for true {
			time.Sleep(time.Duration(ttl) * time.Second)
			Delete(ctx, pkg)
		}
	}()

	return nil
}
