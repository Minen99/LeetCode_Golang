# 解法ヒント

- BFS（Breadth-First Search, 幅優先探索）
  - グリッドを行方向・列方向に舐めるのではなく、キューを介して島の形に沿って広がる

```go 骨格
queue := [][2]int{{row, column}}   // 出発点をキューに入れておく

for len(queue) > 0 {               // キューが空になるまで
    cur := queue[0]                // 先頭を取り出す(peek)
    queue = queue[1:]              // 先頭を捨てる(pop)
    r, c := cur[0], cur[1]

    // ここで「今取り出したマス (r,c) の隣」を見て、
    // 条件を満たす隣だけを queue = append(queue, ...) で入れる
}
```
