package main

/*
#include <mach/mach_time.h>
*/
import "C"
import (
	"fmt"
	"time"
)

// Go에서 구현 방법 (CGO 활용)
//
// Go 소스 코드 내에서 C 패키지를 사용하여 macOS 시스템 API에 직접 접근할 수 있습니다.
//
// 역할: Go 코드 안에서 C 언어 함수를 쓰겠다는 선언입니다.
//   <mach/mach_time.h>: macOS(Darwin) 전용 커널 헤더 파일입니다.
//   여기에 mach_absolute_time() 같은 하이 레벨 타이머 함수들이 정의되어 있습니다.
//
// 네, mach/mach_time.h는 macOS 시스템(Darwin 커널)에 기본적으로 포함되어 있는 헤더 파일이 맞습니다
//
// Go 표준 라이브러리: 사실 Go 1.9 버전 이후부터
//   time.Now()는 내부적으로 macOS에서 단조 클록을 자동으로 사용하므로,
//   특수한 커널 수준의 정밀도가 필요한 것이 아니라면
//   Go 공식 문서의 time 패키지를 활용하는 것이 이식성 면에서 유리합니다.

// getMachAbsoluteTimeNano는 현재 mach_absolute_time()을 나노초로 변환합니다.
func getMachAbsoluteTimeNano() uint64 {
	// 1. Mach timebase info 가져오기 (비율 설정)
	var info C.mach_timebase_info_data_t
	C.mach_timebase_info(&info)

	// 2. mach_absolute_time() 호출
	tick := C.mach_absolute_time()

	// 3. 틱을 나노초로 변환
	// (tick * numer) / denom
	nano := uint64(tick) * uint64(info.numer) / uint64(info.denom)
	return nano
}

func main() {
	// 고루틴 내에서의 측정 예시
	start := getMachAbsoluteTimeNano()

	// 작업 수행 (예: 100ms 대기)
	time.Sleep(100 * time.Millisecond)

	end := getMachAbsoluteTimeNano()

	durationNano := end - start
	fmt.Printf("작업 소요 시간: %v ns\n", durationNano)
	fmt.Printf("작업 소요 시간: %v ms\n", durationNano/uint64(time.Millisecond))
	// 101 038 750
	// 작업 소요 시간: 101038750 ns
	// 작업 소요 시간: 101 ms
}
