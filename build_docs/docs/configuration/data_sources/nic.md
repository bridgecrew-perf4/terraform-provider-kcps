#  Nic

## 使用例

```hcl
data kcps_nic "my_nic"{
    networkid        = "d0f15c14-4d94-4a69-b39a-994721bfc809" 
    virtualmachineid = "be4a143f-3cdd-4091-9a8c-d82c78e49ddf"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`virtualmachineid` |◯|キーワード | - | string | - |
|`nic_id` |-|NicのID | - | string | - |
|`networkid` |-|NetworkのID  | - | string | APIにはない機能 |




## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`nic_id`  |NicのID  | `id`と同じ値 |
|`networkid` |NetworkのID  | - |
|`ipaddress`  |IPアドレス  | - |
|`ip6address`  |IPv6アドレス   | - |
|`macaddress`  |MACアドレス   | - |
|`secondaryip`  |セカンダリIPのマップ   | 詳細は後述 |


## secondaryip

|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`secondaryip_id`  |セカンダリIPのID   | - | 
|`secondaryip_ipaddress`  |セカンダリIPのIPアドレス  | `id`と同じ値 |

 
