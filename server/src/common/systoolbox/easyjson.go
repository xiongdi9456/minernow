// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package systoolbox

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC5a4559bDecodeCommonSystoolbox(in *jlexer.Lexer, out *TaskResult) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "app_name":
			out.AppName = string(in.String())
		case "region_id":
			out.RegionId = int64(in.Int64())
		case "server_id":
			out.ServerId = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "code":
			out.Code = int32(in.Int32())
		case "detail":
			out.Detail = string(in.String())
		case "begin_time":
			out.BeginTime = uint32(in.Uint32())
		case "end_time":
			out.EndTime = uint32(in.Uint32())
		case "ctime":
			out.Ctime = uint32(in.Uint32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeCommonSystoolbox(out *jwriter.Writer, in TaskResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"app_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AppName))
	}
	{
		const prefix string = ",\"region_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RegionId))
	}
	{
		const prefix string = ",\"server_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ServerId))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"code\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Code))
	}
	{
		const prefix string = ",\"detail\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Detail))
	}
	{
		const prefix string = ",\"begin_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.BeginTime))
	}
	{
		const prefix string = ",\"end_time\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.EndTime))
	}
	{
		const prefix string = ",\"ctime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.Ctime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TaskResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeCommonSystoolbox(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TaskResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeCommonSystoolbox(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TaskResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeCommonSystoolbox(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TaskResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeCommonSystoolbox(l, v)
}
func easyjsonC5a4559bDecodeCommonSystoolbox1(in *jlexer.Lexer, out *TaskMsgList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "list":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]TaskMsg, 0, 1)
					} else {
						out.List = []TaskMsg{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v1 TaskMsg
					(v1).UnmarshalEasyJSON(in)
					out.List = append(out.List, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeCommonSystoolbox1(out *jwriter.Writer, in TaskMsgList) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"list\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.List {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TaskMsgList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeCommonSystoolbox1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TaskMsgList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeCommonSystoolbox1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TaskMsgList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeCommonSystoolbox1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TaskMsgList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeCommonSystoolbox1(l, v)
}
func easyjsonC5a4559bDecodeCommonSystoolbox2(in *jlexer.Lexer, out *TaskMsg) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = uint32(in.Uint32())
		case "name":
			out.Name = string(in.String())
		case "spec":
			out.Spec = string(in.String())
		case "func_name":
			out.FuncName = string(in.String())
		case "switch":
			out.Switch = uint8(in.Uint8())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeCommonSystoolbox2(out *jwriter.Writer, in TaskMsg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint32(uint32(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Spec))
	}
	{
		const prefix string = ",\"func_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FuncName))
	}
	{
		const prefix string = ",\"switch\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint8(uint8(in.Switch))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TaskMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeCommonSystoolbox2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TaskMsg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeCommonSystoolbox2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TaskMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeCommonSystoolbox2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TaskMsg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeCommonSystoolbox2(l, v)
}
func easyjsonC5a4559bDecodeCommonSystoolbox3(in *jlexer.Lexer, out *TaskDetailItem) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "status":
			out.Status = int(in.Int())
		case "switch":
			out.Switch = int(in.Int())
		case "spec":
			out.Spec = string(in.String())
		case "prev":
			out.Prev = int64(in.Int64())
		case "next":
			out.Next = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeCommonSystoolbox3(out *jwriter.Writer, in TaskDetailItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"switch\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Switch))
	}
	{
		const prefix string = ",\"spec\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Spec))
	}
	{
		const prefix string = ",\"prev\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Prev))
	}
	{
		const prefix string = ",\"next\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Next))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TaskDetailItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeCommonSystoolbox3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TaskDetailItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeCommonSystoolbox3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TaskDetailItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeCommonSystoolbox3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TaskDetailItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeCommonSystoolbox3(l, v)
}
func easyjsonC5a4559bDecodeCommonSystoolbox4(in *jlexer.Lexer, out *TaskDetail) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "app_name":
			out.AppName = string(in.String())
		case "region_id":
			out.RegionId = int64(in.Int64())
		case "server_id":
			out.ServerId = int64(in.Int64())
		case "list":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]TaskDetailItem, 0, 1)
					} else {
						out.Items = []TaskDetailItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 TaskDetailItem
					(v4).UnmarshalEasyJSON(in)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC5a4559bEncodeCommonSystoolbox4(out *jwriter.Writer, in TaskDetail) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"app_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AppName))
	}
	{
		const prefix string = ",\"region_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.RegionId))
	}
	{
		const prefix string = ",\"server_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ServerId))
	}
	{
		const prefix string = ",\"list\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Items == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Items {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TaskDetail) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeCommonSystoolbox4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TaskDetail) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeCommonSystoolbox4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TaskDetail) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeCommonSystoolbox4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TaskDetail) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeCommonSystoolbox4(l, v)
}