package main

import (
	"github.com/beevik/etree"
	"os"
)

func main() {
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	response := doc.CreateElement("Response")

	message := response.CreateElement("Message")
	body := message.CreateElement("Body")
	body.CreateText("Test")
	media := message.CreateElement("Media")
	body.CreateText("Test")
	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}