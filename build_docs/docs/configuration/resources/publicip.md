# Public IP

## 使用例

```hcl
resource "kcps_publicip" "my_publicip" {
    networkid = "7b921a2d-8c82-4016-bbd0-cf0dd6877408"
    
    staticnat {
        virtualmachineid = "b107347a-3447-4e9f-8778-ed49941b5821"
        vmguestip        = "10.1.1.13"
    }
}
```



## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`networkid`    |◯|NetworkID  | - | string | - |
|`staticnat` |-|静的NATの設定   | - | list(map) | 詳細は後述 |

## staticnat

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`virtualmachineid`      |◯|静的NATを適用するVMのID | - | string | - |
|`vmguestip`   |-|VMのセカンダリIP | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |リソースID              | - | 
|`ipaddress`    |IPアドレス | - | 