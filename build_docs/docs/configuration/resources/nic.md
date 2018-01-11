# Nic

## 使用例

```hcl
resource kcps_nic "my_nic"{
    networkid           = "d0f15c14-4d94-4a69-b39a-994721bfc809"
    virtualmachineid    = "be4a143f-3cdd-4091-9a8c-d82c78e49ddf"

    secondaryip         = ["198.18.57.252","198.18.57.254"]
}
```



## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`networkid`           |◯|NetworkのID                | - | string | - |
|`virtualmachineid`   |◯|Nicを搭載させるVMのID  | - | string | - |
|`secondaryip`          |-|Nicに割り当てるセカンダリーIP        | - | list(string) | - |



## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |リソースID              | - | 