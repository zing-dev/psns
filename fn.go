package psns

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

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

	productSerialNumber := &ProductSerialNumber{
		CM:   *cm,
		C1:   *c1,
		C2:   *c2,
		C3:   *c3,
		C4:   *c4,
		Date: date,
		Code: code,
	}
	productSerialNumber.Name = productSerialNumber.String()
	productSerialNumber.ZhName = productSerialNumber.ZhString()
	return productSerialNumber, err
}

func Generate(cm CM, c1 C1, c2 C2, c3 C3, c4 C4, code int) (*ProductSerialNumber, error) {
	productSerialNumber := new(ProductSerialNumber)
	for _, value := range cms {
		if strings.HasPrefix(cm.String(), value.K.String()) {
			productSerialNumber.CM = value
		}
	}
	if productSerialNumber.CM.K == "" {
		return nil, ErrCM(cm)
	}

	if v, ok := c1s[cm]; ok {
		for _, value := range v {
			if strings.HasPrefix(c1.String(), value.K.String()) {
				productSerialNumber.C1 = value
			}
		}

		if productSerialNumber.C1.K == "" {
			return nil, ErrC1(cm, c1)
		}

	} else {
		return nil, ErrCM(cm)
	}

	if v, ok := c2s[cm]; ok {
		for _, value := range v {
			if strings.HasPrefix(c2.String(), value.K.String()) {
				productSerialNumber.C2 = value
			}
		}

		if productSerialNumber.C2.K == "" {
			return nil, ErrC2(cm, c2)
		}

	} else {
		return nil, ErrCM(cm)
	}

	if v, ok := c3s[cm]; ok {
		for _, value := range v {
			if strings.HasPrefix(c3.String(), value.K.String()) {
				productSerialNumber.C3 = value
			}
		}

		if productSerialNumber.C3.K == "" {
			return nil, ErrC3(cm, c3)
		}

	} else {
		return nil, ErrCM(cm)
	}

	if v, ok := c4s[cm]; ok {
		for _, value := range v {
			if strings.HasPrefix(c4.String(), value.K.String()) {
				productSerialNumber.C4 = value
			}
		}

		if productSerialNumber.C4.K == "" {
			return nil, ErrC4(cm, c4)
		}

	} else {
		return nil, ErrCM(cm)
	}

	productSerialNumber.Date = Date(time.Now().Format("20060102"))

	c := Code(fmt.Sprintf("%04d", code))
	if code > MaxCode {
		return nil, ErrCode(cm, c)
	}
	productSerialNumber.Code = c
	productSerialNumber.Name = productSerialNumber.String()
	productSerialNumber.ZhName = productSerialNumber.ZhString()
	return productSerialNumber, nil
}
