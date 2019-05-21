package go_apachetizer

import (
	"encoding/json"
	"reflect"
	"testing"
	"fmt"
	"os"
)

func TestVHostConfDetector(t *testing.T) {
	got, err := VHostConfDetector("./etc/apache2/sites-available")
	if err != nil{
		fmt.Println(err)
	}
	if reflect.TypeOf(got).Kind() != reflect.Slice{
		panic(got)
	}else{
		fmt.Println("VHostConfDetector successfuly returns a slice (string array)")
	}

}
func TestVHostConfParser(t *testing.T) {
	Reader, err := os.Open("./etc/apache2/testconfig-le-ssl.conf")
	if err != nil {
		panic(err)
	}

	config, _ := VHostConfParser(Reader)
	jsonEncoded, _ := json.Marshal(config)

	fmt.Println(string(jsonEncoded))
}
