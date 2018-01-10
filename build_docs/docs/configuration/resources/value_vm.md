# Value VirtualMachine

## 使用例

```hcl
resource kcps_value_vm "my_vm" {
    name              = "my_vm"
    serviceofferingid = "e3060950-1b4f-4adb-b050-bbe99694da19" 
    templateid        = "0eb72664-f7ad-4d36-be3e-4f4c32ffe0e5"
    zoneid            = "593697b6-c123-4025-b412-ef83822733e5"

    diskoffering{
        diskofferingid = "10cc47d1-2e04-4aeb-aec9-fb08d273198e"
        size           = 100
    }
}
```

## 注意事項

作成したVMはKCPSの内部ネットワーク`PublicFrontSegment`と`MonitoringNetwork`に自動的に接続されます。


## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`name`                |◯|VMの名前                | - | string | - |
|`serviceofferingid`   |◯|サービスオファリングのID  | - | string | - |
|`templateid`          |◯|テンプレートのID        | - | string | - |
|`zoneid`              |◯|ゾーンのID               | - | string | - |
|`diskoffering`        |-|ディスクオファリング    | - | list(map) | 詳細は後述 |

## diskoffering

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`diskofferingid`      |◯|ディスクオファリングのID  | - | string | - |
|`size`   |◯|ディスクのサイズ | - | int | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |リソースID              | - | 
|`password`    |VMのパスワード | 設定されているならば取得できます | 
|`publicip`    | StaticNATルールを使用してVMに関連付けられたパブリックIPアドレスのID     | StaticNATを有効化する必要があります |