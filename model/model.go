package model

type Val = string
type Type = string

type App struct {
	TypeInstances []Type // 型パラメーターの値
	Args          []Val  // 引数にわたす値
}

type GenFunc struct {
	FName string
	Apps  []*App
}

type TemplData struct {
	PkgPath  string
	PkgName  string
	GenFuncs []*GenFunc
}
