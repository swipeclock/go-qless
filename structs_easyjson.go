// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package qless

import (
	json "encoding/json"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjson6a975c40DecodeGithubComRyverappGoQless(in *jlexer.Lexer, out *Failure) {
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
		case "group":
			out.Group = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "when":
			out.When = int64(in.Int64())
		case "worker":
			out.Worker = string(in.String())
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
func easyjson6a975c40EncodeGithubComRyverappGoQless(out *jwriter.Writer, in Failure) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"group\":")
	out.String(string(in.Group))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"message\":")
	out.String(string(in.Message))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"when\":")
	out.Int64(int64(in.When))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"worker\":")
	out.String(string(in.Worker))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Failure) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComRyverappGoQless(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Failure) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComRyverappGoQless(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Failure) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComRyverappGoQless(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Failure) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComRyverappGoQless(l, v)
}
func easyjson6a975c40DecodeGithubComRyverappGoQless1(in *jlexer.Lexer, out *QueueInfo) {
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
		case "paused":
			out.Paused = bool(in.Bool())
		case "waiting":
			out.Waiting = int(in.Int())
		case "running":
			out.Running = int(in.Int())
		case "stalled":
			out.Stalled = int(in.Int())
		case "scheduled":
			out.Scheduled = int(in.Int())
		case "recurring":
			out.Recurring = int(in.Int())
		case "depends":
			out.Depends = int(in.Int())
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
func easyjson6a975c40EncodeGithubComRyverappGoQless1(out *jwriter.Writer, in QueueInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	out.String(string(in.Name))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"paused\":")
	out.Bool(bool(in.Paused))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"waiting\":")
	out.Int(int(in.Waiting))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"running\":")
	out.Int(int(in.Running))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"stalled\":")
	out.Int(int(in.Stalled))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"scheduled\":")
	out.Int(int(in.Scheduled))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"recurring\":")
	out.Int(int(in.Recurring))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"depends\":")
	out.Int(int(in.Depends))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v QueueInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComRyverappGoQless1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v QueueInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComRyverappGoQless1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *QueueInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComRyverappGoQless1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *QueueInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComRyverappGoQless1(l, v)
}
func easyjson6a975c40DecodeGithubComRyverappGoQless2(in *jlexer.Lexer, out *History) {
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
		case "when":
			out.When = int64(in.Int64())
		case "q":
			out.Queue = string(in.String())
		case "what":
			out.What = string(in.String())
		case "worker":
			out.Worker = string(in.String())
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
func easyjson6a975c40EncodeGithubComRyverappGoQless2(out *jwriter.Writer, in History) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"when\":")
	out.Int64(int64(in.When))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"q\":")
	out.String(string(in.Queue))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"what\":")
	out.String(string(in.What))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"worker\":")
	out.String(string(in.Worker))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v History) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComRyverappGoQless2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v History) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComRyverappGoQless2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *History) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComRyverappGoQless2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *History) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComRyverappGoQless2(l, v)
}
func easyjson6a975c40DecodeGithubComRyverappGoQless3(in *jlexer.Lexer, out *jobData) {
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
		case "jid":
			out.JID = string(in.String())
		case "klass":
			out.Class = string(in.String())
		case "state":
			out.State = string(in.String())
		case "queue":
			out.Queue = string(in.String())
		case "worker":
			out.Worker = string(in.String())
		case "tracked":
			out.Tracked = bool(in.Bool())
		case "priority":
			out.Priority = int(in.Int())
		case "expires":
			out.Expires = int64(in.Int64())
		case "retries":
			out.Retries = int(in.Int())
		case "remaining":
			out.Remaining = int(in.Int())
		case "data":
			(out.Data).UnmarshalEasyJSON(in)
		case "tags":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Tags).UnmarshalJSON(data))
			}
		case "history":
			if in.IsNull() {
				in.Skip()
				out.History = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.History = make([]History, 0, 1)
				} else {
					out.History = []History{}
				}
				for !in.IsDelim(']') {
					var v1 History
					(v1).UnmarshalEasyJSON(in)
					out.History = append(out.History, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "failure":
			if in.IsNull() {
				in.Skip()
				out.Failure = nil
			} else {
				if out.Failure == nil {
					out.Failure = new(Failure)
				}
				(*out.Failure).UnmarshalEasyJSON(in)
			}
		case "dependents":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Dependents).UnmarshalJSON(data))
			}
		case "dependencies":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Dependencies).UnmarshalJSON(data))
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
func easyjson6a975c40EncodeGithubComRyverappGoQless3(out *jwriter.Writer, in jobData) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"jid\":")
	out.String(string(in.JID))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"klass\":")
	out.String(string(in.Class))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"state\":")
	out.String(string(in.State))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"queue\":")
	out.String(string(in.Queue))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"worker\":")
	out.String(string(in.Worker))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tracked\":")
	out.Bool(bool(in.Tracked))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"priority\":")
	out.Int(int(in.Priority))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"expires\":")
	out.Int64(int64(in.Expires))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"retries\":")
	out.Int(int(in.Retries))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"remaining\":")
	out.Int(int(in.Remaining))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"data\":")
	(in.Data).MarshalEasyJSON(out)
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tags\":")
	if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Tags {
			if v2 > 0 {
				out.RawByte(',')
			}
			out.String(string(v3))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"history\":")
	if in.History == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v4, v5 := range in.History {
			if v4 > 0 {
				out.RawByte(',')
			}
			(v5).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
	if in.Failure != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"failure\":")
		if in.Failure == nil {
			out.RawString("null")
		} else {
			(*in.Failure).MarshalEasyJSON(out)
		}
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"dependents\":")
	if in.Dependents == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v6, v7 := range in.Dependents {
			if v6 > 0 {
				out.RawByte(',')
			}
			out.String(string(v7))
		}
		out.RawByte(']')
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"dependencies\":")
	if in.Dependencies == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in.Dependencies {
			if v8 > 0 {
				out.RawByte(',')
			}
			out.String(string(v9))
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v jobData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComRyverappGoQless3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v jobData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComRyverappGoQless3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *jobData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComRyverappGoQless3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *jobData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComRyverappGoQless3(l, v)
}
