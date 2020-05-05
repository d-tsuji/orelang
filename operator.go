package orelang

import "log"

type IOperator interface {
	call(engine *Engine, args []interface{}) interface{}
}

type AddOperator struct{}

func (AddOperator) call(engine *Engine, args []interface{}) interface{} {
	var retValue float64
	for _, arg := range args {
		v := engine.Eval(arg)
		iv, ok := v.(float64)
		if !ok {
			log.Fatalf("cannot convert float64 value, %v (%T)", v, v)
		}
		retValue = retValue + iv
	}

	return retValue
}

type MultiplyOperator struct{}

func (MultiplyOperator) call(engine *Engine, args []interface{}) interface{} {
	retValue := 1.0
	for _, arg := range args {
		v := engine.Eval(arg)
		iv, ok := v.(float64)
		if !ok {
			log.Fatalf("cannot convert float64 value, %v (%T)", v, v)
		}
		retValue = retValue * iv
	}

	return retValue
}

type EqualOperator struct{}

func (EqualOperator) call(engine *Engine, args []interface{}) interface{} {
	return engine.Eval(args[0]) == engine.Eval(args[1])
}

type SetOperator struct{}

func (SetOperator) call(engine *Engine, args []interface{}) interface{} {
	value := engine.Eval(args[1])
	engine.variables[engine.Eval(args[0])] = value
	return value
}

type GetOperator struct{}

func (GetOperator) call(engine *Engine, args []interface{}) interface{} {
	return engine.variables[engine.Eval(args[0])]
}

type UntilOperator struct{}

func (UntilOperator) call(engine *Engine, args []interface{}) interface{} {

	var v interface{}
	for {
		b, ok := engine.Eval(args[0]).(bool)
		if !ok {
			log.Fatalf("cannot convert bool value, %v (%T)", b, b)
		}
		if b {
			break
		}
		v = engine.Eval(args[1])
	}
	return v
}

type StepOperator struct{}

func (StepOperator) call(engine *Engine, args []interface{}) interface{} {
	var v interface{}
	for _, arg := range args {
		v = engine.Eval(arg)
	}
	return v
}
