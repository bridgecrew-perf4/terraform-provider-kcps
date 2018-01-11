# Template

## 使用例

```hcl
resource kcps_template "my_template" {
    name                  = "my_template"
    displaytext           = "test desu"
    ostypeid              = "817437d2-7aae-11e4-b5b5-c45444131635"
    snapshotid            = "a48081bd-aa4d-48f3-a53a-75ef960c0bec"

    isdynamicallyscalable = true
    passwordenabled       = true
    ispublic              = true

    #volumeid = "453983a0-5c2e-4645-9468-5e7e63ebf045" 
}

```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`name`          |◯|Templateの名前          | - | string | - |
|`displaytext`   |◯|表示テキスト  | - | string | - |
|`ostypeid`      |◯|OSタイプのID   | - | string | - |
|`snapshotid`    |注1|Template作成に使用するSnapshotのID               | - | string | - |
|`volumeid`      |注1|Template作成に使用するVolumeのID     | - | string | - |
|`isdynamicallyscalable`   |-|動的スケーリング  | false | bool | - |
|`passwordenabled`   |-|パスワード有効化  | false | bool | - |
|`ispublic`   |-|公開レベル  | false | bool | - |


注1: `snapshotid`か`volumeid`のどちらかを設定する必要があります。


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |リソースID              | - | 