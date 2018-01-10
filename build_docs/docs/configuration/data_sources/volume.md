#  Volume

## 使用例

```hcl
data kcps_volume "my_volume"{
    name = "my_volume"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`volume_id` |-|ボリュームのID | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`name` |-|ボリュームの名前  | - | string | - |
|`type` |-|ボリュームのタイプ  | - | string | - |
|`virtualmachineid` |-|ボリュームがアタッチされているVMのID  | - | string | - |
|`zoneid` |-|ゾーンのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`volume_id`  |ボリュームのID  | `id`と同じ値 |
|`name`  |ボリュームの名前  | - |
|`type` |ボリュームのタイプ  | - |
|`virtualmachineid` |ボリュームがアタッチされているVMのID  | - |
|`zoneid`  |ゾーンのID   | - |
|`diskofferingid`  |ディスクオファリングのID   | - |
|`serviceofferingid`  |サービスオファリングのID   | - |
 
