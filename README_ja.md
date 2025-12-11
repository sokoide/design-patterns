# Go Design Patterns (Clean Architecture)

このリポジトリは、GoF (Gang of Four) のデザインパターンを **Go 言語** で実装したサンプル集です。  
各例は Clean Architecture の考え方に沿って、ビジネスロジックと実装詳細を分離した構成になっています。

- 英語版 README: `README.md`
- 詳細解説: `book/00_introduction_ja.md`

## 目的

- GoF パターンの意図と使いどころを、Go の型/インターフェース設計に落とし込んで理解する
- Clean Architecture の層分離（`domain` / `usecase` / `adapter`）と依存性注入（DI）の流れを体感する

## 使い方

### 個別の例を実行

各 `*-example` ディレクトリは独立した Go モジュールです。ディレクトリへ移動して実行します。

```bash
# 例: Factory Method パターンの実行
cd factory-example
go run main.go
```

### すべての例をまとめて実行

リポジトリ直下で以下を実行すると、全モジュールを順番に `go run` します。

```bash
make run
```

### テスト

特定の例、またはリポジトリ全体に対して Go 標準のテストを実行できます。

```bash
cd strategy-example
go test ./...

# ルートから全例まとめて
go test ./...
```

## ディレクトリ構成（共通）

各例は Clean Architecture を意識して、概ね次の構成を取ります。

- `domain/`: 依存のないドメイン層。ドメインモデルとインターフェース（契約）を定義します。
- `usecase/`: アプリケーション層。`domain` のインターフェースを使ってユースケースを記述します。
- `adapter/`: 具体実装層。外部 I/O、アルゴリズム、リポジトリ等の実装が入ります。
- `main.go`: エントリーポイント。依存性を組み立てて実行します。

## パターン一覧

### 1. 生成に関するパターン (Creational Patterns)

オブジェクトの生成プロセスを柔軟にするためのパターンです。

- [**Abstract Factory**](./abstract-factory-example) (`abstract-factory-example`): 関連する一連の生成を、具体型に依存せず切り替えられるようにします。
- [**Builder**](./builder-example) (`builder-example`): 複雑な生成手順を分離し、同じ手順で異なる表現を構築します。
- [**Factory Method**](./factory-example) (`factory-example`): 生成の責務をインターフェース/実装側へ委ね、呼び出し側の依存を減らします。
- [**Prototype**](./prototype-example) (`prototype-example`): 既存インスタンスを複製して新しいインスタンスを作ります。
- [**Singleton**](./singleton-example) (`singleton-example`): インスタンスが一つだけになるよう管理します。

### 2. 構造に関するパターン (Structural Patterns)

クラスやオブジェクトを組み合わせてより大きな構造を作るパターンです。

- [**Adapter**](./adapter-example) (`adapter-example`): 互換性のないインターフェース同士をつなぎます。
- [**Bridge**](./bridge-example) (`bridge-example`): 抽象（利用側）と実装（提供側）を分離し、独立に拡張できるようにします。
- [**Composite**](./composite-example) (`composite-example`): 部分と全体を同一視し、木構造を同じ操作で扱えるようにします。
- [**Decorator**](./decorator-example) (`decorator-example`): 既存オブジェクトに機能を動的に重ねます。
- [**Facade**](./facade-example) (`facade-example`): 複雑なサブシステムに対して単純な窓口を提供します。
- [**Flyweight**](./flyweight-example) (`flyweight-example`): 共有により多数のインスタンスを効率よく扱います。
- [**Proxy**](./proxy-example) (`proxy-example`): 本体の代わりに代理が処理し、アクセス制御や遅延初期化を行います。

### 3. 振る舞いに関するパターン (Behavioral Patterns)

オブジェクト間の責任分担やアルゴリズムの抽象化に関するパターンです。

- [**Chain of Responsibility**](./chain-of-responsibility-example) (`chain-of-responsibility-example`): 処理可能なハンドラが見つかるまで要求を連鎖的に渡します。
- [**Command**](./command-example) (`command-example`): 操作をオブジェクト化し、履歴管理や取り消しを可能にします。
- [**Interpreter**](./interpreter-example) (`interpreter-example`): 文法規則を構造化して表現し、解釈・実行します。
- [**Iterator**](./iterator-example) (`iterator-example`): 内部構造を隠したまま要素を順に走査します。
- [**Mediator**](./mediator-example) (`mediator-example`): 相互作用を仲介役に集約し、依存関係を整理します。
- [**Memento**](./memento-example) (`memento-example`): 状態をスナップショットとして保存し、後で復元できます。
- [**Observer**](./observer-example) (`observer-example`): 状態変化を購読者へ通知します。
- [**State**](./state-example) (`state-example`): 状態の切り替えで振る舞いを変えます。
- [**Strategy**](./strategy-example) (`strategy-example`): アルゴリズムを差し替え可能にします。
- [**Template Method**](./template-method-example) (`template-method-example`): 処理の骨格を共通化し、差分だけを実装側に委ねます。
- [**Visitor**](./visitor-example) (`visitor-example`): データ構造と処理を分離し、構造を変えずに新しい処理を追加します。

### その他のパターン

- [**Functional Options**](./functional-options-example): Go でよく使われる柔軟な初期化パターン。
- [**Entitlement (Gateway)**](./entitlement-example): Clean Architecture に基づく権限管理の例。Gateway（= Adapter）で外部リソース（AD/Cache）アクセスを抽象化します。

## 備考

- 各例は学習用の最小構成です。実運用では要件に応じて設計を調整してください。
- 追加の背景や図解は `book/` を参照してください。
