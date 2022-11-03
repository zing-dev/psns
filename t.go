package psns

const (
	CMT   CM = "T" // 测温
	CMTzh Zh = "测温"

	CT101   C1 = "01" //
	CT101zh Zh = "小型DTS基础款"

	CT102   C1 = "02" //
	CT102zh Zh = "小型DTS升级款"

	CT103   C1 = "03" //
	CT103zh Zh = "机架式"

	CT201   C2 = "01" //
	CT201zh Zh = "单通道"

	CT204   C2 = "04" //
	CT204zh Zh = "四通道"

	CT208   C2 = "08" //
	CT208zh Zh = "八通道"

	CT301   C3 = "01" //
	CT301zh Zh = "单模"

	CT302   C3 = "02" //
	CT302zh Zh = "多模"

	CT401   C4 = "01" //
	CT401zh Zh = "通用款（2米分辨率）"

	CT402   C4 = "02" //
	CT402zh Zh = "高精（1米分辨）"

	CT403   C4 = "03" //
	CT403zh Zh = "超高精（0.5米分辨）"
)

func init() {
	RegisterCM(CMMap{K: CMT, V: CMTzh})

	RegisterC1(CMT, C1Map{K: CT101, V: CT101zh}, C1Map{K: CT102, V: CT102zh}, C1Map{K: CT103, V: CT103zh})

	RegisterC2(CMT, C2Map{K: CT201, V: CT201zh}, C2Map{K: CT204, V: CT204zh}, C2Map{K: CT208, V: CT208zh})

	RegisterC3(CMT, C3Map{K: CT301, V: CT301zh}, C3Map{K: CT302, V: CT302zh})

	RegisterC4(CMT, C4Map{K: CT401, V: CT401zh}, C4Map{K: CT402, V: CT402zh}, C4Map{K: CT403, V: CT403zh})
}
