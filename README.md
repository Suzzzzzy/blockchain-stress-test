# 프로젝트명 
> blockchain-stress-test

# 프로젝트 설명
블록체인 네트워크에 대한 스트레스 테스트 및 성능 최적화를 위한 도구를 개발한 프로젝트 입니다.

## 개발
- 개발 언어: Golang:1.21
- GORM
- docker-compose:3.8

## 폴더 구조
```bash
blockchain-stress-test
├── cmd
    └── main.go
├── internal
    └── blockchain
    └── config
    └── database
    └── stress
    └── util
├── initdb.d

```
- cmd: main.go 파일이 위치해 있으며 프로그램의 진입점입니다.
- blockchain: 블록체인을 정의합니다.
- config: 설정을 로드하고 관리합니다.
- database: 데이터베이스 연결을 관리하고, 데이터베이스 모델을 정의합니다.
- stress: 스트레스 테스트를 진행합니다.
- util: 로그를 기록하고 관리합니다.

# 내용
## stress test
목적: 고루틴과 채널을 활용하여 블록체인 네트워크에 대한 스트레스 테스트를 수행
- 고루틴을 사용하여 블록 및 트랜잭션 생성을 병렬로 처리합니다.
- 생성된 블록과 트랜잭션을 데이터베이스에 저장합니다.
- 생성된 블록 및 트랜잭션의 수와 각각의 생성에 소요된 시간을 계산합니다.
- 결과를 반환합니다.
- sync.WaitGroup을 사용하여 고루틴이 모두 종료될 때까지 기다립니다. 이를 통해 모든 작업이 완료된 후에 결과를 반환합니다.
- 고루틴을 통해 동시성을 활용하여 블록 및 트랜잭션 생성과 데이터베이스 저장을 병렬로 처리하고, 결과를 취합하여 반환하는 방식으로 스트레스 테스트를 수행합니다.


## blockchain 자료구조 최적화
목적: 블록 및 트랜잭션 데이터를 효율적으로 저장하고 검색하기 위한 자료구조를 구현
- calculateTransactionHash 함수를 통해 주어진 각각의 트랜잭션에 대해 해시 값을 계산했습니다.
- 트랜잭션 해시 값을 가지고 머클 트리를 구성했습니다.
- 주어진 해시값들을 가지고 머클 트리를 구성하고 최종적으로 머클 루트 해시값을 반환합니다.
- 새로운 블록을 생성할 때 이전 블록의 해시값과 트랜잭션 데이터를 가지고 머클 루트 해시값을 계산하고 새로운 블록을 생성했습니다.
- 이렇게 머클 트리를 활용하여 블록의 무결성을 검증하고 블록에 포함된 트랜잭션들의 머클 루트 해시값을 계산하는 과정이 코드에서 구현되었습니다.
