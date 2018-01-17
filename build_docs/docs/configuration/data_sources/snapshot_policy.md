# Snapshot Policy

## 使用例

```hcl
data kcps_snapshot_policy "my_snapshot_policy"{
    volumeid          = "ffa4fba4-7468-42f6-a86c-85fe274fb8c1"
    snapshotpolicy_id = "2f7319f7-4a75-4c6d-acdc-ca50a36822c1"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`volumeid` |◯|VolumeのID | - | string | - |
|`snapshot_policy_id` |-|定期SnapshotのID | - | string | - |




## 属性

|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`snapshot_policy_id`  |定期SnapshotのID  | `id`と同じ値 |
|`intervaltype`  | Snapshotを作成する間隔  | KCPSのAPIではinterval typeのIDが返されますが、これらをProvider Plugin内部でWEEKLYなどの文字列に変換して提供します |
|`maxsnaps` | Snapshotの保持数 | - | 
|`schedule` |Snapshotを作成する時間  | - |
|`timezone` |タイムゾーン  | - |
