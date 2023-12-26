package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bahadrdsr/go-redis-cc/internal/store"
)

type CommandHandler struct {
	store *store.Store
}

func New(store *store.Store) *CommandHandler {
	return &CommandHandler{
		store: store,
	}
}

func (h *CommandHandler) HandleCommand(cmd string) string {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return "Invalid command"
	}

	switch parts[0] {
	case "PING":
		return "PONG"
	case "SET":
		if len(parts) != 3 {
			return "Usage: SET key value"
		}
		h.store.Set(parts[1], parts[2])
		return "OK"
	case "GET":
		if len(parts) != 2 {
			return "Usage: GET key"
		}
		value, ok := h.store.Get(parts[1])
		if !ok {
			return "Key not found"
		}
		return value
	case "INCR":
		if len(parts) != 2 {
			return "Usage: INCR key"
		}
		value, err := h.store.Incr(parts[1])
		if err != nil {
			return fmt.Sprintf("Error incrementing value: %v", err)
		}
		return strconv.Itoa(value)
	case "DECR":
		if len(parts) != 2 {
			return "Usage: DECR key"
		}
		value, err := h.store.Decr(parts[1])
		if err != nil {
			return fmt.Sprintf("Error decrementing value: %v", err)
		}
		return strconv.Itoa(value)
	case "DEL":
		if len(parts) != 2 {
			return "Usage: DEL key"
		}
		h.store.Del(parts[1])
		return "OK"
	case "KEYS":
		if len(parts) != 2 {
			return "Usage: KEYS pattern"
		}
		keys := h.store.Keys()
		return strings.Join(keys, "\n")
	case "FLUSH":
		h.store.Flush()
		return "OK"

	default:
		return "Unknown command"
	}
}
