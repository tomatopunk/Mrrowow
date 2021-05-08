package executor

import (
	"os"
	"testing"
)

func createTestExecContext() *AliExecutor {
	exec := &AliExecutor{
		AccessKeySecret: os.Getenv("ali_access_secret"),
		AccessKey: os.Getenv("ali_access_key"),
		Scheme:"https",
	}
	return exec
}
func TestCreateClient(t *testing.T) {
	exec := createTestExecContext()
	_,err := exec.CreateClient("cn-hangzhou")
	if err != nil {
		t.Errorf("%s",err)
	}
}

func TestGetRegions(t *testing.T){
	exec := createTestExecContext()
	res,err := exec.DescribeRegions("instance","PostPaid")
	if err != nil {
		t.Errorf("%s",err)
	}

	_=res
	//t.Log()
}