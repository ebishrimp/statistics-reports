package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/benoitmasson/plotters/piechart"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Bloodtype_mainprocess(path string) {
	_, ratio := importData_b(path)
	plotBar(ratio)
	plotBand(ratio)
	pieChart(ratio)
}

func importData_b(path string) ([]int, []float64) {
	var A int
	var B int
	var AB int
	var O int
	var others int

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		switch line {
		case "A":
			A++
		case "B":
			B++
		case "AB":
			AB++
		case "O":
			O++
		default:
			others++
		}
	}
	scanner.Err()

	slice := []int{A, B, AB, O, others, A + B + AB + O + others}
	ratio := []float64{float64(slice[0]) / float64(slice[5]), float64(slice[1]) / float64(slice[5]), float64(slice[2]) / float64(slice[5]), float64(slice[3]) / float64(slice[5]), float64(slice[4]) / float64(slice[5])}
	var strslice []string
	var strratio []string
	for i := 0; i < 5; i++ {
		strslice = append(strslice, strconv.Itoa(slice[i]))
		strratio = append(strratio, strconv.FormatFloat(ratio[i], 'f', -1, 64))
	}
	fmt.Printf("%v\n", strslice)
	fmt.Printf("%v\n", strratio)
	return slice, ratio
}

func plotBar(data []float64) {
	grA := plotter.Values{data[0], data[1], data[2], data[3], data[4]}

	p := plot.New()
	p.Title.Text = "Blood Type Distribution"
	p.Y.Label.Text = "Ratio"

	w := vg.Points(20)
	barsA, err := plotter.NewBarChart(grA, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)

	p.Y.Max = 0.6

	p.Add(barsA)
	p.NominalX("A", "B", "AB", "O", "Others")

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/bloodtype_bar.png"); err != nil {
		panic(err)
	}
}

func plotBand(data []float64) {
	vA := plotter.Values{data[0]}
	vB := plotter.Values{data[0] + data[1]}
	vAB := plotter.Values{data[0] + data[1] + data[2]}
	vO := plotter.Values{data[0] + data[1] + data[2] + data[3]}
	vOther := plotter.Values{data[0] + data[1] + data[2] + data[3] + data[4]}

	w := vg.Points(20)

	p := plot.New()
	p.Title.Text = "Blood Type Distribution"
	p.Y.Label.Text = "ratio"

	barOther, _ := plotter.NewBarChart(vOther, w)
	barAB, _ := plotter.NewBarChart(vAB, w)
	barO, _ := plotter.NewBarChart(vO, w)
	barB, _ := plotter.NewBarChart(vB, w)
	barA, _ := plotter.NewBarChart(vA, w)
	barA.Color = plotutil.Color(0)
	barB.Color = plotutil.Color(1)
	barO.Color = plotutil.Color(2)
	barAB.Color = plotutil.Color(3)
	barOther.Color = plotutil.Color(4)

	p.Add(barOther)
	p.Add(barO)
	p.Add(barAB)
	p.Add(barB)
	p.Add(barA)
	p.NominalX("distribtion")

	p.Legend.Add("Others", barOther)
	p.Legend.Add("Type AB", barAB)
	p.Legend.Add("Type O", barO)
	p.Legend.Add("Type B", barB)
	p.Legend.Add("Type A", barA)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/bloodtype_band.png"); err != nil {
		panic(err)
	}
}

func pieChart(data []float64) {
	p := plot.New()
	p.Legend.Top = true
	p.HideAxes()

	var pies []*piechart.PieChart

	for i := 0; i < 5; i++ {
		pie, err := piechart.NewPieChart(plotter.Values{data[i]})
		if err != nil {
			panic(err)
		}
		pie.Color = plotutil.Color(i)
		var s float64
		for j := 0; j < len(data); j++ {
			s += data[j]
		}
		pie.Total = s
		for k := 0; k < len(pies); k++ {
			pie.Offset.Value += data[k]
		}
		pie.Labels.Values.Show = true
		pie.Labels.Values.Percentage = true
		pies = append(pies, pie)
	}

	pies[0].Labels.Nominal = []string{"Type A"}
	pies[1].Labels.Nominal = []string{"Type B"}
	pies[2].Labels.Nominal = []string{"Type AB"}
	pies[3].Labels.Nominal = []string{"Type O"}
	pies[4].Labels.Nominal = []string{"Others"}

	p.Add(pies[0], pies[1], pies[2], pies[3], pies[4])

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/bloodtype_pie.png"); err != nil {
		panic(err)
	}
}
