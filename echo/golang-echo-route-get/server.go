package main

//* @author Hamid Poorhashemi
//* @Email Hamidpoorhashemi@gmail.com
//creted website by go lang using echo freamwork
//I used another languae Already and now :) php for website and webservices
// and try to Changing the way
// and it can help u ,if u are like me or not
// even if u don't work in website develping
// Good luck


import (
	"net/http"
	"github.com/labstack/echo"
)
//main function
func main() {
	e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.HTML(http.StatusOK, "<h1>Home page </h1><br> <a href='/about' > About us</a><br> <a href='/contact' > Contact us</a>")
		})

// Some Route for example
//This is not the right way, but for those who used to work on the website, it's very happy.
//but for those who programming in Languages like php ,Can hope
e.GET("/about", aboutpage)
e.GET("/contact", contactpage)


// server port
// U can open it in root address and port
// Example: http://localhost:8000
	e.Logger.Fatal(e.Start(":8000"))
}

func aboutpage(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>About page  </h1> <h4>my name is Hamid Poorhashemi . </h4><p>this is  just example<bold>")

}

func contactpage(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Contact us page  </h1><br><h4>My Email: Hamidpoorhashemi@gmail.com</h4> <p>this is  just example<bold>")

}
