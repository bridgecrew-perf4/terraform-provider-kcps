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
|`volumeid` |◯|ボリュームのID | - | string | - |
|`snapshot_policy_id` |-|定期スナップショットのID | - | string | - |




## 属性

|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`snapshot_policy_id`  |定期スナップショットのID  | `id`と同じ値 |
|`intervaltype`  | スナップショットを作成する間隔  | - |
|`maxsnaps` | スナップショットの保持数 | - | 
|`schedule` |スナップショットを作成する時間  | - |
|`timezone` |タイムゾーン  | - |
