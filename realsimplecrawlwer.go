package main
import (
	"net/http"
	"fmt"
	"os"
	"golang.org/x/net/html"
	"strings"
	"io"
	"log"
)
//fetchHTML fetches the provided URL and returns the response body or an error
func fetchHTML(URL string) (io.ReadCloser, error) {
    resp, err := http.Get(URL)

//if there was an error, report it and exit
if err != nil {
    //.Fatalf() prints the error and exits the process
    log.Fatalf("error fetching URL: %v\n", err)
}
if resp.StatusCode != http.StatusOK {
    log.Fatalf("response status code was %d\n", resp.StatusCode)
}

//check response content type
ctype := resp.Header.Get("Content-Type")
if !strings.HasPrefix(ctype, "text/html") {
    log.Fatalf("response content type was %s not text/html\n", ctype)
}
return resp.Body , err
}

//extractTitle returns the content within the <title> element
//or an error
func extractTitle(body io.ReadCloser) (string, error) {
	//the page title or an error
	tokenizer := html.NewTokenizer(body)
	for {
		//get the next token type
		tokenType := tokenizer.Next()
	
		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			//since this is a simple command-line utility,
			//we can just use log.Fatalf() to report the error
			//and exit the process with a non-zero status code
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}
	
		//process the token according to the token type...
	
	if tokenType == html.StartTagToken {
		//get the token
		token := tokenizer.Token()
		//if the name of the element is "title"
		if "title" == token.Data {
			//the next token should be the page title
			tokenType = tokenizer.Next()
			//just make sure it's actually a text token
			if tokenType == html.TextToken {
				//report the page title and break out of the loop
				fmt.Println(tokenizer.Token().Data)
				break
			}
		}
	}}
	return tokenizer.Token().Data,tokenizer.Err()
}

//fetchTitle fetches the page title for a URL
func fetchTitle(URL string) (string, error) {
  
	
	//TODO: fetch the HTML, extract the title, and make sure the body gets closed
 resp,err:= fetchHTML(URL)
 
 a,err := extractTitle(resp)
return a ,err
}
func extractBrand(body io.ReadCloser) (string, error) {

	tokenizer := html.NewTokenizer(body)
	for {
		//get the next token type
		tokenType := tokenizer.Next()
	
		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			//since this is a simple command-line utility,
			//we can just use log.Fatalf() to report the error
			//and exit the process with a non-zero status code
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}
	
		//process the token according to the token type...
	
	if tokenType == html.StartTagToken {
		//get the token
		token := tokenizer.Token()
		//if the name of the element is "title"
		if "Brand Name" == token.Data {
			//the next token should be the page title
			tokenType = tokenizer.Next()
			//just make sure it's actually a text token
			if tokenType == html.TextToken {
				//report the page title and break out of the loop
				fmt.Println(tokenizer.Token().Data)
				break
			}
		}
	}}
	return tokenizer.Token().Data,tokenizer.Err()
}

//fetchTitle fetches the page title for a URL
func fetchBrand(URL string) (string, error) {
  
	
	//TODO: fetch the HTML, extract the title, and make sure the body gets closed
 resp,err:= fetchHTML(URL)
 
 a,err := extractBrand(resp)
return a ,err
}

func extractASIN(body io.ReadCloser) (string, error) {

	tokenizer := html.NewTokenizer(body)
	for {
		//get the next token type
		tokenType := tokenizer.Next()
	
		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			//since this is a simple command-line utility,
			//we can just use log.Fatalf() to report the error
			//and exit the process with a non-zero status code
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}
	
		//process the token according to the token type...
	
	if tokenType == html.StartTagToken {
		//get the token
		token := tokenizer.Token()
		//if the name of the element is "title"
		if "ASIN" == token.Data {
			//the next token should be the page title
			tokenType = tokenizer.Next()
			//just make sure it's actually a text token
			if tokenType == html.TextToken {
				//report the page title and break out of the loop
				fmt.Println(tokenizer.Token().Data)
				break
			}
		}
	}}
	return tokenizer.Token().Data,tokenizer.Err()
}

//fetchTitle fetches the page title for a URL
func fetchASIN(URL string) (string, error) {
  
	
	//TODO: fetch the HTML, extract the title, and make sure the body gets closed
 resp,err:= fetchHTML(URL)
 
 a,err := extractASIN(resp)
return a ,err
}
func extractPrice(body io.ReadCloser) (string, error) {

	tokenizer := html.NewTokenizer(body)
	for {
		//get the next token type
		tokenType := tokenizer.Next()
	
		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			//since this is a simple command-line utility,
			//we can just use log.Fatalf() to report the error
			//and exit the process with a non-zero status code
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		}
	
		//process the token according to the token type...
	
	if tokenType == html.StartTagToken {
		//get the token
		token := tokenizer.Token()
		//if the name of the element is "title"
		if "Price" == token.Data {
			//the next token should be the page title
			tokenType = tokenizer.Next()
			//just make sure it's actually a text token
			if tokenType == html.TextToken {
				//report the page title and break out of the loop
				fmt.Println(tokenizer.Token().Data)
				break
			}
		}
	}}
	return tokenizer.Token().Data,tokenizer.Err()
}

//fetchTitle fetches the page title for a URL
func fetchPrice(URL string) (string, error) {
  
	
	//TODO: fetch the HTML, extract the title, and make sure the body gets closed
 resp,err:= fetchHTML(URL)
 
 a,err := extractPrice(resp)
return a ,err
}
func main() {
	file , err := os.Create("Result.txt")
	if err !=nil {
		log.Fatal("Cannot create file",err)

	}
	defer file.Closed()
    //if the caller didn't provide a URL to fetch...
    if len(os.Args) < 2 {
        //print the usage and exit with an error
        fmt.Printf("usage:\n  pagetitle <url>\n")
        os.Exit(1)
    }

	title, err := fetchTitle(os.Args[1])
	brand,err :=fetchBrand(os.Args[1])
	ASIN,err := fetchASIN(os.Args[1])
	Price,err := fetchPrice(os.Args[1])
    if err != nil {
        log.Fatalf("error fetching page title: %v\n", err)
    }

    //print the title
	fmt.Fprintf(file,title)
	fmt.Fprintf(file,brand)
	fmt.Fprintf(file,ASIN)
	fmt.Fprintf(file,Price)
	
}