# 推送服务相互转换工具

今天被阿里云只支持钉钉机器人报警这种垄断机制搞得很烦- -，不想收那么多短信，习惯用 pushover ，机智如我开发了这么个小东西，可以实现不同推送服务之间的伪造和转换

## 支持情况

| 推送服务   | 解析 | 推送 |
| ---------- | ---- | ---- |
| pushover   | √    | √    |
| 钉钉机器人 | √    | √    |

我够用了

## 用法

在那些只支持单个推送的sb服务商处的url填上你部署的地址，后面token根据不同的服务写：

### pushover

> URL: xxx/pushover?user=xxx&token=xxx

### 钉钉

> URL: xxx/dingtalk?token=xxx

