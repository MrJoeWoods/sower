// Code generated by "stringer -type=netType util.go"; DO NOT EDIT.

package proxy

import "strconv"

const _netType_name = "QUICKCPTCP"

var _netType_index = [...]uint8{0, 4, 7, 10}

func (i netType) String() string {
	if i < 0 || i >= netType(len(_netType_index)-1) {
		return "netType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _netType_name[_netType_index[i]:_netType_index[i+1]]
}
