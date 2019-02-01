package main

import (
	"fmt"
	"math"
	"net/http"
	"log"
	_"os"
	"io"

)
const (
	width,height = 600,300
	cells = 100
	xyrang = 30.0
	xyscale = width / 2 / xyrang
	zscale = height * 0.4
	angle = math.Pi / 6
)

var sin30,cos30 = math.Sin(angle),math.Cos(angle)
func main()  {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8888",nil))
}

func handler(w http.ResponseWriter,r *http.Request) {
	log.Printf("URL.Path = %q\n",r.URL.Path)
	w.Header().Set("Content-Type","image/svg+xml")
	draw(w)
}
func draw(out io.Writer)  {

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'" +"style='stroke:gery;fill:white;stroke-width:0.7'"+"width:'%d' height='%d'>",width,height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax,ay := corner(i+1,j)
			bx,by := corner(i,j)
			cx,cy := corner(i,j+1)
			dx,dy := corner(i+1,j+1)
			fmt.Printf("<polygon points= '%g,%g %g,%g %g,%g %g, %g'/>\n",ax,ay,bx,by,cx,cy,dx,dy)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int)(float64,float64) {
	x := xyrang *(float64(i)/cells - 0.5)
	y := xyrang *(float64(j)/cells - 0.5)

	z := f(x,y)

	sx := width/2 + (x-y) *cos30 *xyscale
	sy := height/2 + (x+y) *sin30 *xyscale - z*zscale
	return sx,sy

}
func f(x,y float64) float64 {
	r := math.Hypot(x,y)
	return math.Sin(r)/r
}