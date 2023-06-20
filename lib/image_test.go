package lib_test

import (
	"os"
	"testing"

	"github.com/miajio/www/lib"
)

func TestImageCompress(t *testing.T) {
	p, err := os.ReadFile("C:\\Users\\SnaroChrisXiao\\Desktop\\证书.png")
	if err != nil {
		t.Fatalf("read png file fail: %v", err)
	}

	o, err := lib.ImageCompress.Compress(p, 360, 485)
	if err != nil {
		t.Fatalf("compress file fail: %v", err)
	}

	of, _ := os.Create("C:\\Users\\SnaroChrisXiao\\Desktop\\域名1.png")
	_, err = of.Write(o)
	if err != nil {
		t.Fatalf("write file fail: %v", err)
	}

}
