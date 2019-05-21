# go-apachetizer
Simple and lightweight Apache2 Config Detector and Parser

##Used external Dependencies
* github.com/msoap/byLine

##Docs
**All code examples are not handling any errors please keep that in mind**

##### VHostConfDetector(cfgPath string)
This function will detect all configs save it into a <br>[]string array and returns the []string
```go
package main

import("github.com/dbzmelvin/go-apachetizer")

func main(){
	list, _ :=  VHostConfDetector("./etc/apache2/sites-available")
    fmt.Printf("This is your list: %s", list)
}
```

##### VHostConfParser(file io.Reader)
This function will parse the given config file, you have to pass an io.Reader
```go
package main

import("github.com/dbzmelvin/go-apachetizer")

func main(){
    Reader, _ := os.Open("./etc/apache2/testconfig-le-ssl.conf") //io.Reader
    config, _ := VHostConfParser(Reader) //Parse the config
    jsonEncoded, _ := json.Marshal(config) //Encode the []string array to json []byte array
    fmt.Println(string(jsonEncoded)) //Print the []byte array as string
}
```

##License
[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)

Creative Commons Corporation (“Creative Commons”) is not a law firm and does not provide legal services or legal advice. Distribution of Creative Commons public licenses does not create a lawyer-client or other relationship. Creative Commons makes its licenses and related information available on an “as-is” basis. Creative Commons gives no warranties regarding its licenses, any material licensed under their terms and conditions, or any related information. Creative Commons disclaims all liability for damages resulting from their use to the fullest extent possible.