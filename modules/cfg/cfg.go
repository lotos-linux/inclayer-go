package cfg

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/goccy/go-json"
	"github.com/akshaybharambe14/go-jsonc"
	// "github.com/tidwall/sjson"
)

var err error

type Config struct {
	CurrentTheme 	string
	IconSize   		int
	Layer     		string
	Position		string
	Margin			int
	Blur			bool
	Spacing			int
	Priority		string		
}

func GetDefaultConfig() Config {
	config := Config{}

	config.CurrentTheme = "lotos"
	config.IconSize = 21
	config.Layer = "bottom"
	config.Position = "bottom"
	config.Margin = 8
	config.Blur = true
	config.Spacing = 8

	return config
}

func ConnectConfig(jsoncFile string) Config {
	file, err := os.Open(jsoncFile)
	if err != nil {
		fmt.Println("Config file not found!\n", err, "\nLoad default config")
		return GetDefaultConfig()
	}
	defer file.Close()

	decoder := jsonc.NewDecoder(file)
	res, err := ioutil.ReadAll(decoder)
	fmt.Println(json.Valid(res), jsoncFile)

	config := Config{}
	if err = json.Unmarshal(res, &config); err != nil {
		fmt.Println("Config is incorrect!\n", err, "\nLoad default config")
		return GetDefaultConfig()
	}

	return config
}


type ItemList struct {
	List			[]map[string]string
}

func ReadItemList(jsonFile string) ItemList {
	file, _ := os.Open(jsonFile)
	defer file.Close()

	decoder := json.NewDecoder(file)

	itemList := ItemList{}
	err := decoder.Decode(&itemList)
	if err != nil {
	  fmt.Println("error: ", err)
	}
	
	return itemList
}