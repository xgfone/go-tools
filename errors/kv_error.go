// Copyright 2019 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package errors supplies an error type implementation based on
// the type inheritance.
package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// KV represents a key-value pair.
type KV interface {
	Key() string
	Value() interface{}
}

// NewKV returns a new KV instance.
func NewKV(key string, value interface{}) KV {
	return kV{key: key, value: value}
}

type kV struct {
	key   string
	value interface{}
}

func (kv kV) Key() string {
	return kv.key
}

func (kv kV) Value() interface{} {
	return kv.value
}

// KvError is an error with key-value contexts.
type KvError struct {
	Kvs []KV
	Msg string
	Err error
}

// NewKvError returns a new KvError.
//
// Example:
//   NewKvError(err)
//   NewKvError(err, "key1", value1, "key2", value2, ...)
func NewKvError(err error, kvs ...interface{}) KvError {
	_len := len(kvs)
	_kvs := make([]KV, 0, _len/2)
	for i := 0; i < _len; i += 2 {
		_kvs = append(_kvs, kV{key: kvs[i].(string), value: kvs[i+1]})
	}
	return KvError{Kvs: _kvs, Err: err}
}

// NewKvErrorByMsg returns a new KvError.
//
// Example:
//   NewKvErrorByMsg("msg")
//   NewKvErrorByMsg("msg", "key1", value1, "key2", value2, ...)
func NewKvErrorByMsg(msg string, kvs ...interface{}) KvError {
	_len := len(kvs)
	_kvs := make([]KV, 0, _len/2)
	for i := 0; i < _len; i += 2 {
		_kvs = append(_kvs, kV{key: kvs[i].(string), value: kvs[i+1]})
	}
	return KvError{Kvs: _kvs, Msg: msg}
}

// NewKvErrorByKv returns a new KvError.
//
// Example:
//   NewKvErrorByKv("msg", err, NewKV("key1", value1), NewKV("key2", value2))
func NewKvErrorByKv(msg string, err error, kvs ...KV) KvError {
	return KvError{Msg: msg, Err: err, Kvs: kvs}
}

// Error implements the error interface.
func (kve KvError) Error() string {
	buf := bytes.NewBuffer(nil)
	buf.Grow(128)
	kve.WriteTo(buf)
	return buf.String()
}

// MarshalJSON implements json.Marshaler.
func (kve KvError) MarshalJSON() ([]byte, error) {
	ms := make(map[string]interface{}, len(kve.Kvs)*2+4)
	for _, kv := range kve.Kvs {
		ms[kv.Key()] = kv.Value()
	}
	if kve.Msg != "" {
		ms["msg"] = kve.Msg
	}
	if kve.Err != nil {
		ms["err"] = kve.Err.Error()
	}
	return json.Marshal(ms)
}

// MarshalText implements encoding.TextMarshaler.
func (kve KvError) MarshalText() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.Grow(128)
	kve.WriteTo(buf)
	return buf.Bytes(), nil
}

// WriteTo implements io.WriterTo.
func (kve KvError) WriteTo(w io.Writer) (n int64, err error) {
	var m int
	for _, kv := range kve.Kvs {
		if m, err = io.WriteString(w, kv.Key()); err != nil {
			return n + int64(m), err
		}
		n += int64(m)

		if m, err = io.WriteString(w, "="); err != nil {
			return n + int64(m), err
		}
		n += int64(m)

		switch v := kv.Value().(type) {
		case io.WriterTo:
			m, err := v.WriteTo(w)
			if err != nil {
				return n + m, err
			}
			n += m
		case fmt.Stringer:
			if m, err = io.WriteString(w, v.String()); err != nil {
				return n + int64(m), err
			}
			n += int64(m)
		default:
			if m, err = fmt.Fprintf(w, "%+v", kv.Value()); err != nil {
				return n + int64(m), err
			}
			n += int64(m)
		}
	}

	if kve.Msg != "" {
		if m, err = io.WriteString(w, "msg="); err != nil {
			return n + int64(m), err
		}
		n += int64(m)

		if m, err = io.WriteString(w, kve.Msg); err != nil {
			return n + int64(m), err
		}
		n += int64(m)
	}

	if kve.Err != nil {
		if m, err = io.WriteString(w, "err="); err != nil {
			return n + int64(m), err
		}
		n += int64(m)

		if m, err = io.WriteString(w, err.Error()); err != nil {
			return n + int64(m), err
		}
		n += int64(m)
	}

	return
}
