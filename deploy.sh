#!/bin/bash
# Script para build e deploy do Beaver

echo "🦫 Construindo imagem Docker do Beaver..."
docker-compose build

if [ $? -eq 0 ]; then
    echo "✅ Build concluído com sucesso!"
    
    echo "🚀 Parando containers antigos..."
    docker-compose down
    
    echo "🦫 Iniciando Beaver..."
    docker-compose up -d
    
    echo "✅ Beaver está rodando em http://localhost:9001"
    echo "📊 Para ver logs: docker-compose logs -f"
    echo "🛑 Para parar: docker-compose down"
else
    echo "❌ Erro no build!"
    exit 1
fi