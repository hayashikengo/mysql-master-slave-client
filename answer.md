# 回答１
> ※ primary・readreplica構成 == master・slave構成

> ※ 解決策はチェックボックスにしてあります。解決済みはチェック。

- ライブラリの依存管理がされていない
  - [x] go.mod追加
  - [x] examples配下は、別途管理
- ライブラリのREADME.mdが無い
  - [x] README.md追加
    - [x] Overview
    - [x] Getting Start
    - [x] Dev tool
    - [x] License
- testが無い
  - [x] test追加
    - [ ] testの依存関係をmock使って分離したい
  - [x] 実行時は、-raceオプションつける
- ヘルスチェックされていない
  - [x] ヘルスチェック実装
- 意図したエラー発生時に、panicを発生させている
  - [x] errを返すようにする
- readreplica が全て死んだ際に Fallback していない
  - [x] readreplica全滅の際は、primaryにつなげる
  - [x] optionalにしたい
  - [x] primaryが死んだ場合エラー返す
- readreplica の負荷分散アルゴリズムがRoundRobinしか無い
  - [x] Random追加
  - [x] optionalにしたい
  - [ ] 遅延の少ないDBを選択するアルゴリズム追加
- ライブラリを使った場合のオーバーヘッドを知りたい
  - [x] `mydb vs database/sql` のbenchmark追加
- ci欲しい
  - [x] `go build` 回す
  - [x] `go test` 回す
  - [x] `reviewdog` 入れたい
- 開発環境ほしい
  - [x] docker-compose 入れる

# 回答３
- 調査面
  - 同じようなことをやろうとしているライブラリを探して、3つほど読んだ
    - [tsenart/nap](https://github.com/tsenart/nap)
    - [linxGnu/mssqlx](https://github.com/linxGnu/mssqlx)
    - [badoux/gorb](https://github.com/badoux/gorb)
  - ライブラリ自体の構成は、gorm　などの有名ライブラリを参考にした
- ドキュメント面
  - 導入検討に必要な情報を`README.md`にまとめた
    - Getting Started... 手早く試せるように
    - Benchmark... database/sql vs mydb のパフォーマンスが分かるように
  - mydbライブラリの開発は、`README.md`を見れば始められるようにした
    - このライブラリを使用するチームの人にも、できたらPRとか出してほしいため
- 実装面
  - readreplica　の負荷分散部分を `dbBalancer` type に抽象化した
  - interfaceを使って、各typeの I/F を定義した。
  - 運用状況によって変更したい内容はoptionalにした
- 開発環境面
  - `docker-compose up` で master-slave 構成を手軽に動かせるようにした
    - master: １つ
    - readreplica: 3つ
  - `Makefile`にて、開発時に使うコマンドをまとめた