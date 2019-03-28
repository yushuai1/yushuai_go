runtime调度器是个很神奇的东西，但是我真是但愿它不存在，我希望显式调度能更为自然些，多核处理默认开启。

关于runtime包几个函数:

Gosched 让出cpu

NumCPU 返回当前系统的CPU核数量

GOMAXPROCS 设置最大的可同时使用的CPU核数

Goexit 退出当前goroutine(但是defer语句会照常执行)