# Volume

## 使用例

```hcl
resource kcps_volume "my_volume" {
    name     = "myvolume"

    diskoffering {
        diskofferingid = "bc1b5c0c-fcb3-4a7b-b8de-2c9d6952e0a5"
        size           = 100
        zoneid         = "593697b6-c123-4025-b412-ef83822733e5"
    }

    attachto = "fa608125-4658-4ff0-aab4-33a1b428988a"
    
    #snapshot {
    #    snapshotid = "3efa21f5-10dd-41a1-ab96-b1f6e0c3dff9"
    #}
}

```

## 注意事項

Snapshotを用いてVolumeリソースを作成する場合、リソース作成時に同時`attachto`の設定をしていると高確率でボリュームのアタッチに失敗します。原因はよくわかっていません。


## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`name`           |◯|Volume名               | - | string | - |
|`diskoffering`   |注1|Disk Offeringの設定(Disk OfferingからVolumeを作成する場合に必要)  | - | list(map) | 詳細は後述 |
|`snapshot`       |注1|Snapshotの設定(SnapshotからVolumeを作成する場合に必要)  | - | list(map) | 詳細は後述 |
|`attachto`       |-|VolumeをアタッチするVMのID      | - | string | - |


注1: `diskoffering`か`snapshot`どちらかを設定する必要があります。


## diskoffering

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`diskofferingid`      |◯|Disk OfferingのID  | - | string | - |
|`size`   |◯|Diskのサイズ | - | int | - |
|`zoneid`   |◯|ZoneのID | - | string | - |


## snapshot

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`snapshotid`      |◯|SnapshotのID  | - | string | - |



## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |リソースID              | - | 