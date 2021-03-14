`log.txt`
```
2 this is a test
3 Are you like awk
This's a test
10 There are orange,apple,mongo
```

# 每行按空格或TAB分割，输出文本中的1、4项
```bash
awk '{print $1,$4}' log.txt
```
```
2 a
3 like
This's
10 orange,apple,mongo
```

# 格式化输出
```shell
awk '{printf "%-8s %-10s\n",$1,$4}' log.txt
```
```
2        a
3        like
This's
10       orange,apple,mongo
```
