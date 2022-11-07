package psns

const (
	CMP   CM = "P" // 便携仪表
	CMPzh Zh = "便携仪表"

	CP101   C1 = "01" //
	CP101zh Zh = "运维仪"

	CP102   C1 = "02" //
	CP102zh Zh = "清管器/路由仪表"

	CP103   C1 = "03" //
	CP103zh Zh = "OTDR测量仪"

	CP201   C2 = "01" //
	CP201zh Zh = "单通道"

	CP202   C2 = "02" //
	CP202zh Zh = "双通道"

	CP301   C3 = "01" //
	CP301zh Zh = "常规灵敏度"

	CP302   C3 = "02" //
	CP302zh Zh = "高灵敏度"

	CP401   C4 = "01" //
	CP401zh Zh = "通用款"

	CP402   C4 = "02" //
	CP402zh Zh = "短矩离(单级放大)"

	CP403   C4 = "03" //
	CP403zh Zh = "定制款"
)

func init() {
	RegisterCM(CMMap{K: CMP, V: CMPzh})

	RegisterC1(CMP, C1Map{K: CP101, V: CP101zh}, C1Map{K: CP102, V: CP102zh}, C1Map{K: CP103, V: CP103zh})

	RegisterC2(CMP, C2Map{K: CP201, V: CP201zh}, C2Map{K: CP202, V: CP202zh})

	RegisterC3(CMP, C3Map{K: CP301, V: CP301zh}, C3Map{K: CP302, V: CP302zh})

	RegisterC4(CMP, C4Map{K: CP401, V: CP401zh}, C4Map{K: CP402, V: CP402zh}, C4Map{K: CP403, V: CP403zh})
}

func (c CM) IsP() bool {
	return c == CMP
}
