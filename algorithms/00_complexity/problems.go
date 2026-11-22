// Package complexity は計算量（Big-O）を意識して書く練習問題です。
// 解説は README.md、模範解答は solutions/ にあります。
package complexity

// SumToN は 1 から n までの整数の合計を返します。
//
// 例:
//
//	SumToN(1)  == 1
//	SumToN(5)  == 15   // 1+2+3+4+5
//	SumToN(0)  == 0    // n が 0 以下なら 0
//
// 制約: -10^9 <= n <= 10^9
//
// 目標: 時間 O(1)
// ヒント: ループで足すと O(n) です。n = 10^9 だと遅すぎます。
// 小学校で習う「1 から 100 までの和」の公式を思い出しましょう。
func SumToN(n int) int {
	// TODO: ここを実装
	return 0
}

// IsPrime は n が素数なら true を返します。
// 素数とは「1 と自分自身以外に約数を持たない、2 以上の整数」です。
//
// 例:
//
//	IsPrime(2)  == true
//	IsPrime(9)  == false  // 3 で割り切れる
//	IsPrime(1)  == false  // 1 は素数ではない
//
// 制約: -10^12 <= n <= 10^12
//
// 目標: 時間 O(√n)
// ヒント: n が a × b と分解できるなら、a と b の小さいほうは √n 以下です。
// つまり 2 から √n まで割ってみれば十分です。
// （math.Sqrt を使わず i*i <= n と書くと、浮動小数点の誤差を気にしなくて済みます）
func IsPrime(n int) bool {
	// TODO: ここを実装
	return false
}

// HasPairWithSum は、nums の異なる位置にある 2 つの数を足して
// target になる組み合わせがあれば true を返します。
//
// 例:
//
//	HasPairWithSum([]int{1, 4, 6, 9}, 10) == true   // 1+9 または 4+6
//	HasPairWithSum([]int{1, 4, 6, 9}, 2)  == false  // 1+1 は同じ要素なので不可
//	HasPairWithSum([]int{3, 3}, 6)        == true   // 位置が違えば同じ値でも OK
//
// 制約: 0 <= len(nums) <= 1000
//
// 目標: まずは二重ループの O(n²) で書いてみましょう。
// 次の 01_array_hashmap の TwoSum で、map を使って O(n) に改善します。
func HasPairWithSum(nums []int, target int) bool {
	// TODO: ここを実装
	return false
}
