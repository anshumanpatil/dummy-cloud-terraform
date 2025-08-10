package main

type InfrastructureTypes struct{}

var typesOfInfrastructure []string = []string{"compute", "storage", "network"}

func (i InfrastructureTypes) checkInfraTypes(infraType string) bool {
	for _, t := range typesOfInfrastructure {
		if t == infraType {
			return true
		}
	}
	return false
}
