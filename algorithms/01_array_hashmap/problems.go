// Package arrayhashmap は配列と map を使う練習問題です。
// 解説は README.md、模範解答は solutions/ にあります。
package arrayhashmap

// ContainsDuplicate は nums の中に同じ値が 2 回以上出てくれば true を返します。
//
// 例:
//
//	ContainsDuplicate([]int{1, 2, 3, 1}) == true
//	ContainsDuplicate([]int{1, 2, 3, 4}) == false
//
// 制約: 0 <= len(nums) <= 10^5
//
// 目標: 時間 O(n)
// ヒント: 「これまでに見た値」を map に覚えておきましょう。
func ContainsDuplicate(nums []int) bool {
	// TODO: ここを実装
	return false
}

// IsAnagram は s と t がアナグラム（文字を並べ替えると一致する）なら true を返します。
//
// 例:
//
//	IsAnagram("listen", "silent") == true
//	IsAnagram("rat", "car")       == false
//	IsAnagram("aab", "abb")       == false  // 文字の種類は同じでも個数が違う
//
// 制約: 0 <= len(s), len(t) <= 5×10^4。s と t は英小文字のみ。
//
// 目標: 時間 O(n)
// ヒント: 各文字の出現回数を数えて比べましょう。
// 長さが違えば、その時点で false です。
func IsAnagram(s, t string) bool {
	// TODO: ここを実装
	return false
}

// TwoSum は nums の中から、足すと target になる 2 つの要素の添字 [i, j]（i < j）を返します。
// 答えはちょうど 1 つ存在するものとします。同じ要素を 2 回使ってはいけません。
//
// 例:
//
//	TwoSum([]int{2, 7, 11, 15}, 9) == []int{0, 1}  // 2 + 7 = 9
//	TwoSum([]int{3, 2, 4}, 6)      == []int{1, 2}  // 2 + 4 = 6（3+3 は同じ要素なので不可）
//	TwoSum([]int{3, 3}, 6)         == []int{0, 1}
//
// 制約: 2 <= len(nums) <= 10^5
//
// 目標: 時間 O(n)
// ヒント: nums[j] を見ているとき、欲しい相手は target - nums[j] です。
// それが「前に出てきたか」「どの位置にあったか」を map で調べられれば…？
func TwoSum(nums []int, target int) []int {
	// TODO: ここを実装
	return nil
}

// FirstUniqueChar は s の中で 1 回しか出てこない文字のうち、
// 最初に出てくるものの位置（添字）を返します。なければ -1 を返します。
//
// 例:
//
//	FirstUniqueChar("leetcode")     == 0   // 'l'
//	FirstUniqueChar("loveleetcode") == 2   // 'v'
//	FirstUniqueChar("aabb")         == -1
//
// 制約: 0 <= len(s) <= 10^5。s は英小文字のみ。
//
// 目標: 時間 O(n)
// ヒント: 1 周目で数を数え、2 周目で「回数が 1 の最初の文字」を探します。
// （map は順番がランダムなので、2 周目は s のほうを回しましょう）
func FirstUniqueChar(s string) int {
	// TODO: ここを実装
	return 0
}
