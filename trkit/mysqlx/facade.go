package mysqlx

import (
	"context"
	"github.com/woaijssss/godbx"
	txrequest "github.com/woaijssss/godbx/tx"
	context2 "github.com/woaijssss/tros/context"
)

// InitMysqlX 初始化myslqx
func InitMysqlX(ctx context.Context) {
	initResources(ctx)
}

/*
	以下各函数的 workFn 中，可以使用初始的 tc *godbx.TransContext 将事务传递下去。
	<注意>：一条事务链中，有且仅有一个 tc，否则多个 tc 可能造成 innodb 死锁！
*/

// ReadonlyWithResult 只读单事务
func ReadonlyWithResult[T any](ctx context.Context, workFn func(tc *godbx.TransContext) (T, error)) (T, error) {
	return godbx.AutoTransWithResult(func() (*godbx.TransContext, error) {
		return godbx.NewTransContext(globalDatasource, txrequest.RequestReadonly, context2.GetTraceIdFromContext(ctx))
	}, workFn)
}

// WriteNoResult 只写单事务，无返回值
func WriteNoResult(ctx context.Context, workFn func(tc *godbx.TransContext) error) error {
	return godbx.AutoTrans(func() (*godbx.TransContext, error) {
		return godbx.NewTransContext(globalDatasource, txrequest.RequestWrite, context2.GetTraceIdFromContext(ctx))
	}, workFn)
}

// WriteWithResult 只写单事务，有返回值
func WriteWithResult[T any](ctx context.Context, workFn func(tc *godbx.TransContext) (T, error)) (T, error) {
	return godbx.AutoTransWithResult(func() (*godbx.TransContext, error) {
		return godbx.NewTransContext(globalDatasource, txrequest.RequestWrite, context2.GetTraceIdFromContext(ctx))
	}, workFn)
}
