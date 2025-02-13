package maths

const (
	_b  = 1
	_kb = _b * 1024
	_mb = _kb * 1024
	_gb = _mb * 1024
	_tb = _gb * 1024
	_pb = _tb * 1024
	_eb = _pb * 1024
)

var Byte = struct {
	B  int64
	KB int64
	MB int64
	GB int64
	TB int64
	PB int64
	EB int64
}{
	B:  _b,
	KB: _kb,
	MB: _mb,
	GB: _gb,
	TB: _tb,
	PB: _pb,
	EB: _eb,
}
