package ruota

type RSerializer interface {
	SerializeString(string) ([]byte, error)

	DeserializeString([]byte) (string, error)
	DeserializeList([][]byte, RType) ([]interface{}, error)
}

// WriteString

// WriteList
