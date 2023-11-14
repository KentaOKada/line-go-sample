package config

// 環境
type Environment string

const (
	EnvironmentLocal Environment = "local"
	EnvironmentTest  Environment = "test" // ローカルテスト用の環境
	EnvironmentDev   Environment = "development"
	EnvironmentQa    Environment = "qa"
	EnvironmentProd  Environment = "production"
)

func (e Environment) String() string {
	return string(e)
}

func (e Environment) IsLocal() bool {
	return e == EnvironmentLocal
}

func (e Environment) IsTest() bool {
	return e == EnvironmentTest
}

func (e Environment) IsDev() bool {
	return e == EnvironmentDev
}

func (e Environment) IsQa() bool {
	return e == EnvironmentQa
}

func (e Environment) IsProd() bool {
	return e == EnvironmentProd
}
