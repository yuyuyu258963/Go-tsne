package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 生成 0 矩阵
func getZeroVec(n int, m int) [][]float64 {
	var zeroVec = make([][]float64, n)
	for i := 0; i < n; i++ {
		var t = make([]float64, m)
		zeroVec[i] = t
		for j := 0; j <m; j++ {
			zeroVec[i][j] = 0.0
		}
	}
	return zeroVec
}

// 生成符合正态分布的随机数
func GetGaussRandomNum() float64 {
	const min = -1
	const max = 1
	σ := (float64(min) + float64(max)) / 2
	μ := (float64(max) - σ) / 3
	time.Sleep(time.Millisecond * 2)
	rand.Seed(time.Now().UnixNano())
	x := rand.Float64()
	x1 := rand.Float64()
	fmt.Println(x)
	a := math.Cos(2*math.Pi*x) * math.Sqrt((-2)*math.Log(x1))
	result := a*μ + σ
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

// 对距离矩阵进行倒数 并让主对角线元素变为 0
func recVec(x [][]float64) {
	n, m := len(x), len(x[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i != j {
				x[i][j] = 1 / x[i][j]
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

// 返回一列的矩阵
func getLine(x [][]float64,lineId int) []float64 {
	n := len(x)
	vec := make([]float64, n)
	for index,v := range x {
		vec[index] = v[lineId]
	}
	return vec
}

// 实现两个矩阵的乘法 各元素相乘
func multiply(x, y []float64) (z []float64) {
	n := len(x)
	for i := 0; i < n; i++ {
		z = append(z, x[i] * y[i])
	}
	return z
}

// 返回一个行向量减去一个矩阵的结果
func lineSubVec(x []float64, y [][]float64) [][]float64 {
	n, b := len(y),len(y[0])
	c := make([][]float64, n)
	for i := 0; i < n; i++ {
		c[i] = make([]float64, b)
		for j := 0; j < b; j++ {
			c[i][j] = x[j] - y[i][j]
		}
	}
	return c
}

// 实现两个矩阵相乘后压缩的
func getSumOneVec(x, y [][]float64) []float64 {
	n, m := len(x), len(x[0])
	res := make([]float64, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i] += x[j][i] * y[j][i]
		}
	}
	return res
}

// 计算以x轴为投影轴的均值向量
func getMean(x [][] float64) []float64 {
	n, m := len(x), len(x[0])
	res := make([]float64, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i] += x[j][i]
		}
		res[i] /= float64(n)
	}
	return res
}