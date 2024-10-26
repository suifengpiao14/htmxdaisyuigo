package attributes_test

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegexp(t *testing.T) {
	fd, err := os.OpenFile("./attributes.go", os.O_RDONLY, fs.ModePerm)
	require.NoError(t, err)
	b, err := io.ReadAll(fd)
	require.NoError(t, err)
	source := string(b)

	// 正则表达式匹配不带 `_` 后缀的函数
	regex := regexp.MustCompile(`(?s)func\s+\w+\([^)]*\)\s+Attribute\s*{.*?\n}\n`)

	// 删除匹配的函数
	result := regex.ReplaceAllStringFunc(source, func(m string) string {
		if !regexp.MustCompile(`func\s+\w+_`).MatchString(m) { // 判断是否没有 `_` 后缀
			return ""
		}
		return m
	})
	fmt.Println(result)

}
