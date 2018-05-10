package devicemodel

// DeviceInfo type
type DeviceInfo struct {
	Model               string `json:"model"`
	Manufacturer        string `json:"manufacturer"`
	Brand               string `json:"brand"`
	AndroidVersion      string `json:"androidVersion"`
	APILevel            string `json:"apiLevel"`
	BuildNumber         string `json:"buildNumber"`
	AndroidDeviceID     string `json:"androidDeviceId"`
	HardwareSerialNo    string `json:"hardwareSerialNo"`
	InstructionSets     string `json:"instructionSets"`
	CPUHardware         string `json:"cpuHardware"`
	DisplayResolution   string `json:"displayResolution"`
	DisplayDensity      string `json:"displayDensity"`
	DisplayPhysicalSize string `json:"displayPhysicalSize"`
}
