package telebot

// TType is an interface for all `telebot` types, every Type should implement TType.
type TType interface {
	Verify() error

	String() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error

	Type() string
	ReflectType() string
}
