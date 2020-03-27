package main

import (
	"encoding/json"
	"fmt"
	"github.com/andymeneely/git-churn/matrics"
)

func main() {
	//dir,_:= os.Getwd()
	//fmt.Println(gitfuncs.LastCommit(dir))
	//t := gitfuncs.Tags("https://github.com/andymeneely/git-churn")
	//
	//for _, tag := range t {
	//	fmt.Println(tag)
	//}

	diffmetrics := metrics.CalculateDiffMetrics("https://github.com/andymeneely/git-churn", "00da33207bbb17a149d99301012006fbd86c80e4", "testdata/file.txt")
	fmt.Println(fmt.Sprintf("%v", diffmetrics))
	out, err := json.Marshal(diffmetrics)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
