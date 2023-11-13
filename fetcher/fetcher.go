package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) (string, error) {
	//<-rateLimiter用于延时，防止网站监测
	/*
		// 指定目标网站的URL
			targetURL := "https://www.lfgvip.com/"

			// 创建一个HTTP客户端
			client := &http.Client{}

			// 创建一个GET请求
			req, err := http.NewRequest("GET", targetURL, nil)
			if err != nil {
				fmt.Println("创建请求时发生错误:", err)
				return
			}

			// 添加Cookie到请求头
			cookie := &http.Cookie{
				Name:  "PHPSESSID",
				Value: "fgiql7jdg518o4o4aom3v5jiu5",
			}
			req.AddCookie(cookie)

			// 发送请求
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("发送请求时发生错误:", err)
				return
			}
			defer resp.Body.Close()

			// 检查响应状态码
			if resp.StatusCode == http.StatusOK {
				// 读取响应内容
				// 这里可以根据需要处理响应内容
				// 例如，将其解析为HTML或其他数据
				fmt.Println("请求成功！")
			} else {
				fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode)
			}
	*/
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code is wrong: %d", resp.StatusCode)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
