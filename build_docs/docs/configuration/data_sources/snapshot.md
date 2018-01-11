#  Snapshot

## 使用例

```hcl
data kcps_snapshot "my_snapshot"{
    name = "my_snapshot"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`snapshot_id` |-|SnapshotのID | - | string | - |
|`name` |-|Snapshotの名前  | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`intervaltype` |-|Snapshotを作成する間隔  | - | string | - |
|`volumeid` |-|Snapshot作成に使用したVolumeのID  | - | string | - |
|`zoneid` |-|ZoneのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`snapshot_id`  |SnapshotのID  | `id`と同じ値 |
|`name`  | Snapshotの名前  | - |
|`intervaltype` |Snapshotを作成する間隔  | - |
|`volumeid` |-|Snapshot作成に使用したVolumeのID  | - | string | - |
 
