package ecs

import (
	"errors"
	"fmt"
	"github.com/cirno42/sbercloud-api/api/endpoints"
	"github.com/cirno42/sbercloud-api/api/models/ecsModels"
	"github.com/cirno42/sbercloud-api/internal/handlers/requestMakers"
	"sort"
	"strconv"
)

type getECSFlavorListResponse struct {
	Flavors []ecsModels.FlavorModel `json:"flavors"`
}

func GetESCFlavorList(projectID, availabilityZone string) ([]ecsModels.FlavorModel, error) {
	endpoint := fmt.Sprintf(endpoints.GetEndpointAddress(endpoints.EscEndpoint)+"/v1/%s/cloudservers/flavors?availability_zone=%s", projectID, availabilityZone)
	var flavorsArray getECSFlavorListResponse
	err := requestMakers.CreateAndDoRequest(endpoint, requestMakers.HTTP_METHOD_GET, nil, &flavorsArray, nil)
	return flavorsArray.Flavors, err
}

func GetFlavorListBySpec(projectID, generation, flavorType, availabilityZone string, ram, vcpus int) ([]ecsModels.FlavorModel, error) {
	flavors, err := GetESCFlavorList(projectID, availabilityZone)
	if err != nil {
		return nil, err
	}
	filteredFlavors := make([]ecsModels.FlavorModel, 0)
	for i := 0; i < len(flavors); i++ {
		flavorVcpu, err := strconv.Atoi(flavors[i].Vcpus)
		if err != nil {
			panic(err)
		}
		if (flavors[i].Ram >= ram) && (flavorVcpu >= vcpus) && (flavors[i].OsExtraSpecs.EcsPerformancetype == flavorType) {
			if (flavors[i].OsExtraSpecs.EcsGeneration == generation) || (generation == "") {
				filteredFlavors = append(filteredFlavors, flavors[i])
			}
		}
	}
	sort.Slice(filteredFlavors, func(i, j int) bool {
		flavorVcpuI, err := strconv.Atoi(filteredFlavors[i].Vcpus)
		if err != nil {
			panic(err)
		}
		flavorVcpuJ, err := strconv.Atoi(filteredFlavors[j].Vcpus)
		if err != nil {
			panic(err)
		}
		return flavorVcpuI < flavorVcpuJ
	})

	sort.Slice(filteredFlavors, func(i, j int) bool { return filteredFlavors[i].Ram < filteredFlavors[j].Ram })

	return filteredFlavors, err
}

func GetMinimumFlavorBySpec(projectID, generation, flavorType, availabilityZone string, ram, vcpus int) (*ecsModels.FlavorModel, error) {
	flavors, err := GetFlavorListBySpec(projectID, generation, flavorType, availabilityZone, ram, vcpus)
	if err != nil {
		return nil, err
	}
	if len(flavors) == 0 {
		return nil, errors.New("{\"error\" : \"No flavor for specified parameters\"}")
	}
	return &flavors[0], err
}
