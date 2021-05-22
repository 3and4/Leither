GetVar和Act(草)
========  
查询系统变量和状态的操作都放入了GetVar  
待定的Api都放入了Act  

## 一、GetVar  
GetVar(sid, name string, args ...string) (interface{}, error)  
GetVarByContext(sid, name string, mapOpt map[string]string)  (interface{}, error)  
GetGobVar(sid, name string, args ...string) ([]byte, error)  

## 二、Act  
Act(sid, name string, args ...string) error
  
  
