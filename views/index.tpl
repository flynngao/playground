<!DOCTYPE html>

<html>
<head>
  <title>PlayGround</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <script src='http://cdn.staticfile.org/jquery/1.9.1/jquery.min.js'></script>
  <script src='http://files.aw20.net/jquery-linedtextarea/jquery-linedtextarea.js'></script>
  <link rel="stylesheet" href="http://files.aw20.net/jquery-linedtextarea/jquery-linedtextarea.css">
  <style>
    html,body{
      padding: 0;
      margin: 0;
      background: #272822
    }
    #container{
      overflow: hidden;
    }

    #wrap{
      float: left;
      width: 50%;
      height: 97vh;
      border-right: 1px solid #666;
      box-sizing:border-box;
    }
    #code{
      width: 100%;
      height: 97vh;
      box-sizing:border-box;
      background: #272822;
      color: #eee;
      border: 0
    }
    .output {
      float: left;
      width: 50%;
      padding: 10px 20px;
      white-space: normal;
      box-sizing: border-box;
      word-wrap: break-word;
      color: #efefef

    }
    .run {
      border: 3px solid #343436;
      background: #343436;
      text-align: center;
      color: white;
      font-size: 14px;
      padding: 5px 10px


    }
    .run:hover{
      background: #666
    }
    .linedwrap{
      height: 100%;
      width:100%;

    }
    .linedwrap .lines{
      width: 3%;
    }
    .linedwrap textarea{
      width: 97%;
    }
  </style>
</head>

<body>
  <header>
    <button class="run">Run</button>
  </header>
  <div id="container">
    
  <div id="wrap">
<textarea itemprop="description" id="code" name="code" autocorrect="off" autocomplete="off" autocapitalize="off" spellcheck="false">
package main

import "fmt"

func main() {
  fmt.Println("Hello, playground")
}
</textarea>
    </div>
  <pre class="output">
    

  </pre>
  </div>
  <footer>
    <script src="https://ace.c9.io/build/src/ace.js" type="text/javascript" charset="utf-8"></script>
   <script src='/static/js/playground.js'></script>
  </footer>
  <div class="backdrop"></div>
</body>
</html>
