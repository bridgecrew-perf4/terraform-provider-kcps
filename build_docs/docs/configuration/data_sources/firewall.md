# Firewall

## 使用例

```hcl
data kcps_value_firewall "my_firewall"{
    firewall_id = "78bfe7c7-202d-4788-bd61-3c8c46ae2b8f"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`firewall_id` |-|FirewallのID | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`ipaddressid` |-|IPアドレスのID | - | string | - |
|`networkid` |-|Firewallが設定されたネットワークのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`firewall_id`  |FirewallのID  | `id`と同じ値 |
|`ipaddressid`  |IPアドレスのID  | - |
|`networkid` |Firewallが設定されたネットワークのID  | - |
|`cidrlist`  |CIDRのリスト   | listで返される |
|`startport`  |開始ポート  | - |
|`endport`  |終了ポート   | - |
|`icmpcode`  |ICMPメッセージのコード   | - |
|`icmptype`  |ICMPメッセージのタイプ   | - |
|`ipaddress`  |IPアドレス   | - |
|`protocol`  |プロトコル  | - |
