package ws_conn

import "sync"

var manager sync.Map

func store(deviceId int64, ctx *WSConnContext) {
	manager.Store(deviceId, ctx)
}

// 获取websocket
func load(deviceId int64) *WSConnContext {
	value, ok := manager.Load(deviceId)
	if ok {
		return value.(*WSConnContext)
	}
	return nil
}

// 删除
func delete(deviceId int64) {
	manager.Delete(deviceId)
}
