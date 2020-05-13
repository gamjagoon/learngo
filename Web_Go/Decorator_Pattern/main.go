package main

import (
	"Deco/cipher2"
	"Deco/lzw"
	"fmt"
)

type Componet interface {
	Operator(string)
}

var sentData string
var receivData string

type SendComponet struct{}

func (self *SendComponet) Operator(data string) {
	// send data
	sentData = data
}

type ZipComponet struct {
	com Componet
}

func (self *ZipComponet) Operator(data string) {
	zipdata, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipdata))
}

type EncryptComponet struct {
	key string
	com Componet
}

func (self *EncryptComponet) Operator(data string) {
	Encodata, err := cipher2.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(Encodata))
}

type DecryptComponet struct {
	key string
	com Componet
}

func (self *DecryptComponet) Operator(data string) {
	decryptData, err := cipher2.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operator(string(decryptData))
}

type UnzipComponet struct {
	com Componet
}

func (self *UnzipComponet) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipData))
}

type ReadComponet struct{}

func (self *ReadComponet) Operator(data string) {
	receivData = data
}

func main() {
	sender := &EncryptComponet{key: "abcde",
		com: &ZipComponet{
			com: &SendComponet{}}}

	sender.Operator("Hello world")

	fmt.Println(sentData)

	receiver := &UnzipComponet{
		com: &DecryptComponet{
			key: "abcde",
			com: &ReadComponet{}}}
	receiver.Operator(sentData)
	fmt.Println(receivData)
}
