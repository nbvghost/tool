package queue

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Pool struct {
	lock      sync.RWMutex
	die       bool
	pools     *Pools
	Input     chan interface{}
	LastInput int64
	Hash      string
}

func (p *Pool) generatorHash() string {
	dest := [8]byte{}
	if _, err := rand.Read(dest[:]); err != nil {
		log.Panic(err)
	}
	p.Hash = time.Now().Format("20060102150405.999999999") + "." + hex.EncodeToString(dest[:])
	return p.Hash
}

func (p *Pool) IsDie() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.die
}
func (p *Pool) Die() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.die = true

}
func (p *Pool) DetectDie() {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			//poolsLen := len(pools)
			if ((time.Now().UnixNano() - p.LastInput) / 1000 / 1000) > int64(Params.PoolTimeOut) {
				if p.IsDie() {
					if len(p.Input) == 0 {
						close(p.Input)
						log.Println(fmt.Sprintf("删除Pool,Hash：%v  ChanLen:%v  ChanCap:%v  删除成功：%v", p.Hash, len(p.Input), cap(p.Input), p.pools.Remove(p)))
						return
					}

				} else {
					p.Die()
				}
			}
		}
	}
}
func (p *Pool) ReceiveMessageTo(c chan<- interface{}) {
	t := time.NewTicker(time.Second * 10)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			//poolsLen := len(pools)
			if ((time.Now().UnixNano() - p.LastInput) / 1000 / 1000) > int64(Params.PoolTimeOut) {
				if p.IsDie() {
					if len(p.Input) == 0 {
						//close(p.Input)
						log.Println(fmt.Sprintf("删除Pool,Hash：%v  ChanLen:%v  ChanCap:%v  删除成功：%v", p.Hash, len(p.Input), cap(p.Input), p.pools.Remove(p)))
						return
					}

				} else {
					p.Die()
				}
			}

		case msg := <-p.Input:
			c <- msg
		}
	}
}
func (p *Pool) CanInput() bool {
	if len(p.Input) == cap(p.Input) || p.IsDie() {
		return false
	} else {
		return true
	}
}
