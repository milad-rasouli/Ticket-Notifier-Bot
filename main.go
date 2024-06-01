package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Milad75Rasouli/Ticket-Notifier-Bot/internal/entity"
)

func main() {
	// body := `{
	// 	"buses": [
	// 		{
	// 			"id": 43486990,
	// 			"from": 31310000,
	// 			"to": 11321006,
	// 			"fromName": "مشهد",
	// 			"toName": "تهران (جنوب)",
	// 			"corporation": "تعاونی 16 جهان گشت مهر",
	// 			"departureTime": "2024-06-03T08:00:00",
	// 			"arrivalTime": "2024-06-03T20:00:00",
	// 			"price": 4000000,
	// 			"serviceNumber": "",
	// 			"capacity": 0,
	// 			"busType": "VIP 2+1 / برق 220 ولت / سیستم تهویه مطبوع / تخت شو",
	// 			"superCorporationID": 25,
	// 			"superCorporation": "جهان گشت",
	// 			"toTerminal": "جنوب",
	// 			"fromCity": "مشهد",
	// 			"toCity": "تهران",
	// 			"distance": 881,
	// 			"providerID": 3,
	// 			"penaltyRates": [
	// 				{
	// 					"actualHoursBefore": 1.0,
	// 					"hoursBefore": 1.0,
	// 					"hoursBeforeText": null,
	// 					"percent": 10,
	// 					"customText": null
	// 				},
	// 				{
	// 					"actualHoursBefore": null,
	// 					"hoursBefore": null,
	// 					"hoursBeforeText": null,
	// 					"percent": 50,
	// 					"customText": "درصورت عدم چاپ یا استرداد بلیت توسط ترمینال به صورت اینترنتی"
	// 				}
	// 			],
	// 			"busTypeID": 4529517,
	// 			"discount": 0,
	// 			"isVIP": true,
	// 			"points": 2,
	// 			"precisionPoints": 2.5,
	// 			"additionalInfo": null,
	// 			"finalDestinationCity": "تهران",
	// 			"freeCancellationMins": null,
	// 			"weekday": "دوشنبه",
	// 			"dateString": "14 خرداد",
	// 			"fromTerminalSpecified": false,
	// 			"toTerminalSpecified": true,
	// 			"fromTerminalName": null,
	// 			"arrivalDateString": "14 خرداد",
	// 			"corportaionID": 1934,
	// 			"canCancelAfterClose": true,
	// 			"corporationFailed": false,
	// 			"corporationFailedUntil": null,
	// 			"preferredProvider": 1,
	// 			"description": "امکانات سرویس: برق 220 ولت / سیستم تهویه مطبوع / تخت شو، ########### آقای زینلی",
	// 			"unmaskedDescription": null,
	// 			"fromDatabase": true,
	// 			"intermediateDestinations": [],
	// 			"penaltyText": "این بلیط تا 1 ساعت قبل از حرکت  با 10٪ جریمه و پس از آن با 50٪ جریمه درصورت عدم چاپ یا استرداد بلیت توسط ترمینال به صورت اینترنتی قابل کنسل می باشد.\nدر صورت کنسلی حضوری بلیط توسط مسافر از طریق پایانه مسافربری، مِستر بلیط مسئولیتی در قبال استرداد مبلغ باقیمانده ندارد.درصورتی که بلیط  خریداری شده برای تعداد دو صندلی و بیشتر باشد، امکان کنسلی جداگانه به‌صورت آنلاین وجود ندارد.",
	// 			"customOrder": 3,
	// 			"reservable": true,
	// 			"isCar": false,
	// 			"needsNationalCode": false,
	// 			"needsSinglePerson": true,
	// 			"needsSelectSeat": true,
	// 			"briefDescription": null,
	// 			"lables": [],
	// 			"hasTempReserve": false,
	// 			"pngLogo": "Logos/png/0.png",
	// 			"svgLogo": "Logos/svg/0.svg",
	// 			"hasNoPassengers": false,
	// 			"isExclusive": false
	// 		},
	// 		{
	// 			"id": 43636888,
	// 			"from": 31310000,
	// 			"to": 11320000,
	// 			"fromName": "مشهد",
	// 			"toName": "تهران",
	// 			"corporation": "شرکت شماره هشت لوان نور مشهد",
	// 			"departureTime": "2024-06-03T09:00:00",
	// 			"arrivalTime": "2024-06-03T21:00:00",
	// 			"price": 4000000,
	// 			"serviceNumber": "",
	// 			"capacity": 5,
	// 			"busType": "VIP شارژر دار",
	// 			"superCorporationID": 7,
	// 			"superCorporation": "لوان نور ",
	// 			"fromCity": "مشهد",
	// 			"toCity": "تهران",
	// 			"distance": 881,
	// 			"providerID": 1,
	// 			"penaltyRates": [
	// 				{
	// 					"actualHoursBefore": 1.0,
	// 					"hoursBefore": 1.0,
	// 					"hoursBeforeText": null,
	// 					"percent": 10,
	// 					"customText": null
	// 				},
	// 				{
	// 					"actualHoursBefore": null,
	// 					"hoursBefore": null,
	// 					"hoursBeforeText": null,
	// 					"percent": 50,
	// 					"customText": null
	// 				}
	// 			],
	// 			"busTypeID": 4526367,
	// 			"discount": 0,
	// 			"isVIP": true,
	// 			"points": 2,
	// 			"precisionPoints": 2.8,
	// 			"additionalInfo": null,
	// 			"finalDestinationCity": "تهران",
	// 			"freeCancellationMins": null,
	// 			"weekday": "دوشنبه",
	// 			"dateString": "14 خرداد",
	// 			"fromTerminalSpecified": false,
	// 			"toTerminalSpecified": false,
	// 			"fromTerminalName": null,
	// 			"arrivalDateString": "14 خرداد",
	// 			"corportaionID": 1682,
	// 			"canCancelAfterClose": true,
	// 			"corporationFailed": false,
	// 			"corporationFailedUntil": null,
	// 			"preferredProvider": 1,
	// 			"description": "پایانه جنوب-زینلی ###########",
	// 			"unmaskedDescription": null,
	// 			"fromDatabase": false,
	// 			"intermediateDestinations": [],
	// 			"penaltyText": "این بلیط تا 1 ساعت قبل از حرکت  با 10٪ جریمه و پس از آن با 50٪ جریمه درصورت عدم چاپ یا استرداد بلیت توسط ترمینال به صورت اینترنتی قابل کنسل می باشد.\nدر صورت کنسلی حضوری بلیط توسط مسافر از طریق پایانه مسافربری، مِستر بلیط مسئولیتی در قبال استرداد مبلغ باقیمانده ندارد.درصورتی که بلیط  خریداری شده برای تعداد دو صندلی و بیشتر باشد، امکان کنسلی جداگانه به‌صورت آنلاین وجود ندارد.",
	// 			"customOrder": 3,
	// 			"reservable": true,
	// 			"isCar": false,
	// 			"needsNationalCode": false,
	// 			"needsSinglePerson": true,
	// 			"needsSelectSeat": true,
	// 			"briefDescription": null,
	// 			"lables": [],
	// 			"hasTempReserve": true,
	// 			"pngLogo": "Logos/png/0.png",
	// 			"svgLogo": "Logos/svg/0.svg",
	// 			"hasNoPassengers": false,
	// 			"isExclusive": false
	// 		}
	// 	]
	// }`

	var (
		url = "https://bus.mrbilit.ir/api/GetBusServices"
	)
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{
		"from": 31310000,
		"to": 11320000,
		"date": "2024-06-03",
		"includeClosed": true,
		"includePromotions": true,
		"loadFromDbOnUnavailability": true,
		"includeUnderDevelopment": true
	}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	var bus entity.Bus
	err = json.Unmarshal(body, &bus)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", bus)
}
