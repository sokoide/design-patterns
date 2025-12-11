# Go Iterator Pattern Example (Clean Architecture)

このプロジェクトは、**Go**言語を用いて**Iterator Pattern（イテレータパターン）**を実装した教育用のサンプルコードです。コレクション（集合体）の内部構造を露出させずに、要素に順番にアクセスする方法を学びます。

## この例で学べること

- コレクションの内部構造（slice）を隠して走査する方法
- `HasNext` / `GetNext` で要素を順に取得する Iterator の実装

## すぐ試す

`iterator-example` ディレクトリで実行します。

```bash
go run main.go
```

## 🔄 シナリオ: ユーザーリストの走査

`User` のリストを持つ `UserCollection` があります。
クライアント（`main.go`）は、このコレクションが内部で配列を使っているのか、連結リストなのか、あるいはマップなのかを知りたくありません。
ただ「次の要素をくれ（Next）」と言いたいだけです。Iteratorを使ってこれを実現します。

### 登場人物
1.  **Iterator (`domain.Iterator`)**: 「次はあるか？(`HasNext`)」「次の要素は？(`GetNext`)」という走査のためのインターフェース。
2.  **Aggregate (`domain.Collection`)**: Iteratorを生成するためのインターフェース（`CreateIterator`）。
3.  **Concrete Iterator (`adapter.UserIterator`)**: 実際に配列のインデックスを管理し、走査を行うクラス。
4.  **Concrete Aggregate (`adapter.UserCollection`)**: 実際のデータ保持者。

## 🏗 アーキテクチャ構成

```mermaid
classDiagram
    direction TB

    %% Domain Layer
    class Iterator {
        <<interface>>
        +HasNext() bool
        +GetNext() interface{}
    }

    class Collection {
        <<interface>>
        +CreateIterator() Iterator
    }

    %% Adapter Layer
    class UserCollection {
        +Users: User[]
        +CreateIterator() Iterator
    }

    class UserIterator {
        -users: User[]
        -index: int
        +HasNext() bool
        +GetNext() interface{}
    }

    %% Relationships
    UserCollection ..|> Collection : Implements
    UserIterator ..|> Iterator : Implements
    UserCollection ..> UserIterator : Creates
```

### 各レイヤーの役割

1.  **Domain (`/domain`)**:
    *   `Iterator`: 走査の標準API定義。
    *   `Collection`: イテレータを提供するものの定義。
2.  **Adapter (`/adapter`)**:
    *   `UserCollection`: 具体的なユーザーリスト。
    *   `UserIterator`: `UserCollection`専用のイテレータ。現在の進捗（`index`）を状態として持ちます。

## 💡 アーキテクチャ設計ノート (Q&A)

### Q1. Goには `for range` ループがありますが、これを使う意味は？

**A. コレクションの内部構造を隠蔽（カプセル化）したい場合に有効です。**

標準の `range` は、スライスやマップといった「Goの組み込み型」に対してしか使えません。
「ツリー構造を深さ優先探索で巡回したい」「グラフ構造を走査したい」といった複雑なデータ構造の場合、Iteratorパターンを実装することで、利用者は内部アルゴリズムを気にせず `HasNext() / GetNext()` だけで統一的に扱えるようになります。

### Q2. `GetNext()` が `interface{}` を返していますが、型安全性は？

**A. この実装例では型安全性は失われています（ジェネリクス以前のGoスタイル）。**

Go 1.18以降の **Generics** を使えば、`Iterator[T]` のように型安全なイテレータを作ることが可能です。
今回の例はGoFパターンの古典的な構造を模倣しているため `interface{}` を使用していますが、実務ではジェネリクス推奨です。

## 🚀 実行方法

```bash
go run main.go
```
