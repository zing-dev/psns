package psns

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type C1 string // I类产品形态

func (c C1) String() string {
	return string(c)
}

type C2 string // II类配置相关

func (c C2) String() string {
	return string(c)
}

type C3 string // III类更多配置

func (c C3) String() string {
	return string(c)
}

type C4 string // 特殊码

func (c C4) String() string {
	return string(c)
}

type Date string // 日期

func (d Date) String() string {
	return string(d)
}

type Code string // 设备编码

type Zh string

func (z Zh) String() string {
	return string(z)
}

type ProductSerialNumber struct {
	CM   CMMap
	C1   C1Map
	C2   C2Map
	C3   C3Map
	C4   C4Map
	Date Date
	Code Code
}

var (
	cms []CMMap
	c1s = make(map[CM][]C1Map)
	c2s = make(map[CM][]C2Map)
	c3s = make(map[CM][]C3Map)
	c4s = make(map[CM][]C4Map)

	ErrSerialNumberLength = errors.New("序列号的长度必须为21个ASCII字符")
	ErrCM                 = func(cm CM) error { return fmt.Errorf("当前的序列号的产品大类<%s>未注册", cm) }
	ErrC1                 = func(cm CM, c1 C1) error {
		return fmt.Errorf("当前的序列号的产品<%s>1类形态<%s>未注册", cm.ZhString(), c1)
	}
	ErrC2 = func(cm CM, c2 C2) error {
		return fmt.Errorf("当前的序列号的产品<%s>2类配置<%s>未注册", cm.ZhString(), c2)
	}
	ErrC3 = func(cm CM, c3 C3) error {
		return fmt.Errorf("当前的序列号的产品<%s>3类详细配置<%s>未注册", cm.ZhString(), c3)
	}
	ErrC4 = func(cm CM, c4 C4) error {
		return fmt.Errorf("当前的序列号的产品<%s>4类特殊码<%s>未注册", cm.ZhString(), c4)
	}
	ErrDate = func(cm CM, date Date) error {
		return fmt.Errorf("当前的序列号的产品<%s>日期<%s>非法", cm.ZhString(), date)
	}
	ErrCode = func(cm CM, code Code) error {
		return fmt.Errorf("当前的序列号的产品<%s>编号<%s>非法", cm.ZhString(), code)
	}
)

func RegisterCM(cm CMMap) {
	cms = append(cms, cm)
}

func RegisterC1(cm CM, c1 ...C1Map) {
	c1s[cm] = c1
}

func RegisterC2(cm CM, c2 ...C2Map) {
	c2s[cm] = c2
}

func RegisterC3(cm CM, c3 ...C3Map) {
	c3s[cm] = c3
}

func RegisterC4(cm CM, c4 ...C4Map) {
	c4s[cm] = c4
}

func (p *ProductSerialNumber) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", p.CM.K, p.C1.K, p.C2.K, p.C3.K, p.C4.K, p.Date, p.Code)
}

func (p *ProductSerialNumber) ZhString() string {
	return fmt.Sprintf("%s%s%s%s%s", p.CM.V, p.C1.V, p.C2.V, p.C3.V, p.C4.V)
}

func Parse(serialNumber string) (*ProductSerialNumber, error) {
	if len(serialNumber) != 21 {
		return nil, ErrSerialNumberLength
	}
	var (
		cm   *CMMap
		c1   *C1Map
		c2   *C2Map
		c3   *C3Map
		c4   *C4Map
		date Date
		code Code
	)

	v := serialNumber[:1]
	for _, value := range cms {
		if strings.HasPrefix(v, value.K.String()) {
			cm = &value
			break
		}
	}
	if cm == nil {
		return nil, ErrCM(CM(serialNumber[:1]))
	}

	v = serialNumber[1:3]
	data1, ok := c1s[cm.K]
	if !ok {
		return nil, ErrCM(cm.K)
	}
	for _, value := range data1 {
		if strings.HasPrefix(serialNumber[1:], value.K.String()) {
			c1 = &value
			break
		}
	}
	if c1 == nil {
		return nil, ErrC1(cm.K, C1(v))
	}

	v = serialNumber[3:5]
	data2, ok := c2s[cm.K]
	if !ok {
		return nil, ErrCM(cm.K)
	}
	for _, value := range data2 {
		if strings.HasPrefix(v, value.K.String()) {
			c2 = &value
			break
		}
	}
	if c2 == nil {
		return nil, ErrC2(cm.K, C2(v))
	}

	v = serialNumber[5:7]
	data3, ok := c3s[cm.K]
	if !ok {
		return nil, ErrCM(cm.K)
	}
	for _, value := range data3 {
		if strings.HasPrefix(v, value.K.String()) {
			c3 = &value
			break
		}
	}
	if c3 == nil {
		return nil, ErrC3(cm.K, C3(v))
	}

	v = serialNumber[7:9]
	data4, ok := c4s[cm.K]
	if !ok {
		return nil, ErrCM(cm.K)
	}
	for _, value := range data4 {
		if strings.HasPrefix(v, value.K.String()) {
			c4 = &value
			break
		}
	}
	if c4 == nil {
		return nil, ErrC4(cm.K, C4(v))
	}

	v = serialNumber[9:17]
	_, err := time.ParseInLocation("20060102", v, time.Local)
	if err != nil {
		return nil, ErrDate(cm.K, Date(v))
	}
	date = Date(v)

	v = serialNumber[17:21]
	_, err = strconv.Atoi(v)
	if err != nil {
		return nil, ErrCode(cm.K, Code(v))
	}
	code = Code(v)

	return &ProductSerialNumber{
		CM:   *cm,
		C1:   *c1,
		C2:   *c2,
		C3:   *c3,
		C4:   *c4,
		Date: date,
		Code: code,
	}, nil
}
