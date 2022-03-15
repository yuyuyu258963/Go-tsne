package main

import (
	"math"
	"math/rand"
)

// 生成 0 矩阵
func getZeroVec(n int, m int) [][]float64 {
	var zeroVec [][]float64
	for i := 0; i < n; i++ {
		zeroVec = append(zeroVec, make([]float64, m))
	}
	return zeroVec
}

// 生成符合正态分布的随机数
func GetGaussRandomNum() float64 {
	const min = -1
	const max = 1
	σ := (float64(min) + float64(max)) / 2
	μ := (float64(max) - σ) / 3
	// rand.Seed(time.Now().UnixNano())
	x := rand.Float64()
	x1 := rand.Float64()
	a := math.Cos(2*math.Pi*x) * math.Sqrt((-2)*math.Log(x1))
	result := a*μ + σ
	// fmt.Println(result)
	return result
}

// 生成随机数矩阵
func randomRandn(n int, m int) [][]float64 {
	var randn [][]float64
	for i := 0; i < n; i++ {
		var t []float64
		for j := 0; j < m; j++ {
			b := GetGaussRandomNum()
			t = append(t, b)
		}
		randn = append(randn, t)
	}
	return randn
}

// 获得一个转置的矩阵
func transpose(x [][]float64)( y [][]float64) {
	n, m := len(x), len(x[0])
	for i := 0; i < m; i++ {
		var t []float64
		for j := 0; j < n; j++ {
			t = append(t, x[j][i])
		}
		y = append(y, t)
	}
	// fmt.Println(y)
	return y
}

// 实现两个矩阵的加法
func add2vec(x, y [][]float64) [][]float64 {
	n, m := len(x), len(x[0])
	res := make([][]float64, n)
	for i := 0; i < n; i++ {
		t := make([]float64, m )
		res[i] = t
		for j := 0; j < m; j++ {
			res[i][j] = x[i][j] + y[i][j]
		}
	}
	return res
}


// 计算实现矩阵中的加法
func sumSelf(x [][]float64) (sum float64) {
	sum = 0.0
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[0]); j++ {
			sum += x[i][j]
		}
	}
	return sum
}

// 实现一个矩阵除以一个数
func division( x [][]float64, k	float64) [][]float64 {
	n, m := len(x), len(x[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x[i][j] /= k
		}
	}
	return x
}

// 计算两个向量的欧式距离
func cal_dist(a []float64, b []float64) float64 {
	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum += (a[i] - b[i]) * (a[i] - b[i])
	}
	return sum
}

// 避免矩阵中的元素过小
func maxminVec(a [][]float64){
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] < 1e-12 {
				a[i][j] = 1e-12
			}
		}
	}
}

// 对距离矩阵进行倒数
func recVec(x [][]float64) {
	n, m := len(x), len(x[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i != j {
				x[i][j] = 1 / ( 1 + x[i][j])
			} else {
				x[i][j] = 0.0
			}
		}
	}
}

// 实现两个举证的减法
func subtraction(x, y [][]float64) [][]float64 {
	n, m := len(x), len(x[0])
	res := make([][]float64, n)
	for i := 0; i < n; i++ {
		t := make([]float64, m)
		res[i] = t
		for j := 0; j < m; j++ {
			res[i][j] = x[i][j] - y[i][j]
		}
	}
	return res
}