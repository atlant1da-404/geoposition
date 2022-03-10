package ip

import (
	"encoding/json"
	"fmt"
	"geoposition/libs/telegram"
	"io"
	"log"
	"net/http"
)

func SendInfoToTelegram(address string, telegram *telegram.Telegram) {

	var data = getData(address)
	var x, y = GetCoordinates(address)
	go telegram.SendMessage(formatMessage(address, data, x, y))
}

func formatMessage(address string, data map[string]interface{}, x, y string) string {

	return fmt.Sprintf(`
	IP Address: %v, Provider: %v, Organization: %v, Country: %v, Region: %v, City: %v, Zip: %v, "Geolocation": %s
`, address, data["isp"], data["org"], data["country"], data["region"], data["city"], data["zip"],
		fmt.Sprintf(`https://www.google.com.ua/maps/@%s,%s,%sfz?hl=en`,
			x, y, "20"))
}

func getData(address string) map[string]interface{} {

	response, httpError := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", address))
	if httpError != nil {
		log.Fatal(httpError.Error())
		return nil
	}

	return readBody(response.Body)
}

func readBody(body io.ReadCloser) map[string]interface{} {

	var data map[string]interface{}

	readErr := json.NewDecoder(body).Decode(&data)
	if readErr != nil {
		log.Fatal(readErr.Error())
		return nil
	}

	return data
}

func GetCoordinates(address string) (string, string) {

	var data = getData(address)
	return fmt.Sprintf("%v", data["lat"]), fmt.Sprintf("%v", data["lon"])
}
