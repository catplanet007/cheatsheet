package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/glamour"
)

const basepath = "/home/aa/project/cheatsheet/data"

func realpath(rel string) string {
	return filepath.Join(basepath, rel)
}

func relpath(real string) string {
	rel, err := filepath.Rel(basepath, real)
	exitIfErr(err)
	return rel
}

func getRelativePath(args []string, completeIndex bool) string {
	switch len(args) {
	case 0:
		return ""
	case 1:
		fpath := fmt.Sprintf("%s.md", filepath.Join("shell", args[0]))
		if _, err := os.Stat(realpath(fpath)); err == nil {
			return fpath
		}
		if completeIndex {
			return fmt.Sprintf("%s.md", filepath.Join(args[0], "index"))
		}
		return args[0]
	default:
		return fmt.Sprintf("%s.md", filepath.Join(args...))
	}
}

func getContent(args []string) []byte {
	fpath := getRelativePath(args, true)
	buf, err := ioutil.ReadFile(realpath(fpath))
	if os.IsNotExist(err) {
		fmt.Println("err:", err)
		buf, err = ioutil.ReadFile(realpath(filepath.Join("ctt", "index.md")))
	}
	exitIfErr(err)
	return buf
}

func exitIfErr(e error) {
	if e != nil {
		fmt.Println("err:", e)
		os.Exit(1)
	}
}

func renderCheatsheet(args []string) {
	if len(args) == 0 {
		args = append(args, "ctt")
	}
	renderContent(getContent(args))
}

func renderContent(buf []byte) {
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
	)
	exitIfErr(err)
	out, err := r.RenderBytes(buf)
	exitIfErr(err)
	fmt.Print(string(out))
}

func listCheatsheet(args []string) {
	fpath := getRelativePath(args, false)
	buf := make([]byte, 0, 30)
	dirsDepth := make(map[string]int, 0)
	filepath.Walk(realpath(fpath), func(path string, info os.FileInfo, err error) error {
		exitIfErr(err)
		if path == basepath {
			return nil
		}
		if info.IsDir() {
			dir, _ := filepath.Split(path)
			dir = filepath.Clean(dir)
			depth, ok := dirsDepth[dir]
			if ok {
				depth++
				dirsDepth[path] = depth
			} else {
				dirsDepth[dir] = 0
				dirsDepth[path] = 1
				depth = 1
			}
			var s string
			for i := 1; i <= depth && i <= 6; i++ {
				s += "#"
			}
			s = fmt.Sprintf("%s %s\n\n", s, relpath(path))
			buf = append(buf, []byte(s)...)
			return nil
		}
		if strings.HasSuffix(path, ".md") {
			dir, filename := filepath.Split(path)
			dir = filepath.Clean(dir)
			s := fmt.Sprintf("> %s\n", filename[:len(filename)-3])
			buf = append(buf, []byte(s)...)
		}
		return nil
	})
	renderContent(buf)
}

func main() {
	opts := make(map[string]struct{})
	args := make([]string, 0, 2)
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			opts[arg] = struct{}{}
			continue
		}
		args = append(args, arg)
	}

	if _, ok := opts["-l"]; ok {
		listCheatsheet(args)
		os.Exit(0)
	}
	renderCheatsheet(args)
}
