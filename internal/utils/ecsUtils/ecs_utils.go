package ecsUtils

import (
	"errors"
	"github.com/cirno42/sbercloud-api/api/ecs"
)

func GetEcsId(projectID, id, name string) (string, error) {
	if id != "" {
		return id, nil
	} else {
		servers, err := ecs.GetECSList(projectID, 0, 1000, "")
		if err != nil {
			return "", err
		}
		for _, server := range servers {
			if server.Name == name {
				return server.ID, nil
			}
		}
	}
	return "", errors.New("{\"error\" : \"No ECS with specified name\"}")
}
