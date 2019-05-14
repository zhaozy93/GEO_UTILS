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
	flag := false
	var p1, p2 *Point
	for i := 0; i < len(polygon); i++ {
		p1 = polygon[i]
		if i == len(polygon)-1 {
			p2 = polygon[0]
		} else {
			p2 = polygon[i+1]
		}

		x1, y1, x2, y2, x, y := p1.Lat, p1.Lng, p2.Lat, p2.Lng, p.Lat, p.Lng

		// 两个点在射线的同一边
		if (y1 > y && y2 > y) || (y1 < y && y2 < y) {
			continue
		}
		if x1 < x && x2 < x {
			continue
		}

		// 点是否是某一个顶点
		if (x == x1 && y == y1) || (x == x2 && y == y2) {
			return 1
		}

		// 射线与边平行
		if y1 == y2 && y1 == y {
			continue
		}

		px := (x2-x1)/(y2-y1)*(y-y1) + x1
		// 点在线段上
		if px == x {
			return 1
		}
		// 未穿过
		if px < x {
			continue
		}

		// 判断端点是否在射线上
		if px == x1 && y == y1 {
			if y2 > y {
				continue
			}
		} else if px == x2 && y == y2 {
			if y1 > y {
				continue
			}
		}

		flag = !flag
	}
	if flag {
		return 0
	}
	return -1
}

// 是否在线段延长线上，是否在线段上
func onLine(x1, y1, x2, y2, x, y float64) (bool, bool) {
	if dcmp((y2-y)*(x1-x) - (y1-y)*(x2-x)) {
		if x1 < x2 && x > x1 && x < x2 {
			return true, true
		}
		if x1 > x2 && x < x1 && x > x2 {
			return true, true
		}
		return true, false
	}
	return false, false
}

func dcmp(x float64) bool {
	return math.Abs(x) < eps
}
