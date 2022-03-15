package main

import (
	"fmt"
	"math"
)

const (
	MINNUM  		float64	 =    1e-12
)

//  x -> x * m维度的矩阵， 表示n个样本m个属性
// 计算出任意两个点的距离的平方
func cal_pairwise_dist(vec [][]float64) (distVector [][]float64) {
	vecLen := len(vec)
	distVector = make([][]float64, vecLen)
	for index, basic := range vec {
		distVector[index] = make([]float64, vecLen)
		for index2, basic2 := range vec {
			distVector[index][index2] = cal_dist(basic, basic2) + 1
		}
	}
	return distVector
}

// 计算困惑度， 最终会选择合适的beta值， 也就是每个点的方差
func cal_perplexity(dist []float64, idx int, beta float64) ( float64, []float64) {
	var prob = make([]float64, len(dist))
	var perp float64
	for i := 0; i < len(dist); i++ {
		prob[i] = math.Exp(-beta * dist[i])
	}
	// 自身的prob为0
	prob[idx] = 0
	var sum = 0.0
	for _, v := range prob {
		sum += v
	}
	// fmt.Printf("%+v\n", sum)
	if sum == 0 {	
		for i, v := range prob {
			prob[i] = math.Max(v, MINNUM)
			// fmt.Println(math.Log(1))
		}
		perp = -12
	} else {
		for i, v := range prob {
			prob[i] = v / sum
		}
		perp = 0
		for _, v := range prob {
			if v != 0 {
				perp -= v * math.Log(v)
			}
		}
	}
	// 困惑度和pi \ j 的概率分布
	return perp, prob
}

// 二分搜索寻找beta,并计算pairwise的prob
func search_prob(x [][]float64, tol float64, perplexity float64) [][]float64 {
	fmt.Println("Computing pairwise distances...")
	n := len(x)
	dist := cal_pairwise_dist(x)
	beta := make([]float64, n)
	for i := 0; i < n; i++ {
		beta[i] = 1.0
	}
	pair_prob := getZeroVec(n,n)
	base_perp := math.Log(perplexity)

	for i := 0; i < n; i++ {
		betamin :=  math.Inf(-1)
		betamax :=  math.Inf(1)

		perp, this_prob := cal_perplexity(x[i], i, beta[i])
		// fmt.Println("this_prob", this_prob)
		perp_diff := perp - base_perp
		tries := 0
		for math.Abs(perp_diff) > tol && tries < 50 { 
			if perp_diff > 0 {
				betamin = beta[i]
				if betamax == math.Inf(1) || betamax == math.Inf(-1) {
					beta[i] *= 2
				} else {
					beta[i] = (beta[i] + betamax) / 2
				}
			} else {
				betamax = beta[i]
				if betamin == math.Inf(-1) || betamin == math.Inf(1) {
					beta[i] /= 2
				} else {
					beta[i] = (beta[i] + betamin) / 2
				}
			}
			perp, this_prob = cal_perplexity(dist[i], i, beta[i])
			perp_diff = perp - base_perp
			tries += 1
		}
		pair_prob[i] = this_prob
	}
	return pair_prob
}

// 实现t-sne算法
func tsne(x [][]float64, no_dims int, initial_dims int, perplexity float64, max_iter int) [][]float64 {
	n := len(x)

	// 动量
	eta := 500.0
	// 初始化Y
	y := randomRandn(n, no_dims)
	dy := getZeroVec(n, no_dims)

	p := search_prob(x, 1e-5, perplexity)
	pT := transpose(p)
	p = add2vec(p, pT)
	// fmt.Println(sumSelf(p))
	_ = division(p, sumSelf(p))
	_ = division(p, float64(0.25))
	maxminVec(p)

	for i := 0; i < max_iter; i++ {
		distVec := cal_pairwise_dist(y)
		recVec(distVec)
		if i == 0 {
			fmt.Println("distVec", distVec)
		}
		q := division(distVec, sumSelf(distVec))
		maxminVec(q)
		
		PQ := subtraction(p, q)
		for j := 0; j < n; j++ {
			PQLine := getLine(PQ,j)
			numLine := getLine(distVec,j)
			var aim_title [][]float64
			mutledLine := multiply(PQLine, numLine)
			for p := 0; p < no_dims; p++ {
				aim_title = append(aim_title, mutledLine)
			}
			aimTileT := transpose(aim_title)
			subY := lineSubVec(y[j], y)
			resY := getSumOneVec(aimTileT, subY)
			dy[j] = resY
		}
		meanY := getMean(y)
		for j := 0; j < n; j++ {
			for k := 0; k < no_dims; k++ {
				y[j][k] -= (eta * dy[j][k] + meanY[k])
			}
		}
		// fmt.Println(y)
		if i == 100 {
			_ = division(p, 4)
		}
	}
	return y
}


func main() {
	a := []float64{1.0, 2, 3.0, 4.0}
	b := []float64{1.0, 2, 2.0, 3.0} 
	c := [][]float64{a, b}
	z := tsne(c, 2, 0, 30.0, 200)
	// z := cal_pairwise_dist(c)
	fmt.Println(z)
}