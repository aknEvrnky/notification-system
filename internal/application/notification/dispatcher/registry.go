package dispatcher

import "errors"

var (
	registry         = make(map[string]EventHandler)
	ErrEventNotFound = errors.New("event handler not found")
)

func RegisterHandler(event string, handler EventHandler) {
	registry[event] = handler
}

func ResolveHandler(event string) (EventHandler, error) {
	if handler, exists := registry[event]; exists {
		return handler, nil
	}
	return nil, ErrEventNotFound
}
