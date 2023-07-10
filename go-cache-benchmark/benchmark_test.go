package main

import (
	"go-cache-benchmark/databases"
	"testing"
)

var num = 1000

func BenchmarkPrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runDBOnly(databases.NewDBClient())
	}
}
