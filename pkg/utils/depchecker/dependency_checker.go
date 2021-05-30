package depchecker

var dependencies map[string]Dependency

//RegisterDependency function take dependecy as parameter and register it
func RegisterDependency(dep Dependency) {
	if dependencies == nil {
		dependencies = make(map[string]Dependency)
	}
	dependencies[dep.GetName()] = dep
}

func getDependencies() map[string]Dependency {
	return dependencies
}

//Dependency represent depedency check interface
type Dependency interface {
	GetPinger() func() (bool, error)
	GetName() string
}

//CheckDependencies run call back method attach while register depedency
func CheckDependencies() map[string]string {
	result := make(map[string]string)
	for _, dep := range getDependencies() {
		result[dep.GetName()] = "OK"
		if isOK, err := dep.GetPinger()(); !isOK {
			result[dep.GetName()] = err.Error()
		}
	}
	if len(result) == 0 {
		//no dependencies registered
		result["NODEP"] = "No Dependencies Registered"
	}
	return result
}
