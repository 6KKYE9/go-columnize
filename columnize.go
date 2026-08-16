package columnize

import "strings"

// Columnize 把若干行（每行用 sep 分隔成多列）按列对齐输出，列间用 gap 个空格隔开。
// 列宽取该列最宽一项。空行保留为空白行。
func Columnize(lines []string, sep string, gap int) string {
	if gap < 1 {
		gap = 1
	}
	pad := strings.Repeat(" ", gap)
	var rows [][]string
	maxCols := 0
	for _, ln := range lines {
		if strings.TrimSpace(ln) == "" {
			rows = append(rows, nil)
			continue
		}
		cols := strings.Split(ln, sep)
		rows = append(rows, cols)
		if len(cols) > maxCols {
			maxCols = len(cols)
		}
	}
	widths := make([]int, maxCols)
	for _, r := range rows {
		if r == nil {
			continue
		}
		for i, c := range r {
			if w := displayLen(c); w > widths[i] {
				widths[i] = w
			}
		}
	}
	var out []string
	for _, r := range rows {
		if r == nil {
			out = append(out, "")
			continue
		}
		var cells []string
		for i, c := range r {
			w := 0
			if i < len(widths) {
				w = widths[i]
			}
			// 最后一列不补尾空格，否则整行末尾会多出无意义的空白，
			// 既难看也影响和期望字符串比对。
			if i == len(r)-1 {
				cells = append(cells, c)
			} else {
				cells = append(cells, padTo(c, w))
			}
		}
		out = append(out, strings.Join(cells, pad))
	}
	return strings.Join(out, "\n")
}

func displayLen(s string) int {
	return len([]rune(s))
}

func padTo(s string, width int) string {
	gap := width - displayLen(s)
	if gap <= 0 {
		return s
	}
	return s + strings.Repeat(" ", gap)
}
