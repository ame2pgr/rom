package rom

// rom is an implementation of read-only map
type rom[key comparable, value any] struct {
	data map[key]value
}

// New ...
func New[kType comparable, vType any](data map[kType]vType) Rom[kType, vType] {
	return &rom[kType, vType]{data: data}
}

// Get ...
func (r *rom[kType, vType]) Get(key kType) Content[kType, vType] {
	value, ok := r.data[key]
	return Content[kType, vType]{Key: key, Value: value, Ok: ok}
}

// GetBatch ...
func (r *rom[kType, vType]) GetBatch(keys ...kType) []Content[kType, vType] {
	var contents = make([]Content[kType, vType], 0, len(keys))
	for _, key := range keys {
		value, ok := r.data[key]
		contents = append(contents, Content[kType, vType]{Key: key, Value: value, Ok: ok})
	}
	return contents
}

// GetAll...
func (r *rom[kType, vType]) GetAll() []Content[kType, vType] {
	var contents = make([]Content[kType, vType], 0, len(r.data))
	for key, value := range r.data {
		contents = append(contents, Content[kType, vType]{Key: key, Value: value, Ok: true})
	}
	return contents
}

// Len...
func (r *rom[kType, vType]) Len() uint {
	return uint(len(r.data))
}
