package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := json.NewEncoder(file)//NewEncoder 返回一个写入 w 的新编码器。
	err := enc.Encode(vc)//Encode 将 v 的 JSON 编码写入流中，
	if err != nil {
		log.Println("Error in encoding json")
	}
}

//解码流
//func NewDecoder(r io.Reader) *Decoder
//func (dec *Decoder) Decode(v interface{}) error

//编码
//func NewEncoder(w io.Writer) *Encoder
//func (enc *Encoder) Encode(v interface{})

