MeiTuanPeiSong SDK for Golang
美团配送

## 安装
##### 启用 Go module

```
go get github.com/royromny/meituan-peisong
```

```
import github.com/royromny/meituan-peisong
```

##### 未启用 Go module

```
go get github.com/royromny/meituan-peisong
```

```
import github.com/royromny/meituan-peisong
```

## 帮助

## 关于各分支（版本)

* master - 和主要维护分支同步；

## 两个回调接口：
```
http.HandleFunc("/meituan", func(rep http.ResponseWriter, req *http.Request) {
	var notify, _ = client.GetOrderNotification(req)
	if notify != nil {
		fmt.Println("订单状态为:", notify.OrderStatus)
	}
	meituan.AckNotification(rep) // 确认收到通知消息
})
```

## License
This project is licensed under the MIT License.
