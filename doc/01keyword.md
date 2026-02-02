
---
간결함이 생명인 Go 언어의 키워드(Keywords)는 총 25개뿐입니다.   
Go는 복잡함을 줄이기 위해 이 키워드들을 엄격하게 제한합니다.  

1. 선언 및 구조 (8개)  
프로그램의 기본 뼈대를 만듭니다.  
- package, import: 패키지 정의 및 호출  
- func: 함수 정의  
- type, struct, interface: 커스텀 타입, 구조체, 인터페이스 선언  
- var, const: 변수 및 상수 선언  

2. 제어 흐름 (10개)  
코드의 실행 방향을 결정합니다.  
- if, else, switch, case, default: 조건문  
- for, range: 반복문 (Go의 유일한 반복 키워드)  
- break, continue, fallthrough: 반복/조건 제어  
- goto: 특정 라벨로 이동 (권장되지 않음)  

3. 함수 제어 및 지연 (3개)    
함수 종료 시점과 관련된 특수 동작을 수행합니다.    
- return: 결과값 반환  
- defer: 함수 종료 직전에 실행 예약 (자원 해제 시 유용)  
- panic: 심각한 오류 발생 시 프로그램 중단  

4. 동시성 (Concurrent) (4개)  
Go의 핵심인 고루틴(Goroutine)과 채널(Channel)을 다룹니다.  
- go: 함수를 비동기(고루틴)로 실행  
- chan: 채널 생성 시 사용  
- select: 여러 채널 작업 중 준비된 것을 선택  
- map: 해시 맵 자료구조 선언  

더 자세한 설명은 Go 언어 사양(The Go Programming Language Specification) 공식 문서에서 확인할 수 있습니다.  

---
## fallthrough

switch 문에서 조건이 일치하는 case를 실행한 후,   
다음 case의 조건 검사 없이 바로 실행하게 만듭니다.  
**Go는 기본적으로 break가 자동 포함된 것처럼 작동하므로**, 
이 키워드가 있어야만 아래로 내려갑니다.  

```go
switch score := 10; score {
case 10:
    fmt.Println("만점입니다.")
    fallthrough // 다음 case인 9로 무조건 진입
case 9:
    fmt.Println("거의 완벽해요.")
default:
    fmt.Println("종료")
}
// 출력: 만점입니다. -> 거의 완벽해요.
```

---
## defer

함수가 종료되기 직전(return 직후)에 특정 코드를 실행하도록 예약합니다.   
주로 파일 닫기나 리소스 해제에 사용됩니다.   
스택 구조라 여러 개 사용 시 역순으로 실행됩니다.  
```go 
func read() {
    fmt.Println("1. 파일을 엽니다.")
    defer fmt.Println("3. 파일을 닫습니다.") // 함수 끝날 때까지 대기
    
    fmt.Println("2. 데이터를 읽습니다.")
}
// 출력: 1 -> 2 -> 3
```

---
## panic

프로그램을 강제로 중단시키는 비상 상황을 발생시킵니다.  
에러 처리가 불가능한 심각한 상황에서 사용하며,   
실행 시 **defer로 예약된 함수들은 다 실행하고** 죽습니다.  

```go
func check(val int) {
    if val < 0 {
        panic("음수는 허용되지 않습니다!") // 여기서 즉시 중단
    }
    fmt.Println("정상:", val)
}
```

---
## panic + recover

panic이 발생했을 때 프로그램을 즉시 종료하지 않고, 다시 제어권을 잡아 복구하는 것이 recover입니다.

**defer 내부에서만 유효**: panic이 터지면 defer로 예약된 함수들만 실행됩니다.  
따라서 recover는 반드시 defer 함수 안에 작성해야 합니다.

**복구 후 흐름**: panic이 발생한 지점 이후의 코드는 실행되지 않지만,  
해당 함수를 호출했던 상위 함수로 돌아가 정상 흐름을 이어갑니다.

```go
package main

import "fmt"

func recoverMe() {
    // defer 내부에서 recover 호출
    if r := recover(); r != nil {
        fmt.Printf("비상 상황 복구 완료: %v\n", r)
    }
}

func startPanic() {
    defer recoverMe() // 1. 복구 함수 예약
    
    fmt.Println("작업 시작...")
    panic("심각한 서버 에러 발생!") // 2. 패닉 발생 (함수 중단)
    
    fmt.Println("이 문구는 출력되지 않습니다.")
}

func main() {
    startPanic()
    fmt.Println("프로그램이 죽지 않고 main으로 돌아왔습니다.") // 3. 정상 실행 계속
}
```

---
## recover()
recover()는 Go 언어의 런타임에 내장된 빌트인 함수(Built-in function)로,  
패닉 상태에 빠진 고루틴을 다시 정상 상태로 되돌리는 역할을 합니다.

**패닉 값 획득**: 패닉이 발생했을 때 panic() 함수에 전달되었던 인자(주로 에러 메시지)를 다시 반환받을 수 있습니다.

**런타임 제어**: 프로그램이 비정상 종료되는 것을 막고, defer를 통해 실행 흐름을 안전하게 복구합니다.

**반환 타입**: any(또는 interface{}) 타입을 반환하며, 패닉 상태가 아니면 nil을 반환합니다.

**실제 사용 예시 (Gin 프레임워크 내부 원리)**  
우리가 사용하는 Gin 프레임워크도 내부적으로 이 빌트인 함수를 사용하여 서버가 죽지 않게 보호합니다.  
```go
func someMiddleware() {
    defer func() {
        if r := recover(); r != nil {
            // 패닉이 발생하면 여기서(defer func 에서) 가로채서 에러 로그를 남김
            fmt.Println("Recovered from panic:", r)
        }
    }()
    // ... 로직 수행 중 panic 발생 가능성 있음
}
```

대충 이해하면,  
defer myFunc() = 실행할 함수 예약  
panic() = 에러 생성 + 함수 중단 + 프로그램 종료 예약  
recover() = 아까 그 에러 가져와 + 프로그램 종료 예약 취소  

---
## 동시성 키워드 4개

go, chan, select, map

1. map (데이터 저장소)
   - **키-값(Key-Value)** 쌍의 해시 테이블입니다. 
   - `make(map[key타입]value타입)` 으로 생성합니다.
2. chan (통로)
   - 고루틴(Goroutine)끼리 데이터를 주고받는 **채널**입니다. 
   - make(chan 타입)으로 선언합니다.
3. select (멀티플렉싱)
   - 여러 채널 중 **데이터가 도착한 채널**을 선택해 실행합니다. 
   - **채널용 switch문** 이라고 생각하면 쉽습니다.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. map: 메뉴와 가격 저장
	menu := map[string]int{
		"Coffee": 4000,
		"Tea":    3500,
	}

	// 2. chan: 주문을 주고받을 통로 생성
	// chan: 데이터를 안전하게 전달하는 파이프 역할을 합니다.
	orderChan := make(chan string)
	abortChan := make(chan bool)

	// 비동기 요리 시작 (고루틴)
	//go: 예제의 go func()처럼 함수를 비동기로 실행하여 멀티태스킹을 가능하게 합니다.
	go func() {
		for {
			// 3. select: 어떤 채널에서 신호가 오는지 감시
			// select: 어떤 채널 작업이 먼저 끝날지 모를 때, 준비된 쪽을 먼저 처리하게 해줍니다.
			select {
			case item := <-orderChan: // 주문 채널에 데이터가 들어오면 실행
				fmt.Printf("[요리] %s(가격:%d원) 제작 중...\n", item, menu[item])
				time.Sleep(time.Second)
				fmt.Printf("[완료] %s 나왔습니다!\n", item)
			case <-abortChan: // 중단 채널에 신호가 오면 종료
				fmt.Println("[알림] 영업을 종료합니다.")
				return
			}
		}
	}()

	// 주문 보내기 (채널 전송)
	orderChan <- "Coffee"
	orderChan <- "Tea"

	time.Sleep(time.Second * 2)
	abortChan <- true // 영업 종료 신호
}
```

---
### `<-` 연산자의 두 얼굴

- `<-` 연산자는 Go의 핵심인 채널(Channel) 전송 및 수신 연산자입니다.
    - 자바에는 이처럼 연산자 수준에서 동시성을 지원하는 문법이 없어서 생소할 수 있습니다.
- 데이터가 흘러가는 방향만 기억하면 쉽습니다.
  - 전송 (Send): `채널 <- 데이터`
    - 데이터를 채널 안으로 밀어 넣습니다.
    - `ch <- 7` (채널 ch에 숫자 7을 보냄)
  - 수신 (Receive): `변수 := <- 채널`
    - 채널에서 데이터를 뽑아옵니다. 화살표가 채널로부터 나가는 모양입니다.
    - `data := <- ch` (채널 ch로부터 값을 꺼내 data에 저장)
- 자세한 활용법은 Go by Example: Channels에서 확인하실 수 있습니다.

---
### go func for select case

Go에서 실시간 이벤트 감시나 백그라운드 워커를 만들 때 가장 흔하게 사용하는 표준 패턴(Idiom)입니다.   
하지만 무조건 고정은 아니며 목적에 따라 변형됩니다.    

왜 이렇게 쓰나요? (무한 루프 + 대기)  

- `for { ... }`: 프로그램을 종료하지 않고 **무한히 반복** 하며 일을 시키기 위함입니다.
- `select { case ... }`: 채널에 신호가 올 때까지 **가만히 대기(Blocking)** 하다가, 신호가 오면 즉시 처리하기 위함입니다.
- 이 둘이 합쳐지면 **"채널에 신호가 올 때마다 반복해서 일을 처리하는 이벤트 리스너"** 가 됩니다.

① 한 번만 기다릴 때 (for 생략)
```go
select {
// func After(d Duration) <-chan Time
// 읽기 전용 시간 채널을 생성해서 리턴
case <-time.After(1 * time.Second):
    fmt.Println("1초 대기 완료")
} // 한 번 실행 후 종료
```

② 비어있는 select (데드락 발생)
```
아무 케이스도 없는 select {}는 고루틴을 영원히 정지시킵니다. 
(주로 메인 함수 종료 방지용으로 사용되나 권장되지는 않습니다.)
```

③ default 케이스 (Non-blocking)  
**신호가 없어도** 기다리지 않고, 바로 넘어가고 싶을 때 사용합니다.  
```go
for {
    select {
    case msg := <-ch:
        fmt.Println("메시지 수신:", msg)
    default:
        fmt.Println("아무 일도 없군...") // 대기 없이 즉시 실행
        time.Sleep(500 * time.Millisecond)
    }
}
```

### 자바와의 비교  

자바에서는 이를 구현하기 위해,  
`while(true)` 안에서 `BlockingQueue` 의 `take()` 메서드를 호출하는 것과 유사합니다.  

Go는 이를 **언어 수준의 키워드**(`select`)로 더 직관적이고 효율적으로 지원하는 것입니다.
