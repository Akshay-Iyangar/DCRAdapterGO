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


for{

    /*aladin logic will be all togther different */
    //All this can be also be ridden if i have a mechanism to store the time and the endpoint.
    //i.e i will only need 2 if condition. this is something i need to think abut. May be Redis
    //This will also prevent me to worry about having only one Docker container.
    //I can have more than one Docker container in that case.

    if previousGET_Ftime == "0001-01-01 00:00:00.000" {
      fmt.Println("Running fullfillmentLog (first iteration)")
      previous_Ftime = DCRAdapter("fulfillmentLog",previousGET_Ftime)
      previousGET_Ftime = previous_Ftime.Format("2006-01-02 15:04:05.000")
      time.Sleep(30000 * time.Millisecond)
    }

    //fmt.Println(time.Since(previous_Ftime).Minutes())
    if time.Since(previous_Ftime).Minutes()>= 5{
      fmt.Println("Running fullfillmentLog")
      previous_Ftime = DCRAdapter("fulfillmentLog",previousGET_Ftime)
      previousGET_Ftime = previous_Ftime.Format("2006-01-02 15:04:05.000")
    }
    if previousGET_CEtime == "0001-01-01 00:00:00.000" {
      fmt.Println("Running commonerroLog (first iteration)")
      previous_CEtime = DCRAdapter("commerrorLog",previousGET_CEtime)
      previousGET_CEtime = previous_CEtime.Format("2006-01-02 15:04:05.000")
      time.Sleep(30000 * time.Millisecond)
    }
    if time.Since(previous_CEtime).Minutes()>= 5{
      fmt.Println("Running commonerroLog")
      previous_CEtime = DCRAdapter("commerrorLog",previousGET_CEtime)
      previousGET_CEtime = previous_CEtime.Format("2006-01-02 15:04:05.000")

    }
    if previousGET_Stime == "0001-01-01 00:00:00.000" {
      fmt.Println("Running siteerrorLog (first iteration)")
      previous_Stime = DCRAdapter("siteerrorLog",previousGET_Stime)
      previousGET_Stime = previous_Stime.Format("2006-01-02 15:04:05.000")
      time.Sleep(30000 * time.Millisecond)
    }
    if time.Since(previous_Stime).Minutes()>= 5{
      fmt.Println("Inside S2")
      previous_Stime = DCRAdapter("siteerrorLog",previousGET_Stime)
      previousGET_Stime = previous_Stime.Format("2006-01-02 15:04:05.000")

    }
    if previousGET_Mtime == "0001-01-01 00:00:00.000" {
      fmt.Println("Running msmqLog (first iteration)")
      previous_Mtime = DCRAdapter("msmqLog",previousGET_Mtime)
      previousGET_Mtime = previous_Mtime.Format("2006-01-02 15:04:05.000")
      time.Sleep(30000 * time.Millisecond)

    }
    if time.Since(previous_Mtime).Minutes()>= 5{
      fmt.Println("Running msmqLog")
      previous_Mtime = DCRAdapter("msmqLog",previousGET_Mtime)
      previousGET_Mtime = previous_Mtime.Format("2006-01-02 15:04:05.000")

    }
    if previousGET_Ctime == "0001-01-01 00:00:00.000"{
      fmt.Println("Running commonpurchaseLog (first iteration)")
      previous_Ctime = DCRAdapter("commonpurchaseLog",previousGET_Ctime)
      previousGET_Ctime = previous_Ctime.Format("2006-01-02 15:04:05.000")
      time.Sleep(30000 * time.Millisecond)

    }
    if  time.Since(previous_Ctime).Minutes()>= 5{
      fmt.Println("Running commonpurchaseLog")
      previous_Ctime = DCRAdapter("commonpurchaseLog",previousGET_Ctime)
      previousGET_Ctime = previous_Ctime.Format("2006-01-02 15:04:05.000")

    }
    /*
    temp := (5 - time.Since(previous_Ftime).Minutes()) * 60
    fmt.Println(temp)
    fmt.Println(time.Duration(temp) * time.Second)
    time.Sleep(time.Duration(temp) * time.Second)*/
    //time.Sleep(30000 * time.Millisecond)
    //x:= previous_Ftime.Add(-time.Since(previous_Ftime).Minutes() * time.Minute)
    //var duration_Seconds time.Duration = 300 * time.Second
    //var duration1_Seconds time.Duration = time.Since(previous_Ftime)
    // := duration_Seconds - time.Since(previous_Ftime)

    //fmt.Println(x)
    //time.Sleep(x * time.Minute) // write it in minutes*/
  }
}
func DCRAdapter(endpoint string,previousGET_time string) (time.Time){
  current,previousGET_time := LogTimer(previousGET_time) //everything has to be a string
  currentGET_time := current.Format("2006-01-02 15:04:05.000")
  body := DashboardGET(endpoint,currentGET_time,previousGET_time)
  //previousGET_time = currentGET_time
  DCRPost(body)
  return current   //this has to be time
}


func DashboardGET(endpoint string ,currentGET_time string,previousGET_time string)([]byte){
  host := os.Args[1]
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
  fmt.Println("The current time",endpoint,":",currentGET_time)
  fmt.Println("The previous time",endpoint,":",previousGET_time)
  return body
}


func LogTimer(previousGET_time string)(time.Time,string){



  loc, _  := time.LoadLocation("US/Pacific")
  current_time := time.Now().Format("2006-01-02 15:04:05.000")
  ct,_ := time.Parse("2006-01-02 15:04:05.000",current_time)
  current:= ct.In(loc) // time format


  //currentGET_time := current.Format("2006-01-02 15:04:05.000")

  if previousGET_time == "0001-01-01 00:00:00.000" {
          //this should be the time we make the thread sleep
          previoustime := time.Now().Add(-5 * time.Minute).Format("2006-01-02 15:04:05.000")
          pt,_ := time.Parse("2006-01-02 15:04:05.000", previoustime) //in time format
          previous := pt.In(loc)
          previousGET_time = previous.Format("2006-01-02 15:04:05.000") //in string format
  }

  return current,previousGET_time
}

func DCRPost(body []byte){
  client := &http.Client{}
  //The DCR endpoint
  url := os.Args[2]
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
