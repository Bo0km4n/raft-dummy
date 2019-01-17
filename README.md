# Raft implementation

I implment Raft by experiment.

## Test commit
`make test`

## About
Raftを実装した分散インメモリKVS

## Project architecture
.
├── Gopkg.lock
├── Gopkg.toml
├── Makefile
├── README.md
├── cmd
│   ├── node.go
│   └── root.go
├── kvs
│   └── kvs.go
├── main.go
├── proto
│   ├── raft.pb.go
│   └── raft.proto
├── raft
│   ├── commit_test.go
│   ├── grpc.go
│   ├── handler.go
│   ├── log.go
│   └── node.go
├── test.sh
└── testdata
    ├── log_1.log
    ├── log_2.log
    ├── log_3.log
    ├── log_4.log
    ├── log_5.log
    └── nodes.json

5 directories, 22 files

### cmd
起動コマンドのオプション及びエントリーポイント

### kvs
インメモリKVSの実体
インターフェイスとしてraftのstate構造体にスタブしている

### raft
raftのコアロジック実装部分

### proto
protobufの定義ファイルとコンパイル済みのgoファイル

### testdata
ローカルで実験する際logファイルの書き出しとコンフィグのテストファイル