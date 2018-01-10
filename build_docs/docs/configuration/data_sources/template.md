#  Template

## 使用例

```hcl
data kcps_template "my_template"{
    templatefilter = "self"
    name           = "Ubuntu16.04 Server"
}
```

## パラメータ

|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`templatefilter` |◯|テンプレートの検索条件 | - | string | - |
|`template_id` |-|テンプレートのID | - | string | - |
|`name` |-|テンプレートの名前  | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`zoneid` |-|ゾーンのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`template_id`  |テンプレートのID  | `id`と同じ値 |
|`name`  |テンプレートの名前  | - |
|`zoneid`  |ゾーンのID   | - |
|`format` |テンプレートのフォーマット | - |
|`hypervisor` |ハイパーバイザ  | - |
|`isdynamicallyscalable` |動的スケーリング  | - |
|`ispublic`  |公開レベル   | - |
|`passwordenabled`  |パスワードが有効であるか   | - |
|`ostypeid`  |OSタイプのID   | - |
 
