package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/av1ppp/dada-ptizza_tg-bot/internal/config"
)

func main() {
	conf, err := config.ParseFile("config.yaml")
	if err != nil {
		panic(err)
	}

	// Send request
	form := url.Values{}
	form.Set("client_id", conf.YooMoney.ClientID)
	form.Set("response_type", "code")
	form.Set("redirect_uri", conf.YooMoney.RedirectURI)
	form.Set("scope", conf.YooMoney.Scope)

	formString := form.Encode()
	body := strings.NewReader(formString)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	request, err := http.NewRequest("POST", "https://yoomoney.ru/oauth/authorize", body)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Conten-Length", fmt.Sprint(len(formString)))

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	redirectURL := response.Header.Get("Location")

	err = exec.Command("xdg-open", redirectURL).Start()
	if err != nil {
		panic(err)
	}

	// Запрос кода
	fmt.Print("Code: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	code := scanner.Text()

	form = url.Values{}
	form.Set("code", code)
	form.Set("client_id", conf.YooMoney.ClientID)
	form.Set("grant_type", "authorization_code")

	formString = form.Encode()
	body = strings.NewReader(formString)

	request, err = http.NewRequest("POST", "https://yoomoney.ru/oauth/token", body)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Conten-Length", fmt.Sprint(len(formString)))

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(data))
}
