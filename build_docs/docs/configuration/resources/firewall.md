# Firewall

## 使用例

```hcl
resource kcps_firewall "my_firewall" {
    ipaddressid = "165bd632-738b-44e4-8449-562fcd6da509"
    protocol    = "TCP"
    cidrlist    = ["1.1.1.1/32", "2.2.2.2/32"] 
    
    port {
        startport   = 1020
        endport     = 1023
    }

    #icmp {
    #    icmpcode = 0
    #    icmptype = 3
    #}
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`ipaddressid` |◯|Firewallルールを設定するIPアドレスのID  | - | string | - |
|`protocol`   |◯|Firewallルールのプロトコル           | - | string | TCP、UDP、ICMPの内のどれか |
|`cidrlist`    |◯|Firewallで許可する送信元のCIDR        | - | list(string) | - |
|`port`        |注1 注2|許可するポートの範囲の設定              | - | list(map) | 詳細は後述 |
|`icmp`        |注1 注3|拒否するICMPメッセージに関する設定    | - | list(map) | 詳細は後述 |

注1: `port`か`icmp`どちらかを設定する必要があります。 <br>
注2: `protocol="TCP"`または`protocol="UDP"`の場合に設定できます。 <br>
注3: `protocol="ICMP"`の場合に設定できます。


## port

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`startport`      |◯|開始ポート  | - | int | - |
|`endport`   |◯|終了ポート | - | int | - |

## icmp

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`icmpcode`      |◯|ICMPメッセージのコード  | - | int | - |
|`icmptype`   |◯|ICMPメッセージのタイプ | - | int | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |リソースID   | - | 
|`ipaddress`    |IPアドレス | - | 
