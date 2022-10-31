package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		},
		{
			name:    "Kurniawan",
			request: "Kurniawan",
		},
	}
	for _, benchmarks := range benchmarks {
		b.Run(benchmarks.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmarks.request)
			}
		})
	}

}

//benchmark
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}
func TestTableHelloWorld(t *testing.T) {
	test := []struct {
		name     string
		request  string
		expected string
	}{{
		name:     "Eko",
		request:  "Eko",
		expected: "Hello Eko",
	}, {
		name:     "Kurniawan",
		request:  "Kurniawan",
		expected: "Hello Kurniawan",
	}}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE unit test")
	m.Run()
	//after
	fmt.Println("AFTER unit test")
}
func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})
	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		assert.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kurniawan'")
	})
}

//testing akan dilewati pada sebuah kondisi
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Tes tidak bisa berjalan di linux")
	}
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Test HelloWorld with assert Done")
}

//assert jika gagal akan memanggil fail
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Test HelloWorld with assert Done")
}

//require jika gagal akan memanggil failnow
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Test HelloWorld with require Done")
}
func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		t.Error("Result must be 'Hello Eko'") //kode di bawahnya tetap di jalankan
	}
	fmt.Println("tes hello eko selesai")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")
	if result != "Hello Khannedy" {
		t.Fatal("Result must be 'Hello Khannedy'") //kode di bawahnya tidak akan dijalankan
	}

	fmt.Println("Tes Hello khannedy selesai")
}

func TestHelloWorld(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, HelloWorld(tt.args.name), "HelloWorld(%v)", tt.args.name)
		})
	}
}
