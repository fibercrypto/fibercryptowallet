package integration

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/fibercrypto/skywallet-go/src/skywallet"

	"github.com/gogo/protobuf/proto"
	messages "github.com/fibercrypto/skywallet-protob/go"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/stretchr/testify/require"
)

var (
	log = logging.MustGetLogger("device-interface-tests")

	binaryPath string
)

const (
	binaryName = "hwgo-cli.test"

	testModeEmulator = "EMULATOR"
	testModeUSB      = "USB"
	defaultSeed      = "cloud flower upset remain green metal below cup stem infant art thank"
)

func execCommand(args ...string) *exec.Cmd {
	args = append(args)
	cmd := exec.Command(binaryPath, args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "AUTO_PRESS_BUTTONS="+autopressButtons())
	return cmd
}

func execCommandCombinedOutput(args ...string) ([]byte, error) {
	cmd := execCommand(args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}
	return output, nil
}

func mode(t *testing.T) string {
	mode := os.Getenv("HW_GO_INTEGRATION_TEST_MODE")
	switch mode {
	case "":
		mode = testModeEmulator
	case testModeUSB, testModeEmulator:
	default:
		t.Fatalf("Invalid test mode %s, must be emulator or wallet", mode)
	}
	return mode
}

func enabled() bool {
	return os.Getenv("HW_GO_INTEGRATION_TESTS") == "1"
}

func autopressButtons() string {
	return os.Getenv("AUTO_PRESS_BUTTONS")
}

func TestMain(m *testing.M) {
	if !enabled() {
		return
	}

	err := os.Setenv("DEVICE_TYPE", os.Getenv("HW_GO_INTEGRATION_TEST_MODE"))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	abs, err := filepath.Abs(binaryName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get binary name absolute path failed: %v\n", err)
		os.Exit(1)
	}

	binaryPath = abs

	// Build cli binary file.
	args := []string{"build", "-ldflags", "-X main.AUTO_PRESS_BUTTONS=true", "-o", binaryPath, "../../../cmd/cli/cli.go"}
	if err := exec.Command("go", args...).Run(); err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Make %v binary failed: %v\n", binaryName, err))
		os.Exit(1)
	}

	ret := m.Run()

	// Remove the generated cli binary file.
	if err := os.Remove(binaryPath); err != nil {
		fmt.Fprintf(os.Stderr, "Delete %v failed: %v", binaryName, err)
		os.Exit(1)
	}

	os.Exit(ret)
}

func doWalletOrEmulator(t *testing.T) bool {
	if enabled() {
		switch mode(t) {
		case testModeUSB, testModeEmulator:
			return true
		}
	}

	t.Skip("usb and emulator tests disabled")
	return false
}

func doWallet(t *testing.T) bool {
	if enabled() {
		switch mode(t) {
		case testModeUSB:
			return true
		}
	}

	t.Skip("usb tests disabled")
	return false
}

func doEmulator(t *testing.T) bool {
	if enabled() {
		switch mode(t) {
		case testModeEmulator:
			return true
		}
	}

	t.Skip("emulator tests disabled")
	return false
}

func bootstrap(t *testing.T, testname string, testType string) *skywallet.Device {
	switch testType {
	case "USB":
		if !doWallet(t) {
			return nil
		}
	case "EMULATOR":
		if !doEmulator(t) {
			return nil
		}
	default:
		if !doWalletOrEmulator(t) {
			return nil
		}
	}

	device := skywallet.NewDevice(skywallet.DeviceTypeFromString(mode(t)))
	require.NotNil(t, device)

	err := device.Connect()
	require.NoError(t, err)

	if !device.Connected() {
		t.Skip(fmt.Sprintf("%s does not work if device is not connected", testname))
		return nil
	}
	require.NoError(t, device.Disconnect())

	if device.Driver.DeviceType() == skywallet.DeviceTypeEmulator && runtime.GOOS == "linux" {
		err := device.SetAutoPressButton(true, skywallet.ButtonRight)
		require.NoError(t, err)
	}

	// bootstrap

	// get features to check if bootstrap needs to be done
	msg, err := device.GetFeatures()
	require.NoError(t, err)
	require.Equal(t, msg.Kind, uint16(messages.MessageType_MessageType_Features))
	features := &messages.Features{}
	err = proto.Unmarshal(msg.Data, features)
	require.NoError(t, err)

	if *features.Initialized == false || *features.NeedsBackup == false {
		msg, err = device.Wipe()
		require.NoError(t, err)
		_, err = device.ButtonAck()
		require.NoError(t, err)

		_, err = device.SetMnemonic(defaultSeed)
		require.NoError(t, err)
		_, err = device.ButtonAck()
		require.NoError(t, err)
	}

	return device
}

func TestAddressGen(t *testing.T) {
	device := bootstrap(t, "TestAddressGen", "")
	if device == nil {
		return
	}

	tt := []struct {
		name         string
		args         []string
		expectOutput string
	}{
		{
			name:         "addressGen -n 2",
			args:         []string{"addressGen", "-addressN", "2"},
			expectOutput: "[2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw zC8GAQGQBfwk7vtTxVoRG7iMperHNuyYPs]",
		},

		{
			name:         "addressGen -n 2 -s 2",
			args:         []string{"addressGen", "-addressN", "2", "--startIndex", "2"},
			expectOutput: "[28L2fexvThTVz6e2dWUV4pSuCP8SAnCUVku 2NckPkQRQFa5E7HtqDkZmV1TH4HCzR2N5J6]",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output, err := execCommandCombinedOutput(tc.args...)
			if err != nil {
				require.EqualError(t, err, "exit status 1")
			}

			require.Contains(t, string(output), tc.expectOutput)
		})
	}
}

func TestApplySettings(t *testing.T) {
	device := bootstrap(t, "TestApplySettings", "")
	if device == nil {
		return
	}

	tt := []struct {
		name  string
		label string
		args  []string
	}{
		{
			name:  "applySettings",
			label: "my custom device label",
			args:  []string{"applySettings", "-usePassphrase", "false", "-label", "my custom device label"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output, err := execCommandCombinedOutput(tc.args...)
			if err != nil {
				require.EqualError(t, err, "exit status 1")
			}

			require.Contains(t, string(output), "Settings applied")

			output, err = execCommandCombinedOutput([]string{"features"}...)
			if err != nil {
				require.Equal(t, err, "exit status 1")
			}

			decoder := json.NewDecoder(bytes.NewReader(output))
			decoder.DisallowUnknownFields()

			var features messages.Features
			require.NoError(t, decoder.Decode(&features))
			require.Equal(t, *features.Label, tc.label)
		})
	}
}

func TestApplySettingsLabelShouldNotBeReset(t *testing.T) {
	device := bootstrap(t, "TestApplySettingsLabelShouldNotBeReset", "")
	if device == nil {
		return
	}

	label := "custom label"
	output, err := execCommandCombinedOutput([]string{"applySettings", "-usePassphrase", "false", "-label", label}...)
	if err != nil {
		require.EqualError(t, err, "exit status 1")
	}

	require.Contains(t, string(output), "Settings applied")

	output, err = execCommandCombinedOutput([]string{"features"}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	decoder := json.NewDecoder(bytes.NewReader(output))
	decoder.DisallowUnknownFields()

	var features messages.Features
	require.NoError(t, decoder.Decode(&features))
	require.Equal(t, *features.Label, label)

	output, err = execCommandCombinedOutput([]string{"applySettings", "-usePassphrase", "false", "-label", ""}...)
	if err != nil {
		require.EqualError(t, err, "exit status 1")
	}

	require.Contains(t, string(output), "Settings applied")

	output, err = execCommandCombinedOutput([]string{"features"}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	decoder = json.NewDecoder(bytes.NewReader(output))
	decoder.DisallowUnknownFields()

	var featuresNotChange messages.Features
	require.NoError(t, decoder.Decode(&featuresNotChange))
	require.Equal(t, *featuresNotChange.Label, label)
}

func TestBackup(t *testing.T) {
	device := bootstrap(t, "TestBackup", "")
	if device == nil {
		return
	}

	output, err := execCommandCombinedOutput([]string{"backup"}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	require.Contains(t, string(output), "Device backed up!")
}

func TestCheckMessageSignature(t *testing.T) {
	device := bootstrap(t, "TestCheckMessageSignature", "")
	if device == nil {
		return
	}

	tt := []struct {
		name    string
		address string
		args    []string
	}{
		{
			name:    "checkMessageSignature -n 2",
			address: "2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw",
			args: []string{"checkMessageSignature",
				"--address", "2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw", "--message", "Hello World", "--signature", "f026001ee2d4e6bf4dfdbbdd33b6622bae24c8f6232fc19e5dd0013f2b68f0f14606d448897df20348c8b63dcad6923e9fecb2fe0969b183dc31eb4796e001d101"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output, err := execCommandCombinedOutput(tc.args...)
			if err != nil {
				require.EqualError(t, err, "exit status 1")
			}

			require.Contains(t, string(output), tc.address)
		})
	}
}

func TestFeatures(t *testing.T) {
	device := bootstrap(t, "TestFeatures", "")
	if device == nil {
		return
	}

	output, err := execCommandCombinedOutput([]string{"features"}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	decoder := json.NewDecoder(bytes.NewReader(output))
	decoder.DisallowUnknownFields()

	var features messages.Features
	err = decoder.Decode(&features)
	require.NoError(t, err)

	ff := skywallet.NewFirmwareFeatures(uint64(*features.FirmwareFeatures))
	err = ff.Unmarshal()
	require.NoError(t, err)
	require.False(t, ff.HasRdpMemProtectEnabled())
}

func TestGenerateMnemonic(t *testing.T) {
	if !doWalletOrEmulator(t) {
		return
	}

	device := skywallet.NewDevice(skywallet.DeviceTypeFromString(mode(t)))
	require.NotNil(t, device)

	err := device.Connect()
	require.NoError(t, err)

	if !device.Connected() {
		t.Skip(fmt.Sprintf("%s does not work if device is not connected", "TestGenerateMnemonic"))
		return
	}
	require.NoError(t, device.Disconnect())

	if device.Driver.DeviceType() == skywallet.DeviceTypeEmulator && runtime.GOOS == "linux" {
		err := device.SetAutoPressButton(true, skywallet.ButtonRight)
		require.NoError(t, err)
	}

	tt := []struct {
		name           string
		args           []string
		isUsageError   bool
		expectedOutput string
	}{
		{
			name:           "generateMnemonic -wc 12",
			args:           []string{"generateMnemonic", "--wordCount", "12"},
			expectedOutput: "Mnemonic successfully configured",
		},

		{
			name:           "generateMnemonic -wc 24",
			args:           []string{"generateMnemonic", "--wordCount", "24"},
			expectedOutput: "Mnemonic successfully configured",
		},

		{
			name:           "generateMnemonic invalid word count",
			args:           []string{"generateMnemonic", "--wordCount", "15"},
			expectedOutput: "word count must be 12 or 24",
			isUsageError:   true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// bootstrap
			_, err = device.Wipe()
			require.NoError(t, err)
			_, err = device.ButtonAck()
			require.NoError(t, err)

			output, err := execCommandCombinedOutput(tc.args...)
			if err != nil {
				require.Equal(t, err, "exit status 1")
			}

			require.Contains(t, string(output), tc.expectedOutput)

			if !tc.isUsageError {
				output, err = execCommandCombinedOutput([]string{"features"}...)
				if err != nil {
					require.Equal(t, err, "exit status 1")
				}

				decoder := json.NewDecoder(bytes.NewReader(output))
				decoder.DisallowUnknownFields()

				var features messages.Features
				require.NoError(t, decoder.Decode(&features))
				require.True(t, *features.NeedsBackup)
				require.False(t, *features.PinProtection)
			}
		})
	}
}

func TestRecovery(t *testing.T) {
	// This test checks for failure on invalid input
	// It is not possible to programmatically check for valid input

	if !doWalletOrEmulator(t) {
		return
	}

	device := skywallet.NewDevice(skywallet.DeviceTypeFromString(mode(t)))
	require.NotNil(t, device)

	err := device.Connect()
	require.NoError(t, err)

	if !device.Connected() {
		t.Skip(fmt.Sprintf("%s does not work if device is not connected", "TestRecovery"))
		return
	}
	require.NoError(t, device.Disconnect())

	if device.Driver.DeviceType() == skywallet.DeviceTypeEmulator && runtime.GOOS == "linux" {
		err := device.SetAutoPressButton(true, skywallet.ButtonRight)
		require.NoError(t, err)
	}

	// bootstrap
	_, err = device.Wipe()
	require.NoError(t, err)
	_, err = device.ButtonAck()
	require.NoError(t, err)

	cmd := execCommand([]string{"recovery"}...)

	stdoutPipe, err := cmd.StdoutPipe()
	require.NoError(t, err)
	stderrPipe, err := cmd.StderrPipe()
	require.NoError(t, err)
	stdInPipe, err := cmd.StdinPipe()
	require.NoError(t, err)
	require.NoError(t, cmd.Start())

	var fail = false
	var stdInDone = false

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)

		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			m := scanner.Text()
			if m == "Word:" {
				time.Sleep(1 * time.Second)
				_, err := stdInPipe.Write([]byte("foobar\n"))
				require.NoError(t, err)

				stdInDone = true
			} else if stdInDone {
				if m == "Wrong" || m == "Word" {
					fail = true
					break
				}
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			log.Errorln(scanner.Text())
		}
	}()

	err = cmd.Wait()
	require.NoError(t, err)
	require.True(t, fail)
}

func TestSetMnemonic(t *testing.T) {
	if !doWalletOrEmulator(t) {
		return
	}

	device := skywallet.NewDevice(skywallet.DeviceTypeFromString(mode(t)))
	require.NotNil(t, device)

	err := device.Connect()
	require.NoError(t, err)

	if !device.Connected() {
		t.Skip(fmt.Sprintf("%s does not work if device is not connected", "TestGenerateMnemonic"))
		return
	}
	require.NoError(t, device.Disconnect())

	if device.Driver.DeviceType() == skywallet.DeviceTypeEmulator && runtime.GOOS == "linux" {
		err := device.SetAutoPressButton(true, skywallet.ButtonRight)
		require.NoError(t, err)
	}

	tt := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			name: "setMnemonic 12",
			args: []string{"setMnemonic", "--mnemonic",
				"cloud flower upset remain green metal below cup stem infant art thank"},
			expectedOutput: "cloud flower upset remain green metal below cup stem infant art thank",
		},

		{
			name: "setMnemonic 24",
			args: []string{"setMnemonic", "--mnemonic",
				"dress fee animal silly multiply demand casino gold pipe matrix latin badge umbrella orbit safe cover glove one dash chicken play obey employ post"},
			expectedOutput: "dress fee animal silly multiply demand casino gold pipe matrix latin badge umbrella orbit safe cover glove one dash chicken play obey employ post",
		},

		{
			name: "setMnemonic invalid mnemonic length",
			args: []string{"setMnemonic", "--mnemonic",
				"dress fee animal silly multiply demand casino gold pipe matrix latin badge umbrella orbit safe cover glove one dash chicken play obey"},
			expectedOutput: "Mnemonic with wrong checksum provided",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// bootstrap
			_, err = device.Wipe()
			require.NoError(t, err)
			_, err = device.ButtonAck()
			require.NoError(t, err)

			output, err := execCommandCombinedOutput(tc.args...)
			if err != nil {
				require.Equal(t, err, "exit status 1")
			}

			require.Contains(t, string(output), tc.expectedOutput)
		})
	}

}

func TestSetPinCode(t *testing.T) {
	// This test checks for failure on invalid input
	// It is not possible to programmatically check for valid input

	device := bootstrap(t, "TestSetPinCode", "")
	if device == nil {
		return
	}

	cmd := execCommand([]string{"setPinCode"}...)

	stdoutPipe, err := cmd.StdoutPipe()
	require.NoError(t, err)
	stderrPipe, err := cmd.StderrPipe()
	require.NoError(t, err)
	stdInPipe, err := cmd.StdinPipe()
	require.NoError(t, err)
	require.NoError(t, cmd.Start())

	var fail = false
	var stdInDone = false

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)

		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			m := scanner.Text()
			if m == "response:" {
				time.Sleep(1 * time.Second)
				_, err := stdInPipe.Write([]byte("123\n"))
				require.NoError(t, err)

				stdInDone = true
			} else if stdInDone {
				if m == "mismatch" {
					fail = true
					break
				}
			}
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			log.Errorln(scanner.Text())
		}
	}()

	err = cmd.Wait()
	require.NoError(t, err)
	require.True(t, fail)
}

func TestRemovePinCode(t *testing.T) {
	device := bootstrap(t, "TestRemovePinCode", "")
	if device == nil {
		return
	}

	output, err := execCommandCombinedOutput([]string{"removePinCode"}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	require.Contains(t, string(output), "PIN removed")
}

func TestSignMessage(t *testing.T) {
	device := bootstrap(t, "TestTransactionSign", "")
	if device == nil {
		return
	}

	output, err := execCommandCombinedOutput([]string{"signMessage", "--message", "Hello World"}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	outputStr := strings.TrimSpace(string(bytes.Replace(output, []byte("PASS"), []byte{}, 1)))

	checkOutput, err := execCommandCombinedOutput([]string{"checkMessageSignature", "--address", "2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw", "--message", "Hello World", "--signature", outputStr}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	require.Contains(t, string(checkOutput), "2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw")
}

func TestTransactionSign(t *testing.T) {
	device := bootstrap(t, "TestTransactionSign", "")
	if device == nil {
		return
	}

	tt := []struct {
		name    string
		args    []string
		message []string
		address []string
	}{
		{
			name: "transactionSign sample 1",
			args: []string{"transactionSign", "--inputHash", "181bd5656115172fe81451fae4fb56498a97744d89702e73da75ba91ed5200f9",
				"--inputIndex", "0", "--outputAddress=K9TzLrgqz7uXn3QJHGxmzdRByAzH33J2ot", "--coin", "100000", "--hour", "2"},
			message: []string{"d11c62b1e0e9abf629b1f5f4699cef9fbc504b45ceedf0047ead686979498218"},
			address: []string{"2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw"},
		},

		{
			name: "transactionSign sample 2",
			args: []string{"transactionSign", "--inputHash", "01a9ef6c25271229ef9760e1536c3dc5ccf0ead7de93a64c12a01340670d87e9",
				"--inputHash", "8c2c97bfd34e0f0f9833b789ce03c2e80ac0b94b9d0b99cee6ea76fb662e8e1c", "--inputIndex", "0", "--inputIndex", "0",
				"--outputAddress=K9TzLrgqz7uXn3QJHGxmzdRByAzH33J2ot", "--coin", "20800000", "--hour", "255"},
			message: []string{"9bbde062d665a8b11ae15aee6d4f32f0f3d61af55160c142060795a219378a54", "f947b0352b19672f7b7d04dc2f1fdc47bc5355878f3c47a43d4d4cfbae07d026"},
			address: []string{"2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw", "2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw"},
		},

		{
			name: "transactionSign sample 3",
			args: []string{"transactionSign", "--inputHash", "da3b5e29250289ad78dc42dcf007ab8f61126198e71e8306ff8c11696a0c40f7", "--inputIndex", "0",
				"--inputHash", "33e826d62489932905dd936d3edbb74f37211d68d4657689ed4b8027edcad0fb", "--inputIndex", "0",
				"--inputHash", "668f4c144ad2a4458eaef89a38f10e5307b4f0e8fce2ade96fb2cc2409fa6592", "--inputIndex", "0",
				"--outputAddress=K9TzLrgqz7uXn3QJHGxmzdRByAzH33J2ot", "--coin", "111000000", "--hour", "6464556",
				"--outputAddress=2iNNt6fm9LszSWe51693BeyNUKX34pPaLx8", "--coin", "1900000", "--hour", "1"},
			message: []string{"ff383c647551a3ba0387f8334b3f397e45f9fc7b3b5c3b18ab9f2b9737bce039",
				"c918d83d8d3b1ee85c1d2af6885a0067bacc636d2ebb77655150f86e80bf4417",
				"0e827c5d16bab0c3451850cc6deeaa332cbcb88322deea4ea939424b072e9b97"},
			address: []string{"2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw", "2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw", "2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw"},
		},

		{
			name: "transactionSign sample 4",
			args: []string{"transactionSign", "--inputHash", "b99f62c5b42aec6be97f2ca74bb1a846be9248e8e19771943c501e0b48a43d82", "--inputIndex", "0",
				"--inputHash", "cd13f705d9c1ce4ac602e4c4347e986deab8e742eae8996b34c429874799ebb2", "--inputIndex", "0",
				"--outputAddress=22S8njPeKUNJBijQjNCzaasXVyf22rWv7gF", "--coin", "23100000", "--hour", "0"},
			message: []string{"42a26380399172f2024067a17704fceda607283a0f17cb0024ab7a96fc6e4ac6",
				"5e0a5a8c7ea4a2a500c24e3a4bfd83ef9f74f3c2ff4bdc01240b66a41e34ebbf"},
			address: []string{"2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw", "2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw"},
		},

		{
			name: "transactionSign sample 5",
			args: []string{"transactionSign", "--inputHash", "4c12fdd28bd580989892b0518f51de3add96b5efb0f54f0cd6115054c682e1f1", "--inputIndex", "0",
				"--outputAddress=2iNNt6fm9LszSWe51693BeyNUKX34pPaLx8", "--coin", "1000000", "--hour", "0"},
			message: []string{"c40e110f5e460532bfb03a5a0e50262d92d8913a89c87869adb5a443463dea69"},
			address: []string{"2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw"},
		},

		{
			name: "transactionSign sample 6",
			args: []string{"transactionSign", "--inputHash", "c5467f398fc3b9d7255d417d9ca208c0a1dfa0ee573974a5fdeb654e1735fc59", "--inputIndex", "0",
				"--outputAddress=K9TzLrgqz7uXn3QJHGxmzdRByAzH33J2ot", "--coin", "10000000", "--hour", "1",
				"--outputAddress=VNz8LR9JTSoz5o7qPHm3QHj4EiJB6LV18L", "--coin", "5500000", "--hour", "0",
				"--outputAddress=22S8njPeKUNJBijQjNCzaasXVyf22rWv7gF", "--coin", "4500000", "--hour", "1"},
			message: []string{"7edea77354eca0999b1b023014eb04638b05313d40711707dd03a9935696ccd1"},
			address: []string{"2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw"},
		},

		{
			name: "transactionSign sample 7",
			args: []string{"transactionSign", "--inputHash", "ae6fcae589898d6003362aaf39c56852f65369d55bf0f2f672bcc268c15a32da", "--inputIndex", "0",
				"--outputAddress=3pXt9MSQJkwgPXLNePLQkjKq8tsRnFZGQA", "--coin", "1000000", "--hour", "1000"},
			message: []string{"47bfa37c79f7960df8e8a421250922c5165167f4c91ecca5682c1106f9010a7f"},
			address: []string{"2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw"},
		},

		{
			name: "transactionSign sample 8",
			args: []string{"transactionSign", "--inputHash", "ae6fcae589898d6003362aaf39c56852f65369d55bf0f2f672bcc268c15a32da", "--inputIndex", "0",
				"--outputAddress=3pXt9MSQJkwgPXLNePLQkjKq8tsRnFZGQA", "--coin", "300000", "--hour", "500",
				"--outputAddress=S6Dnv6gRTgsHCmZQxjN7cX5aRjJvDvqwp9", "--coin", "700000", "--hour", "500"},
			message: []string{"e0c6e4982b1b8c33c5be55ac115b69be68f209c5d9054954653e14874664b57d"},
			address: []string{"2EU3JbveHdkxW6z5tdhbbB2kRAWvXC2pLzw"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			output, err := execCommandCombinedOutput(tc.args...)
			if err != nil {
				require.Equal(t, err, "exit status 1")
			}

			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			resp := removeCharacters(lines[1], "[]")
			signatures := strings.Fields(resp)

			for n, signature := range signatures {
				checkOutput, err := execCommandCombinedOutput([]string{"checkMessageSignature", "--address", tc.address[n], "--message", tc.message[n], "--signature", signature}...)
				if err != nil {
					require.Equal(t, err, "exit status 1")
				}

				require.Equal(t, tc.address[n], strings.TrimSpace(string(checkOutput)))
			}
		})
	}
}

func TestWipe(t *testing.T) {
	output, err := execCommandCombinedOutput([]string{"wipe"}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	require.Contains(t, string(output), "Device wiped")

	output, err = execCommandCombinedOutput([]string{"features"}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	decoder := json.NewDecoder(bytes.NewReader(output))
	decoder.DisallowUnknownFields()

	var features messages.Features
	require.NoError(t, decoder.Decode(&features))
	require.NotNil(t, features.Initialized)
	require.False(t, *features.Initialized)
}

func TestGetMixedEntropy(t *testing.T) {
	device := bootstrap(t, "TestGetMixedEntropy", "USB")
	if device == nil {
		return
	}

	bytesNum := 4
	output, err := execCommandCombinedOutput([]string{"getMixedEntropy", "--entropyBytes", fmt.Sprintf("%d", bytesNum)}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	stripped := strings.ReplaceAll(lines[1], "]", "")
	require.Len(t, strings.Fields(stripped), bytesNum)
}

func TestGetRawEntropy(t *testing.T) {
	device := bootstrap(t, "TestGetRawEntropy", "USB")
	if device == nil {
		return
	}

	bytesNum := 4
	output, err := execCommandCombinedOutput([]string{"getRawEntropy", "--entropyBytes", fmt.Sprintf("%d", bytesNum)}...)
	if err != nil {
		require.Equal(t, err, "exit status 1")
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	stripped := strings.ReplaceAll(lines[1], "]", "")
	require.Len(t, strings.Fields(stripped), bytesNum)
}

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if !strings.ContainsRune(characters, r) {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)

}
