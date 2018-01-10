#  GuestOS

## 使用例

```hcl
data kcps_guestos "my_guestos"{
    description = "Ubuntu 10.04 (64-bit)" 
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`guestos_id` |-|OSタイプのID | - | string | - |
|`description` |-|説明  | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`oscategoryid` |-|OSのカテゴリーのID  | - | string | - |



## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`guestos_id`  |OSタイプのID  | `id`と同じ値 |
|`description`  |説明  | - |
|`oscategoryid` |OSのカテゴリーのID  | - |
 
