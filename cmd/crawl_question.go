package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/anaskhan96/soup"
	"github.com/parnurzeal/gorequest"
)

const dir = "~/go/src/github.com/zhiruchen/500Algorithms/all"

func main() {
	runtime.GOMAXPROCS(8)
	url := "https://techiedelight.quora.com/500-Data-Structures-and-Algorithms-interview-questions-and-their-solutions?__filter__&__nsrc__=2&__snid3__=1594232728&amp&share=1"
	crawl(url)
}

// <p class="ui_qtext_para"><b>Array:</b></p>
func crawl(url string) {
	_, body, errs := gorequest.New().Get(url).End()
	if len(errs) > 0 {
		fmt.Println(errs)
		return
	}

	var pkgName, pkgPath string
	doc := soup.HTMLParse(body)
	ps := doc.FindAll("p", "class", "ui_qtext_para")

	for _, p := range ps {
		title := p.FindAll("b")
		if len(title) > 0 {
			titleText := title[0].Text()
			pkgName = strings.ToLower(strings.Replace(strings.Replace(titleText, ":", "", -1), " ", "", -1))
			pkgName = strings.Replace(pkgName, "&", "", -1)
			pkgPath = dir + "/" + pkgName
			if err := os.Mkdir(pkgPath, os.ModePerm); err != nil {
				fmt.Println(err)
				return
			}
			continue
		}

		if pkgName != "" && pkgPath != "" {
			var wg sync.WaitGroup
			// <span class="qlink_container"><a href="http://www.techiedelight.com/introduction-linked-lists/" rel="noopener nofollow" target="_blank" onclick="return Q.openUrl(this);" class="external_link" data-qt-tooltip="techiedelight.com">Introduction to Linked Lists</a></span>
			qs := p.FindAll("span", "class", "qlink_container")
			wg.Add(len(qs))
			for _, q := range qs {
				questionLink := q.FindAll("a")[0]
				go genGoFiles(pkgPath, pkgName, questionLink.Text(), questionLink.Attrs()["href"], &wg)
			}

			wg.Wait()
			pkgName, pkgPath = "", ""
		}
	}
}

func genGoFiles(pkgPath, pkgName, name, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	tmpl := `package %s

// %s %s
func %s() {

}

`
	tmplTest := `package %s
import (
	"testing"
)

func Test%s(t testing.T) {

}

`
	funcName := strings.Title(name)
	for _, syb := range []string{" ", "|", "-", "&", "(", ")", "/"} {
		funcName = strings.Replace(funcName, syb, "", -1)
	}
	funcName = strings.TrimRight(funcName, "_")

	fileName := strings.ToLower(name)
	for _, syb := range []string{" ", "|", "-", "&", "(", ")", "/"} {
		fileName = strings.Replace(fileName, syb, "_", -1)
	}
	fileName = strings.TrimRight(fileName, "_")

	content := fmt.Sprintf(tmpl, pkgName, funcName, url, funcName)
	if err := ioutil.WriteFile(pkgPath+"/"+fileName+".go", []byte(content), os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}

	testContent := fmt.Sprintf(tmplTest, pkgName, funcName)
	if err := ioutil.WriteFile(pkgPath+"/"+fileName+"_test.go", []byte(testContent), os.ModePerm); err != nil {
		fmt.Println(err)
		return
	}
}
