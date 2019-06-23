# holdmybeer
`hmb`, for short, is your buddy to ease your environment-setup phase. 

## Usage
`hmb` has two subcommands - `setup` and `install`

Subcommand: `setup`, can be used to development tools. Currently, it only supports flag: `--bare-minimum` (or `-b` for shorthand) to install tools like `vim, curl and htop`

Subcommand: `setup`, can be used to install languages. Currently, only go and node installation is supported by `hmb`, with `java` support in the pipeline

To install go, simply run:   
`hmb install lang=go`
Note that above command will install version `1.12.5`.

To install specific version of go, attach the version after the colon, like:   
`hmb install lang=go:1.12.5`
Note that absolute version of go is compulsory, as the setup for the same will be downloaded from the internet.

To install node, simply run:   
`hmb install lang=node`
Note that above command will install version `19.x`. To customise installation, attach major-version like `10.x` or `12.x` after the colon. Please note that only major-version-number is required for node followed by '.x' to indicate whichever is the latest for that major release.