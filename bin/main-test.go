package main
//
//import (
//    "io/ioutil"
//    "fmt"
//)
//
//
//////
//////
//////
//////
//////
////////import (
////////	"net/url"
////////	"fmt"
////////)
////////
////////
////////func main() {
////////	u, err := url.Parse("http://example.com:8080/test?first=true")
////////	if err == nil {
////////		fmt.Print(u.Host, "\n", u.Path, "\n", u.Scheme, "\n", u.RequestURI())
////////	}
////////}
////////
////////
////import (
////    "bytes"
////    "fmt"
////    "net/http"
////    "net/url"
////    "crypto/hmac"
////    "crypto/sha1"
////    "time"
////    "encoding/hex"
////    "strings"
////    "net/http/httputil"
////    "log"
////)
////func CreateSign(greenwich_date, access_key, secret_key, http_path, http_method string)(string){
////    sign := hmac.New(sha1.New, []byte(secret_key))
////    sign.Write([]byte (http_method + "\n"))
////    sign.Write([]byte(http_path + "\n"))
////    sign.Write([]byte(greenwich_date + "\n"))
////    //fmt.Println(sign.Sum(nil))
////    return hex.EncodeToString(sign.Sum(nil)[:])
////}
////
////
////func main() {
////    apiUrl := "http://api.speedycloud.cn/api/v1/products/cloud_servers/provision"
////    data := url.Values{}
////    data.Set("az", "SPC-BJ-15-A")
////    data.Add("bandwidth", "2")
////    data.Add("cpu", "1")
////    data.Add("disk", "20")
////    data.Add("disk_type", "Normal")
////    data.Add("image", "Ubuntu 14.04")
////    data.Add("isp", "Private")
////    data.Add("memory", "1024")
////
////    u, _ := url.ParseRequestURI(apiUrl)
////    urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/"
////
////    client := &http.Client{}
////    method := "POST"
////    access_key := "A44384F4AE3998DFEEAFD5E6CC6B1D56"
////    secret_key :="26685451e2a302886fb2a53958a787f65d3433731bee56d8f7fe96a714f677db"
////    http_path :="/api/v1/products/cloud_servers/provision"
////
////    greenwich := strings.Replace(time.Now().UTC().Format(time.RFC1123), "UTC", "GMT", -1)
////    r, _ := http.NewRequest(method, urlStr, bytes.NewBufferString(data.Encode()))
////    r.Header.Add("Authorization",
////        fmt.Sprintf("%s,%s", access_key, CreateSign(greenwich,access_key,secret_key, http_path, method)))
////    r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
////    r.Header.Add("Date", greenwich)
////    r.Header.Add("Accept", "*/*")
////    //fmt.Println(CreateSign(greenwich,access_key,secret_key, http_path, method))
////    //fmt.Println(CreateSign(
////    //    "fri",
////    //    "26685451e2a302886fb2a53958a787f65d3433731bee56d8f7fe96a714f677db",
////    //    "A44384F4AE3998DFEEAFD5E6CC6B1D56",
////    //    "/",
////    //    "POST"))
////    fmt.Println(r)
////    out, err := httputil.DumpRequestOut(r, true)
////    if err != nil {
////        log.Fatal(err)
////    }
////    fmt.Println(strings.Replace(string(out), "\r", "", -1))
////    resp, _ := client.Do(r)
////    fmt.Println(resp.Status)
////    //out := make([]byte, 500)
////    //resp.Body.Read(out)
////    //fmt.Println(out)
////}
////////
////////type Person struct {
////////    Name   string
////////    Age    int
////////    Emails []string
////////    Extra  map[string]string
////////}
//////
//////// This input can come from anywhere, but typically comes from
//////// something like decoding JSON where we're not quite sure of the
//////// struct initially.
//////
//////import (
//////    "encoding/json"
//////    "fmt"
//////    "log"
//////    "strings"
//////    "github.com/mitchellh/mapstructure"
//////    "golang.org/x/tools/go/gcimporter15/testdata"
//////    "github.com/Azure/go-autorest/autorest/validation"
//////)
//////
//////func main() {
//////    const jsonStream = `
//////		[
//////			{"Name": "Ed", "Text": "Knock knock."},
//////			{"Name": "Sam", "Text": "Who's there?"},
//////			{"Name": "Ed", "Text": "Go fmt."},
//////			{"Name": "Sam", "Text": "Go fmt who?"},
//////			{"Name": "Ed", "Text": "Go fmt yourself!"}
//////		]
//////	`
//////    type Message struct {
//////        Name, Text string
//////    }
//////    dec := json.NewDecoder(strings.NewReader(jsonStream))
//////
//////    var m interface{}
//////    err := dec.Decode(&m)
//////    if err != nil {
//////        log.Fatal(err)
//////    }
//////    var res []struct {
//////        Name, Text string
////// }
//////    mapstructure.Decode(m, &res)
//////
//////    for _, val := range res {
//////        fmt.Println(val.(type))
//////
//////    }
//////
//////}
//
//func main() {
//
//    userdata, _ := ioutil.ReadFile("/tmp/userdata")
//    fmt.Println(string(userdata))
//}