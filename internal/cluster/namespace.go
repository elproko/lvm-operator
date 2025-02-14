package cluster

import (
	"fmt"
	"os"
)

const OperatorNamespaceEnvVar = "POD_NAMESPACE"

// GetOperatorNamespace returns the Namespace the operator should be watching for changes
func GetOperatorNamespace() (string, error) {
	// The env variable POD_NAMESPACE which specifies the Namespace the pod is running in
	// and hence will watch.

	ns, found := os.LookupEnv(OperatorNamespaceEnvVar)
	if !found {
		return "", fmt.Errorf("%s not found", OperatorNamespaceEnvVar)
	}
	return ns, nil
}
