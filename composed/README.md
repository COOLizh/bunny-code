### Installation
1. Clone the repo
```sh
git clone https://gitlab.com/greenteam1/composed.git
```
2. Move to directory
```sh
cd composed/
```
3. Download required files
```sh
make download
```
4. Setup your environment. Environment variables list could be found in `env.example`.
Also, you can use default settings:
```sh
make set-env
```
5. Run app
```sh
make run
```
It will be available on your host machine on the port, specified in `API_PORT` variable (by default `localhost:8808`)

