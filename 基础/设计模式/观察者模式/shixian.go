package main

import "fmt"

type ConcreteObserver struct {
	Id int
}

func (co *ConcreteObserver) Update(e *Event) {
	fmt.Printf("observer [%d] recieved msg: %s.\n", co.Id, e.Data)
}

type ConcreteSubject struct {
	Observers map[Observer]struct{}
}

func (cs *ConcreteSubject) Regist(ob Observer) {
	cs.Observers[ob] = struct{}{}
}

func (cs *ConcreteSubject) Deregist(ob Observer) {
	delete(cs.Observers, ob)
}

// 通知每个观察者事件
func (cs *ConcreteSubject) Notify(e *Event) {
	for ob, _ := range cs.Observers {
		ob.Update(e)
	}
}
