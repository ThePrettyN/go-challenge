
# Go Docker アプリ

このプロジェクトは、2つのシンプルなGoプログラムを示しています：

1. **IPアドレスフォーマッタ** (`main1.go`)： `fmt.Stringer`インタフェースを実装して、IPアドレスをドット区切り形式（例：`127.0.0.1`）でフォーマットします。
2. **ROT13リーダー** (`main2.go`)： `io.Reader`インタフェースを実装して、ROT13換字式暗号を用いて文字列をデコードします。

---

## プロジェクト構成

```plaintext
go-docker-app/
├── cmd/
│   ├── main1.go        # IPアドレスフォーマッタ
│   ├── main2.go        # ROT13リーダー
├── docker-compose.yml  # Docker Composeの設定ファイル
├── Dockerfile          # アプリをビルド・実行するDockerfile
├── go.mod              # Goモジュール定義
├── go.sum              # Goの依存関係
```

---

## プログラムの説明

### 1. **IPアドレスフォーマッタ (`main1.go`)**
このプログラムは、`fmt.Stringer`インタフェースを使用してIPアドレスをドット区切り形式にフォーマットし、名前に対応するIPアドレスのリストを出力します。

#### 実行例:
```plaintext
loopback: 127.0.0.1
googleDNS: 8.8.8.8
```

---

### 2. **ROT13リーダー (`main2.go`)**
このプログラムは、`io.Reader`を実装して文字列を読み取り、ROT13暗号をすべてのアルファベット文字に適用します。ROT13エンコードされたメッセージを読み取り、デコードする動作を示します。

#### 実行例:
```plaintext
You cracked the code!
```

---

## Dockerセットアップ

### 必要条件

- Docker
- Docker Compose

### プロジェクトのビルドと実行

1. **Dockerコンテナをビルドして起動**:
   ```bash
   docker-compose up --build
   ```

2. **出力を確認**:
   両プログラム（`main1.go`と`main2.go`）の出力がターミナルに表示されます。

---

## 動作の仕組み

1. `Dockerfile`はGoプロジェクトをビルドし、2つのバイナリ（`ipaddr`と`rot13`）を生成します。
2. `docker-compose.yml`はコンテナを起動し、これら2つのバイナリを順次実行します。

---

## 注意点

- このプロジェクトはGoの基本的なコンセプトを示しています：
  - インタフェースの実装（`fmt.Stringer`と`io.Reader`）。
  - Dockerを使用したGoアプリケーションのビルドと実行。
- `main1.go`や`main2.go`を変更・拡張して、Goの機能をさらに学ぶことができます！
