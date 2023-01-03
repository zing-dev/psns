package psns

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	MaxCode = 9999
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

func (c Code) ToInt() int {
	code, _ := strconv.Atoi(string(c))
	return code
}

type Zh string

func (z Zh) String() string {
	return string(z)
}

// ProductSerialNumber 产品序列号组成
type ProductSerialNumber struct {
	CM     CMMap  `json:"-"`    // 产品大类
	C1     C1Map  `json:"-"`    // 一级分类
	C2     C2Map  `json:"-"`    // 二级分类
	C3     C3Map  `json:"-"`    // 三级分类
	C4     C4Map  `json:"-"`    // 四级分类
	Date   Date   `json:"date"` // 日期
	Code   Code   `json:"code"` // 当天编号
	ZhName string `json:"zh_name"`
	Name   string `json:"name"`
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
		return fmt.Errorf("当前的序列号的产品<%s>1类形态<%s>编码未注册", cm.ZhString(), c1)
	}
	ErrC2 = func(cm CM, c2 C2) error {
		return fmt.Errorf("当前的序列号的产品<%s>2类配置<%s>编码未注册", cm.ZhString(), c2)
	}
	ErrC3 = func(cm CM, c3 C3) error {
		return fmt.Errorf("当前的序列号的产品<%s>3类详细配置<%s>编码未注册", cm.ZhString(), c3)
	}
	ErrC4 = func(cm CM, c4 C4) error {
		return fmt.Errorf("当前的序列号的产品<%s>4类特殊码<%s>编码未注册", cm.ZhString(), c4)
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

func (p *ProductSerialNumber) Equal(p2 *ProductSerialNumber) bool {
	return p.CM.K == p2.CM.K &&
		p.C1.K == p2.C1.K &&
		p.C2.K == p2.C2.K &&
		p.C3.K == p2.C3.K &&
		p.C4.K == p2.C4.K &&
		p.Date == p2.Date &&
		p.Code == p2.Code
}
