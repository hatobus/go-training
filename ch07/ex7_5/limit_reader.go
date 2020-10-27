package ex7_5

import (
	"io"
)

type LimitReader struct {
	R io.Reader
	N int64
}

func (lr *LimitReader) Read(p []byte) (int, error) {
	// Nが0以下だったらEOFを返す
	if lr.N <= 0 {
		return 0, io.EOF
	}

	// 指定したN以上の場合には読み込める残りのN byteまで読み出す
	if int64(len(p)) > lr.N {
		p = p[0:lr.N]
	} /*else { } ← 0 <= len(p) <= N の間なら何もしない*/

	n, err := lr.R.Read(p)
	if err != nil {
		return 0, err
	}

	// Nを更新して読み込める残りのバイト長を変える
	lr.N -= int64(n)

	return n, nil
}

func NewLimitReader(r io.Reader, n int64) io.Reader {
	return &LimitReader{r, n}
}
