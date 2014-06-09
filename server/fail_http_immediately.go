package server

import (
  "net/http"
  "fmt"
  "strconv"
  "io/ioutil"
  "encoding/json"
)

//returns an empty byte array. intended for use as a parameter
//to NewFailHttpImmediately or HttpImmediately.succMsg
func EmptyBody() []byte {
  return []byte{}
}

//the action that each HttpStep will take
type Runner func() error

//a Runner to extract a value from a query string key. intended for
//use as a HttpStep.runner. the returned Runner will:
//- gets the query string from the given http.Request
//- gets the value(s) for the given key from that query string
//- of the values for that key, attempts to get the value at index idx
//- writes that value to *target
//- returns nil
// if any of the above steps fails, the function returns an appropriate error
//and writes nothing to *target. idx is 0 based
func QueryString(req *http.Request, key string, idx int, target *string) Runner {
  return func() error {
    queryString := req.URL.Query()
    vals := queryString[key]
    if vals == nil {
      return fmt.Errorf("no key %s in the query string", key)
    } else if idx >= len(vals) {
      return fmt.Errorf("fewer than %d values for key %s in the query string", idx+1, key)
    }
    *target = vals[idx]
    return nil
  }
}

//a Runner that:
//- executes QueryString
//- attempts to convert the output of QueryString to a bool
//- writes that value to *target
//- returns nil
//if anything went wrong along the way, return an appropriate error and write
//nothing to *target
func QueryStringParseBool(req *http.Request, key string, idx int, target *bool) Runner {
  return func() error {
    var valStr string
    err := QueryString(req, key, idx, &valStr)()
    if err != nil {
      return err
    }
    b, err := strconv.ParseBool(valStr)
    if err != nil {
      return err
    }
    *target = b
    return nil
  }
}

//a Runner that:
//- reads the entire body from req
//- writes the body to *target
//if there were errors at any step, returns the error and does not write
//to *target
func ReadBody(req *http.Request, target *[]byte) Runner {
  return func() error {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
      return err
    }
    *target = body
    return nil
  }
}

//a Runner that:
//- reads the entire body from req
//- attempts to decode json from that body to target
//if there were errors at any step, returns the error.
//make sure to pass an address to the data structure to
//read, similar to what you'd do with the encoding/json package
//for example:
//  var targetMap map[string]string
//  server.ReadJson(req, &targetMap)
//when the ReadJson runner is executed, the json will be written into
//targetMap if there were no errors
func ReadJson(req *http.Request, target interface{}) Runner {
  return func() error {
    var body []byte
    err := ReadBody(req, &body)()
    if err != nil {
      return err
    }
    err = json.Unmarshal(body, target)
    if err != nil {
      return err
    }
    return nil
  }
}

//a Runner that:
//- encodes val to json
//- writes the resulting bytes to *target
//- returns nil
//if there were errors at any step, does not write to *target and returns
//the error
func EncodeJson(val interface{}, target *[]byte) Runner {
  return func() error {
    bytes, err := json.Marshal(val)
    if err != nil {
      return err
    }
    *target = bytes
    return nil
  }
}

//convert a standard error into a JSON format.
//intended for use in HttpStep.failMsg
func JsonErr(err error) []byte {
  str := fmt.Sprintf(`{"error":"%s"}`, err.Error())
  return []byte(str)
}

//a single step to run in a series to implement a server endpoint
type HttpStep struct {
  Runner Runner
  FailCode int
  FailMsg func(error)[]byte
}

//a server that executes all of its steps and writes the output,
//failure or success to the provided http.ResponseWriter.
//if all steps succeeded, writes SuccCode and SuccMsg to the ResponseWriter.
//otherwise, stops execution immediately after the failed HttpStep
//and writes details on the failed HttpStep to the same ResponseWriter.
//example usage:
//  steps := []*server.HttpStep {
//    &server.HttpStep {
//      Runner: func() error {
//        //do nothing
//        return nil
//      },
//      FailCode: http.StatusInternalServerError,
//      FailMsg: server.JsonErr
//    },
//  }
//  NewFailImmediately(resp, steps, http.StatusOK, []byte("OK!")).Execute()
type FailHttpImmediately struct {
  resp http.ResponseWriter
  steps []*HttpStep
  succCode int
  succMsg []byte
}

//create a new FailHttpImmediately server
func NewFailHttpImmediately(resp http.ResponseWriter, steps []*HttpStep, succCode int, succMsg []byte) *FailHttpImmediately {
  return &FailHttpImmediately {
    resp: resp,
    steps: steps,
    succCode: succCode,
    succMsg: succMsg,
  }
}

//execute the series of steps in the FailHttpImmediately
func (fh *FailHttpImmediately) Execute() []error {
  for _, step := range fh.steps {
    err := step.Runner()
    if err != nil {
      fh.resp.WriteHeader(step.FailCode)
      fh.resp.Write(step.FailMsg(err))
      return []error{err}
    }
  }

  fh.resp.WriteHeader(fh.succCode)
  fh.resp.Write(fh.succMsg)
  return []error{}
}
