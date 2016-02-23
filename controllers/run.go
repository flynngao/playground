package controllers

import (
  "github.com/astaxie/beego"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "path/filepath"
  "bufio"
)

type RunController struct {
  beego.Controller
}


type getResult struct {
  Message string `json:"message"`
}


func (this *RunController) Get(){
  res := getResult{"nothing"}
  res = getResult{Message:"nothing"}
  this.Data["json"] = &res
  this.ServeJSON()
}

/** run post data structure
  {
    error
    message
  }

**/

type Output struct{
  Message string `json:"message"`
  Error string `json:"error"`
}

func RunCode(code string) (string,error){
  tmpDir, err := ioutil.TempDir("", "sandbox")
  if err != nil {
    return "",err
  }
  // delete after return

  fmt.Println("temp dir path:", tmpDir)
  defer os.RemoveAll(tmpDir)

  in := filepath.Join(tmpDir, "main.go")
  if err := ioutil.WriteFile(in, []byte(code), 0400); err != nil {
    return "", err
    }

  fmt.Println("files path", in)
  cmd := exec.Command("go","run",in)

  stdpipe, err := cmd.StdoutPipe()
  if err != nil {
    return "", err
  }
  errpipe, err := cmd.StderrPipe()
  if err != nil {
    return "", err
  }
  
  if err := cmd.Start(); err != nil {
    return "", err
  }

  normalOut := bufio.NewScanner(stdpipe)
  errOut := bufio.NewScanner(errpipe)

  no := ""
  for normalOut.Scan() {
    no += normalOut.Text()
    no += "\n"
  }

  fmt.Println(no)
  eo := ""
  for errOut.Scan() {
    eo += errOut.Text()
    eo += "\n"
  }
  fmt.Println(eo)
  if len(no) > 0 {
    return no, nil
  } else{
    return eo, nil
  }
}

func (this *RunController) Post() {

  code := this.GetString("code")
  fmt.Println("request code: ", code)
  message, err := RunCode(code)
  if err != nil {
    // this.Data["json"] = &Output{error:err}
  } else {
    this.Data["json"] = &Output{"",message}
  }

  this.ServeJSON()
}