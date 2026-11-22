# 01. 配列とハッシュマップ

コーディング面接で**一番よく出る**のがこのテーマです。
キーワードは **「二重ループを map で1重にする」** です。

## Go の map の基本

```go
// 作成
count := make(map[string]int)

// 追加・更新 … 平均 O(1)
count["apple"] = 3
count["apple"]++ // 存在しないキーでも 0 から始まるので ++ できる

// 参照 … 平均 O(1)
v := count["banana"] // 存在しないキーはゼロ値（int なら 0）

// 存在チェック（カンマ ok イディオム）
if v, ok := count["apple"]; ok {
    fmt.Println("ある:", v)
}

// 削除
delete(count, "apple")

// 「集合」として使う（値は使わない）
seen := make(map[int]struct{}) // struct{} はメモリを使わない
seen[5] = struct{}{}
_, exists := seen[5]

// 初心者のうちは map[int]bool のほうが読みやすいので、これでも OK
visited := make(map[int]bool)
visited[5] = true
if visited[5] { /* ... */ }
```

**注意:** map を `range` で回すと、**順番は毎回ランダム**です。順番が大事なときはスライスで回しましょう。

## なぜ map で速くなるのか

「配列の中に x があるか？」を調べるとき:

| 方法 | 1回あたり | n 回調べると |
|------|----------|-------------|
| スライスを先頭から探す | O(n) | O(n²) |
| map で引く | O(1) | **O(n)** |

つまり **「前に見た値を map に覚えておく」** だけで、二重ループの多くは1重になります。
代わりに map の分だけメモリを使います（空間 O(n)）。
この「**メモリを使って時間を買う**」という考え方は面接で頻出です。

## 典型パターン

### パターン1: 出現したかを覚える（集合）

```go
seen := make(map[int]bool)
for _, x := range nums {
    if seen[x] {
        // x は前にも出てきた
    }
    seen[x] = true
}
```

### パターン2: 出現回数を数える（カウンタ）

```go
count := make(map[rune]int)
for _, r := range s {
    count[r]++
}
```

### パターン3: 値 → 位置を覚える

```go
index := make(map[int]int) // 値 → 添字
for i, x := range nums {
    index[x] = i
}
```

## 文字列の扱い（Go 特有の注意）

- `s[i]` は **byte**（1バイト）です。英小文字だけなら問題ありません
- 日本語などを含むなら `for _, r := range s` で **rune**（1文字）単位で回します
- 英小文字だけと分かっているなら、map の代わりに `[26]int` 配列で数えると速くて簡潔です

```go
var count [26]int
for i := 0; i < len(s); i++ {
    count[s[i]-'a']++
}
```

## 問題

`problems.go` を開いて実装してください。易しい順に並んでいます。

| 問題 | 内容 | 使うパターン | 目標 |
|------|------|-------------|------|
| `ContainsDuplicate` | 重複する値があるか | 集合 | 時間 O(n) |
| `IsAnagram` | 2つの文字列が並べ替えで一致するか | カウンタ | 時間 O(n) |
| `TwoSum` | 足して target になる2つの添字 | 値 → 位置 | 時間 O(n) |
| `FirstUniqueChar` | 最初の「1回しか出ない文字」の位置 | カウンタ | 時間 O(n) |

```bash
go test -v ./algorithms/01_array_hashmap/
```

`TwoSum` は [00_complexity](../00_complexity/) の `HasPairWithSum` の改良版です。
二重ループ版と見比べてみてください。

## 面接で言えるようにしておくこと

- 「素直に書くと二重ループで O(n²) ですが、map で見た値を覚えておけば O(n) になります。
  その代わり、map の分だけ空間 O(n) を使います」
- map の操作は**平均** O(1)（最悪の場合は遅くなることもある）
- 「ソートしてから比べる」解法もあります（O(n log n)、追加メモリが少ない）。
  **複数の解法とトレードオフを話せる**と評価が上がります
