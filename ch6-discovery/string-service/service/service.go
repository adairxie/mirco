package service

import (
	"errors"
	"strings"
)

// Service constants
const (
	StrMaxSize = 1024
)

// Service errors
var (
	ErrMaxSize  = errors.New("maximum size of 1024 bytes exceeded")
	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

// Service Define a service interface
type Service interface {
	// 连接字符串a, b
	Concat(a, b string) (string, error)

	// 获取字符串a, b 公共字符
	Diff(a, b string) (string, error)

	// 健康检查
	HealthCheck() bool
}

// StringService implement Service interface
type StringService struct {
}

// Concat concat a, b string
func (s StringService) Concat(a, b string) (string, error) {
	if len(a)+len(b) > StrMaxSize {
		return "", ErrMaxSize
	}
	return a + b, nil
}

// Diff get the common characters between a and b string
func (s StringService) Diff(a, b string) (string, error) {
	if len(a) < 1 || len(b) < 1 {
		return "", nil
	}

	res := ""

	if len(a) >= len(b) {
		for _, char := range b {
			if strings.Contains(a, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range a {
			if strings.Contains(b, string(char)) {
				res = res + string(char)
			}
		}
	}
	return res, nil
}

// HealthCheck implements Service method
func (s StringService) HealthCheck() bool {
	return true
}

// ServiceMiddleware define Service middleware
type ServiceMiddleware func(Service) Service
