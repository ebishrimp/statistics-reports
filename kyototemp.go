package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type Data struct {
	date string
	ave  float64
	max  float64
	min  float64
}

type Analysis struct {
	label string
	ave   float64
	v     float64
	s     float64
	min   float64
	q1    float64
	cen   float64
	q3    float64
	max   float64
}

var temperatures []Data

var avepr Analysis
var maxpr Analysis
var minpr Analysis

func temp_mainprocess(path string) {
	temp_importdata(path)
	avepr = analyze(temperatures, "ave")
	maxpr = analyze(temperatures, "max")
	minpr = analyze(temperatures, "min")
	fmt.Printf("%+v\n", avepr)
	fmt.Printf("%+v\n", maxpr)
	fmt.Printf("%+v\n", minpr)
	plot_hist(temperatures)
	plot_line(temperatures)
}

func temp_importdata(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	temp_scanner := bufio.NewScanner(f)

	flag := true
	for temp_scanner.Scan() {
		var tempdata Data

		if flag {
			flag = false
			continue
		}

		line := temp_scanner.Text()

		strs := strings.Split(line, ",")

		tempdata.date = strs[0]
		tempdata.ave, err = strconv.ParseFloat(strs[1], 64)
		if err != nil {
			panic(err)
		}
		tempdata.max, err = strconv.ParseFloat(strs[2], 64)
		if err != nil {
			panic(err)
		}
		tempdata.min, err = strconv.ParseFloat(strs[3], 64)
		if err != nil {
			panic(err)
		}
		temperatures = append(temperatures, tempdata)
	}
	temp_scanner.Err()
}

func analyze(data []Data, sub string) Analysis {
	var d []float64
	var res Analysis

	len := len(data)
	switch sub {
	case "ave":
		for i := 0; i < len; i++ {
			d = append(d, data[i].ave)
		}
	case "max":
		for i := 0; i < len; i++ {
			d = append(d, data[i].max)
		}
	case "min":
		for i := 0; i < len; i++ {
			d = append(d, data[i].min)
		}
	}

	res.label = sub
	res.ave = stat.Mean(d, nil)
	res.v = stat.Variance(d, nil)
	res.s = math.Sqrt(res.v)
	res.min = floats.Min(d)
	res.max = floats.Max(d)
	sort.Float64s(d)
	res.q1 = stat.Quantile(0.25, stat.Empirical, d, nil)
	res.cen = stat.Quantile(0.50, stat.Empirical, d, nil)
	res.q3 = stat.Quantile(0.75, stat.Empirical, d, nil)

	return res
}

func plot_hist(data []Data) {
	length := len(data)
	ave := make(plotter.Values, length)
	max := make(plotter.Values, length)
	min := make(plotter.Values, length)
	diff := make(plotter.Values, length)
	for i := 0; i < len(data); i++ {
		ave[i] = data[i].ave
		max[i] = data[i].max
		min[i] = data[i].min
		diff[i] = data[i].max - data[i].min
	}

	p_temp1 := plot.New()
	p_temp2 := plot.New()
	p_temp3 := plot.New()
	p_temp4 := plot.New()
	var p []plot.Plot = []plot.Plot{*p_temp1, *p_temp2, *p_temp3, *p_temp4}

	p[0].Title.Text = "Average Histogram"
	p[1].Title.Text = "Max Histogram"
	p[2].Title.Text = "Min Histogram"
	p[3].Title.Text = "difference between max and min"

	var h []plotter.Histogram

	h_temp, err := plotter.NewHist(ave, 16)
	if err != nil {
		panic(err)
	}
	h = append(h, *h_temp)
	h_temp, err = plotter.NewHist(max, 16)
	if err != nil {
		panic(err)
	}
	h = append(h, *h_temp)
	h_temp, err = plotter.NewHist(min, 16)
	if err != nil {
		panic(err)
	}
	h = append(h, *h_temp)
	h_temp, err = plotter.NewHist(diff, 16)
	if err != nil {
		panic(err)
	}
	h = append(h, *h_temp)

	for i := 0; i < 4; i++ {
		h[i].Normalize(1)
		temp := h[i]
		p[i].Add(&temp)
	}

	if err := p[0].Save(4*vg.Inch, 4*vg.Inch, "./images/temp_average.png"); err != nil {
		panic(err)
	}
	if err := p[1].Save(4*vg.Inch, 4*vg.Inch, "./images/temp_max.png"); err != nil {
		panic(err)
	}
	if err := p[2].Save(4*vg.Inch, 4*vg.Inch, "./images/temp_min.png"); err != nil {
		panic(err)
	}
	if err := p[3].Save(4*vg.Inch, 4*vg.Inch, "./images/temp_diff.png"); err != nil {
		panic(err)
	}
}

func plot_line(data []Data) {
	length := len(data)
	date := make([]string, length)
	ave := make(plotter.Values, length)
	max := make(plotter.Values, length)
	min := make(plotter.Values, length)
	diff := make(plotter.Values, length)
	for i := 0; i < length; i++ {
		date[i] = data[i].date
		ave[i] = data[i].ave
		max[i] = data[i].max
		min[i] = data[i].min
		diff[i] = data[i].max - data[i].min
	}

	p := plot.New()
	pd := plot.New()

	p.Title.Text = "Temperature Trend"
	pd.Title.Text = "Difference Trend"
	p.X.Label.Text = "Date"
	pd.X.Label.Text = "Date"
	p.Y.Label.Text = "Temperature (C)"
	pd.Y.Label.Text = "Temperature (C)"

	pts1 := make(plotter.XYs, length)
	pts2 := make(plotter.XYs, length)
	pts3 := make(plotter.XYs, length)
	pts4 := make(plotter.XYs, length)

	var ticks []plot.Tick
	var lastMonth time.Month

	for i := 0; i < length; i++ {

		t, err := time.Parse("2006/1/2", date[i])
		if err != nil {
			panic(err)
		}

		pts1[i].X = float64(t.Unix())
		pts2[i].X = float64(t.Unix())
		pts3[i].X = float64(t.Unix())
		pts4[i].X = float64(t.Unix())

		pts1[i].Y = ave[i]
		pts2[i].Y = max[i]
		pts3[i].Y = min[i]
		pts4[i].Y = diff[i]

		if t.Month() != lastMonth {
			ticks = append(ticks, plot.Tick{
				Value: float64(t.Unix()),
				Label: t.Format("2006/01"),
			})
			lastMonth = t.Month()
		}
	}

	l1, err := plotter.NewLine(pts1)
	l2, err := plotter.NewLine(pts2)
	l3, err := plotter.NewLine(pts3)
	l4, err := plotter.NewLine(pts4)
	if err != nil {
		panic(err)
	}
	l1.Color = plotutil.Color(0)
	l2.Color = plotutil.Color(1)
	l3.Color = plotutil.Color(2)
	l4.Color = plotutil.Color(0)
	p.Legend.Add("ave for each day", l1)
	p.Legend.Add("max for each day", l2)
	p.Legend.Add("min for each day", l3)
	pd.Legend.Add("difference", l4)
	p.Add(l1, l2, l3)
	pd.Add(l4)

	p.X.Tick.Marker = plot.ConstantTicks(ticks)
	pd.X.Tick.Marker = plot.ConstantTicks(ticks)

	if err := p.Save(8*vg.Inch, 4*vg.Inch, "./images/temperature_line.png"); err != nil {
		panic(err)
	}
	if err := pd.Save(8*vg.Inch, 4*vg.Inch, "./images/temperature_diff_line.png"); err != nil {
		panic(err)
	}
}
