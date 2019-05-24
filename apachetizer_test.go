package apachetizer

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"os"
)

func TestVHostConfDetector(t *testing.T) {
	got, err := VHostConfDetector("./etc/apache2/sites-available")
	if err != nil{
		log.Fatal(err)
	}
	if reflect.TypeOf(got).Kind() != reflect.Slice{
		panic(got)
	}else{
		log.Println("VHostConfDetector successfully returns a slice (string array)")
	}

}
func TestVHostConfParser(t *testing.T) {
	Reader, err := os.Open("./etc/apache2/testconfig-le-ssl.conf")
	if err != nil {
		panic(err)
	}

	config, _ := VHostConfParser(Reader)
	jsonEncoded, _ := json.Marshal(config)

	log.Println(string(jsonEncoded))
}
