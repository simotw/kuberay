package support

import (
	"os"
	"runtime"
	"strings"
)

const (
	// KuberayTestOutputDir is the testing output directory, to write output files into.
	KuberayTestOutputDir = "KUBERAY_TEST_OUTPUT_DIR"

	// KuberayTestRayVersion is the version of Ray to use for testing.
	KuberayTestRayVersion = "KUBERAY_TEST_RAY_VERSION"

	// KuberayTestRayImage is the Ray image to use for testing.
	KuberayTestRayImage = "KUBERAY_TEST_RAY_IMAGE"
)

func GetRayVersion() string {
	return lookupEnvOrDefault(KuberayTestRayVersion, RayVersion)
}

func GetRayImage() string {
	return lookupEnvOrDefault(KuberayTestRayImage, RayImage)
}

func lookupEnvOrDefault(key, value string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return value
}

func GetRayVersionForAutoScalerV2() string {
	rayVersion := os.Getenv("E2E_AUTOSCALER_V2_TEST_RAY_VERSION")
	if strings.TrimSpace(rayVersion) == "" {
		rayVersion = "2.10.0"
	}
	return rayVersion
}

func GetRayImageForAutoScalerV2() string {
	rayImage := os.Getenv("E2E_AUTOSCALER_V2_TEST_RAY_IMAGE")
	if strings.TrimSpace(rayImage) == "" {
		rayImage = "rayproject/ray:2.10.0.ee9422-py310"
	}
	// detect if we are running on arm64 machine, most likely apple silicon
	// the os name is not checked as it also possible that it might be linux
	// also check if the image does not have the `-aarch64` suffix
	if runtime.GOARCH == "arm64" && !strings.HasSuffix(rayImage, "-aarch64") {
		rayImage = rayImage + "-aarch64"
	}
	return rayImage
}
