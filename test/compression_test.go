package azuretls

import (
	"strings"
	"testing"

	"github.com/DCH81/azuretls-client"
)

func TestDecompressBody_Gzip(t *testing.T) {
	session := azuretls.NewSession()

	session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept-Encoding", "gzip"},
	}

	response, err := session.Get("https://httpbin.org/gzip")

	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(response.Body), "\"gzipped\": true") {
		t.Fatal("TestDecompressBody_Gzip failed, expected: ", "\"gzipped\": true", ", got: ", string(response.Body))
	}
}

func tTestDecompressBody_Deflate(t *testing.T) {
	session := azuretls.NewSession()

	session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept-Encoding", "gzip, deflate, br"},
	}

	response, err := session.Get("https://httpbin.org/deflate")

	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(response.Body), "\"deflated\": true") {
		t.Fatal("TestDecompressBody_Deflate failed, expected: ", "\"deflated\": true", ", got: ", string(response.Body))
	}
}

func TestDecompressBody_Brotli(t *testing.T) {
	session := azuretls.NewSession()

	session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept-Encoding", "br"},
	}

	response, err := session.Get("https://httpbin.org/brotli")

	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(response.Body), "\"brotli\": true") {
		t.Fatal("TestDecompressBody_Brotli failed, expected: ", "\"brotli\": true", ", got: ", string(response.Body))
	}
}
