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
|`templatefilter` |◯|Templateの検索条件 | - | string | - |
|`template_id` |-|TemplateのID | - | string | - |
|`name` |-|Templateの名前  | - | string | - |
|`keyword` |-|キーワード | - | string | - |
|`zoneid` |-|ZoneのID  | - | string | - |


## 属性
|属性名 |説明      |補足 |
|:----------|:------|:---------|
|`id`          |データソースID   | - | 
|`template_id`  |TemplateのID  | `id`と同じ値 |
|`name`  |Templateの名前  | - |
|`zoneid`  |ZoneのID   | - |
|`format` |Templateのフォーマット | - |
|`hypervisor` |ハイパーバイザ  | - |
|`isdynamicallyscalable` |動的スケーリング  | - |
|`ispublic`  |公開レベル   | - |
|`passwordenabled`  |パスワードが有効であるか   | - |
|`ostypeid`  |OSタイプのID   | - |
|`displaytext`  |表示テキスト   | - |
 
