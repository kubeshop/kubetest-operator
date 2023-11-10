package bus

import (
	"sync"

	"github.com/kubeshop/testkube-operator/api/events/v1"
)

var _ Bus = (*EventBusMock)(nil)

type EventBusMock struct {
	events sync.Map
}

func NewEventBusMock() *EventBusMock {
	return &EventBusMock{}
}

func (b *EventBusMock) ListQueues() []string {
	var keys []string
	b.events.Range(func(key, value any) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}

func (b *EventBusMock) Publish(event events.Event) error {
	b.events.Range(func(key, e interface{}) bool {
		e.(chan events.Event) <- event
		return true
	})
	return nil
}
func (b *EventBusMock) Subscribe(queue string, handler Handler) error {

	ch := make(chan events.Event)
	b.events.Store(queue, ch)

	go func() {
		for e := range ch {
			handler(e)
		}
	}()
	return nil
}

func (b *EventBusMock) PublishTopic(topic string, event events.Event) error {
	b.events.Range(func(key, e interface{}) bool {
		e.(chan events.Event) <- event
		return true
	})
	return nil
}
func (b *EventBusMock) SubscribeTopic(topic, queue string, handler Handler) error {

	ch := make(chan events.Event)
	b.events.Store(queue, ch)

	go func() {
		for e := range ch {
			handler(e)
		}
	}()
	return nil
}

func (b *EventBusMock) Unsubscribe(queue string) error {
	b.events.Delete(queue)
	return nil
}

func (b *EventBusMock) Close() error {
	b.events.Range(func(key, e interface{}) bool {
		b.events.Delete(key)
		return true
	})
	return nil
}
