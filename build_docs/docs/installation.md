# Installation

## 1. Terraformのインストール

MacであればHomebrewでインストール可能です。

```sh
$ brew install terraform
```


## 2. APIキーと秘密キーの取得

TerraformからKDDI Cloud Platform Serviceを操作するには、APIを利用するための各種キーが必要です。

KDDIクラウドプラットフォームナレッジサイトの[API利用準備](https://iaas.cloud-platform.kddi.ne.jp/developer/api/cloud-stack-api/use/)を参考にして入手してください。

## 3. Provider Pluginのインストール

Provider PluginはGoバイナリです。ソースコードをダウンロードしてバイナリファイルを作成します。

```sh
$ git clone https://github.com/ezoiwana/terraform-provider-kcps
$ go build
```

`~/.terraformrc`にProvider Pluginのバイナリファイルまでのパスを設定し、プラグインの有効化を行います。以下のように編集してください。

```hcl
providers {
    kcps = "/path/to/terraform-provider-kcps"
}
```

