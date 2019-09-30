package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		panic("ヘルスチェックするURLを引数にしてください")
	}

	url := flag.Arg(0)

	apiRes, apiErr := http.Get(
		url,
	)

	googleRes, err := http.Get(
		"https://google.com",
	)

	if err != nil {
		fmt.Print(err)
	}

	if apiErr != nil || (apiRes.StatusCode != 200 && googleRes.StatusCode == 200) {
		name := "API ヘルスチェッカー"
		text := "<@UHQRFR9KK>" + url + " が死んだ"
		channel := "server"

		jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + text + `"}`

		req, err := http.NewRequest(
			"POST",
			os.Getenv("HOOKS_URL"),
			bytes.NewBuffer([]byte(jsonStr)),
		)

		if err != nil {
			fmt.Print(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Print(err)
		}

		fmt.Print(resp)
		defer resp.Body.Close()
	}
	fmt.Print("success")
}
