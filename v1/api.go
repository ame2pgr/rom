package rom

// Content ...
type Content[kType comparable, vType any] struct {
	Key   kType
	Value vType
	Ok    bool
}

// Rom ...
type Rom[kType comparable, vType any] interface {
	Get(key kType) Content[kType, vType]
	GetBatch(keys ...kType) []Content[kType, vType]
	GetAll() []Content[kType, vType]
	Len() uint
}
