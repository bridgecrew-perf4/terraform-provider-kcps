# Load Balancer

## 使用例

```hcl
data kcps_value_loadbalancer "my_loadbalancer"{
    virtualmachineid = "be4a143f-3cdd-4091-9a8c-d82c78e49ddf"
    publicipid = "b2a3a2b6-67ff-46d2-877b-25618824b1ae"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`loadbalancer_id` |-|Load BalancerのID | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`name` |-|Load Balancerの名前 | - | string | - |
|`publicipid` |-|Public IPのID | - | string | - |
|`networkid` |-|ネットワークのID  | - | string | - |
|`virtualmachineid` |-|VMのID  | - | string | - |
|`zoneid` |-|ZoneのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`loadbalancer_id`  |Load BalancerのID  | `id`と同じ値 |
|`name` |Load Balancerの名前 | - | 
|`publicipid` |Public IPのID | - | 
|`networkid` |ネットワークのID  | - |
|`zoneid` |ZoneのID  | - | 
|`algorithm`  |Load Balancerのアルゴリズム  | - |
|`publicip`  |Public IPのIPアドレス   | - |
|`privateort`  |プライベートポート   | - |
|`publicport`  |パブリックポート   | - |
