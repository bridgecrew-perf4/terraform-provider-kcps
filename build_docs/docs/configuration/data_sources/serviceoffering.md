# Service Offering

## 使用例

```hcl
data kcps_service_offering "my_service_offering"{
    name = "Small1(1vCPU,Mem2GB)"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`serviceoffering_id` |-|Service OfferingのID | - | string | - |
|`name` |-|Service Offeringの名前 | - | string | - |
|`keyword` |-|キーワード | - | string | - |


## 属性

|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`serviceoffering_id`  |Service OfferingのID  | `id`と同じ値 |
|`name` |Service Offeringの名前 | - |


 
