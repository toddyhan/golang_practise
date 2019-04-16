package main

import (
	"os"
	"bufio"
	"io"
	"strings"
	"fmt"
)

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	originalMap := map[string]string{}
	oidResultMap := map[string]string{}
	result := []string{}
	lineNo :=0
	f, err := os.Open("oid.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		//fmt.Println(line)

		lineArry := strings.Split(strings.TrimSpace(line),"\t")

		originalMap[lineArry[0]] = lineArry[1]
		oidResultMap[lineArry[0]] = "0"
		if lineArry[2] != "0" {
			if strings.Contains(lineArry[2],"#") {
				ids := strings.Split(lineArry[2], "#")
				for _,id := range ids {
					oidResultMap[id] = lineArry[1]
				}
			} else {
				oidResultMap[lineArry[2]] = lineArry[1]
			}
		} else {
			oidResultMap[lineArry[0]] =  "0"
		}
		lineNo = lineNo + 1
	}
	f.Seek(0,0)
	//for k,v := range oidResultMap {
	//	fmt.Println(k,v)
	//}

	rd2 := bufio.NewReader(f)
	for {
		line, err := rd2.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		lineArry := strings.Split(strings.TrimSpace(line),"\t")
		oidNew,ok := oidResultMap[lineArry[0]]
		if !ok {
			line = strings.Split(line,"\n")[0] + "\t0"
		} else {
			line = strings.Split(line,"\n")[0] + "\t" + oidNew
		}
		result = append(result, line)

	}
	for _,v := range result {
		fmt.Println(v)
	}

	errW := writeLines(result,"new_oid.txt")
	if err != nil {
		fmt.Println(errW)
	}

}
