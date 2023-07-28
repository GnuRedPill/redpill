package utils

import (
	"fmt"
	"log"
	"net"
	"testing"
)

func TestAuth(t *testing.T) {
	arr, err := Encrypt([]byte("{'a':1,'b':2,'c':3}"), []byte("keysaisa"))
	if err != nil {
		t.Error(err)
	}
	brr, err := Decrypt(arr, []byte("keysaisa"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(brr))
}

func TestAes(t *testing.T) {

	a := AuthTool{}.Aes(1, "a", "b")
	t.Log(a)
	e := AuthTool{}.UnAes(a)
	t.Log(e)

}

func TestLog(t *testing.T) {
	dial, err := net.Dial("udp", "0.0.0.0:9999")
	if err != nil {
		return
	}
	defer dial.Close()
	log.SetOutput(dial)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("asdasdad")
}
