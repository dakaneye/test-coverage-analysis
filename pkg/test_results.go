package pkg

type Results struct {
	TestEnvs []TestEnv `json:"testenvs"`
}

type TestEnv struct {
	Py38 Py38Env `json:"py38"`
}

type Py38Env struct {
	Test []Test `json:"test"`
}

type Test struct {
	Command []string `json:"command"`
	Output string `json:"output"`
}

