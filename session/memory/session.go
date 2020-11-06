package session

import "sync"

/*
存放session数据的结构体
*/
type SessionData struct {
	Id     string                 //sessionId
	Data   map[string]interface{} //data
	RwLock sync.RWMutex           //读写锁保证获取的修改的线程安全
}

/*
存放 sessionId 和 sessionData 对应关系的结构体
*/
type SessionMgr struct {
	Session map[string]SessionData //SESSION
	RwLock  sync.RWMutex           //保证线程安全
}
