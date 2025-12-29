# Go Design Patterns (Clean Architecture)

このリポジトリは、GoF (Gang of Four) のデザインパターンを **Go 言語** で実装したサンプル集です。
各例は Clean Architecture の考え方に沿って、ビジネスロジックと実装詳細を分離した構成になっています。

- [英語版 README](./README.md)
- [詳細解説](./book/00_introduction_ja.md)

## 目的

- GoF パターンの意図と使いどころを、Go の型/インターフェース設計に落とし込んで理解する
- Clean Architecture の層分離（`domain` / `usecase` / `adapter`）と依存性注入（DI）の流れを体感する

## 使い方

### 個別の例を実行

各 `*-example` ディレクトリは独立した Go モジュールです。ディレクトリへ移動して実行します。

```bash
cd strategy-example
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

オブジェクトの生成プロセスを柔軟にするためのパターンですが、Go の特性上、標準的なコンストラクタで十分な場合が多く、本リポジトリでは以下の理由により実装を省略（または無視）しています。

- **Builder** (`builder-example`): 複雑な生成手順を分離し、同じ手順で異なる表現を構築します。
  - Go では **Functional Options** を使うことが多いため、省略します。
- **Abstract Factory** (`abstract-factory-example`): 関連する一連の生成を、具体型に依存せず切り替えられるようにします。
  - 実際使われることはまずありません。省略します。
- **Factory Method** (`factory-example`): 生成の責務をインターフェース/実装側へ委ね、呼び出し側の依存を減らします。
  - Go ではあまり見かけません。
  - `NewServer(cfg Config) (*Server, error)` や `Open(path string) (*File, error)` などのコンストラクタで事足りることがほとんどです。
  - 動的にインスタンスを作り分ける必要がなければ使いません。
- **Singleton** (`singleton-example`): インスタンスが一つだけになるよう管理します。
  - **避けるべき** です。テスト性や並行性の観点からグローバルな状態は推奨されません。
- **Prototype** (`prototype-example`): 既存インスタンスを複製して新しいインスタンスを作ります。
  - 構造体のコピーコストが小さいため（`b := a`）、省略します。

### 2. 構造に関するパターン (Structural Patterns)

クラスやオブジェクトを組み合わせてより大きな構造を作るパターンです。

- [**Adapter**](./adapter-example) (`adapter-example`): 互換性のないインターフェース同士をつなぎます。
- [**Decorator**](./decorator-example) (`decorator-example`): 既存オブジェクトに機能を動的に重ねます。
- **Composite** (`composite-example`): 部分と全体を同一視し、木構造を同じ操作で扱えるようにします。
- [**Facade**](./facade-example) (`facade-example`): 複雑なサブシステムに対して単純な窓口を提供します。
- [**Proxy**](./proxy-example) (`proxy-example`): 本体の代わりに代理が処理し、アクセス制御や遅延初期化を行います。
- [**Flyweight**](./flyweight-example) (`flyweight-example`): 共有により多数のインスタンスを効率よく扱います。

以下は無視します。

- **Bridge** (`bridge-example`): 抽象（利用側）と実装（提供側）を分離し、独立に拡張できるようにします。
  - `interface` があるため不要です。

### 3. 振る舞いに関するパターン (Behavioral Patterns)

オブジェクト間の責任分担やアルゴリズムの抽象化に関するパターンです。

- [**Strategy**](./strategy-example) (`strategy-example`): アルゴリズムを差し替え可能にします。
- [**Observer**](./observer-example) (`observer-example`): 状態変化を購読者へ通知します。
- [**Command**](./command-example) (`command-example`): 操作をオブジェクト化し、履歴管理や取り消しを可能にします。
- [**Chain of Responsibility**](./chain-of-responsibility-example) (`chain-of-responsibility-example`): 処理可能なハンドラが見つかるまで要求を連鎖的に渡します。
- [**State**](./state-example) (`state-example`): 状態の切り替えで振る舞いを変えます。
- [**Memento**](./memento-example) (`memento-example`): 状態をスナップショットとして保存し、後で復元できます。

以下は無視します。

- **Iterator** (`iterator-example`): 内部構造を隠したまま要素を順に走査します。
  - `for range` があるし、自分で Iterate するときは `Next() (T, bool)` するだけなので、飛ばします。
- **Mediator** (`mediator-example`): 相互作用を仲介役に集約し、依存関係を整理します。
  - 使うと巨大な God Object になりがちなので、飛ばします。
- **Template Method** (`template-method-example`): 処理の骨格を共通化し、差分だけを実装側に委ねます。
  - Go では継承がないので、関数・小さいインターフェースとコンポジションで達成します。
- **Visitor** (`visitor-example`): データ構造と処理を分離し、構造を変えずに新しい処理を追加します。
  - 言語処理系を作るとき便利ですが、それ以外は `switch n := n.(type)` で代替されることが多いため、飛ばします。
- **Interpreter** (`interpreter-example`): 文法規則を構造化して表現し、解釈・実行します。
  - インタプリタやコンパイラを作る際には重要ですが、本リポジトリでは対象外とします。

### その他のパターン

Goでは頻出です。

- [**Functional Options**](./functional-options-example): Go でよく使われる柔軟な初期化パターン。

## 備考

- 各例は学習用の最小構成です。実運用では要件に応じて設計を調整してください。
- 追加の背景や図解は [book](./book/00_introduction_ja.md) を参照してください。
