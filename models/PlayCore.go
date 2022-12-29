package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

var AudioSourceFile string = "./tmp/"

type JSONData struct {
	Code int `json:"code"`
	Data struct {
		ID      string `json:"id"`
		Song    string `json:"song"`
		Sing    string `json:"sing"`
		Album   string `json:"album"`
		Cover   string `json:"cover"`
		URL     string `json:"url"`
		Name    string `json:"name"`
		Content string `json:"content"`
	} `json:"data"`
}

var a int = 0

func PlanCore() {
	go DOME1()
	time.Sleep(1000 * time.Second)
}

func Dowfile(durl string, SoneName string) {
	res, err := http.Get(durl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	f, err := os.Create(AudioSourceFile + SoneName + ".mp3")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)

}

func Music() {
	var API_info JSONData
	ApiInfo, err := http.Get("https://tenapi.cn/comment")
	if err != nil {
		fmt.Println("查询失败！")

	}
	defer ApiInfo.Body.Close()
	body, err := ioutil.ReadAll(ApiInfo.Body)
	json.Unmarshal(body, &API_info)
	Dowfile(string(API_info.Data.URL), string(API_info.Data.Song))
	if err != err {
		fmt.Println("cloa")
	}
	// time.Sleep(5 * time.Second)
	fmt.Println("----------------------------------")
	fmt.Println("正在播放：" + string(API_info.Data.Song))
	fmt.Println("歌手：" + string(API_info.Data.Sing))
	fmt.Println("网友：" + string(API_info.Data.Name))
	fmt.Println("优秀评论：" + string(API_info.Data.Content))
	fmt.Println("----------------------------------")
	ChikeMusic(API_info.Data.Song)
	defer PlayframeCore(string(API_info.Data.Song))

}
func ChikeMusic(FileName string) {
	cmd0 := exec.Command("file", AudioSourceFile+FileName+".mp3")
	var outputBuf1 bytes.Buffer
	cmd0.Stdout = &outputBuf1
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The first command can not be startup %s\n", err)
		return
	}
	if err := cmd0.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the second command: %s\n", err)
		return
	}

	cmd := outputBuf1.Bytes()
	undome := strings.Contains(string(cmd), "Audio")
	if !undome {
		fmt.Println("Server Error!")
		os.Remove(AudioSourceFile + FileName + ".mp3")
		os.Exit(0)

	}

}
func DOME1() {

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		keys := string(char)
		switch {
		case keys == "K" || keys == "k":
			go Music()
		// case keys == "p":

		case keys == "Q" || keys == "q":
			os.Exit(0)
		case keys == "S" || keys == "s":

		default:
			fmt.Println("请继续！！！！", keys)
			a++
			fmt.Println(a)
		}
	}
}
