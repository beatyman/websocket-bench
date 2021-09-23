package impl

import (
	"context"
	"reflect"
	"time"
	"websocket-bench/api"
)

type CommonAPI struct {
	//some other
}
func (n *CommonAPI) GetTime(context.Context) (time.Time, error) {
	return time.Now(),nil
}

func proxy(in interface{}, outstr interface{}) {
	outs := api.GetInternalStructs(outstr)
	for _, out := range outs {
		rint := reflect.ValueOf(out).Elem()
		ra := reflect.ValueOf(in)

		for f := 0; f < rint.NumField(); f++ {
			field := rint.Type().Field(f)
			fn := ra.MethodByName(field.Name)

			rint.Field(f).Set(reflect.MakeFunc(field.Type, func(args []reflect.Value) (results []reflect.Value) {
				ctx := args[0].Interface().(context.Context)
				// upsert function name into context
				// pass tagged ctx back into function call
				args[0] = reflect.ValueOf(ctx)
				return fn.Call(args)
			}))
		}
	}
}

func NewCommonAPI(a api.Common) api.Common {
	var out api.CommonStruct
	proxy(a, &out)
	return &out
}