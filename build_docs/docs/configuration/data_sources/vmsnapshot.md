#  VMSnapshot

## 使用例

```hcl
data kcps_vmsnapshot "my_vmsnapshot"{
    displayname = "my_vmsnapshot"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`vmsnapshot_id` |-|スナップショットのID | - | string | - |
|`displayname` |-|VMスナップショットの名前  | - | string | - |
|`state` |-|状態 | - | string | - |
|`virtualmachineid` |-|VMスナップショット作成に使用したVMのID | - | string | - |



## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`vmsnapshot_id`  |スナップショットのID  | `id`と同じ値 |
|`displayname`  | VMスナップショットの名前  | - |
|`state` | 状態 | - | 
|`virtualmachineid` |VMスナップショット作成に使用したVMのID  | - |
