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
			distVector[index][index2] = cal_dist(basic, basic2)
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
	fmt.Printf("%+v\n", sum)
	if sum == 0 {	
		for i, v := range prob {
			prob[i] = math.Max(v, MINNUM)
		}
		perp = -12
	} else {
		for i, v := range prob {
			prob[i] = v / sum
		}
		perp = 0
		for _, v := range prob {
			if v != 0 {
				perp -= math.Log(v)
			}
		}
	}
	// 困惑度和pi \ j 的概率分布
	return perp, prob
}

// 二分搜索寻找beta,并计算pairwise的prob
func search_prob(x [][]float64, tol float64, perplexity float64) [][]float64 {
	fmt.Println("Computing pairwise distances...")
	n, _ := len(x), len(x[0])
	dist := cal_pairwise_dist(x)
	beta := make([]float64, n)
	for i := 0; i < n; i++ {
		beta[i] = 1.0
	}
	pair_prob := make([][]float64, n)
	for i := 0; i < n; i++ {
		t := make([]float64, n)
		for j := 0; j < n; j++ {
			t[j] = 0.0
		}
		pair_prob[i] = t
	}
	
	base_perp := math.Log(perplexity)
	for i := 0; i < n; i++ {
		var this_prob []float64
		betamin :=  math.Inf(-1)
		betamax :=  math.Inf(1)

		perp, _ := cal_perplexity(x[i], i, beta[i])

		perp_diff := perp - base_perp
		tries := 0
		for math.Abs(perp_diff) > tol && tries < 50 { 
			if perp_diff > 0 {
				betamin = beta[i]
				if betamax == math.Inf(1) {
					beta[i] *= 2
				} else {
					beta[i] = (beta[i] + betamax) / 2
				}
			} else {
				betamax = beta[i]
				if betamin == math.Inf(-1) {
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

// 
// func tsne(x [][]float64, no_dims int, initial_dims int, perplexity float64, max_iter int) {
// 	n, d := len(x), len(x[0])

// 	// 动量
// 	eta := 500
// 	// 初始化Y
// 	y := randomRandn(n, no_dims)
// 	dy := getZeroVec(n, no_dims)

// 	p := search_prob(x, 1e-5, perplexity)
// 	pT := transpose(p)
// 	p = add2vec(p, pT)
// 	p = division(p, sumSelf(p))
// 	p = division(p, 1 / 4)
// 	maxminVec(p)

// 	for i := 0; i < max_iter; i++ {
// 		distVec := cal_pairwise_dist(y)
// 		recVec(distVec)
// 		q := division(distVec, sumSelf(distVec))
// 		maxminVec(q)
// 		PQ := subtraction(p,q)
		
		
// 	}
	
// }


func main() {

	a := []float64{1.0, 2}
	b := []float64{1.0, 2}
	c := [][]float64{a,b}
	maxminVec(c)
	fmt.Printf("%+v\n", c)
}