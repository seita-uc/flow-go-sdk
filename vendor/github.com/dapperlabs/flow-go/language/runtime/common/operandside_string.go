// Code generated by "stringer -type=OperandSide"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OperandSideUnknown-0]
	_ = x[OperandSideLeft-1]
	_ = x[OperandSideRight-2]
}

const _OperandSide_name = "OperandSideUnknownOperandSideLeftOperandSideRight"

var _OperandSide_index = [...]uint8{0, 18, 33, 49}

func (i OperandSide) String() string {
	if i < 0 || i >= OperandSide(len(_OperandSide_index)-1) {
		return "OperandSide(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperandSide_name[_OperandSide_index[i]:_OperandSide_index[i+1]]
}
