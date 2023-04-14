package websocket

type mediaServer struct {
	connection Socket
	name       string
}

type Socket struct {
}

// func getAnswer(socket Socket, description string)  {

// 	answer := ""
// 	socket.send(description)
// 	// socket.listen(msg -> {
// 	// 	answer = msg
// 	// })
// 	socket.getResponse()
// }

// func Socket ping(){
// 	socket.send("ping")
// 	// wait 10 servonds
// 	response = socket.receive()
// 	if (respinse == "pong" ){
// 		return
// 	}else{
// 		shutdown()
// 	}
// }

// func Socket connect(){
// 	socket.send("send jwt")
// 	// wait 10 servonds
// 	response = socket.receive()
// 	if (checkJwt(response) ){
// 		return
// 	}else{
// 		shutdown()
// 	}
// }
