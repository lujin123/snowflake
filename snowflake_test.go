package snowflake

import (
	"fmt"
	"testing"
)

func TestSnowflake(t *testing.T) {
	snow := New(1)
	for i := 0; i < 100; i++ {
		fmt.Println(snow.NextID())
	}
}

func BenchmarkSnowflake(b *testing.B) {
	snow := New(1)
	for i := 0; i < b.N; i++ {
		snow.NextID()
	}
}
