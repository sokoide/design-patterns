# Day 3: 複雑さの隠蔽と制御 (Structural Patterns Part 2)

折り返し地点の Day 3 です！
今日は「構造」パターンの後半として、システムが大規模・複雑になるにつれて重要性を増す 3 つのパターンを学びます。
煩雑な裏側をどうシンプルに見せるか、リソースをどう効率化するか、そしてアクセスをどう制御するか。
それが今日のテーマです。

本日は以下の 3 つのパターンを学びます。

1.  **Facade**: 複雑な裏側を隠す「建物の正面」
2.  **Proxy**: 代理人がアクセスを制御
3.  **Flyweight**: 軽量化して大量生産

---

## 4. Facade (ファサード)

### 📖 ストーリー：ホテルのコンシェルジュ

高級ホテルに泊まったとしましょう。
レストランの予約、タクシーの手配、演劇のチケット購入…これら全てを自分で各業者に電話するのは大変です。
コンシェルジュ（Facade）がいれば、「これとこれをお願い」と言うだけで、裏ですべて手配してくれます。
客（Client）は、裏側の複雑なシステム（サブシステム）を知る必要がありません。

### 💡 コンセプト

複雑なサブシステムに対するシンプルで統一されたインターフェース（窓口）を提供します。

```mermaid
classDiagram
    class Facade {
        +Operation()
    }
    class SubSystemA {
        +Action()
    }
    class SubSystemB {
        +Action()
    }
    Facade --> SubSystemA
    Facade --> SubSystemB
```

### 🐹 Go 実装の極意

Go では、パッケージ設計そのものが Facade の考え方に近いです。
パッケージ内の複雑な構造体や関数は小文字（private）にして隠蔽し、
使いやすい大文字（public）の関数やインターフェースだけを公開することで、パッケージ利用者に優しい設計になります。

### 🧪 ハンズオン

`facade-example` を見てみましょう。
`SmartHomeFacade` が、内部の `Lighting`, `Audio`, `Projector` などのサブシステムを隠蔽し、
`WatchMovie` というシンプルなメソッドだけを提供している様子を確認してください。

### ❓ クイズ

**Q1. Facade パターンを使うと、サブシステムに直接アクセスできなくなる？**
A. はい、完全に禁止されます。
B. いいえ、必要であれば直接アクセスすることも可能です。
C. サブシステムが消滅します。

<details>
<summary>正解</summary>
**B**. Facadeはあくまで「便利な窓口」を提供するだけで、直接アクセスを禁止するものではありません（言語機能で制限しない限り）。
</details>

---

## 5. Proxy (プロキシ)

### 📖 ストーリー：クレジットカード

現金（RealSubject）を持ち歩くのは重いし危険です。
クレジットカード（Proxy）を使えば、支払いの機能は果たせますし、
「限度額チェック（アクセス制御）」や「明細記録（ロギング）」もできます。
本当にお金が必要になった時だけ、銀行口座（実体）から引き落とされます。

### 💡 コンセプト

あるオブジェクトへのアクセスを制御するための代理オブジェクトを提供します。

```mermaid
classDiagram
    class Subject {
        <<interface>>
        +Request()
    }
    class RealSubject {
        +Request()
    }
    class Proxy {
        -realSubject RealSubject
        +Request()
    }
    Subject <|.. RealSubject
    Subject <|.. Proxy
    Proxy o-- RealSubject
```

### 🐹 Go 実装の極意

`Subject` インターフェースを定義し、`RealSubject` と `Proxy` がそれを実装します。
`Proxy` は内部に `RealSubject` を持ちますが、メソッドが呼ばれた時に初めて生成する（遅延初期化）ことも可能です。
Nginx などのリバースプロキシも、アーキテクチャレベルでの Proxy パターンと言えます。

### 🧪 ハンズオン

`proxy-example` で、Nginx のようなアクセス制限機能を持つ Proxy を実装してみましょう。
特定の URL へのアクセスだけを許可し、それ以外は拒否するロジックを追加してみてください。

### ❓ クイズ

**Q2. Proxy パターンの用途として適切でないものは？**
A. 重いオブジェクトの遅延初期化 (Virtual Proxy)
B. アクセス権限のチェック (Protection Proxy)
C. オブジェクトの生成手順の分離 (Builder)

<details>
<summary>正解</summary>
**C**. それは Builder パターンの役割です。
</details>

---

## 6. Flyweight (フライウェイト)

### 📖 ストーリー：森の木々

ゲームで広大な森を描画するとします。木が 100 万本あります。
1 本 1 本の木が「葉のテクスチャ」「幹のモデルデータ」を持っていたら、メモリがパンクしてしまいます。
「木のモデルデータ」は 1 つだけ共有し、100 万本の木は「座標」だけを持つようにすれば、メモリは劇的に節約できます。
これが「Flyweight（フライ級＝軽量級）」です。

### 💡 コンセプト

多数の細かいオブジェクトを効率よく扱うために、共有できる部分（Intrinsic State）を共有し、メモリ使用量を削減します。

```mermaid
classDiagram
    class FlyweightFactory {
        -pool map[string]Flyweight
        +GetFlyweight(key)
    }
    class Flyweight {
        <<interface>>
        +Operation(extrinsicState)
    }
    FlyweightFactory o-- Flyweight
```

### 🐹 Go 実装の極意

`map` を使ったキャッシュ（Factory）を用意し、すでに作成済みのオブジェクトがあればそれを返し、なければ作って登録します。
Go の並行処理環境では、このマップへのアクセスを `sync.RWMutex` などで保護することを忘れないでください。

### 🧪 ハンズオン

`flyweight-example` で、同じ種類のオブジェクトが再利用されていることを、アドレス（ポインタ）を表示して確認してみましょう。

### ❓ クイズ

**Q3. Flyweight パターンで共有すべき情報は？**
A. オブジェクトごとに異なる情報（座標、色など）
B. 全てのオブジェクトで共通の不変な情報（テクスチャ、形状データなど）
C. データベースの接続情報

<details>
<summary>正解</summary>
**B**. 変化しない情報（Intrinsic State）を共有します。
</details>

---

Day 3 はここまでです！
「効率」と「制御」を意識した設計は、大規模システムを支える土台となります。
明日からは、いよいよオブジェクト同士の「振る舞い」の世界に深く飛び込みます。お楽しみに！