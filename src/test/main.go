package main

import(
	"fmt"
	"math"
	"net/http"
	"log"
	"image" 
	"image/color" 
	"image/png" 
	"math/cmplx" 
)
const(
	width,height=600,320
	cells=100
	xyrange=30.0
	xyscale=width/2/xyrange
	zscale=height*0.4
	angle=math.Pi/6
)
var sin30,cos30=math.Sin(angle),math.Cos(angle)
type map_value  map[string][]string
func test3() {
	m:=map_value{"lang":{"ch","en"}}
	m["item"]=append(m["item"],"item1")
	fmt.Println(m["lang"])
	fmt.Println(m["item"])
	m=nil
	m["item"]=append(m["item"],"item2")
	fmt.Println(m["item"])
}
type ByteCounter int
func (c *ByteCounter)Write(p []byte)(int,error) {
	*c+=ByteCounter(len(p))
	return len(p),nil	
}
func test4() {
	var c ByteCounter
	c.Write([]byte("hellooo"))
	fmt.Println(c)
	var name="Dolly"
	c=0
	fmt.Fprintf(&c,"hello,%s",name)
	fmt.Println(c)
}
func main() {

	test4()
	// startHttpServer()
	
}
func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5) 
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
func svg(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")
	body1:=[]byte(fmt.Sprintf("<div><svg xmlns='http://www.w3.org/2000/svg' "+
				"style='stroke: grey; fill: white; stroke-width: 0.7' "+
				"width='%d' height='%d'>",width,height))
	w.Write(body1)
	for i:=0;i<cells;i++{
		for j:=0;j<cells;j++{
			ax,ay:=corner(i+1,j)
			bx,by:=corner(i,j)
			cx,cy:=corner(i,j+1)
			dx,dy:=corner(i+1,j+1)
			body2:=[]byte(fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",ax,ay,bx,by,cx,cy,dx,dy))
			w.Write(body2)
		}
	}
	w.Write([]byte("</svg></div>"))
	// fmt.Fprintln(w, "finish")
}
func mandelbrots(w http.ResponseWriter, r *http.Request) {
	const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
		x := float64(px)/width*(xmax-xmin) + xmin 
		z := complex(x, y)
		// Image point (px, py) represents complex value z.
		img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ { v = v*v + z
	if cmplx.Abs(v) > 2 {
	return color.Gray{255 - contrast*n}
	}
	}
	return color.Black
}
func startHttpServer() {
    http.HandleFunc("/svg", svg)
    http.HandleFunc("/mandelbrot", mandelbrots)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}