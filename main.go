package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"strings"
)

var filename_map map[string]int // global downloaded html filename map

func fetch_and_search(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: get content error: url: " + url)
		return
	}
	defer resp.Body.Close()
	resp_content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error: readutil from content error: url: " + url)
	}
	resp_content_str := string(resp_content)
	//fmt.Printf("resp_content: ", resp_content_str)
	fmt.Println()
	mark_str := "a class=\"xref\""

	cnt := strings.Count(resp_content_str, mark_str) // add all sub-link into visiting list
	pos_begin := 0
	for i := 0; i < cnt; i++ {
		tmp_substr := resp_content_str[pos_begin : len(resp_content_str)-1]
		pos_begin += strings.Index(tmp_substr, mark_str)
		pos_begin += 21
		tmp_substr2 := resp_content_str[pos_begin : len(resp_content_str)-1]
		pos_end := strings.Index(tmp_substr2, "\"")
		ans := resp_content_str[pos_begin : pos_begin+pos_end]
		if !strings.HasPrefix(ans, "https://") && !strings.HasPrefix(ans, "/") { //exclude some annoying link
			fmt.Println("ans: " + ans)
			tmp_pos := strings.Index(ans, "?")
			filename := ans[:tmp_pos]
			fmt.Println("filename: " + filename)
			_, ok := filename_map[filename]
			if !ok {
				filename_map[filename] = 0
				fmt.Println("log: filename map added: " + filename)
				modify_and_save(resp_content_str, filename)

				fetch_and_search("https://docs.microsoft.com/en-us/dotnet/api/" + filename)

			} else {
				fmt.Println("log: filename map exist: " + filename)
			}
		}
		pos_begin += pos_end
	}

}

func modify_and_save(content string, name string) {
	m_str1 := "script src=\"/"
	m_str2 := "script src=\"https://docs.microsoft.com/"
	m_str3 := "<link rel=\"stylesheet\" href=\"/"
	m_str4 := "<link rel=\"stylesheet\" href=\"https://docs.microsoft.com/"
	content = strings.Replace(content, "?view=azure-dotnet", ".html", -1) // link structure
	content = strings.Replace(content, m_str1, m_str2, -1)                // js relative link correction
	content = strings.Replace(content, m_str3, m_str4, -1)                // css relative link correction

	r := strings.NewReader(content)
	g, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println("error: delete html node error")
	}
	g.Find("div.header-holder").Remove()
	g.Find("div.sidebar").Remove()
	g.Find("footer#footer").Remove()
	content, err = g.Html()
	if err != nil {
		fmt.Println("error: convert html node to string")
	}

	tmp := []byte(content)
	err = ioutil.WriteFile(name+".html", tmp, 777)
	if err != nil {
		fmt.Println("error: save file error. filename: " + name)
	}
}

func main() {
	fmt.Println("log: start")

	json_string := "https://docs.microsoft.com/api/apibrowser/dotnet/namespaces?moniker=azure-dotnet&api-version=0.2"
	resp_list, err := http.Get(json_string)
	if err != nil {
		fmt.Println("error: get json error")
	}
	defer resp_list.Body.Close()
	resp_list_result, err := ioutil.ReadAll(resp_list.Body)
	if err != nil {
		fmt.Println("error: readutil from json error")
	}
	fmt.Printf("log: json response: %s", resp_list_result)

	json_result, err := simplejson.NewJson(resp_list_result)
	if err != nil {
		fmt.Println("error: new json error")
	}

	json_result_array := json_result.Get("apiItems").MustArray()

	filename_map = make(map[string]int)

	for i, _ := range json_result_array {
		json_result_array_url := json_result_array[i].(map[string]interface{})
		fmt.Println(json_result_array_url["url"])
		fetch_and_search(json_result_array_url["url"].(string))
	}

}
