# Public IP

## 使用例

```hcl
data kcps_publicip "my_publicip"{
    ipaddress = "27.85.233.111"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`publicip_id` |-|Public IPのID | - | string | - |
|`associatednetworkid` |-|結びつけられたNetworkのID | - | string | - |
|`forloadbalancing` |-|Load balancerに使用されているか | - | bool | - |
|`ipaddress` |-|IPアドレス | - | string | - |
|`issourcenat` |-|ソースNATに使用されているか | - | bool | - |
|`isstaticnat` |-|StaticNATに使用されているか | - | bool | - |
|`keyword` |-|キーワード | - | string | - |
|`zoneid` |-|ZoneのID | - | string | - |


## 属性

|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`publicip_id`  |Public IPのID  | `id`と同じ値 |
|`associatednetworkid` |結びつけられたNetworkのID | - |
|`ipaddress` |IPアドレス | - |
|`issourcenat` |ソースNATに使用されているか | - |
|`isstaticnat` |StaticNATに使用されているか | - |
|`zoneid` |ZoneのID | - | 
|`networkid` |NetworkのID  | - |
|`virtualmachineid`  |VMのID  | StaticNATに使用されている場合取得可能 |
|`vmipaddress`  |VMのIPアドレス  | StaticNATに使用されている場合取得可能 |



 
