// Package solutions は 00_complexity の模範解答です。
// 自分で考えてから見るようにしましょう。
package solutions

// SumToN: 等差数列の和の公式 n(n+1)/2 を使います。
//
// 時間 O(1) / 空間 O(1)
//
// ループで 1 から n まで足すと O(n) です。
// n = 10^9 だと 10 億回の足し算になり、テストの時間制限に引っかかります。
//
// 補足: n = 10^9 のとき n*(n+1) は約 10^18 で、
// int（64bit 環境では int64、最大約 9.2×10^18）に収まります。
// 面接では「オーバーフローしないか」を自分から確認すると好印象です。
func SumToN(n int) int {
	if n <= 0 {
		return 0
	}
	return n * (n + 1) / 2
}

// IsPrime: 2 から √n まで割り切れるか調べます。
//
// 時間 O(√n) / 空間 O(1)
//
// n = a × b（a <= b）と分解できるなら a <= √n です。
// なので √n までに約数が見つからなければ素数です。
// n = 10^12 でも √n = 10^6 回で済みます。
//
// 条件を i*i <= n と書くことで、math.Sqrt の浮動小数点誤差を避けています。
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// HasPairWithSum: 全ペアを二重ループで調べます。
//
// 時間 O(n²) / 空間 O(1)
//
// j を i+1 から始めることで、
//   - 同じ要素を 2 回使うこと（i == j）
//   - 同じペアを 2 回調べること（(i,j) と (j,i)）
//
// を防いでいます。
// map を使った O(n) の解法は 01_array_hashmap の TwoSum で学びます。
func HasPairWithSum(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}
