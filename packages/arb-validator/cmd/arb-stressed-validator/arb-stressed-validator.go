/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"flag"
	zerolog "github.com/rs/zerolog/log"
	"log"
	"os"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/cmdhelper"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

// Launches the rollup validator with the following command line arguments:
// 1) Compiled Arbitrum bytecode file
// 2) private key file
// 3) Global EthBridge addresses json file
// 4) ethURL
func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	zerolog.Logger = zerolog.With().Caller().Logger()

	// Check number of args
	flag.Parse()
	switch os.Args[1] {
	case "validate":
		if err := cmdhelper.ValidateRollupChain("evil-arb-validator", createStressedManager); err != nil {
			log.Fatal(err)
		}
	default:
	}
}

func createStressedManager(ctx context.Context, rollupAddress common.Address, client arbbridge.ArbClient, contractFile string, dbPath string) (*rollupmanager.Manager, error) {
	return rollupmanager.CreateManager(
		context.Background(),
		rollupAddress,
		arbbridge.NewStressTestClient(client, time.Second*10),
		contractFile,
		dbPath,
	)
}
