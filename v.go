package psns

const (
	CMV   CM = "V" // 测振
	CMVzh Zh = "测振"

	CV101   C1 = "01" //
	CV101zh Zh = "小型低功耗(MDVS)"

	CV102   C1 = "02" //
	CV102zh Zh = "机架式1U(RDVS)"

	CV103   C1 = "04" //
	CV103zh Zh = "机架式4U(RDVS)"

	CV201   C2 = "01" //
	CV201zh Zh = "单通道"

	CV202   C2 = "02" //
	CV202zh Zh = "双通道"

	CV204   C2 = "04" //
	CV204zh Zh = "四通道"

	CV208   C2 = "08" //
	CV208zh Zh = "八通道"

	CV301   C3 = "01" //
	CV301zh Zh = "常规灵敏度"

	CV302   C3 = "02" //
	CV302zh Zh = "高灵敏度"

	CV303   C3 = "03" //
	CV303zh Zh = "超高灵敏度"

	CV401   C4 = "01" //
	CV401zh Zh = "通用款"

	CV402   C4 = "02" //
	CV402zh Zh = "短矩离(单级放大)"

	CV403   C4 = "03" //
	CV403zh Zh = "加长距离(内置三级放大)"
)

func init() {
	RegisterCM(CMMap{K: CMV, V: CMVzh})

	RegisterC1(CMV, C1Map{K: CV101, V: CV101zh}, C1Map{K: CV102, V: CV102zh}, C1Map{K: CV103, V: CV103zh})

	RegisterC2(CMV, C2Map{K: CV201, V: CV201zh}, C2Map{K: CV202, V: CV202zh}, C2Map{K: CV204, V: CV204zh}, C2Map{K: CV208, V: CV208zh})

	RegisterC3(CMV, C3Map{K: CV301, V: CV301zh}, C3Map{K: CV302, V: CV302zh}, C3Map{K: CV303, V: CV303zh})

	RegisterC4(CMV, C4Map{K: CV401, V: CV401zh}, C4Map{K: CV402, V: CV402zh}, C4Map{K: CV403, V: CV403zh})
}

func (c CM) IsV() bool {
	return c == CMV
}
