package arrayhashmap

import (
	"slices"
	"strings"
	"testing"
	"time"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"重複あり", []int{1, 2, 3, 1}, true},
		{"重複なし", []int{1, 2, 3, 4}, false},
		{"全部同じ", []int{7, 7, 7}, true},
		{"負の数", []int{-1, 0, 1, -1}, true},
		{"空", []int{}, false},
		{"要素 1 つ", []int{42}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDuplicate(tt.nums); got != tt.want {
				t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestContainsDuplicate_Fast(t *testing.T) {
	// 重複なしの 10^5 要素。O(n²) だと約 50 億回の比較になります。
	nums := make([]int, 100_000)
	for i := range nums {
		nums[i] = i
	}
	got, ok := runWithin(500*time.Millisecond, func() bool { return ContainsDuplicate(nums) })
	if !ok {
		t.Fatalf("ContainsDuplicate が遅すぎます。map を使って O(n) にしましょう")
	}
	if got {
		t.Errorf("ContainsDuplicate(0..99999) = true, want false")
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, t string
		want bool
	}{
		{"listen", "silent", true},
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"aab", "abb", false},
		{"a", "ab", false},
		{"", "", true},
	}
	for _, tt := range tests {
		if got := IsAnagram(tt.s, tt.t); got != tt.want {
			t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
		}
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"基本", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"同じ要素は使えない", []int{3, 2, 4}, 6, []int{1, 2}},
		{"同じ値 2 つ", []int{3, 3}, 6, []int{0, 1}},
		{"負の数", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"末尾にある", []int{1, 5, 8, 2, 10}, 12, []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.nums, tt.target); !slices.Equal(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestTwoSum_Fast(t *testing.T) {
	// 答えが末尾にある 10^5 要素。O(n²) だと間に合いません。
	n := 100_000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i * 2 // 偶数だけ
	}
	nums[n-1] = 1                                                                       // 奇数は末尾の 1 だけ → 答えは (0, n-1) の組のみ
	got, ok := runWithin(500*time.Millisecond, func() []int { return TwoSum(nums, 1) }) // 0 + 1
	if !ok {
		t.Fatalf("TwoSum が遅すぎます。map を使って O(n) にしましょう")
	}
	if want := []int{0, n - 1}; !slices.Equal(got, want) {
		t.Errorf("TwoSum(...) = %v, want %v", got, want)
	}
}

func TestFirstUniqueChar(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
		{"z", 0},
		{"aabbc", 4},
		{"", -1},
	}
	for _, tt := range tests {
		if got := FirstUniqueChar(tt.s); got != tt.want {
			t.Errorf("FirstUniqueChar(%q) = %d, want %d", tt.s, got, tt.want)
		}
	}
}

func TestFirstUniqueChar_Fast(t *testing.T) {
	// 10^5 文字で、ユニークな文字は末尾だけ。
	s := strings.Repeat("ab", 50_000) + "c"
	got, ok := runWithin(500*time.Millisecond, func() int { return FirstUniqueChar(s) })
	if !ok {
		t.Fatalf("FirstUniqueChar が遅すぎます。出現回数を先に数えて O(n) にしましょう")
	}
	if got != len(s)-1 {
		t.Errorf("FirstUniqueChar(...) = %d, want %d", got, len(s)-1)
	}
}

// runWithin は f を実行し、limit 以内に終われば (結果, true) を返します。
// 間に合わなければ (ゼロ値, false) を返します（遅い実装でテストが固まらないようにするため）。
func runWithin[T any](limit time.Duration, f func() T) (T, bool) {
	ch := make(chan T, 1)
	go func() { ch <- f() }()
	select {
	case v := <-ch:
		return v, true
	case <-time.After(limit):
		var zero T
		return zero, false
	}
}
