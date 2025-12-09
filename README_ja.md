# Go Design Patterns (Clean Architecture)

このリポジトリは、GoF (Gang of Four) のデザインパターンを **Go 言語** と **Clean Architecture** の原則に基づいて実装したサンプル集です。
ビジネスロジック（Usecase）と実装詳細（Adapter）を分離し、疎結合な設計を行う方法を学ぶことができます。

詳しく学ぶには、[book を参照してください。](./book/00_introduction_ja.md)

## 📂 プロジェクト構成

各パターンは独立したディレクトリ（Go モジュール）として管理されています。

### 1. 生成に関するパターン (Creational Patterns)

オブジェクトの生成プロセスを柔軟にするためのパターンです。

- [**Abstract Factory**](./abstract-factory-example) (`abstract-factory-example`): 関連する一連のインスタンスを、具体的なクラスを指定せずに生成します。
- [**Builder**](./builder-example) (`builder-example`): 複雑なインスタンスの生成過程を隠蔽し、同じ生成過程で異なる表現を作ります。
- [**Factory Method**](./factory-example) (`factory-example`): インスタンス生成をサブクラス（または実装）に任せます。
- [**Prototype**](./prototype-example) (`prototype-example`): 原型となるインスタンスをコピー（クローン）して新しいインスタンスを作ります。
- [**Singleton**](./singleton-example) (`singleton-example`): インスタンスが一つしか存在しないことを保証します。

### 2. 構造に関するパターン (Structural Patterns)

クラスやオブジェクトを組み合わせてより大きな構造を作るパターンです。

- [**Adapter**](./adapter-example) (`adapter-example`): 互換性のないインターフェースを持つクラス同士をつなぎ合わせます。
- [**Bridge**](./bridge-example) (`bridge-example`): 機能のクラス階層と実装のクラス階層を分離して、それぞれ独立に拡張できるようにします。
- [**Composite**](./composite-example) (`composite-example`): 部分と全体を同一視して、再帰的な構造（ツリー構造）を作ります。
- [**Decorator**](./decorator-example) (`decorator-example`): オブジェクトに動的に責任（機能）を追加します。
- [**Facade**](./facade-example) (`facade-example`): 複雑なシステムに対するシンプルな窓口（インターフェース）を提供します。
- [**Flyweight**](./flyweight-example) (`flyweight-example`): 多数のインスタンスを効率よく扱うために、インスタンスを共有します。
- [**Proxy**](./proxy-example) (`proxy-example`): 本人（オブジェクト）の代わりに代理人が処理を行います（アクセス制御や遅延初期化など）。

### 3. 振る舞いに関するパターン (Behavioral Patterns)

オブジェクト間の責任分担やアルゴリズムの抽象化に関するパターンです。

- [**Chain of Responsibility**](./chain-of-responsibility-example) (`chain-of-responsibility-example`): 要求を処理できるオブジェクトが見つかるまで、チェーン状につないだオブジェクトを順に渡していきます。
- [**Command**](./command-example) (`command-example`): 要求をオブジェクトとしてカプセル化し、履歴管理やアンドゥなどを可能にします。
- [**Interpreter**](./interpreter-example) (`interpreter-example`): 文法規則をクラスで表現し、言語を解釈・実行します。
- [**Iterator**](./iterator-example) (`iterator-example`): 集合体の内部構造を露出させずに、要素を順に走査します。
- [**Mediator**](./mediator-example) (`mediator-example`): 複数のオブジェクト間の相互作用を仲介者（Mediator）に集約し、複雑な依存関係を整理します。
- [**Memento**](./memento-example) (`memento-example`): オブジェクトの状態を保存し、後でその状態に戻せるようにします。
- [**Observer**](./observer-example) (`observer-example`): オブジェクトの状態変化を、依存している他のオブジェクトに通知します。
- [**State**](./state-example) (`state-example`): オブジェクトの内部状態が変化したときに、振る舞いを変えるようにします。
- [**Strategy**](./strategy-example) (`strategy-example`): アルゴリズムを交換可能にします。
- [**Template Method**](./template-method-example) (`template-method-example`): 処理の枠組みを親クラス（または共通構造）で定め、具体的な処理をサブクラス（または実装）に任せます。
- [**Visitor**](./visitor-example) (`visitor-example`): データ構造と処理を分離します。データ構造を変更せずに新しい処理を追加できます。

### その他のパターン

- [**Functional Options**](./functional-options-example): Go 言語特有の、柔軟な構造体初期化パターン。
- [**Entitlement (Gateway)**](./entitlement-example): Clean Architecture に基づいた権限管理の例。Gateway (=Adapter)パターンを用いて、外部リソース（AD/Cache）へのアクセスを抽象化しています。

## 🏗 共通アーキテクチャ構成

各ディレクトリ内は、Clean Architecture を意識した以下の構成になっています。

- **`domain/`**: ビジネスルールの核心。インターフェースやドメインモデルを定義します。外部ライブラリには依存しません。
- **`usecase/`**: アプリケーション固有のビジネスロジック。`domain`のインターフェースを使って処理を記述します。
- **`adapter/`**: インターフェースの実装（Concrete Class）。DB アクセス、API クライアント、具体的なアルゴリズムなどがここに含まれます。
- **`main.go`**: アプリケーションのエントリーポイント。依存性の注入（DI）を行い、各コンポーネントを組み立てて実行します。

## 🚀 実行方法

各ディレクトリに移動して `go run main.go` を実行してください。

```bash
# 例: Factory Methodパターンの実行
cd factory-example
go run main.go
```
