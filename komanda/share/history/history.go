package history

import (
	"sync"

	"github.com/davecgh/go-spew/spew"
	"github.com/mephux/komanda-cli/komanda/logger"
)

// History struct
type History struct {
	Max   int
	Data  []string
	Index int

	mu sync.Mutex
}

// New history struct
func New() *History {
	return &History{
		Max:   3000,
		Data:  []string{},
		Index: 0,
	}
}

// Add line to history
func (h *History) Add(line string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	logger.Logger.Print("in Add\n")

	prev := h.Prev()

	if line == prev {
		return
	}

	if len(h.Data) > h.Max {
		h.Data = append(h.Data[:0], h.Data[1:]...)
	}

	h.Data = append(h.Data, line)
	h.Index = len(h.Data) - 1

	logger.Logger.Printf("ADD %s %d\n", h.Data, h.Index)
}

// HasLine will check is the history buffer contains the given string
func (h *History) HasLine(line string) bool {
	for _, l := range h.Data {
		if line == l {
			return true
		}
	}

	return false
}

// Get history at index
func (h *History) Get(index int) string {
	return h.Data[index]
}

// Empty history
func (h *History) Empty() bool {
	if len(h.Data) <= 0 {
		return true
	}

	return false
}

// Prev returns the previous history line fvrom the current index
func (h *History) Prev() string {
	logger.Logger.Printf("Prev from %d %d\n", h.Index, len(h.Data)-1)

	h.Index--

	if h.Index < 0 {
		h.Index = 0
	}

	// if h.Index == -1 {
	// h.Index = 0
	// } else if h.Index < -1 {
	// h.Index = len(h.Data) - 1
	// }

	logger.Logger.Printf("Set Prev Index %d\n", h.Index)

	if h.Empty() {
		return ""
	}

	logger.Logger.Printf("PREV %s\n", spew.Sdump(h.Data))

	return h.Data[h.Index]
}

// Next returns the history line after the current index
func (h *History) Next() string {
	logger.Logger.Print("Next\n")

	h.Index++

	if h.Index >= len(h.Data) {
		h.Index = len(h.Data) - 1
	}

	logger.Logger.Printf("Set Next Index %d\n", h.Index)

	if h.Empty() {
		return ""
	}

	logger.Logger.Printf("NEXT %s\n", h.Data[h.Index])

	return h.Data[h.Index]
}

// Current history line
func (h *History) Current() {
	h.Index = len(h.Data)
}
