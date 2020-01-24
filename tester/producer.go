package tester

import "github.com/lovoo/goka"

type producerMock struct {
	emitter EmitHandler
}

func newProducerMock(emitter EmitHandler) *producerMock {
	return &producerMock{
		emitter: emitter,
	}
}

// Emit emits messages to arbitrary topics.
// The mock simply forwards the emit to the KafkaMock which takes care of queueing calls
// to handled topics or putting the emitted messages in the emitted-messages-list
func (p *producerMock) Emit(topic string, key string, value []byte) *goka.Promise {
	return p.emitter(topic, key, value)
}

// Close closes the producer mock
// No action required in the mock.
func (p *producerMock) Close() error {
	logger.Printf("Closing producer mock")
	return nil
}

// flushingProducer wraps the producer and
// waits for all consumers after the Emit.
type flushingProducer struct {
	tester   *Tester
	producer goka.Producer
}

// Emit using the underlying producer
func (e *flushingProducer) Emit(topic string, key string, value []byte) *goka.Promise {
	prom := e.producer.Emit(topic, key, value)
	// e.tester.waitForConsumers()
	return prom
}

// Close using the underlying producer
func (e *flushingProducer) Close() error {
	return e.producer.Close()
}