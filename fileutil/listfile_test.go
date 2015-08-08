package fileutil

import "testing"
import "reflect"
import "fmt"

func TestGetCurrentDirFiles(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		//结构体中包含[]string，初始化时在值得前面加上类型[]string，
		//否则出现missing type in composite literal
		{"D:\\lqbz\\test\\go", []string{".gitignore", "ips.go", "server.go"}},
	}

	for _, c := range cases {
		got, _ := GetCurrentDirFiles(c.in)
		//使用reflect.DeepEqual可进行数组、slice和结构体的比较
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("getCurrentDir(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGetCurrentDirAllFiles(t *testing.T) {
	dirname := "D:\\lqbz\\test\\go"
	filenames, _ := GetCurrentDirAllFiles(dirname, ".go")
	fmt.Println(len(filenames))
}
