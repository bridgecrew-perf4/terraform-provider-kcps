# Disk Offering

## 使用例

```hcl
data kcps_disk_offering "my_disk_offering"{
    name = "MIDDLE_STORAGE"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`diskoffering_id` |-|Disk OfferingのID | - | string | - |
|`name` |-|Disk Offeringの名前 | - | string | - |
|`keyword` |-|キーワード | - | string | - |


## 属性

|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`diskoffering_id`  |Disk OfferingのID  | `id`と同じ値 |
|`name` |Disk Offeringの名前 | - |


 
