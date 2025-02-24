// TODO: Working with basic Go Latency Checker

package main

import (
 "bufio"
 "flag"
 "fmt"
 "io"
 "net/http"
 "os"
 "regexp"
 "time"
)

func getEndpointsFromURL(url string) ([]string, error) {
 resp, err := http.Get(url)
 if err != nil {
  return nil, err
 }
 defer resp.Body.Close()

 // Read the response body as a byte slice
 bodyBytes, err := io.ReadAll(resp.Body)
 if err != nil {
  return nil, err
 }

 // Convert the byte slice to a string
 bodyString := string(bodyBytes)

 // Define a regular expression pattern to match URLs
 pattern := `href=["'](http[^"'\s]+)["']`
 re := regexp.MustCompile(pattern)

 // Find all matches of the pattern in the HTML content
 matches := re.FindAllStringSubmatch(bodyString, -1)

 // Extract URLs from the matches and filter out CSS and image URLs
 urls := make([]string, 0)
 for _, match := range matches {
  url := match[1]
  if !isCSSOrImageURL(url) {
   urls = append(urls, url)
  }
 }

 return urls, nil
}

func isCSSOrImageURL(url string) bool {
 // Define regular expression patterns for CSS and image file extensions
 cssPattern := `\.(css)$`
 imagePattern := `\.(jpg|jpeg|png|gif|svg|ico|webp)$`

 // Compile regular expressions
 cssRe := regexp.MustCompile(cssPattern)
 imageRe := regexp.MustCompile(imagePattern)

 // Check if the URL matches any of the patterns
 if cssRe.MatchString(url) || imageRe.MatchString(url) {
  return true
 }
 return false
}

func checkLatency(url string, timeout time.Duration) (time.Duration, error) {
 client := http.Client{
  Timeout: timeout,
 }
 start := time.Now()
 _, err := client.Get(url)
 if err != nil {
  return 0, err
 }
 return time.Since(start), nil
}

func main() {
 // Parse command-line arguments
 url := flag.String("url", "", "URL to search for URLs")
 outputFile := flag.String("output", "endpoints.txt", "Output file to store endpoints")
 flag.Parse()

 if *url == "" {
  fmt.Println("Please provide a URL to search for URLs using the -url flag.")
  return
 }

 // Get endpoints from the URL
 endpoints, err := getEndpointsFromURL(*url)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }

 // Write endpoints to the output file
 file, err := os.Create(*outputFile)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }
 defer file.Close()

 for _, endpoint := range endpoints {
  _, err := file.WriteString(endpoint + "\n")
  if err != nil {
   fmt.Println("Error:", err)
   return
  }
 }

 fmt.Printf("Endpoints written to %s\n", *outputFile)

 // Check latency for each endpoint in the file
 file, err = os.Open(*outputFile)
 if err != nil {
  fmt.Println("Error:", err)
  return
 }
 defer file.Close()

 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
  endpoint := scanner.Text()

  // Check latency of the endpoint
  latency, err := checkLatency(endpoint, 5*time.Second)
  if err != nil {
   fmt.Printf("Error checking latency for %s: %v\n", endpoint, err)
  } else {
   fmt.Printf("Latency for %s: %v\n", endpoint, latency)
  }
 }

 if err := scanner.Err(); err != nil {
  fmt.Println("Error:", err)
  return
 }
 }
