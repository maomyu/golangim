package sqlid

import (
	"database/sql"
	"errors"
	"time"
)

// ErrTimeOut 获取uid超时错误
var ErrTimeOut = errors.New("get uid timeout")

// 定义一个Uid
type Uid struct {
	db         *sql.DB
	businessId string
	ch         chan int64
	min, max   int64
}

// 新建一个Uid
func NewUid(db *sql.DB, businessId string, len int) (*Uid, error) {
	lid := Uid{
		db:         db,
		businessId: businessId,
		ch:         make(chan int64, len),
	}
	go lid.produceId()
	return &lid, nil
}

// 获得ID

func (u Uid) Get() (int64, error) {
	select {
	case <-time.After(1 * time.Second):
		return 0, ErrTimeOut
	case uid := <-u.ch:
		return uid, nil
	}
}

// 生产Id
// 当ch达到最大容量的时候，这个方法会阻塞，直到ch中的id被消耗
func (u *Uid) produceId() {
	// 从数据库中获取id
	u.reload()

	for {
		if u.min >= u.max {
			// 从数据库中获取id
			u.reload()
		}
		u.min++
		u.ch <- u.min
	}
}

// 从数据库中获取id段，如果失败，会每隔一秒尝试一次
func (u *Uid) reload() error {
	var err error
	for {
		err = u.getFromDB()
		if err == nil {
			return nil
		}

		// 如果获取失败，等待
		time.Sleep(time.Second)
	}
}

// 从数据库中获取id 段
func (u *Uid) getFromDB() error {
	var (
		maxId int64
		step  int64
	)
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//sql语句
	sqlquery := "select max_id,step from uid where business_id = ? FOR UPDATE "
	err = tx.QueryRow(sqlquery, u.businessId).Scan(&maxId, &step)
	if err != nil {
		return err
	}
	// 更新数据库中uid的最大值
	update := "update uid set max_id = ? where business_id = ?"
	_, err = tx.Exec(update, maxId+step, u.businessId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	u.min = maxId
	u.max = maxId + step
	return nil
}
