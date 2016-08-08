package main

import (
    "io/ioutil"
    "net/http"
    "net/url"
	  "path"
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
  var previous_time time.Time
  var previous_Ftime time.Time
  var previous_CEtime time.Time
  var previous_Stime time.Time
  var previous_Mtime time.Time
  var previous_Ctime time.Time

  //This is needed for matching the SQL format
  var previousGET_Ftime string = previous_time.Format("2006-01-02 15:04:05.000")
  var previousGET_Mtime string = previous_time.Format("2006-01-02 15:04:05.000")
  var previousGET_Ctime string = previous_time.Format("2006-01-02 15:04:05.000")
  var previousGET_Stime string = previous_time.Format("2006-01-02 15:04:05.000")
  var previousGET_CEtime string = previous_time.Format("2006-01-02 15:04:05.000")
  loc, _ := time.LoadLocation("US/Pacific")

for{

    /*aladin logic will be all togther different */

    if previousGET_Ftime == "0001-01-01 00:00:00.000" {
      previousGET_Ftime = DCRAdapterFulfillment("fulfillmentLog",previousGET_Ftime)
      previous_Ftime,_ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_Ftime, loc)
      time.Sleep(30000 * time.Millisecond)
    }
    if time.Since(previous_Ftime).Minutes()>= 5{
      previousGET_Ftime = DCRAdapterFulfillment("fulfillmentLog",previousGET_Ftime)
      previous_Ftime,_ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_Ftime, loc)
    }
    if previousGET_CEtime == "0001-01-01 00:00:00.000" {
      previousGET_CEtime = DCRAdapterCommError("commerrorLog",previousGET_CEtime)
      previous_CEtime, _ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_CEtime, loc)
      time.Sleep(30000 * time.Millisecond)
    }
    if time.Since(previous_CEtime).Minutes()>= 5{
      previousGET_CEtime = DCRAdapterCommError("commerrorLog",previousGET_CEtime)
      previous_CEtime, _ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_CEtime, loc)
    }
    if previousGET_Stime == "0001-01-01 00:00:00.000" {
      previousGET_Stime = DCRAdapterSiteError("siteerrorLog",previousGET_Stime)
      previous_Stime, _ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_Stime, loc)
      time.Sleep(30000 * time.Millisecond)
    }
    if time.Since(previous_Stime).Minutes()>= 5{
      previousGET_Stime = DCRAdapterSiteError("siteerrorLog",previousGET_Stime)
      previous_Stime, _ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_Stime, loc)
    }

    if previousGET_Mtime == "0001-01-01 00:00:00.000" {
      previousGET_Mtime = DCRAdapterMSMQ("msmqLog",previousGET_Mtime)
      previous_Mtime, _ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_Mtime, loc)
      time.Sleep(30000 * time.Millisecond)
    }
    if time.Since(previous_Mtime).Minutes()>= 5{
      previousGET_Mtime = DCRAdapterMSMQ("msmqLog",previousGET_Mtime)
      previous_Mtime, _ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_Mtime, loc)
    }

    if previousGET_Ctime == "0001-01-01 00:00:00.000"{
      previousGET_Ctime = DCRAdapterCommonpurchase("commonpurchaseLog",previousGET_Ctime)
      previous_Ctime, _ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_Ctime, loc)
      time.Sleep(30000 * time.Millisecond)
    }
    if  time.Since(previous_Ctime).Minutes()>= 5{
      previousGET_Ctime = DCRAdapterCommonpurchase("commonpurchaseLog",previousGET_Ctime)
      previous_Ctime, _ = time.ParseInLocation("2006-01-02 15:04:05.000", previousGET_Ctime, loc)
    }

    time.Sleep(120000 * time.Millisecond) // write it in minutes

  }
}
//func <<function_name>>(input type) return type {}
func DCRAdapterFulfillment(endpoint string,previousGET_Ftime string) (string){
  currentGET_Ftime,previousGET_Ftime := FulfillmentLogTimer(previousGET_Ftime)
  body := DashboardGET(endpoint,currentGET_Ftime,previousGET_Ftime) // This is common for everything
  previousGET_Ftime = currentGET_Ftime
  fmt.Println(previousGET_Ftime)
  DCRPost(body)
  return previousGET_Ftime
      //time.Sleep(30000 * time.Millisecond)
}
func DCRAdapterMSMQ(endpoint string, previousGET_Mtime string)(string){
  currentGET_Mtime,previousGET_Mtime := MSMQLogTimer(previousGET_Mtime)
  body := DashboardGET(endpoint,currentGET_Mtime,previousGET_Mtime) // This is common for everything
  previousGET_Mtime = currentGET_Mtime
  DCRPost(body)
  return previousGET_Mtime
}
func DCRAdapterCommonpurchase(endpoint string,previousGET_Ctime string)(string){
  currentGET_Ctime,previousGET_Ctime := CommonPurchaseLogTimer(previousGET_Ctime)
  body := DashboardGET(endpoint,currentGET_Ctime,previousGET_Ctime) // This is common for everything
  previousGET_Ctime = currentGET_Ctime
  DCRPost(body)
  return previousGET_Ctime
}

func DCRAdapterCommError(endpoint string,previousGET_CEtime string)(string){
  currentGET_CEtime,previousGET_CEtime := CommErrorLogTimer(previousGET_CEtime)
  body := DashboardGET(endpoint,currentGET_CEtime,previousGET_CEtime) // This is common for everything
  previousGET_CEtime = currentGET_CEtime
  DCRPost(body)
  return previousGET_CEtime
}

func DCRAdapterSiteError(endpoint string,previousGET_Stime string)(string){
  currentGET_Stime,previousGET_Stime := SiteErrorLogTimer(previousGET_Stime)
  body := DashboardGET(endpoint,currentGET_Stime,previousGET_Stime) // This is common for everything
  previousGET_Stime = currentGET_Stime
  DCRPost(body)
  return previousGET_Stime
}


//We only need to keep adding this function w.r.t all the endpoints we have.



/*
func CommonPurchaseLogTimer(previousGET_time string){
  fmt.Println("Inside CommonPurchaseLog")
  //For test the timer is of 19 minutes
  ticker := time.NewTicker(time.Minute * 19)
  for timer := range ticker.C {
    endpoint := "commonpurchaseLog"
    currentGET_time,previousGET_time := TimeInformation(previousGET_time)
    body := DashboardGET(endpoint,currentGET_time,previousGET_time) // This is common for everything
    previousGET_time = currentGET_time
    DCRPost(body) //This is common for everything
    debug("Timer for fullfillment Log", timer)
  }
}
*/



func DashboardGET(endpoint string ,currentGET_time string,previousGET_time string)([]byte){
  host := os.Args[1]
  //"http://localhost:9179/"
  URL,_ := url.Parse(host)
	URL.Path = path.Join(URL.Path, endpoint)
  client := &http.Client{}
  //Dont go by the time entered in  the format its just a format for to send
  Getquery, _ := http.NewRequest("GET",URL.String(),nil)
  query := Getquery.URL.Query()
  query.Add("current", currentGET_time)
  query.Add("previous", previousGET_time)
  Getquery.URL.RawQuery = query.Encode()
  req, err := client.Do(Getquery)
  if err != nil {
      debug("Could not request Fulfillment Log")
      panic(err)
  }
  defer req.Body.Close()
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    debug(" The response body from fullfillment log failed!!!")
    panic(err)
  }
  //fmt.Println(body)
  return body
}


func FulfillmentLogTimer(previousGET_Ftime string)(string,string){
  current_time := time.Now()
  currentGET_Ftime := current_time.Format("2006-01-02 15:04:05.000")

  if previousGET_Ftime == "0001-01-01 00:00:00.000" {
          //this should be the time we make the thread sleep
          previous_time := current_time.Add(-5 * time.Minute)
          previousGET_Ftime = previous_time.Format("2006-01-02 15:04:05.000")
  }
  fmt.Println("The current time fulfillment:",currentGET_Ftime)
  fmt.Println("The previous time fulfillment :",previousGET_Ftime)
  return currentGET_Ftime,previousGET_Ftime
}

func MSMQLogTimer(previousGET_Mtime string)(string,string){
  current_time := time.Now()
  currentGET_Mtime := current_time.Format("2006-01-02 15:04:05.000")

  if previousGET_Mtime == "0001-01-01 00:00:00.000" {
          previous_time := current_time.Add(-5 * time.Minute)
          previousGET_Mtime = previous_time.Format("2006-01-02 15:04:05.000")
  }
  fmt.Println("The current time msmqLog :",currentGET_Mtime)
  fmt.Println("The previous time msmqLog :",previousGET_Mtime)
  return currentGET_Mtime,previousGET_Mtime
}

func CommonPurchaseLogTimer(previousGET_Ctime string)(string,string){
  current_time := time.Now()
  currentGET_Ctime := current_time.Format("2006-01-02 15:04:05.000")

  if previousGET_Ctime == "0001-01-01 00:00:00.000" {
          previous_time := current_time.Add(-5 * time.Minute)
          previousGET_Ctime = previous_time.Format("2006-01-02 15:04:05.000")
  }
  fmt.Println("The current time commonpurchaseLog :",currentGET_Ctime)
  fmt.Println("The previous time commonpurchaseLog :",previousGET_Ctime)
  return currentGET_Ctime,previousGET_Ctime
}

func CommErrorLogTimer(previousGET_CEtime string)(string,string){
  current_time := time.Now()
  currentGET_CEtime := current_time.Format("2006-01-02 15:04:05.000")

  if previousGET_CEtime == "0001-01-01 00:00:00.000" {
          //this should be the time we make the thread sleep
          previous_time := current_time.Add(-5 * time.Minute)
          previousGET_CEtime = previous_time.Format("2006-01-02 15:04:05.000")
  }
  fmt.Println("The current time commerror:",currentGET_CEtime)
  fmt.Println("The previous time commerror :",previousGET_CEtime)
  return currentGET_CEtime,previousGET_CEtime
}

func SiteErrorLogTimer(previousGET_Stime string)(string,string){
  current_time := time.Now()
  currentGET_Stime := current_time.Format("2006-01-02 15:04:05.000")

  if previousGET_Stime == "0001-01-01 00:00:00.000" {
          //this should be the time we make the thread sleep
          previous_time := current_time.Add(-5 * time.Minute)
          previousGET_Stime = previous_time.Format("2006-01-02 15:04:05.000")
  }
  fmt.Println("The current time siteerror:",currentGET_Stime)
  fmt.Println("The previous time siteerror :",previousGET_Stime)
  return currentGET_Stime,previousGET_Stime
}


func DCRPost(body []byte){
  client := &http.Client{}
  //The DCR endpoint
  url := os.Args[2]
  //"http://localhost:12285/v1/dc/logs/ecomm/logs"
  debug("URL:>", url)

  resp, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
  resp.Header.Set("X-Custom-Header", "Dashboard")
  resp.Header.Set("Content-Type", "application/json")

  //clientPost := &http.Client{}
  response, err := client.Do(resp)
  if err != nil {
          panic(err)
          debug("The response was not successful after pushing logs to DCR")
  }
  defer response.Body.Close()
  fmt.Println("response Status:", response.Status)
  debug("response Status:", response.Status)
  debug("response Headers:", response.Header)
}
