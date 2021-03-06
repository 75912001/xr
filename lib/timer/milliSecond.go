package timer

import (
	"container/list"
	"time"
)

// Millisecond 毫秒级定时器
type Millisecond struct {
	Arg      interface{} //参数
	Function OnTimerFun  //超时调用的函数
	expire   int64       //过期时间戳
	valid    bool        //有效(false:不执行,扫描时自动删除)
}

// 判断是否有效
func (p *Millisecond) IsValid() bool {
	return p.valid
}

//AddMillisecond 添加毫秒级定时器
func (p *TimerMgr) AddMillisecond(cb OnTimerFun, arg interface{}, expireMillisecond int64) (t *Millisecond) {
	t = &Millisecond{
		Arg:      arg,
		Function: cb,
		expire:   expireMillisecond,
		valid:    true,
	}

	p.milliSecondMutex.Lock()
	defer func() {
		p.milliSecondMutex.Unlock()
	}()

	p.millisecondList.PushBack(t)
	return
}

//DelMillisecond 删除毫秒级定时器(必须与该timerOutChan线性处理.如:在同一个goroutine select中处理数据.)
func DelMillisecond(t *Millisecond) {
	t.inValid()
}

// 设为无效
func (p *Millisecond) inValid() {
	p.valid = false
	p.Arg = nil
	p.Function = nil
	p.expire = 0
}

// 扫描毫秒级定时器
func (p *TimerMgr) scanMillisecond() {
	t := time.Now()
	millisecond := t.UnixNano() / 1000000

	var next *list.Element
	for e := p.millisecondList.Front(); e != nil; e = next {
		timerMillisecond := e.Value.(*Millisecond)
		if !timerMillisecond.IsValid() {
			next = e.Next()
			p.millisecondList.Remove(e)
			continue
		}
		if timerMillisecond.expire <= millisecond {
			p.timerOutChan <- timerMillisecond
			next = e.Next()
			p.millisecondList.Remove(e)
		} else {
			next = e.Next()
		}
	}
}
