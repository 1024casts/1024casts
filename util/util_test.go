package util

import (
	"testing"
)

func TestGenShortId(t *testing.T) {
	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		t.Error("GenShortId failed!")
	}

	t.Logf("GenShortId test pass, shortId: %s", shortId)
}

func TestEncodeUid(t *testing.T) {
	strUid := EncodeUid(123456)
	if strUid == "" {
		t.Error("EncodeUid failed")
	}

	t.Logf("EncodeUid test pass, strUid: %s", strUid)
}

//func TestDecodeUid(t *testing.T) {
//	intUid := DecodeUid("Q14WkgdBlyovD73WabD3Z72nRxbwzP")
//	if intUid == 0 {
//		t.Error("EncodeUid failed")
//	}
//
//	t.Logf("DecodeUid test pass, intUid: %d", intUid)
//}

func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer() //重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}
