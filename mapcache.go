package mapcache

import (
	"context"
	"sync"
)

var GlobalMapMutex = &sync.RWMutex{}

var GlobalMap = make(map[interface{}]interface{})

// var GlobalMap = sync.Map{}

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

// func Load(ctx context.Context, pkg string, attribute interface{}) (interface{}, bool) {
// 	childMapItf, ok := GlobalMap.Load(pkg)
// 	if !ok {
// 		return nil, false
// 	}

// 	childMap, ok := childMapItf.(sync.Map)
// 	if !ok {
// 		return nil, false
// 	}

// 	val, ok := childMap.Load(attribute)
// 	if !ok {
// 		return nil, false
// 	}

// 	return val, true
// }

// func Save(ctx context.Context, pkg string, attribute interface{}, vrb interface{}) {
// 	GlobalMapMutex.Lock()

// 	x, _ := GlobalMap.LoadOrStore(pkg, sync.Map{})

// 	y, ok := x.(sync.Map)
// 	if ok {
// 		y.Store(attribute, vrb)
// 	}

// 	GlobalMap.Store(pkg, y)

// 	GlobalMapMutex.Unlock()
// }
