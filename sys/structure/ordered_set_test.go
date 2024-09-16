package structure

import (
	"github.com/woaijssss/tros/context"
	trlogger "github.com/woaijssss/tros/logx"
	"testing"
)

func TestListSet(t *testing.T) {
	ctx := context.GetContextWithTraceId()
	ls := NewOrderedSetWithCap[int64](100)
	trlogger.Infof(ctx, "ls.Cap == 100? [%+v]", ls.Cap() == 100)
	ls.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	trlogger.Infof(ctx, "ls.Len == 10? [%+v]", ls.Len() == 10)
	ls.Add(20, 11, 33, 15, 23)
	trlogger.Infof(ctx, "ls.Len == 15? [%+v]", ls.Len() == 15)
	elems := ls.AllElements()
	trlogger.Infof(ctx, "elems == {1,2,3,4,5,6,7,8,9,10,11,15,20,23,33}? [%+v]", elems)
}
