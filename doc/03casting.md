
---
### 형변환

```go
select {
  case c.(chan Time) <- Now().Add(Duration(-delta)):
  default:
}
```
이 코드에서 c.(chan Time) 부분이 이해가 안됨

이 문법은 Go의 **타입 단언(Type Assertion)** 입니다. 
c라는 변수가 인터페이스(interface{} 또는 any) 타입일 때, 
이를 실제 사용 가능한 **채널(chan Time) 타입으로 변환** 하는 과정입니다.

- `c`: 현재 이 변수는 무엇이든 담을 수 있는 `interface{}` 타입으로 선언되어 있을 것입니다.
- `.(chan Time)`: "너(`c`) 안에 들어있는 실제 데이터가 `chan Time` 타입이 맞지 ? 
  - 맞다면 그 타입으로 동작하게 해줘"라고 요청하는 것입니다.
- 만약 `c`가 채널이 아니라면 이 지점에서 **패닉(Panic)** 이 발생합니다.

주로 `time.After` 나 `Timer` 의 내부 로직에서 볼 수 있는데,   
여러 타입을 공통으로 다루는 인터페이스 구조에서 **특정 채널에 시간을 전송하기 위해** 타입을 강제하는 것입니다.  

코드 해석
```go
select {
    // c를 채널로 변환한 뒤, 현재 시간에서 delta를 뺀 값을 전송 시도
    case c.(chan Time) <- Now().Add(Duration(-delta)):
    
    // 만약 채널이 가득 찼거나(Buffer full) 받을 준비가 안 되어 있다면
    default:
        // 대기하지 않고 그냥 넘어감 (Non-blocking)
}
```

---
### 자바와 비교 (Casting)

- 자바의 형변환(Casting)과 유사
- Java: `((Channel<Time>) c).send(time);`
- Go: `c.(chan Time) <- time`

---
### 콤마 ok 패턴

실무에서는 패닉을 방지하기 위해 다음과 같이 콤마 ok 패턴을 주로 사용합니다.  

```go
if ch, ok := c.(chan Time); ok {
    ch <- Now()
} else {
    fmt.Println("c는 채널이 아닙니다!")
}
```






