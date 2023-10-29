# go-architecture-multiple-packages

## 概要
Goチームが紹介しているパッケージ構成のうちの一つである`multiple-packages`の実装例を紹介します([参考リンク](https://go.dev/doc/modules/layout))
interanalという名前のディレクトリ内のパッケージは、異なるモジュールからの参照をできない仕様にGo言語ではなっています。internalと同じようにディレクトリを作成しているauthディレクトリ内のパッケージは別モジュールからの参照が可能です。

## 実行方法
```go
go run ./...
```
