# Snapshot Policy

## 使用例

```hcl
resource kcps_snapshot_policy "my_snapshot_policy"{
    intervaltype = "WEEKLY"
    maxsnaps     = 3
    schedule     = "10:10:6"
    timezone     = "JST"
    volumeid     = "2eaf5912-78a5-41ef-825f-9f4558511450"
}
```


## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |備考|
|:----------|:------|:---------|:--------|:--------|:--|
|`intervaltype` |◯|スナップショットを作成する間隔  | - | string | DAILY、WEEKLY、MONTHLYの内のどれか |
|`maxsnaps`   |◯|スナップショットをいくつ保持するか           | - | int | - |
|`schedule`    |◯| スナップショットを作成する時間       | - | string | フォーマットについては後述 |
|`timezone` |◯|タイムゾーン     | - | string | - |
|`volumeid` |◯|スナップショット作成に使用するボリュームのID  | - | string | - |


## scheduleのフォーマット

`intervaltype`によってフォーマットが変化します。

|intervaltype |scheduleのフォーマット    |DDについて    |
|:----------|:------|:---------|
|DAILY |MM:HH*|-  |
|WEEKLY   |MM:HH:DD |WEEKLYのDDは曜日を示す(日曜～土曜：1～7)        |
|MONTHLY    |MM:HH:DD| MONTHLYのDDは日付(1～31)       | 




## 属性
|属性名 |説明      |備考 |
|:----------|:------|:---------|
|`id`          |リソースID   | - | 
