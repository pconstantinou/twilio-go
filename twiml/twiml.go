//This is the file where all the twiml is getting converted
package main

import (
	"encoding/xml"
	"fmt"
	"github.com/beevik/etree"
)

//<?xml version="1.0" encoding="UTF-8"?>
//<Response>
	//<Message>
	//	<Body>Hello World!</Body>
	//</Message>
//</Response>



//        :param body: Message Body
//        :param to: Phone Number to send Message to
//        :param from: Phone Number to send Message from
//        :param action: Action URL
//        :param method: Action URL Method
//        :param status_callback: Status callback URL. Deprecated in favor of action.
//        :param kwargs: additional attributes
//        :returns: <Message> element


//create struct
type Noun struct{
	XMLName xml.Name `xml:"Message"`
	Body    string   `xml:"Body"`
	Media   string   `xml:"Media"`
}

func main() {
	twiml := &Noun{Body: "Hello world", Media: "http.hello.com"}
	out, _ := xml.MarshalIndent(twiml, " ", "  ")
	fmt.Println(xml.Header + string(out))
}

func nest(Body string, To string, From string, Action string, Method string) string{
	fmt.Println(Body)
	fmt.Println(To)
	fmt.Println(From)
	fmt.Println(Action)
	fmt.Println(Method)

	return Body+To+From
}

func BuildXml(){
	twiml := etree.NewDocument()
	twiml.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)

	//all XML must be nested under response (https://www.twilio.com/docs/voice/twiml#twiml-elements)
	response := twiml.CreateElement("Response")
	//from here we need to determine the verbs

	//from here we need determine the noun
}

