package psns

const (
	Ezh Zh = ""
)

type CM string // 产品大类

type CMMap struct {
	K CM `json:"k"`
	V Zh `json:"v"`
}

type C1Map struct {
	K C1 `json:"k"`
	V Zh `json:"v"`
}

type C2Map struct {
	K C2 `json:"k"`
	V Zh `json:"v"`
}

type C3Map struct {
	K C3 `json:"k"`
	V Zh `json:"v"`
}
type C4Map struct {
	K C4 `json:"k"`
	V Zh `json:"v"`
}

func (c CM) ZhString() Zh {
	for _, cm := range cms {
		if c == cm.K {
			return cm.V
		}
	}
	return Ezh
}

func (c CM) String() string {
	return string(c)
}
