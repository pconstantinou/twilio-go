package main

// This doc will correspond to creating message twiml


//Create a <Message> element
//        :param body: Message Body
//        :param to: Phone Number to send Message to
//        :param from: Phone Number to send Message from
//        :param action: Action URL
//        :param method: Action URL Method
//        :param status_callback: Status callback URL. Deprecated in favor of action.
//        :param kwargs: additional attributes
//        :returns: <Message> element

//<?xml version="1.0" encoding="UTF-8"?>
//<Response>
//<Message><Body>Hello World!</Body></Message>
//</Response>
func Message(Body string, To string, From string, Action string, Method string) string{
/*
   Create a <Message> element
   :param body: Message Body
   :param to: Phone Number to send Message to
   :param from: Phone Number to send Message from
   :param action: Action URL
   :param method: Action URL Method
   :param status_callback: Status callback URL. Deprecated in favor of action.
   :param kwargs: additional attributes

   :returns: <Message> element
 */

	test := nest(Body, To, From, Action, Method)
	return test
}

func Redirect(Url string, Method string){
/*
	Create a <Redirect> element
   :param url: Redirect URL
   :param method: Redirect URL method
   :param kwargs: additional attributes

   :returns: <Redirect> element
 */

}
