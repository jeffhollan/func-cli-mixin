package functions

import "fmt"

// This is an example. Replace thie following with whatever steps are needed to
// install required components into
// const dockerfileLines = `RUN apt-get update && \
// apt-get install gnupg apt-transport-https lsb-release software-properties-common -y && \
// echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ stretch main" | \
//    tee /etc/apt/sources.list.d/azure-cli.list && \
// apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
// 	--keyserver packages.microsoft.com \
// 	--recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
// apt-get update && apt-get install azure-cli
// `

// curl -s https://api.github.com/repos/azure/azure-functions-core-tools/releases/latest \
// | grep "browser_download_url.*linux-x64.*zip\"" \
// | cut -d : -f 2,3 \
// | tr -d \" \
// | wget -qi -
//
// unzip Azure.Functions.*.zip -d /func-cli
// chmod +x /func-cli/func
//
// ln -s /func-cli/func /usr/bin/func

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {

	fmt.Fprintln(m.Out, `RUN curl -s https://api.github.com/repos/azure/azure-functions-core-tools/releases/latest \`)
	fmt.Fprintln(m.Out, `| grep "browser_download_url.*linux-x64.*zip\"" \`)
	fmt.Fprintln(m.Out, `| cut -d : -f 2,3 \`)
	fmt.Fprintln(m.Out, `| tr -d \" \`)
	fmt.Fprintln(m.Out, `| wget -qi -`)
	fmt.Fprintln(m.Out, `RUN unzip Azure.Functions.*.zip -d /func-cli`)
	fmt.Fprintln(m.Out, `RUN chmod +x /func-cli/func`)
	fmt.Fprintln(m.Out, `RUN ln -s /func-cli/func /usr/bin/func`)

	return nil
}
