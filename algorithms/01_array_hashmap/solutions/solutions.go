// Package solutions は 01_array_hashmap の模範解答です。
// 自分で考えてから見るようにしましょう。
package solutions

// ContainsDuplicate: 見た値を map（集合）に覚えておきます。
//
// 時間 O(n) / 空間 O(n)
//
// 別解: slices.Sort でソートしてから隣同士を比べる方法もあります。
// 時間 O(n log n) で、追加メモリはほぼ不要です（ただし元の配列を書き換えます）。
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))
	for _, x := range nums {
		if seen[x] {
			return true
		}
		seen[x] = true
	}
	return false
}

// IsAnagram: 英小文字だけなので [26]int で出現回数を数えます。
// s の文字で +1、t の文字で -1 して、最後に全部 0 ならアナグラムです。
//
// 時間 O(n) / 空間 O(1)（配列の大きさは 26 で固定なので）
//
// 日本語なども含む場合は map[rune]int を使い、for _, r := range s で回します。
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

// TwoSum: 「値 → 添字」の map を作りながら 1 周します。
// nums[j] を見たとき、相手 target-nums[j] が既に map にあれば答えです。
//
// 時間 O(n) / 空間 O(n)
//
// ポイント: 先に map を全部作ってから探すのではなく、
// 「探す → 自分を追加」の順にすることで、同じ要素を 2 回使うことを自然に防げます。
func TwoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums)) // 値 → 添字
	for j, x := range nums {
		if i, ok := index[target-x]; ok {
			return []int{i, j}
		}
		index[x] = j
	}
	return nil // 問題の前提上ここには来ない
}

// FirstUniqueChar: 1 周目で出現回数を数え、2 周目で回数 1 の最初の文字を探します。
//
// 時間 O(n) / 空間 O(1)（配列の大きさは 26 で固定なので）
//
// 2 周目で count のほうではなく s のほうを回しているのは、
// 「s の中で最初に出てくる」位置を求めたいからです。
func FirstUniqueChar(s string) int {
	var count [26]int
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if count[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
