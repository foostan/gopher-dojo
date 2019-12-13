package try

type Stringer interface {
	String() string
}

type CnvErr string

func (e CnvErr) Error() string {
	return string(e)
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, CnvErr("Can not convert")
}
