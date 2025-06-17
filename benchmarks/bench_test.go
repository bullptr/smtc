package benchmarks

import (
	"testing"
	// ap "github.com/smtx/smtx_parser"
)

// func BenchmarkTypeChecking(b *testing.B) {
// 	testFile := "../tests/hello.go"
// 	b.Run("TypeChecker", func(b *testing.B) {
// 		gosf := compiler.BuildGoSourceFile(testFile)
// 		b.ReportAllocs()

// 		for b.Loop() {
// 			types.CheckSourceFile(gosf)
// 		}
// 	})

// 	b.Run("TypeCheckerPlusSourceBuilder", func(b *testing.B) {
// 		b.ReportAllocs()

// 		for b.Loop() {
// 			gosf := compiler.BuildGoSourceFile(testFile)
// 			types.CheckSourceFile(gosf)
// 		}
// 	})
// }

func BenchmarkParsers(b *testing.B) {
	// smallTest := "../tests/test-002.smt2"
	// largeTest := "../tests/_large.smt2"

	// b.Run("TreeSitter Small", func(b *testing.B) {
	// 	b.ReportAllocs()

	// 	for b.Loop() {
	// 		ts.ParseFile(smallTest)
	// 	}
	// })

	// b.Run("Antlr Small", func(b *testing.B) {
	// 	b.ReportAllocs()

	// 	for b.Loop() {
	// 		ap.ParseFile(smallTest)
	// 	}
	// })

	// b.Run("TreeSitter Large", func(b *testing.B) {
	// 	b.ReportAllocs()

	// 	for b.Loop() {
	// 		ts.ParseFile(largeTest)
	// 	}
	// })

	// b.Run("Antlr Large", func(b *testing.B) {
	// 	b.ReportAllocs()

	// 	for b.Loop() {
	// 		ap.ParseFile(largeTest)
	// 	}
	// })

	// b.Run("Participle", func(b *testing.B) {
	// 	parser := p.NewParticipleParser
	// 	b.ReportAllocs()

	// 	for b.Loop() {
	// 		compiler.BuildSourceFileList(filenames, &parser)
	// 	}
	// })

}
