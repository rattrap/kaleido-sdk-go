// Copyright 2018 Kaleido, a ConsenSys business

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package profile

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage on-chain indentity profile",
}

func init() {
	// rootCmd adds registrCmd manually

	// create persistent flags for all sub-commands as this is not a complete command by itself
	profileCmd.PersistentFlags().StringP("id", "i", "", "Service ID (optional)")
	profileCmd.PersistentFlags().StringP("profile", "r", "default", "Profile to use for network connections")
	viper.BindPFlag("profile", profileCmd.PersistentFlags().Lookup("profile"))

	viper.BindPFlag("services.idregistry.id", profileCmd.PersistentFlags().Lookup("id"))
}

// NewProfileCmd registry cmd
func NewProfileCmd() *cobra.Command {
	return profileCmd
}