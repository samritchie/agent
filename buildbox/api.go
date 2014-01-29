package buildbox

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

type Build struct {
  State string
  Script string
  Output string
}

func (b Build) String() string {
  return fmt.Sprintf("Build(%s)", b.State)
}

type Response struct {
  Build *Build
}

func Get(url string) (*http.Response) {
  log.Printf("GET %s", url)

  resp, err := http.Get(url)

  // Check to make sure no error returned from the get request
  if err != nil {
    log.Fatal(err)
  }

  // Check the status code
  if resp.StatusCode != http.StatusOK {
    log.Fatal(resp.Status)
  }

  // io.Copy(os.Stdout, resp.Body)

  return resp
}

func GetNextBuild() (*Build) {
  var url string = "http://agent.buildbox.dev/v1/e6296371ed3dd3f24881b0866506b8c6/builds/queue/next"
  resp := Get(url)

  var r *Response = new(Response)
  err := json.NewDecoder(resp.Body).Decode(r)
  if err != nil {
    log.Fatal(err)
  }

  return r.Build
}
