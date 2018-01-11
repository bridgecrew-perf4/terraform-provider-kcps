# NAT/PortForward

## 使用例

```hcl
data kcps_value_nat_portforward "my_nat_portforward"{
    natportforward_id = "b22eb8f6-e0d2-4d35-a8cb-ec850f912e4f"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`natportforward_id` |-|NAT/PortForwardのID | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`ipaddressid` |-|IPアドレスのID | - | string | - |
|`networkid` |-|ネットワークのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`natportforward_id`  |NAT/PortForwardのID  | `id`と同じ値 |
|`ipaddressid`  |IPアドレスのID  | - |
|`networkid` |ネットワークのID  | - |
|`ipaddress`  |IPアドレス   | - |
|`privateport`  |プライベートポートの開始ポート  | - |
|`publicport`  |パブリックポートの開始ポート   | - |
|`privateendport`  |プライベートポートの終了ポート   | - |
|`publicendport`  |パブリックポートの終了ポート   | - |
|`protocol`  |プロトコル  | - |
|`virtualmachineid`  |VMのID   | - |
|`vmguestip`  |VMのセカンダリIP   | - |