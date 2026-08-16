# go-columnize

把"用分隔符隔开的行"按列对齐排好。CSV 看着歪歪扭扭的时候、或者想给日志里的多列数据做个表格感，这个就派上用场。每列取最宽那一项当宽度，列间用空格隔开，中文按一个字算。

空行原样保留，不会把表格从中间戳断。

## 装

```bash
go build -o columnize ./cmd/columnize
```

## 用

```bash
printf "name,age\nAlice,30\nBob,5\n" | ./columnize -sep , -gap 2
# name  age
# Alice 30
# Bob   5
```

参数：
- `-sep`：列分隔符，默认 `,`
- `-gap`：列间空格数，默认 2

## 当库用

```go
import "columnize"

columnize.Columnize([]string{"a,b", "cc,dd"}, ",", 1)
// "a  b\ncc dd"
```

## License

MIT
