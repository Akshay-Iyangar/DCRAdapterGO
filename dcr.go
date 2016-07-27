package main

import (
    "io/ioutil"
    "net/http"
    "time"
    "bytes"
    "log"
    "os"
    "fmt"
)

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func main() {

    ticker := time.NewTicker(time.Millisecond * 30000)
    for time := range ticker.C {
      resp, err := http.Get("http://dashboard.ecom.int.godaddy.com/fulfillment")
      if err != nil {
          debug("Could not request Fulfillment Log")
          panic(err)
      }
      defer resp.Body.Close()
      body, err := ioutil.ReadAll(resp.Body)
     	if err != nil {
     		debug(" The response body from fullfillment log failed!!!")
        panic(err)
     	}
    	raw_message := string(body)
    	fmt.Println(raw_message)

			if err != nil {
				debug("flushHttp - Error encoding JSON: ", err)
			}
      debug("The JSON encoded message that would be sent to the DCR: ",message)
      debug("Tick at", time)
      //This should not be hard-coded and this should be taken from the script.
      url := os.Args[1]
      //"http://localhost:12285/v1/dc/logs/ecomm/logs"
      debug("URL:>", url)

      //var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
      req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
      req.Header.Set("X-Custom-Header", "Dashboard")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      response, err := client.Do(req)
      if err != nil {
              panic(err)
              debug("The response was not successful after pushing logs to DCR")
      }
      defer response.Body.Close()

      debug("response Status:", response.Status)
      debug("response Headers:", response.Header)

    }
}
