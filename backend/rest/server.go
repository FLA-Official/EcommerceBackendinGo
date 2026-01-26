package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {

	//creating object for Middleware Manager
	manager := middlewares.NewManager()
	// appending middlewares in globalMiddlware array
	manager.Use(
		middlewares.Cors,
		middlewares.Preflight,
		middlewares.Logger,
	)

	//after handling CORS and Preflight go to main business logic

	// declaring Router
	mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(int(cnf.HttpPort))

	fmt.Println("Server is running at the Port", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error loading the server", err)
		os.Exit(1)
	}
}

/*
here, the variable 'mux' works as Router. So we first declare the Router to access the functions inside http.

First NewServeMux, it doesn't take any input rather it enables us to create a struct name ServeMux.

Second HandleFunc, it takes a string which will consider as pattern, by which the client will able to see his/her website. It also takes another function which will act as Handler. In Handlefunc function, it is defined that the function will take http.ResponseWriter and *http.Request. which is why we created a function related to the page it will handle(fe. here we declared helloHandler & aboutHandler). in the function we declared w http.ResponseWriter, r *http.Request as the HandleFunc expects. Then we simply used Fprintln to print our messege into localhost but in future it may serve another purpose.

Third http.ListenandServe, it will take an address as String and handler(in this case it's "mux" I declared). This function will return an error. If error is "nil" then then the server will host smoothly otherwise it will return an error. To handle the error we used an if condition to understand the error and it will only trigger if there is a string inside error.

*/
