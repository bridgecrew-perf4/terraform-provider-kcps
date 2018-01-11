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
|`volume_id` |-|VolumeのID | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`name` |-|Volumeの名前  | - | string | - |
|`type` |-|Volumeのタイプ  | - | string | - |
|`virtualmachineid` |-|VolumeがアタッチされているVMのID  | - | string | - |
|`zoneid` |-|ZoneのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`volume_id`  |VolumeのID  | `id`と同じ値 |
|`name`  |Volumeの名前  | - |
|`type` |Volumeのタイプ  | - |
|`virtualmachineid` |VolumeがアタッチされているVMのID  | - |
|`zoneid`  |ZoneのID   | - |
|`diskofferingid`  |Disk OfferingのID   | - |
|`serviceofferingid`  |Service OfferingのID   | - |
 
