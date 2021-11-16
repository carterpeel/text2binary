package main

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"text2binary/core"
	"time"
)

const (
	dictionary = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
)

var (
	conv = core.NewEncoder()
	rnd  = rand.New(rand.NewSource(time.Now().UnixNano()))
	data []byte
)

func init() {
	length, err := strconv.ParseInt(os.Getenv("LEN"), 10, 64)
	if err != nil {
		panic(err)
	}

	data = randBytes(length)
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = conv.Encode(data)
	}
}

// Helper functions
func randBytes(length int64) (ret []byte) {
	ret = make([]byte, length)
	for i := range ret {
		ret[i] = dictionary[rnd.Intn(len(dictionary))]
	}
	return
}