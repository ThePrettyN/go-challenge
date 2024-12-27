// Exercise: rot13Reader
// よくあるパターンは、別の io.Reader をラップし、ストリームの内容を何らかの方法で変換するio.Readerです。

// 例えば、 gzip.NewReader は、 io.Reader (gzipされたデータストリーム)を引数で受け取り、 *gzip.Reader を返します。 その *gzip.Reader は、 io.Reader (展開したデータストリーム)を実装しています。

// io.Reader を実装し、 io.Reader でROT13 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用して読み出すように rot13Reader を実装してみてください。

// rot13Reader 型は提供済みです。 この Read メソッドを実装することで io.Reader インタフェースを満たしてください。

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r13.r.Read(p)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, nil
}

func rot13(b byte) byte {
	if 'a' <= b && b <= 'z' {
		return 'a' + (b-'a'+13)%26
	}
	if 'A' <= b && b <= 'Z' {
		return 'A' + (b-'A'+13)%26
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
