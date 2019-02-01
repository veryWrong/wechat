package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/veryWrong/wechat/wxsdk"
	"log"
	"net/http"
	"strings"
)

const token = "weChat"

var weChat = wxsdk.New(wxsdk.WxConfig{
	AppID:     "wxeb49f2b196ee4fe9",
	AppSecret: "8bb3cf212e78bd2eb2abdce2e9685d7d",
})

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/addKF", addKf)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	signature := strings.Join(r.Form["signature"], "")
	log.Println(signature)
	timestamp := strings.Join(r.Form["timestamp"], "")
	nonce := strings.Join(r.Form["nonce"], "")
	str := nonce + timestamp + token
	h := sha1.New()
	h.Write([]byte(str))
	hashed := hex.EncodeToString(h.Sum(nil))
	log.Println(hashed)
	if hashed == signature {
		fmt.Fprintf(w, strings.Join(r.Form["echostr"], ""))
	} else {
		fmt.Fprintf(w, "error")
	}
}

func addKf(w http.ResponseWriter, r *http.Request) {
	if err := weChat.AddKFAccount("test1@test", "客服1", "pswmd5"); err != nil {
		fmt.Fprintf(w, "error")
	} else {
		fmt.Fprintf(w, "success")
	}
}
