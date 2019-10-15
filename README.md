## goyacc-carcdr-evaluator

cons, car and cdr evaluator for [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc) practice

```
$ go run main.go
> (car (cdr (cons 1 (cons (cons 20 300.4) (cons 50 60)))))
(20 . 300.4)
```

[goyaccでparserを生成しLispのcons,car,cdrの式を評価する - sambaiz-net](https://www.sambaiz.net/article/244/)