package main

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var x_values []string
var systolic []opts.LineData
var diastolic []opts.LineData

func httpserver(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Method : ", req.Method)
	fmt.Println("URL : ", req.URL)
	fmt.Println("Header : ", req.Header)

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "My Blood Pressure",
			Subtitle: "Hojoon Son",
		}))

	switch req.Method {
	case http.MethodPost:
		fmt.Print(req.FormValue(("date")))
		fmt.Print(req.FormValue(("systolic")))
		fmt.Print(req.FormValue(("diastolic")))
		x_values = append(x_values, req.FormValue("date"))
		systolic = append(systolic, opts.LineData{Value: req.FormValue("systolic")})
		diastolic = append(diastolic, opts.LineData{Value: req.FormValue("diastolic")})
	}

	line.SetXAxis(x_values).
		AddSeries("Systolic", systolic).
		AddSeries("Diastolic", diastolic).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8981", nil)
}
