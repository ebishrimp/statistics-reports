package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/benoitmasson/plotters/piechart"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

var data []string

type Cnt struct {
	kinoko   float64
	takenoko float64
	sumkt    float64
	both     float64
	notboth  float64
	sumb     float64
	l5       float64
	l4       float64
	l3       float64
	l2       float64
	l1       float64
	sumc     float64
}

var cnt Cnt

func Kinokotakenoko_mainprocess(path string) {
	data = importData_kt(path)
	cnt = kt_count(data)
	kt_plotbar(cnt)
	kt_plotBand(cnt)
	kt_pieChart(cnt)
	fmt.Printf("%g\n ", cnt.kinoko)
	fmt.Printf("%g\n ", cnt.takenoko)
	fmt.Printf("%g\n ", cnt.both)
	fmt.Printf("%g\n ", cnt.notboth)
}

func importData_kt(path string) []string {
	var strs []string

	var cross []int = []int{0, 0, 0, 0} // k&b k&nb t&b t&nb

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	kt_scanner := bufio.NewScanner(f)

	flag := true
	for kt_scanner.Scan() {
		if flag {
			flag = false
			continue
		}
		line := kt_scanner.Text()

		temp := strings.Split(line, ",")

		if temp[0] == "きのこの山" {
			if temp[1] == "両方いける" {
				cross[0]++
			} else {
				cross[1]++
			}
		}
		if temp[0] == "たけのこの里" {
			if temp[1] == "両方いける" {
				cross[2]++
			} else {
				cross[3]++
			}
		}

		strs = append(strs, temp...)
	}
	kt_scanner.Err()

	fmt.Printf("%v\n", cross)

	return strs
}

func kt_count(data []string) Cnt {
	temp := Cnt{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < len(data); i++ {
		switch data[i] {
		case "きのこの山":
			temp.kinoko++
		case "たけのこの里":
			temp.takenoko++
		case "両方いける":
			temp.both++
		case "片方だけ":
			temp.notboth++
		case "5":
			temp.l5++
		case "4":
			temp.l4++
		case "3":
			temp.l3++
		case "2":
			temp.l2++
		case "1":
			temp.l1++
		}
	}
	temp.sumkt = temp.kinoko + temp.takenoko
	temp.sumb = temp.both + temp.notboth
	temp.sumc = temp.l1 + temp.l2 + temp.l3 + temp.l4 + temp.l5
	return temp
}

func kt_plotbar(cnt Cnt) {
	sum1 := cnt.kinoko + cnt.takenoko
	grA := plotter.Values{cnt.kinoko / sum1, cnt.takenoko / sum1}

	p := plot.New()
	p.Title.Text = "Kinoko or Takenoko Distribution"
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
	p.NominalX("Kinoko", "Takenoko")

	if err := p.Save(3*vg.Inch, 4*vg.Inch, "./images/kinokoOrTakenoko_bar.png"); err != nil {
		panic(err)
	}

	sum2 := cnt.both + cnt.notboth
	grB := plotter.Values{cnt.both / sum2, cnt.notboth / sum2}

	p1 := plot.New()
	p1.Title.Text = "Both or Not Distribution"
	p1.Y.Label.Text = "Ratio"

	barsB, err := plotter.NewBarChart(grB, w)
	if err != nil {
		panic(err)
	}

	barsB.LineStyle.Width = vg.Length(0)
	barsB.Color = plotutil.Color(0)

	p1.Y.Max = 0.6

	p1.Add(barsB)
	p1.NominalX("Both", "Not")

	if err := p.Save(3*vg.Inch, 4*vg.Inch, "./images/bothOrNot_bar.png"); err != nil {
		panic(err)
	}

	sum3 := cnt.l1 + cnt.l2 + cnt.l3 + cnt.l4 + cnt.l5
	grC := plotter.Values{cnt.l1 / sum3, cnt.l2 / sum3, cnt.l3 / sum3, cnt.l4 / sum3, cnt.l5 / sum3}

	p3 := plot.New()
	p3.Title.Text = "Chocolate Distribution"
	p3.Y.Label.Text = "Ratio"

	w = vg.Points(20)
	barsC, err := plotter.NewBarChart(grC, w)
	if err != nil {
		panic(err)
	}
	barsC.LineStyle.Width = vg.Length(0)
	barsC.Color = plotutil.Color(0)

	p3.Y.Max = 0.7

	p3.Add(barsC)
	p3.NominalX("1", "2", "3", "4", "5")

	if err := p3.Save(4*vg.Inch, 4*vg.Inch, "./images/chocolate_bar.png"); err != nil {
		panic(err)
	}
}

func kt_plotBand(cnt Cnt) {
	k := plotter.Values{cnt.kinoko / cnt.sumkt}
	kt := plotter.Values{(cnt.kinoko + cnt.takenoko) / cnt.sumkt}
	w := vg.Points(20)

	p := plot.New()
	p.Title.Text = "Kinoko or Takenoko Distribution"
	p.Y.Label.Text = "Ratio"

	barK, _ := plotter.NewBarChart(k, w)
	barKT, _ := plotter.NewBarChart(kt, w)
	barK.Color = plotutil.Color(0)
	barKT.Color = plotutil.Color(1)

	p.Add(barKT)
	p.Add(barK)
	p.NominalX("distribution")

	p.Legend.Add("Kinoko", barK)
	p.Legend.Add("Takenoko", barKT)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "./images/kinokoOrTakenoko_band.png"); err != nil {
		panic(err)
	}

	b := plotter.Values{cnt.both / cnt.sumb}
	nb := plotter.Values{(cnt.both + cnt.notboth) / cnt.sumb}

	p2 := plot.New()
	p2.Title.Text = "Both or Not Distribution"
	p2.Y.Label.Text = "Ratio"

	barB, _ := plotter.NewBarChart(b, w)
	barNB, _ := plotter.NewBarChart(nb, w)
	barB.Color = plotutil.Color(0)
	barNB.Color = plotutil.Color(1)
	p2.Add(barNB)
	p2.Add(barB)
	p2.NominalX("Distribution")

	p2.Legend.Add("Both", barB)
	p2.Legend.Add("not Both", barNB)

	if err := p2.Save(4*vg.Inch, 4*vg.Inch, "./images/bothOrNot_band.png"); err != nil {
		panic(err)
	}

	l1 := plotter.Values{cnt.l1 / cnt.sumc}
	l2 := plotter.Values{(cnt.l1 + cnt.l2) / cnt.sumc}
	l3 := plotter.Values{(cnt.l1 + cnt.l2 + cnt.l3) / cnt.sumc}
	l4 := plotter.Values{(cnt.l1 + cnt.l2 + cnt.l3 + cnt.l4) / cnt.sumc}
	l5 := plotter.Values{(cnt.l1 + cnt.l2 + cnt.l3 + cnt.l4 + cnt.l5) / cnt.sumc}

	p3 := plot.New()
	p3.Title.Text = "How much like chocolate"
	p3.Y.Label.Text = "Ratio"

	bar5, _ := plotter.NewBarChart(l5, w)
	bar4, _ := plotter.NewBarChart(l4, w)
	bar3, _ := plotter.NewBarChart(l3, w)
	bar2, _ := plotter.NewBarChart(l2, w)
	bar1, _ := plotter.NewBarChart(l1, w)
	bar1.Color = plotutil.Color(0)
	bar2.Color = plotutil.Color(1)
	bar3.Color = plotutil.Color(2)
	bar4.Color = plotutil.Color(3)
	bar5.Color = plotutil.Color(4)

	p3.Add(bar5)
	p3.Add(bar4)
	p3.Add(bar3)
	p3.Add(bar2)
	p3.Add(bar1)
	p3.NominalX("distribution")

	p3.Legend.Add("1", bar1)
	p3.Legend.Add("2", bar2)
	p3.Legend.Add("3", bar3)
	p3.Legend.Add("4", bar4)
	p3.Legend.Add("5", bar5)

	if err := p3.Save(4*vg.Inch, 4*vg.Inch, "./images/chocolate_band.png"); err != nil {
		panic(err)
	}
}

func kt_pieChart(cnt Cnt) {
	p1 := plot.New()
	p1.Legend.Top = true
	p1.HideAxes()

	var pies1 []*piechart.PieChart

	pie, err := piechart.NewPieChart(plotter.Values{cnt.kinoko})
	if err != nil {
		panic(err)
	}
	pie.Color = plotutil.Color(0)
	pie.Total = cnt.sumkt
	pie.Labels.Values.Show = true
	pie.Labels.Values.Percentage = true
	pies1 = append(pies1, pie)
	pie, err = piechart.NewPieChart(plotter.Values{cnt.takenoko})
	if err != nil {
		panic(err)
	}
	pie.Color = plotutil.Color(1)
	pie.Total = cnt.sumkt
	pie.Offset.Value += cnt.kinoko
	pie.Labels.Values.Show = true
	pie.Labels.Values.Percentage = true
	pies1 = append(pies1, pie)

	pies1[0].Labels.Nominal = []string{"Kinoko"}
	pies1[1].Labels.Nominal = []string{"Takenoko"}

	p1.Add(pies1[0], pies1[1])

	if err := p1.Save(4*vg.Inch, 4*vg.Inch, "./images/kinokoOrTakenoko_pie.png"); err != nil {
		panic(err)
	}

	p2 := plot.New()
	p2.Legend.Top = true
	p2.HideAxes()

	var pies2 []*piechart.PieChart

	pie, err = piechart.NewPieChart(plotter.Values{cnt.both})
	if err != nil {
		panic(err)
	}
	pie.Color = plotutil.Color(0)
	pie.Total = cnt.sumb
	pie.Labels.Values.Show = true
	pie.Labels.Values.Percentage = true
	pies2 = append(pies2, pie)
	pie, err = piechart.NewPieChart(plotter.Values{cnt.notboth})
	if err != nil {
		panic(err)
	}
	pie.Color = plotutil.Color(1)
	pie.Total = cnt.sumb
	pie.Offset.Value += cnt.both
	pie.Labels.Values.Show = true
	pie.Labels.Values.Percentage = true
	pies2 = append(pies2, pie)

	pies2[0].Labels.Nominal = []string{"Both"}
	pies2[1].Labels.Nominal = []string{"not Both"}

	p2.Add(pies2[0], pies2[1])

	if err := p2.Save(4*vg.Inch, 4*vg.Inch, "./images/bothOrNot_pie.png"); err != nil {
		panic(err)
	}

	p3 := plot.New()
	p3.Legend.Top = true
	p3.HideAxes()

	var pies3 []*piechart.PieChart

	data := []float64{cnt.l1, cnt.l2, cnt.l3, cnt.l4, cnt.l5}

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
		for k := 0; k < len(pies3); k++ {
			pie.Offset.Value += data[k]
		}
		pie.Labels.Values.Show = true
		pie.Labels.Values.Percentage = true
		pies3 = append(pies3, pie)
	}

	pies3[0].Labels.Values.Percentage = false

	pies3[1].Labels.Nominal = []string{"2"}
	pies3[2].Labels.Nominal = []string{"3"}
	pies3[3].Labels.Nominal = []string{"4"}
	pies3[4].Labels.Nominal = []string{"5"}

	p3.Add(pies3[0], pies3[1], pies3[2], pies3[3], pies3[4])

	if err := p3.Save(4*vg.Inch, 4*vg.Inch, "./images/chocolate_pie.png"); err != nil {
		panic(err)
	}
}
