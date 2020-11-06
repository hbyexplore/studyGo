package session

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

/*
查找sessionData
*/
func (s *SessionData) Get(key string) (value interface{}, err error) {
	//获取读锁
	s.RwLock.RLock()
	//准备释放读锁
	defer s.RwLock.RUnlock()
	value = s.Data[key]
	//没获取到value 就报错
	if value == nil || value == "" {
		err = fmt.Errorf("invalid Key")
		return nil, err
	}
	//获取到就返回value
	return value, nil
}

/*
增加/修改sessionData
*/
func (s *SessionData) Set(key string, value interface{}) {
	s.RwLock.Lock()
	defer s.RwLock.Unlock()
	s.Data[key] = value
}

/*
删除sessionData
*/
func (s *SessionData) Delete(key string) {
	s.RwLock.Lock()
	defer s.RwLock.Unlock()
	delete(s.Data, key)
}

/*
创建sessionData
*/
func (session *SessionMgr) CreateSessionData(id string) (sessionData *SessionData) {
	session.RwLock.Lock()
	defer session.RwLock.Unlock()
	sessionData = &SessionData{
		Id:   id,
		Data: make(map[string]interface{}, 10),
	}
	return
}

/*
查找session
*/
func (session *SessionMgr) Get(sessionId string) (sessionData SessionData) {
	session.RwLock.RLock()
	defer session.RwLock.RUnlock()
	sessionData = session.Session[sessionId]
	return sessionData
}

/*
创建session
*/
func (session *SessionMgr) Create(key string, value interface{}) string {
	session.RwLock.Lock()
	defer session.RwLock.Unlock()
	//生成一个uuid
	uuid := uuid.NewV4().String()
	//创建session
	sessionData := session.CreateSessionData(uuid)
	sessionData.Data[key] = value
	session.Session[uuid] = *sessionData
	return uuid
}

/*
删除session
*/
func (session *SessionMgr) Delete(id string) {
	session.RwLock.Lock()
	defer session.RwLock.Unlock()
	delete(session.Session, id)
}
