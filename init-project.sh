#!/usr/bin/env bash

# ============================================================================
# Gentleman Stack Initializer
# Un script para dominar el mundo (o al menos tus proyectos)
# ============================================================================

set -e

# Configuración de colores
RED='\033[0;31m'
GREEN='\033[0;32m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m'

log_info() { echo -e "${CYAN}${BOLD}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}${BOLD}[OK]${NC} $1"; }
log_error() { echo -e "${RED}${BOLD}[ERROR]${NC} $1"; }

# 1. Validar nombre del proyecto
PROJECT_NAME=$1
if [[ -z "$PROJECT_NAME" ]]; then
    log_error "Falta el nombre del proyecto. Uso: ./init-project.sh <nombre>"
    exit 1
fi

# 2. Auto-update del Toolbox (Opcional pero recomendado)
TOOLBOX_DIR="/Users/sebailla/Documents/Proyectos/gentleman-toolbox"
if [ -d "$TOOLBOX_DIR/.git" ]; then
    log_info "Verificando actualizaciones del Toolbox..."
    # Guardamos el dir actual para volver después del check
    pushd "$TOOLBOX_DIR" > /dev/null
    git fetch -q upstream main
    
    UPSTREAM="upstream/main"
    LOCAL=$(git rev-parse @)
    REMOTE=$(git rev-parse "$UPSTREAM")
    BASE=$(git merge-base @ "$UPSTREAM")

    if [ "$LOCAL" = "$REMOTE" ]; then
        log_success "Toolbox actualizado con lo último de Alan."
    elif [ "$LOCAL" = "$BASE" ]; then
        log_info "Se encontraron cambios nuevos en Upstream (Alan). Actualizando..."
        echo -e "${CYAN}Cambios detectados:${NC}"
        git log ..upstream/main --oneline -n 10
        git pull --rebase -q upstream main
        log_success "Toolbox actualizado con éxito."
    elif [ "$REMOTE" = "$BASE" ]; then
        log_info "Tenés cambios locales no pusheados. Continuando..."
    else
        log_error "El repositorio ha divergido. Por favor, resolvé los conflictos manualmente en $TOOLBOX_DIR"
    fi
    popd > /dev/null
fi

# 3. Crear y entrar al directorio
log_info "Creando proyecto: $PROJECT_NAME..."
mkdir -p "$PROJECT_NAME"
cd "$PROJECT_NAME"

# 3. Inicializar repositorio Git
log_info "Inicializando Git..."
git init -q

# 4. Crear Next.js App (Gentleman Stack)
log_info "Lanzando create-next-app (Next.js latest, TS, Tailwind v4, App Router)..."
# Usamos --no-install para controlar nosotros las dependencias
bunx create-next-app@latest . \
    --typescript \
    --tailwind \
    --eslint \
    --app \
    --src-dir \
    --import-alias "@/*" \
    --use-bun \
    --no-install

# 5. Instalar dependencias core del stack
log_info "Instalando dependencias del Gentleman Stack..."
bun add @prisma/client lucide-react clsx tailwind-merge date-fns zod \
    react-hot-toast ioredis bcryptjs jsonwebtoken

bun add -d prisma vitest @testing-library/react @testing-library/dom \
    jsdom @playwright/test husky lint-staged tsx @types/node @types/react \
    @types/react-dom @types/bcryptjs @types/jsonwebtoken

# 6. Inicializar Prisma (PostgreSQL por defecto)
log_info "Inicializando Prisma..."
bunx prisma init

# 7. Configurar Gentleman Guardian Angel (GGA)
log_info "Configurando GGA (Guardian Angel)..."
cat > .gga << 'EOF'
# Gentleman Guardian Angel Configuration
# https://github.com/Gentleman-Programming/gentle-ai

# AI Provider (required)
# Usando OpenCode (usa tu modelo por defecto)
PROVIDER="opencode"

# File patterns to review (comma-separated globs)
FILE_PATTERNS="*.ts,*.tsx,*.js,*.jsx"

# Patterns to exclude
EXCLUDE_PATTERNS="*.test.ts,*.spec.ts,*.test.tsx,*.spec.tsx,*.d.ts"
EOF

if command -v gga &>/dev/null; then
    gga install
else
    log_warn "GGA no encontrado en el sistema, pero generamos .gga por si lo instalás luego."
fi

# 8. Estructura de carpetas pro
log_info "Creando estructura de carpetas pro..."
mkdir -p src/lib src/hooks src/services src/types src/components/ui
mkdir -p .docs .agent/skills plans/2.0

# 9. Crear AGENTS.md base (Reglas para la IA)
log_info "Generando AGENTS.md base..."
cat > AGENTS.md <<EOF
# Project Rules (Gentleman Standard)

## Architecture
- Use Clean/Hexagonal Architecture.
- Keep business logic in \`src/services\` or \`src/lib\`.
- UI components should be in \`src/components\`.

## Standards
- Atomic Design for components.
- Zod for schema validation.
- Prisma for all DB operations.
- Tailwind CSS v4 for styling.

## Quality
- All new features must have unit tests (Vitest).
- Critical flows must have E2E tests (Playwright).
- Use Conventional Commits.
EOF

# 10. Configurar Husky
log_info "Finalizando configuración de Husky..."
bunx husky init
# Agregar lint-staged al pre-commit
cat > .husky/pre-commit <<'EOF'
bun test
bunx lint-staged
gga run
EOF
# 11. Configurar package.json scripts
log_info "Actualizando scripts en package.json..."
# Bun no tiene un equivalente exacto a 'npm pkg set', mantenemos npm para edición de metadatos o usamos sed/node
npm pkg set scripts.test="vitest"
npm pkg set scripts.db:seed="tsx prisma/seed.ts"
npm pkg set scripts.db:reset="prisma migrate reset --force && bun run db:seed"

# 12. Inyectar inteligencia (Gentle AI Skills & Ecosystem)
log_info "Configurando ecosistema Gentle AI (Skills, Engram, Persona)..."
if command -v gentle-ai &>/dev/null; then
    # Instalamos/Actualizamos globalmente para asegurar que existen
    gentle-ai install --agent opencode --preset full-gentleman
    
    # Copiamos los skills al proyecto para que el agente los vea localmente
    SKILLS_SOURCE="$HOME/.config/opencode/skills"
    if [ -d "$SKILLS_SOURCE" ]; then
        log_info "Copiando skills desde $SKILLS_SOURCE..."
        cp -r "$SKILLS_SOURCE/"* .agent/skills/
    else
        log_warn "No se encontraron skills en $SKILLS_SOURCE. Corré 'gentle-ai install' manualmente."
    fi
else
    log_warn "gentle-ai no encontrado. Corré 'gentle-ai install' manualmente."
fi

log_success "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
log_success "  ¡PROYECTO '$PROJECT_NAME' LISTO PARA LA GUERRA!"
log_success "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
log_info "Pasos siguientes:"
echo "  1. cd $PROJECT_NAME"
echo "  2. bun install"
echo "  3. Configura tu .env con la URL de la base de datos"
echo "  4. Empezá a laburar con el Guardian Angel cuidándote la espalda."
