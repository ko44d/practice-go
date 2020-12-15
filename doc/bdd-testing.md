# Go言語でのビヘイビア駆動開発（BDD）について

## 概要
Go言語のBDD用のフレームワークとしてGinkgoを使っているため、Ginkgoでのユニットテストの実行方法や実装方法についてまとめる。

## ビヘイビア駆動開発（BDD）とは
- https://ja.wikipedia.org/wiki/ビヘイビア駆動開発

## Ginkgoとは
- https://onsi.github.io/ginkgo/
- Go言語のBDD用テストフレームワークである
- Gomaga というMatcherライブラリと組み合わせるのが最適である
    - 但し、Matcherライブラリには依存しない設計となっている
