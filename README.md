## goyacc-carcdr-evaluator

Lisp car and cdr evaluator using parser generated by [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc)

```
$ go run main.go
> (car (cdr (cons 1 (cons (cons 20 300.4) (cons 50 60)))))
(20 . 300.4)
```

