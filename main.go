package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	token := enc("20CSN0000251033")
	cekresi("POS", "20CSN0000251033", token)
}
func enc(noresi string) string {
	cmd := exec.Command("php", "enc.php", "-f", noresi)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return string(out)
}
func cekresi(kurir, noresi, enc string) {
	datax := `e=` + kurir + `&noresi=` + noresi + `&timers=` + enc
	client := &http.Client{}
	var data = strings.NewReader(datax)
	req, err := http.NewRequest("POST", "https://apix3.cekresi.com/cekresi/resi/initialize.php", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", "https://cekresi.com")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://cekresi.com/")
	req.Header.Set("Accept-Language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)
	document, _ := goquery.NewDocumentFromReader(strings.NewReader(string(bodyText)))
	document.Find("table tr").Each(func(idx int, rowhtml *goquery.Selection) {
		fmt.Println("Cell: ", idx, rowhtml.Text())
	})
}
