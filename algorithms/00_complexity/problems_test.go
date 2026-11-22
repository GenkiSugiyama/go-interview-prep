package complexity

import (
	"testing"
	"time"
)

func TestSumToN(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1},
		{"5", 5, 15},
		{"100", 100, 5050},
		{"0 は 0", 0, 0},
		{"負の数は 0", -3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumToN(tt.n); got != tt.want {
				t.Errorf("SumToN(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestSumToN_Fast(t *testing.T) {
	// O(n) のループだと 10^9 回かかり、時間切れになります。
	n, want := 1_000_000_000, 500_000_000_500_000_000
	got, ok := runWithin(50*time.Millisecond, func() int { return SumToN(n) })
	if !ok {
		t.Fatalf("SumToN(%d) が遅すぎます。O(1) の公式を使いましょう", n)
	}
	if got != want {
		t.Errorf("SumToN(%d) = %d, want %d", n, got, want)
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{-7, false},
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{9, false},
		{17, true},
		{25, false},
		{97, true},
		{7919, true},
		{7917, false}, // 3 × 7 × 13 × 29
	}
	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestIsPrime_Fast(t *testing.T) {
	// 2 から n-1 まで全部割る O(n) だと 10^12 回かかり、終わりません。
	tests := []struct {
		n    int
		want bool
	}{
		{1_000_000_007, true},          // よく使われる大きな素数
		{999_999_999_989, true},        // 10^12 未満で最大の素数
		{1_000_003 * 1_000_033, false}, // 大きな数 × 大きな数（小さい約数を持たない）
	}
	for _, tt := range tests {
		got, ok := runWithin(time.Second, func() bool { return IsPrime(tt.n) })
		if !ok {
			t.Errorf("IsPrime(%d) が遅すぎます。√n まで調べれば十分です", tt.n)
			continue
		}
		if got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}

func TestHasPairWithSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"見つかる", []int{1, 4, 6, 9}, 10, true},
		{"同じ要素は 2 回使えない", []int{1, 4, 6, 9}, 2, false},
		{"同じ値でも位置が違えば OK", []int{3, 3}, 6, true},
		{"見つからない", []int{1, 2, 3}, 100, false},
		{"負の数", []int{-5, 2, 8}, 3, true},
		{"空", []int{}, 0, false},
		{"要素 1 つ", []int{5}, 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPairWithSum(tt.nums, tt.target); got != tt.want {
				t.Errorf("HasPairWithSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
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
