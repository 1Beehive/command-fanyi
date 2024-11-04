package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// ANSI 颜色代码
const (
	ColorReset = "\033[0m"
	ColorGreen = "\033[32m" // 绿色
	ColorRed   = "\033[31m" // 红色
	ColorBlue  = "\033[34m" // 蓝色
)

const (
	BaiduAPIURL = "https://fanyi-api.baidu.com/api/trans/vip/translate"
)

var AppID string
var SecretKey string

func translate(text string) (string, error) {
	salt := "1435660288" // 随机数，可以是任意字符串
	sign := generateSign(text, salt)
	escape := url2.QueryEscape(text)

	url := fmt.Sprintf("%s?q=%s&from=en&to=zh&appid=%s&salt=%s&sign=%s", BaiduAPIURL, escape, AppID, salt, sign)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		TransResult []struct {
			Dst string `json:"dst"`
		} `json:"trans_result"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.TransResult) > 0 {
		return result.TransResult[0].Dst, nil
	}
	return "", fmt.Errorf("翻译结果为空")
}

func generateSign(text string, salt string) string {
	str := fmt.Sprintf("%s%s%s%s", AppID, text, salt, SecretKey)
	hash := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", hash)
}

func getCommandOutput(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func printWithColor(line string) {
	// 使用正则表达式匹配 -- 及其后所有内容
	re := regexp.MustCompile(`(--\n)`)
	line = re.ReplaceAllStringFunc(line, func(s string) string {
		return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorReset)
	})

	// 打印最终的行
	fmt.Printf("%s%s%s\n", ColorGreen, line, ColorReset)
}

func main() {
	args := os.Args[1:] // 获取命令行参数
	if len(args) < 1 {
		fmt.Println("请提供要执行的命令.")
		return
	}

	AppID = os.Getenv("APPID")
	SecretKey = os.Getenv("SECRETKEY")
	if AppID == "" || SecretKey == "" {
		fmt.Println("APPID和SECRETKEY未设置\n")
		return
	}

	isTranslate := false
	if args[len(args)-1] == "fanyi" {
		isTranslate = true
		args = args[:len(args)-1] // 去掉最后的 fanyi 参数
	}

	command := strings.Join(args, " ") // 组合成完整命令

	output, err := getCommandOutput(command)
	if err != nil {
		fmt.Printf("命令执行错误: %s\n", err)
		return
	}

	lines := strings.Split(output, "\n")
	if isTranslate {
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				// 打印英文，添加颜色
				printWithColor(line)
				//fmt.Printf("%s%s%s\n", ColorGreen, line, ColorReset)

				// 保留前导空格
				leadingSpaces := strings.Repeat(" ", len(line)-len(strings.TrimLeft(line, " ")))

				// 翻译并打印中文
				translated, err := translate(line)
				if err != nil {
					fmt.Println("翻译错误:", err)
					continue
				}
				// 打印中文并与英文对齐
				fmt.Printf("%s%s%s%s\n", leadingSpaces, ColorRed, translated, ColorReset)
				fmt.Println("")
			}
		}
	} else {
		// 如果没有 fanyi 参数，按原样输出
		printWithColor(output)
	}
}
