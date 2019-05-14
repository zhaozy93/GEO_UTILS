package main

import (
	"fmt"
	"math"
)

// 参考文档
// https://blog.csdn.net/liminlu0314/article/details/8553926
const R = 6378137.0

func main() {
	distance(116.40742, 39.90417, 121.4737021, 31.2303904)
}

// lat 精度 lng 维度
func distance(lng1, lat1, lng2, lat2 float64) float64 {
	OA, OB := R, R
	AOC, BOD, COD := lat1*math.Pi/180, lat2*math.Pi/180, (lng2-lng1)*math.Pi/180 // 公式1

	AC, OC, BD, OD := math.Sin(lat1*math.Pi/180)*R, math.Cos(lat1*math.Pi/180)*R, math.Sin(lat2*math.Pi/180)*R, math.Cos(lat2*math.Pi/180)*R // 公式2

	_, _, _, _, _, _, _, _, _ = OA, OB, AOC, BOD, COD, AC, OC, BD, OD
	// AC = ED
	// AE = CD
	// BE = BD-ED = BD-AC  // 公式3

	// AB *AB = AE*AE + BE*BE   // 公式4

	// AB * AB = AC *AC + BC *BC - 2*AC*BC*cos(ACB) // 公式5
	// CD * CD = OC*OC + OD*OD - 2*OC*OD*cos(COD)

	// 公式1235代入公式4
	// AB * AB = CD*CD + (BD-ED) * (BD-ED) = CD*CD + (BD-AC) * (BD-AC)

	// AB * AB = (OC*OC + OD*OD - 2*OC*OD*cos(COD)) * (OC*OC + OD*OD - 2*OC*OD*cos(COD))
	//           + 			(math.Sin(lat2)*R - 	math.Sin(lat1)*R)	 *	(math.Sin(lat2)*R - 	math.Sin(lat1)*R)
	// AB * AB = math.Pow(math.Pow(math.Cos(lat2)*R, 2)+math.Pow(math.Cos(lat2)*R, 2)-2*math.Cos(lat2)*R*math.Cos(lat2)*R*math.Cos(COD), 2) +
	// +math.Pow((math.Sin(lat2)*R-math.Sin(lat1)*R), 2)

	// CD*CD = OC*OC + OD*OD - 2*OC*OD*cos(COD)
	//       = math.Cos(lat1)*R * math.Cos(lat1)*R    +   math.Cos(lat2)*R * math.Cos(lat2)*R - 2 * math.Cos(lat1)*R *  math.Cos(lat2)*R * math.Cos(lng2-lng1)
	//       = R*R *( math.Cos(lat1)*math.Cos(lat1) - math.Cos(lat2)*math.Cos(lat2) - 2*math.Cos(lat1)*math.Cos(lat2)*math.Cos(lng2-lng1) )

	// (BD-ED) * (BD-ED) = (BD-AC) * (BD-AC)
	//                   = (math.Sin(lat2)*R - 	math.Sin(lat1)*R)	 *	(math.Sin(lat2)*R - 	math.Sin(lat1)*R)
	//                   = (R*R * (math.Sin(lat2)-math.Sin(lat1))) * (R*R * (math.Sin(lat2)-math.Sin(lat1)))
	//                   = R*R*math.Sin(lat2)*math.Sin(lat2) + R*R*math.Sin(lat1)*math.Sin(lat1) - 2*R*R*math.Sin(lat2)*math.Sin(lat1)
	//                   = R*R*(math.Sin(lat2)*math.Sin(lat2) + math.Sin(lat1)*math.Sin(lat1) - 2*math.Sin(lat2)*math.Sin(lat1))
	// 根据上面其实可以得出AB*AB的结果

	// AB*AB = 2*R*R*(1-math.Cos(AOC)*math.Cos(BOD)*math.Cos(COD)-math.Sin(AOC)*math.Sin(BOD))

	// 同时AB*AB = AO*AO+BO*BO-2*AO*BO*cos(AOB)
	// cos(AOB) = (AB*AB- AO*AO-BO*BO) / (2*AO*BO)
	//          = (2*R*R*(1-math.Cos(AOC)*math.Cos(BOD)*math.Cos(COD)-math.Sin(AOC)*math.Sin(BOD)) - 2*R*R)  / (-2*R*R)
	//          = math.Cos(AOC)*math.Cos(BOD)*math.Cos(COD)+math.Sin(AOC)*math.Sin(BOD)
	AOB := math.Cos(AOC)*math.Cos(BOD)*math.Cos(COD) + math.Sin(AOC)*math.Sin(BOD)
	AB := math.Acos(AOB) * R
	fmt.Println(AB)
	return AB
}
