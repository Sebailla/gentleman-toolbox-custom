#!/bin/bash

# install-sai.sh - Instalador de SAI (Super AI / Seba AI)
# Este script compila sai y lo hace disponible en el sistema.

set -e

# Colores
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

# Asegurar que estamos en el directorio correcto
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$DIR"

echo -e "${BLUE}============================================${NC}"
echo -e "${BLUE} 🚀 SAI Stack Installer${NC}"
echo -e "${BLUE}============================================${NC}"
echo ""

# -----------------------------------------------
# 1. Compilar SAI
# -----------------------------------------------
echo -e "${BLUE}[1/3] Compilando SAI...${NC}"
go build -o sai ./cmd/sai/main.go

mkdir -p ./bin
mv sai ./bin/sai
echo -e "${GREEN}✅ SAI compilado en ./bin/sai${NC}"

# -----------------------------------------------
# 2. Instalar symlinks globales en /usr/local/bin
# -----------------------------------------------
echo ""
echo -e "${BLUE}[2/3] Instalando comandos globales...${NC}"

INSTALL_DIR="/usr/local/bin"

install_symlink() {
  local SOURCE="$1"
  local TARGET="$2"
  local NAME="$3"

  if [ ! -f "$SOURCE" ]; then
    echo -e "${RED}❌ No se encontró: $SOURCE — saltando $NAME${NC}"
    return
  fi

  if sudo ln -sf "$SOURCE" "$TARGET" 2>/dev/null; then
    echo -e "${GREEN}✅ '$NAME' instalado en $TARGET${NC}"
  else
    echo -e "${YELLOW}⚠️  No se pudo crear el symlink para '$NAME' en $TARGET.${NC}"
    echo -e "   Intentá manualmente: ${GREEN}sudo ln -sf \"$SOURCE\" \"$TARGET\"${NC}"
  fi
}

# Instalar 'sai'
install_symlink "$DIR/bin/sai" "$INSTALL_DIR/sai" "sai"

# Instalar 'init-project' (mapeado desde init-project.sh)
install_symlink "$DIR/init-project.sh" "$INSTALL_DIR/init-project" "init-project"

# Asegurar que init-project.sh sea ejecutable
chmod +x "$DIR/init-project.sh"

# -----------------------------------------------
# 3. Verificar
# -----------------------------------------------
echo ""
echo -e "${BLUE}[3/3] Verificando instalación...${NC}"

echo -ne "  sai:          "
if command -v sai &>/dev/null; then
  echo -e "${GREEN}$(sai version 2>&1 | head -1)${NC}"
else
  echo -e "${YELLOW}No encontrado en PATH (reiniciá el shell)${NC}"
fi

echo -ne "  init-project: "
if command -v init-project &>/dev/null; then
  echo -e "${GREEN}OK → $(command -v init-project)${NC}"
else
  echo -e "${YELLOW}No encontrado en PATH (reiniciá el shell)${NC}"
fi

echo ""
echo -e "${GREEN}============================================${NC}"
echo -e "${GREEN} ✅ Instalación completa${NC}"
echo -e "${GREEN}============================================${NC}"
echo ""
echo -e "Usá ${GREEN}sai <comando>${NC}       para tus herramientas personalizadas."
echo -e "Usá ${GREEN}init-project <nombre>${NC} para arrancar un nuevo proyecto."
echo ""
