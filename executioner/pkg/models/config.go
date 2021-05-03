/*Package models consists of the structures necessary for the executor to work*/
package models

//BuildConfig consists of the necessary information to run the solution in a specific programming language
type BuildConfig struct {
	FileExtension  string `json:"file_extension"`
	BuildCommands  string `json:"build_commands"`
	ContainerName  string `json:"container_name"`
	BinPath        string `json:"container_bin_path"`
	SrcPath        string `json:"container_src_path"`
	DockerFilePath string `json:"docker_file_path"`
}
