package config

import(
	"github.com/0xor1/json"
	"os"
	"strings"
	"time"
)

type config struct{
	defaults              *json.PJson
	fileValues            *json.PJson
	envVarsStringSeparator string
}

//create a new config object based on a env vars and/or a json file and/or programmatic defaults
//pass in empty file path to not use a config file, pass in an empty envVarSeparator to ignore environment variables
func New(file string, envVarSeparator string) *config {
	ret := &config{
		defaults: json.PNew(),
		fileValues: json.PNew(),
		envVarsStringSeparator: envVarSeparator,
	}
	if file != "" {
		ret.fileValues = json.PFromFile(file)
	}
	return ret
}

func (c *config) SetDefault(path string, val interface{}) {
	jsonPath, _, _ := makeJsonPathAndGetEnvValAndExists(path, "")
	c.defaults.Set(val, jsonPath...)
}

func (c *config) GetString(path string) string {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return envVal
	} else {
		if val, err := c.fileValues.Json.String(jsonPath...); err == nil {
			return val
		}
		return c.defaults.String(jsonPath...)
	}
}

func (c *config) GetStringSlice(path string) []string {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return json.PFromString(envVal).StringSlice()
	} else {
		if val, err := c.fileValues.Json.StringSlice(jsonPath...); err == nil {
			return val
		}
		return c.defaults.StringSlice(jsonPath...)
	}
}

func (c *config) GetMap(path string) map[string]interface{} {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return json.PFromString(envVal).Map()
	} else {
		if val, err := c.fileValues.Json.Map(jsonPath...); err == nil {
			return val
		}
		return c.defaults.Map(jsonPath...)
	}
}

func (c *config) GetStringMap(path string) map[string]string {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return json.PFromString(envVal).MapString()
	} else {
		if val, err := c.fileValues.Json.MapString(jsonPath...); err == nil {
			return val
		}
		return c.defaults.MapString(jsonPath...)
	}
}

func (c *config) GetInt(path string) int {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return json.PFromString(envVal).Int()
	} else {
		if val, err := c.fileValues.Json.Int(jsonPath...); err == nil {
			return val
		}
		return c.defaults.Int(jsonPath...)
	}
}

func (c *config) GetInt64(path string) int64 {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return json.PFromString(envVal).Int64()
	} else {
		if val, err := c.fileValues.Json.Int64(jsonPath...); err == nil {
			return val
		}
		return c.defaults.Int64(jsonPath...)
	}
}

func (c *config) GetBool(path string) bool {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return json.PFromString(envVal).Bool()
	} else {
		if val, err := c.fileValues.Json.Bool(jsonPath...); err == nil {
			return val
		}
		return c.defaults.Bool(jsonPath...)
	}
}

func (c *config) GetTime(path string) time.Time {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return json.PFromString(envVal).Time()
	} else {
		if val, err := c.fileValues.Json.Time(jsonPath...); err == nil {
			return val
		}
		return c.defaults.Time(jsonPath...)
	}
}

func (c *config) GetDuration(path string) time.Duration {
	if jsonPath, envVal, envValExists := makeJsonPathAndGetEnvValAndExists(path, c.envVarsStringSeparator); envValExists {
		return json.PFromString(envVal).Duration()
	} else {
		if val, err := c.fileValues.Json.Duration(jsonPath...); err == nil {
			return val
		}
		return c.defaults.Duration(jsonPath...)
	}
}

func makeJsonPathAndGetEnvValAndExists(path, envVarsStringSeparator string) ([]interface{}, string, bool) {
	parts := strings.Split(path, ".")
	envName := ""
	if envVarsStringSeparator != "" {
		envName = strings.ToUpper(strings.Join(parts, envVarsStringSeparator))
	}
	if envName != "" {
		if val, exists := os.LookupEnv(envName); exists {
			return nil, val, true
		}
	}

	is := make([]interface{}, 0, len(parts))
	for _, str := range parts {
		is = append(is, str)
	}
	return is, "", false
}