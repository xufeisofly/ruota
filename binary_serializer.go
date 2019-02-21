package ruota

type RBinarySerializer struct {
}

func NewRBinarySerializer() (RSerializer, error) {
	return &RBinarySerializer{}, nil
}

func (p *RBinarySerializer) SerializeString(value string) []byte {
	return []byte(value)
}

func (p *RBinarySerializer) SerializeList(value []interface{}, elemType RType, size int32) [][]byte {
	var ret [][]byte

	switch elemType {
	case STRING:
		for _, item := range value {
			ret = append(ret, []byte(item.(string)))
		}
		break
	case INT8:
		break
	case INT16:
		break
	case INT32:
		break
	case BOOL:
		break
	default:
	}
	return ret
}

func (p *RBinarySerializer) DeserializeString(value []byte) (string, error) {
	return string(value), nil
}

func (p *RBinarySerializer) DeserializeList(value []interface{}, elemType RType) ([]string, error) {
	var ret []string

	switch elemType {
	case STRING:
		for _, item := range value {
			ret = append(ret, string(item.([]byte)))
		}
		break
	default:
	}

	return ret, nil
}
