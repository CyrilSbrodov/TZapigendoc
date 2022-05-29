package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Url = "https://sycret.ru/service/apigendoc/forma_025u.xml"

type Path struct {
	FilePath    string
	FilePathNew string
}

func FileName() Path {
	var file Path
	file.FilePath = "data/file.xml"
	time := time.Now()
	file.FilePathNew = "data/" + time.Format("2006-01-02 15-04-05") + ".doc"
	return file
}

//type WordDocument struct {
//	XMLName xml.Name `xml:"wordDocument"`
//	Body    struct {
//		Sect struct {
//			Use struct {
//				Tbl []struct {
//					Tr []struct {
//						Tc []struct {
//							P []struct {
//								Text struct {
//									//Field string `xml:"field,attr"`
//									R struct {
//										T string `xml:"t"`
//									} `xml:"r"`
//								} `xml:"text"`
//								Use struct {
//									Text []struct {
//										//Field string `xml:"field,attr"`
//										R struct {
//											T string `xml:"t"`
//										} `xml:"r"`
//									} `xml:"text"`
//								} `xml:"use"`
//							} `xml:"p"`
//						} `xml:"tc"`
//					} `xml:"tr"`
//				} `xml:"tbl"`
//			} `xml:"use"`
//		} `xml:"sect"`
//	} `xml:"body"`
//}
//}

func NewFileFromServer(filePathNew string) {
	resp, err := http.Get(Url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("Read body: %v", err)
	}
	fmt.Println(data)
	err = ioutil.WriteFile(filePathNew, data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func getXML() ([]byte, error) {
	resp, err := http.Get(Url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
