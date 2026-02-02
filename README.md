# start go lang


---
# 언어 특징

- 2009년 구글에서 개발
- 인프라, 서버 시스템 개발에 최적화
- 미니멀리즘
  - 키워드가 25개 내외
- 명확한 코드
  - 상속, 제네릭 사용 지양
- 뛰어난 동시성 지원
  - 고루틴 : OS 스레드보다 가벼운 논리적 스레드
  - 채널 : 데이터를 공유하는 대신, 메시지를 주고 받는 방식으로 동기화 문제 해결
- 빠른 성능
  - 컴파일 언어, 빠른 컴파일 속도가 설계 단계부터 고려
- 정적 타입
  - 컴파일 시점에 변수 타입이 결정
- 단일 실행 파일
  - 컴파일 시 필요한 모든 라이브러리 포함
  - 별도의 jvm, python 인터프리터 같은 런타임 설치 없이 배포 가능
- 강력한 표준 도구
  - 코드 포맷팅 go fmt
  - 테스트 go test
  - 패키지 관리 go mod
- 주요 활용 분야
  - docker, kubernetes, msa, 데이터 처리 파이프라인


---
# todo task

- 자료 조사 (doc) + 예제 코드
- [x] 키워드 : [01keyword.md](doc/01keyword.md)
- [ ] 연산자
- [ ] 패키지, 변수, 클래스, 타입, 구조체
- [ ] 형변환
- [ ] 함수, 인터페이스
- [ ] 상속, 인터페이스
- [ ] 제어문 if, for, while, switch
- [ ] 문자
  - [ ] ' (작은따옴표), " (큰따옴표), ` (백틱)
  - [ ] 자르기 : sub
  - [ ] 비교 : startWith, endWith, equal, has
  - [ ] 정보 : length
  - [ ] 인코딩
- [ ] 자료 구조
    - [ ] 배열, 리스트, 맵
- [ ] 파일
  - [ ] 파일 읽기, 쓰기, 폴더 탐색
- [ ] http client
- [ ] json
- [ ] sql
  - [ ] mysql CRUD
  - [ ] mysql 설정
- [x] 웹 서버 (Gin) 테스트 실행 : [cmd/hello_gin/main.go](cmd/hello_gin/main.go)
- [ ] 샘플 도메인 rest api 작성
- [ ] 기타 특징 (defer, go, 등)


---
# 문서 doc 목록

- [01keyword.md](doc/01keyword.md)
- [02filename_package_name.md](doc/02filename_package_name.md)
- [03casting.md](doc/03casting.md)
- [04nanosec.md](doc/04nanosec.md)


---
# 기타 (이런것 주의, 참고 ?!)

- go.mod 파일
  - go.mod 파일은 Go 프로젝트의 패키지 관리자 설정 파일입니다.
  - Java의 pom.xml(Maven)이나 Python의 requirements.txt, Node.js의 package.json과 비슷한 역할을 합니다. 

