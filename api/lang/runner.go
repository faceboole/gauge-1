// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package lang

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"github.com/getgauge/gauge/api"
	"github.com/getgauge/gauge/config"
	"github.com/getgauge/gauge/conn"
	gm "github.com/getgauge/gauge/gauge_messages"
	"github.com/getgauge/gauge/logger"
	"github.com/getgauge/gauge/manifest"
	"github.com/getgauge/gauge/runner"
	"github.com/getgauge/gauge/util"
	"github.com/sourcegraph/go-langserver/pkg/lsp"
)

type langRunner struct {
	runner    runner.Runner
	killChan  chan bool
	lspID     string
	conn      *grpc.ClientConn
	lspClient gm.LspServiceClient
}

var lRunner langRunner

func startRunner() error {
	lRunner.killChan = make(chan bool)
	var err error
	lRunner.runner, err = connectToRunner(lRunner.killChan)
	if err != nil {
		return fmt.Errorf("Unable to connect to runner : %s", err.Error())
	}
	return nil
}

func connectToRunner(killChan chan bool) (runner.Runner, error) {
	logInfo(nil, "Starting language runner")
	outFile, err := util.OpenFile(logger.ActiveLogFile)
	if err != nil {
		return nil, err
	}
	address, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	os.Setenv("GAUGE_LSP_GRPC_PORT", strconv.Itoa(address.Port))
	runner, err := api.ConnectToRunner(killChan, false, outFile)
	if err != nil {
		return nil, err
	}
	return runner, nil
}

func cacheFileOnRunner(uri lsp.DocumentURI, text string) error {
	r := &gm.CacheFileRequest{Content: text, FilePath: string(util.ConvertURItoFilePath(uri)), IsClosed: false}
	_, err := lRunner.lspClient.CacheFile(context.Background(), r)
	return err
}

func sendMessageToRunner(cacheFileRequest *gm.Message) error {
	err := conn.WriteGaugeMessage(cacheFileRequest, lRunner.runner.Connection())
	if err != nil {
		return fmt.Errorf("Error while connecting to runner : %v", err)
	}
	return nil
}

var GetResponseFromRunner = func(message *gm.Message) (*gm.Message, error) {
	if lRunner.runner == nil {
		return nil, fmt.Errorf("Error while connecting to runner")
	}
	return conn.GetResponseForMessageWithTimeout(message, lRunner.runner.Connection(), config.RunnerRequestTimeout())
}

func getStepPositionResponse(uri lsp.DocumentURI) (*gm.StepPositionsResponse, error) {
	stepPositionsRequest := &gm.StepPositionsRequest{FilePath: util.ConvertURItoFilePath(uri)}
	response, err := lRunner.lspClient.GetStepPositions(context.Background(), stepPositionsRequest)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to runner : %s", err)
	}
	if response.GetError() != "" {
		return nil, fmt.Errorf("error while connecting to runner : %s", response.GetError())
	}
	return response, nil
}

func getImplementationFileList() (*gm.ImplementationFileListResponse, error) {
	implementationFileListRequest := &gm.Empty{}
	response, err := lRunner.lspClient.GetImplementationFiles(context.Background(), implementationFileListRequest)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to runner : %s", err.Error())
	}
	return response, nil
}

func putStubImplementation(filePath string, codes []string) (*gm.FileDiff, error) {
	stubImplementationCodeRequest := &gm.StubImplementationCodeRequest{ImplementationFilePath: filePath, Codes: codes}
	response, err := lRunner.lspClient.ImplementStub(context.Background(), stubImplementationCodeRequest)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to runner : %s", err.Error())
	}
	return response, nil
}

func getAllStepsResponse() (*gm.StepNamesResponse, error) {
	getAllStepsRequest := &gm.StepNamesRequest{}
	response, err := lRunner.lspClient.GetStepNames(context.Background(), getAllStepsRequest)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to runner : %s", err.Error())
	}
	return response, nil
}

func killRunner() {
	if lRunner.runner != nil {
		lRunner.runner.Kill()
	}
}

func getLanguageIdentifier() (string, error) {
	m, err := manifest.ProjectManifest()
	if err != nil {
		return "", err
	}
	info, err := runner.GetRunnerInfo(m.Language)
	if err != nil {
		return "", err
	}
	return info.LspLangId, nil
}
