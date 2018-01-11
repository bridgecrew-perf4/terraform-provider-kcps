# NAT/PortForward

## 使用例

```hcl
resource kcps_nat_portforward "my_nat_portforward" {
    ipaddressid      = "165bd632-738b-44e4-8449-562fcd6da509"
    protocol         = "TCP"
    virtualmachineid = "be4a143f-3cdd-4091-9a8c-d82c78e49ddf"

    port {
        privateport    = 1020
        publicport     = 1020
        privateendport = 1023
        publicendport  = 1023
    }

    vmguestip        = "10.1.1.56"
}
```


## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`ipaddressid` |◯|PortForwardingルールを設定するIPアドレスのID  | - | string | - |
|`protocol`   |◯|プロトコル           | - | string | TCPまたはUDP |
|`virtualmachineid`    |◯| PortForwardingルールを適用するVMのID       | - | string | - |
|`port` |◯|プライベートポートとパブリックポートの範囲の設定            | - | list(map) | 詳細は後述 |
|`vmguestip` |-|VMのセカンダリIP  | - | string | - |



## port

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`privateport` |◯|プライベートポートの開始ポート  | - | int | - |
|`publicport`   |◯|パブリックポート開始ポート | - | int | - |
|`privateendport`   |-|プライベートの終了ポート | - | int | - |
|`publicendport`   |-|パブリックポートの終了ポート | - | int | - |

## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |リソースID   | - | 
|`ipaddress`    |IPアドレス | - | 
