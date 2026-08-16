package columnize

import (
	"strings"
	"testing"
)

func TestColumnize(t *testing.T) {
	lines := []string{"name,age,city", "alice,30,beijing", "bob,5,sh"}
	got := Columnize(lines, ",", 1)
	// 列宽按该列最宽项对齐：name 列被 alice(5) 撑到 5 宽，所以首行 name 后留 2 空格。
	want := "name  age city\nalice 30  beijing\nbob   5   sh"
	if got != want {
		t.Errorf("Columnize=\n%q\nwant\n%q", got, want)
	}

	// 空行保留
	lines = []string{"a,b", "", "c,d"}
	got = Columnize(lines, ",", 1)
	if strings.Count(got, "\n") != 2 {
		t.Errorf("blank line not kept: %q", got)
	}
}
