package optstruct

type SimpleMail struct {
	internalMap map[string]string
}

func NewSimpleMail(options ...optStruct) *SimpleMail {
	m := &SimpleMail{
		internalMap: make(map[string]string),
	}
	for _, option := range options {
		option.Apply(m)
	}
	return m
}

func (m *SimpleMail) Add(key, value string) {
	m.internalMap[key] = value
}

type optStruct interface {
	Apply(m *SimpleMail)
}

func WithMessage(message string) optStruct {
	return &withMessage{message: message}
}

type withMessage struct{ message string }

func (w *withMessage) Apply(m *SimpleMail) {
	m.Add("message", w.message)
}

func WithSender(sender string) optStruct {
	return &withSender{sender: sender}
}

type withSender struct{ sender string }

func (w *withSender) Apply(m *SimpleMail) {
	m.Add("sender", w.sender)
}
