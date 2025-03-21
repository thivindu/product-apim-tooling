/*
*  Copyright (c) WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 LLC. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package get

import (
	"github.com/spf13/cobra"
	"github.com/wso2/product-apim-tooling/import-export-cli/impl"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

const defaulEnvsTableFormat = "table {{.Name}}\t{{.MiManagementEndpoint}}"

var envsCmdFormat string

// GetEnvsCmd related info
const GetEnvsCmdLiteral = "envs"
const getEnvsCmdShortDesc = "Display the list of environments"

const getEnvsCmdLongDesc = `Display a list of environments defined in '` + utils.MainConfigFileName + `' file`

var getEnvsCmdExamples = utils.GetMICmdName() + " list envs"

// getEnvsCmd represents the envs command
var getEnvsCmd = &cobra.Command{
	Use:     GetEnvsCmdLiteral,
	Short:   getEnvsCmdShortDesc,
	Long:    getEnvsCmdLongDesc,
	Example: getEnvsCmdExamples,
	Deprecated: "instead refer to https://mi.docs.wso2.com/en/latest/observe-and-manage/managing-integrations-with-micli/ for updated usage.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Logln(utils.LogPrefixInfo + GetEnvsCmdLiteral + " called")
		envs := utils.GetMainConfigFromFile(utils.MainConfigFilePath).Environments
		impl.PrintEnvs(envs, envsCmdFormat, defaulEnvsTableFormat)
	},
}

func init() {
	if utils.GetMICmdName() == "" {
		GetCmd.AddCommand(getEnvsCmd)
		getEnvsCmd.Flags().StringVarP(&envsCmdFormat, "format", "", defaulEnvsTableFormat, "Pretty-print "+
			"environments using go templates")
	}
}
