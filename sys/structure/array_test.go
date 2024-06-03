package structure

import (
	"gitee.com/idigpower/tros/context"
	trlogger "gitee.com/idigpower/tros/logx"
	"testing"
)

func TestArray(t *testing.T) {
	ctx := context.GetContextWithTraceId()
	a := NewFromArray[int64]([]int64{0, 1, 2, 2, 3, 4, 9, 5, 3, 4, 2, 2, 3, 45, 5, 6, 6})

	trlogger.Infof(ctx, "a: [%+v]", a.Array())
	a.RemoveDuplicates()
	trlogger.Infof(ctx, "a: [%+v]", a.Array())
}
