# Network

## 使用例

```hcl
data kcps_network "my_network"{
    name = "PublicFrontSegment"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`network_id` |-|NetworkのID | - | string | - |
|`name` |-|Networkの名前 | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`zoneid` |-|ZoneのID | - | string | - |


## 属性

|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`network_id`  |NetworkのID  | `id`と同じ値 |
|`name` |Networkの名前 | - |
|`zoneid` |ZoneのID | - | 


 
