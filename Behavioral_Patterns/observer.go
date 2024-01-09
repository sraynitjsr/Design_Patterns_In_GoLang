package behavioral

import (
	"fmt"
	"sync"
)

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}

type Observer interface {
	Update(message string)
}

type ConcreteSubject struct {
	observers []Observer
	mu        sync.Mutex
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("Observer %s received message: %s\n", o.name, message)
}

func Observer() {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}
	
	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	subject.NotifyObservers("State change 1")

	subject.RemoveObserver(observer1)

	subject.NotifyObservers("State change 2")
}
