package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

// Message is a message from user
type Message struct {
	Content string

	once sync.Once
	sig  string // cached signature
}

// Signature returns the digital signature of the message
func (m *Message) Signature() string {
	m.once.Do(m.calcSig)
	return m.sig
}

func (m *Message) calcSig() {
	log.Printf("calculating signature")
	h := sha1.New()
	io.Copy(h, strings.NewReader(m.Content))
	m.sig = fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	m := Message{
		Content: "There is nothing more deceptive than an obvious fact.",
	}
	fmt.Println(m.Signature())
	// 2021/05/03 20:24:45 calculating signature
	// b931605bbbcdd058f9c33b11d7093fe8030b5413
	fmt.Println(m.Signature())
	// b931605bbbcdd058f9c33b11d7093fe8030b5413
}
