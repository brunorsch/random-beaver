# Random Beaver 🦫

Um serviço web que serve fotos aleatórias de castores para cada IP.

🌐 **Deploy:** [https://beaver.brunorsch.dev.br](https://beaver.brunorsch.dev.br)

## 🚀 Executando com Docker

### Pré-requisitos
- Docker
- Docker Compose

### Setup inicial (apenas uma vez)
```bash
# Adicionar usuário ao grupo docker (executar com sudo)
sudo ./setup-docker.sh

# Aplicar permissões do grupo
newgrp docker
```

### Deploy
```bash
# Build e deploy automatizado
./deploy.sh
```

### Comandos úteis
```bash
# Build manual
docker-compose build

# Executar
docker-compose up -d

# Ver logs
docker-compose logs -f

# Parar
docker-compose down

# Rebuild completo
docker-compose build --no-cache
```

## 🛠️ Desenvolvimento

### Executando localmente (sem Docker)
```bash
go mod download
go run main.go
```

### Estrutura do projeto
- `main.go` - Aplicação principal em Go
- `castores/` - Imagens dos castores
- `Dockerfile` - Configuração Docker
- `docker-compose.yml` - Orquestração Docker
