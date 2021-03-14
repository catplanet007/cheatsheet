# Usage:
```shell
ctt [-l] <arg> [<arg>...]
```
# Example:
```shell
# render ./go/sql.md
ctt go sql

# if only one <arg> given
ctt awk # render ./shell/awk.md
ctt go  # render ./go/index.md

# list all available cheatsheet
ctt -l

# list available go cheatsheet
ctt -l go
```