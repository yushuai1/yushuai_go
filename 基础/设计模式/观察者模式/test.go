package main

import (
	"time"
	"fmt"
)

func main() {
	cs := &ConcreteSubject{
		Observers: make(map[Observer]struct{}),
	}

	//实例化两个观察者
	cobserver1 := &ConcreteObserver{1}
	cobserver2 := &ConcreteObserver{2}

	//注册观察者
	cs.Regist(cobserver1)
	cs.Regist(cobserver2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg [%d]", i)}
		cs.Notify(e)

		time.Sleep(time.Duration(1) * time.Second)
	}
}
