# LoadBalancer

## 使用例

```hcl
resource kcps_loadbalancer "my_loadbalancer" {
    name        = "my_loadbalancer" 
    algorithm   = "source"
    privateport = 80
    publicport  = 8080
    publicipid  = "b2a3a2b6-67ff-46d2-877b-25618824b1ae"

    assignto    = [ "fa608125-4658-4ff0-aab4-33a1b428988a",
                    "be4a143f-3cdd-4091-9a8c-d82c78e49ddf" ]  
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`name`         |◯|Load Balancerの名前  | - | string | - |
|`algorithm`   |◯|Load Balancerのアルゴリズム | - | string | source、roundrobin、leastconnの内のどれか |
|`privateport`    |◯|プライベートポート    | - | int | - |
|`publicport`     |◯|パブリックポート     | - | int | - |
|`publicipid` |◯|Load Balancerを設定するPublic IPのID  | - | string | - |
|`assignto`        |-| Load Balancerに接続するVMのID | - | list(string) | - |



## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |リソースID   | - | 
