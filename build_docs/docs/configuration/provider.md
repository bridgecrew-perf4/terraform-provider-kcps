# Provider

使用例

```hcl
provider "kcps" {
    api_url     = "https://portal2-east.cloud-platform.kddi.ne.jp:10443/client/api"
    api_key     = "XXXXXXXXXXXXXXXXXXXXXXXXXX"
    secret_key  = "XXXXXXXXXXXXXXXXXXXXXXXXXX"

    # verify_ssl = true
}
```


パラメータ


|パラメータ名 |必須    |説明      |初期値    |タイプ    |補足|
|:----------|:------|:---------|:--------|:--------|:--|
|`api_url`      |◯|APIのエンドポイントURL              | - | string | - |
|`api_key`      |◯|APIキー                            | - | string | - |
|`secret_key`   |◯|秘密キー                             | - | string | - |
|`verify_ssl`   |-|trueなら証明書エラーを無視する            | false | bool | - |
