#  Value VirtualMachine

## 使用例

```hcl
data kcps_value_vm "my_vm"{
    name   = "my_vm"
    zoneid = "593697b6-c123-4025-b412-ef83822733e5"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`valuevm_id` |-|VMのID | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`name` |-|VMの名前  | - | string | - |
|`networkid` |-|VMが接続しているNetworkのID  | - | string | - |
|`state` |-|VMの状態  | - | string | - |
|`templateid` |-|VM作成時に使用したTemplateのID  | - | string | - |
|`zoneid` |-|ZoneのID  | - | string | - |



## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`valuevm_id`  |VMのID  | `id`と同じ値 |
|`name`  |VMの名前  | - |
|`templateid`  |VM作成時に使用したTemplateのID   | - |
|`zoneid`  |ZoneのID   | - |
|`diskofferingid`  |Disk OfferingのID   | - |
|`hypervisor`  |ハイパーバイザ   | - |
|`publicip`  |StaticNATルールを使用してVMに関連付けられたPublic IPのIPアドレス   | StaticNATを有効化する必要があります |
|`publicipid`  |StaticNATルールを使用してVMに関連付けられたPublic IPのID   | StaticNATを有効化する必要があります |
|`serviceofferingid`  |Service OfferingのID   | - |
|`isoid`  |VMにアタッチされているISOのID   | - |
 
