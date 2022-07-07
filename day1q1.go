package main

import "fmt"

type Matrix struct {
	rows int
	col int
	mat [][]int

}

func (r Matrix)number_of_rows() int  {
	return r.rows
}
func (r Matrix)number_of_cols() int  {
	return r.col
}
func (r Matrix)set(i,j,val int) [][]int  {
	r.mat[i][j]=val
	return r.mat
}
func(r Matrix)add(m2 [][]int) [][]int{
	for i:=0;i<r.rows;i++{
		for j:=0;j<r.col;j++{
			r.mat[i][j]+=m2[i][j]
		}
	}
	return	r.mat
}
func main() {
	//arr:=[3][3]int{}
	//fmt.Println(arr)
	r:=Matrix{rows: 2,col: 2,mat: [][]int{{1,1},{1,1}}}
	fmt.Println(r.add([][]int{{1,2},{2,1}}))
}
