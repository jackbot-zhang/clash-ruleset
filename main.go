package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFile(url string, filepath string) error {
	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}
	defer resp.Body.Close()

	// 创建本地文件
	outFile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	// 将下载的数据写入本地文件
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	fmt.Println("File downloaded successfully!")
	return nil
}

func main() {
	urls := []string{
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/direct.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/proxy.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/reject.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/private.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/apple.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/icloud.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/google.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/gfw.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/tld-not-cn.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/telegramcidr.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/lancidr.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/cncidr.txt",
		"https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/applications.txt",
	}
	err := os.MkdirAll("ruleset", 0755)
	if err != nil {
		fmt.Println("Error creating folders:", err)
	} else {
		fmt.Println("Folders created successfully")
	}
	for _, url := range urls {
		arr := strings.Split(url, "/")
		arr1 := strings.Split(arr[len(arr)-1], ".")
		err := downloadFile(url, "ruleset/"+arr1[0]+".yaml")

		if err != nil {
			fmt.Println("Error:", err)
		}
	}

}
