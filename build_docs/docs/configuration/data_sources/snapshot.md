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
|`snapshot_id` |-|スナップショットのID | - | string | - |
|`name` |-|スナップショットの名前  | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`intervaltype` |-|スナップショットを作成する間隔  | - | string | - |
|`volumeid` |-|スナップショット作成に使用したボリュームのID  | - | string | - |
|`zoneid` |-|ゾーンのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`snapshot_id`  |スナップショットのID  | `id`と同じ値 |
|`name`  | スナップショットの名前  | - |
|`intervaltype` |スナップショットを作成する間隔  | - |
|`volumeid` |-|スナップショット作成に使用したボリュームのID  | - | string | - |
 
