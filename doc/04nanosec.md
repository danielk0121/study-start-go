
---
### 곱하기 연산자

time.After() 함수를 보면 Duration 을 인자로 받음  
```go
func After(d Duration) <-chan Time {
	return NewTimer(d).C
}
```

After(1 * time.Second) 코드를 보면 1 과 타입을 곱하기 연산자를 사용함  
After(Duration(1, time.Second)) 와 같은 형식 일 것이라고 생각하는데,  
곱하기 연산자를 사용함...
```go
select {
// 읽기 전용 시간 채널을 생성해서 리턴
case <-time.After(1 * time.Second):
    fmt.Println("1초 대기 완료")
} // 한 번 실행 후 종료
```

그 이유는 Go 언어의 **상수(Untyped Constant)와 `Duration` 타입의 내부 구조** 때문입니다.  

`Duration`은 사실 `int64` 입니다.  
Go의 `time` 패키지를 보면 `Duration` 은 다음과 같이 정의되어 있습니다.  

```go
type Duration int64

const (
Nanosecond  Duration = 1
Microsecond          = 1000 * Nanosecond
Millisecond          = 1000 * Microsecond
Second               = 1000 * Millisecond 
// Second = 결과적으로 1,000,000,000 (나노초)
)
```

즉, `time.Second` 는 **1,000,000,000** 이라는 숫자를 가진 `Duration` 타입의 값일 뿐입니다. 
Go 패키지 문서: Duration에서 이 수치들을 확인할 수 있습니다.

---
### 나노초

`time` 패키지가 아닌 **Go 런타임(Runtime) 스케줄러**가   
CPU의 **하드웨어 타이머** 와 직접 대화하기 때문에 가능합니다.

### 연결 고리: `runtime.nanotime()`

`time.Sleep` 이나 `time.After` 를 호출하면 결국 **runtime/time.go** 에   
정의된 런타임 함수로 연결됩니다.

- Go 런타임은 OS로부터 현재 시간을 나노초 단위의 정수로 가져오는 `nanotime()` 함수를 가지고 있습니다.
- Linux 환경이라면 `clock_gettime` 시스템 콜을 사용해 하드웨어 클럭 값을 나노초로 읽어옵니다.

### 관련 소스 코드 위치

- 사용자 인터페이스: `src/time/sleep.go` (우리가 호출하는 `Sleep` 정의)
- 런타임 연결: `src/runtime/time.go` (타이머를 관리하는 실제 로직)
- 시스템 콜: `src/runtime/sys_linux_amd64.s` 또는 `sys_darwin_arm64.s` 등 (OS에 시간을 물어보는 어셈블리 코드)

결국 "10 = 10ns"라는 약속이 지켜지는 이유는,   
Go 런타임이 OS에게 "나노초 단위의 시스템 시계"를 기준으로 대기 명령을 내리기 때문입니다.

### 쉽게 요약
- 1 Duration = 1 나노초, Go 에서 규칙으로 정함
- 나노초 구하는 과정
    - 1.. OS 시스템 콜을 통해 어셈블리 코드 실행하여 나노초 구하는 값 얻음
    - 2.. cpu 특성을 고려하여, 나노초 보정 (틱 계산)
    - 3.. runtime.nanotime() 나노초 정수값 리턴







