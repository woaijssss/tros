package structure

import (
	"fmt"
	"gitee.com/idigpower/tros/context"
	trlogger "gitee.com/idigpower/tros/logx"
	"testing"
)

func TestStack(t *testing.T) {
	ctx := context.GetContextWithTraceId()
	st := Stack[int32]{}
	fmt.Println(st)
	st.Push(10)
	st.Push(20)
	st.Push(30)
	trlogger.Infof(ctx, "stack size == 3? [%+v]", st.Len() == 3)
	v := st.Pop()
	trlogger.Infof(ctx, "pop v == 30? [%+v]", v == 30)
}
