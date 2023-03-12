package api

import "os"

func InitApi(projectID, accessKey, secretKey, region string) {
	_ = os.Setenv("PROJECT_ID", projectID)
	_ = os.Setenv("ACCESS_KEY", accessKey)
	_ = os.Setenv("SECRET_KEY", secretKey)
	_ = os.Setenv("REGION", region)
}
