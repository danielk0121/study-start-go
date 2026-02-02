package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 변수 선언 및 타입 추론 (:=)
	// var r *gin.Engine = gin.Default() 와 동일한 문장
	// var 생략하고, 타입추론으로 타입 생략 하고
	// 고랭은 세미콜론(;) 사용안함
	//
	// gin 은 왜 소문자 인가 ? Default() 는 왜 대문자로 시작하는가 ?
	// gin : 패키지 이름이라서 소문자
	//
	// Go 언어에는 public이나 private 같은 키워드가 없는 대신, 첫 글자의 대소문자로 접근 권한을 결정합니다.
	//
	// 대문자 시작 (Exported): 패키지 외부에서 호출할 수 있는 함수나 변수를 의미합니다.
	//  gin 패키지 외부에서 Default() 함수를 사용해야 하므로 대문자로 시작합니다.
	//
	// 소문자 시작 (Unexported): 패키지 내부에서만 사용할 수 있습니다.
	//  만약 default() 였다면 우리가 작성하는 코드에서 불러다 쓸 수 없습니다.
	r := gin.Default()

	// 익명 함수 = 클로저
	// Java: 람다 표현식 (c) -> { ... } 이나 익명 클래스를 사용하여 구현합니다.
	// Go: func(c *gin.Context) { ... }: 인라인으로 정의된 익명 함수입니다.
	//
	// c *gin.Context: 함수의 매개변수 선언입니다.
	// Go에서는 "변수 이름이 먼저", "타입이 나중에"
	// *는 Java의 참조 변수와 유사한 포인터를 의미하며, gin.Context 객체의 메모리 주소를 가리킵니다.
	//
	// Go 스타일 가이드(Effective Go)와 커뮤니티 표준에 따르면,
	//  JSON, HTTP, URL, ID와 같은 약어(Initialisms)는 전체를 대문자로 쓰는 것이 원칙입니다.
	//
	//
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Java: 예외 처리를 위해 try-catch-finally 블록과 throw new Exception()을 사용합니다.
	// if err := r.Run(); err != nil { ... }: Go는 예외(Exception) 개념이 없으며,
	//  대신 함수의 반환값으로 오류 객체(error 타입)를 반환하는 방식을 사용합니다.
	//
	// 관용적으로 함수는 여러 값을 반환할 수 있으며, 마지막 반환 값은 보통 error 타입입니다.
	// 오류가 없으면 nil (Java의 null)을 반환합니다.
	//
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

// 원본 코드
//func main() {
//	// Create a Gin router with default middleware (logger and recovery)
//	r := gin.Default()
//
//	// Define a simple GET endpoint
//	r.GET("/ping", func(c *gin.Context) {
//		// Return JSON response
//		c.JSON(http.StatusOK, gin.H{
//			"message": "pong",
//		})
//	})
//
//	// Start server on port 8080 (default)
//	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
//	if err := r.Run(); err != nil {
//		log.Fatalf("failed to run server: %v", err)
//	}
//}
