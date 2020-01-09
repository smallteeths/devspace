package examples

import (
	"bytes"

	"github.com/devspace-cloud/devspace/pkg/util/log"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// RunDependencies runs the test for the dependencies example
func RunDependencies(f *customFactory, logger log.Logger) error {
	buff := &bytes.Buffer{}
	f.cacheLogger = log.NewStreamLogger(buff, logrus.InfoLevel)

	var buffString string
	buffString = buff.String()

	if f.verbose {
		f.cacheLogger = logger
		buffString = ""
	}

	logger.Info("Run sub test 'dependencies' of test 'examples'")
	logger.StartWait("Run test...")
	defer logger.StopWait()

	err := beforeTest(f, "../examples/dependencies")
	defer afterTest(f)
	if err != nil {
		return errors.Errorf("sub test 'dependencies' of 'examples' test failed: %s %v", buffString, err)
	}

	err = RunTest(f, nil)
	if err != nil {
		return errors.Errorf("sub test 'dependencies' of 'examples' test failed: %s %v", buffString, err)
	}

	return nil
}