package main

import "math"

func main() {
	pointInPolygon(&Point{0.5, 0.5},
		[]*Point{&Point{0, 0}, &Point{1, 0}, &Point{1, 1}, &Point{0, 1}})
	pointInPolygon(&Point{0, 0.5},
		[]*Point{&Point{0, 0}, &Point{1, 0}, &Point{1, 1}, &Point{0, 1}})
	pointInPolygon(&Point{-1, 0.5},
		[]*Point{&Point{0, 0}, &Point{1, 0}, &Point{1, 1}, &Point{0, 1}})
}

type Point struct {
	Lng float64
	Lat float64
}

const eps = 1e-6

// -1 表示在外部
// 0  表示内部
// 1  表示点重合或者在某一条边上
func pointInPolygon(p *Point, polygon []*Point) int {
	if polygon == nil || len(polygon) < 3 {
		return -1
	}
	sum := float64(0)
	var p1, p2 *Point
	for i := 0; i < len(polygon); i++ {
		p1 = polygon[i]
		if i == len(polygon)-1 {
			p2 = polygon[0]
		} else {
			p2 = polygon[i+1]
		}

		x1, y1, x2, y2, x, y := p1.Lat, p1.Lng, p2.Lat, p2.Lng, p.Lat, p.Lng

		// 点是否是某一个顶点
		if (x == x1 && y == y1) || (x == x2 && y == y2) {
			return 1
		}
		// 点在线段上
		if y2 == y1 && y == y1 {
			if x2 > x1 {
				if x > x1 && x < x2 {
					return 1
				}
			}
			if x2 < x1 {
				if x < x1 && x > x2 {
					return 1
				}
			}
		}

		if y2 != y1 && dcmp(x-(x2-x1)/(y2-y1)*(y-y1)-x1) {
			return 1
		}

		// 点与相邻顶点连线的夹角
		angle := math.Atan2(y2-y, x2-x) - math.Atan2(y1-y, x1-x)

		// 确保夹角不超出取值范围（-π 到 π）
		if angle >= math.Pi {
			angle = angle - math.Pi*2
		} else if angle <= -math.Pi {
			angle = angle + math.Pi*2
		}

		sum += angle
	}
	if math.Round(sum/math.Pi) == 0 {
		return -1
	}
	return 0
}

func dcmp(x float64) bool {
	return math.Abs(x) < eps
}
