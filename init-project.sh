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
    git fetch -q origin main 2>/dev/null
    git fetch -q upstream main 2>/dev/null
    
    LOCAL=$(git rev-parse HEAD)
    ORIGIN_REMOTE=$(git rev-parse origin/main 2>/dev/null || echo "$LOCAL")
    UPSTREAM_REMOTE=$(git rev-parse upstream/main 2>/dev/null || echo "$LOCAL")
    
    # 1. Verificar si hay novedades de Alan (Upstream)
    UPSTREAM_BASE=$(git merge-base HEAD "$UPSTREAM_REMOTE" 2>/dev/null || echo "$LOCAL")
    if [ "$LOCAL" != "$UPSTREAM_REMOTE" ] && [ "$UPSTREAM_BASE" != "$UPSTREAM_REMOTE" ]; then
        log_info "Se encontraron cambios nuevos en Upstream (Alan). Actualizando..."
        echo -e "${CYAN}Cambios detectados:${NC}"
        git log HEAD..upstream/main --oneline -n 10
        git pull --rebase -q upstream main
        log_success "Toolbox actualizado con lo último de Alan."
        LOCAL=$(git rev-parse HEAD) # Actualizar LOCAL después del rebase
    fi

    # 2. Verificar estado con tu repo personal (Origin)
    if [ "$LOCAL" = "$ORIGIN_REMOTE" ]; then
        log_success "Toolbox sincronizado con tu repo personal."
    else
        ORIGIN_BASE=$(git merge-base HEAD "$ORIGIN_REMOTE" 2>/dev/null || echo "$LOCAL")
        if [ "$LOCAL" = "$ORIGIN_BASE" ]; then
            log_info "Tu repo local está atrasado respecto a tu origin. Actualizando..."
            git pull -q origin main
        elif [ "$ORIGIN_REMOTE" = "$ORIGIN_BASE" ]; then
            log_warn "Tenés cambios locales no pusheados a tu repo personal (origin)."
        else
            log_warn "Tu historial ha divergido de 'origin' (probablemente por el rebase de Alan)."
            log_info "Sugerencia: Revisá los cambios y tirá un 'git push --force origin main' para sincronizar tu repo."
        fi
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
# Usamos --yes para evitar promts interactivos (ej: React Compiler)
bunx create-next-app@latest . \
    --typescript \
    --tailwind \
    --eslint \
    --app \
    --src-dir \
    --import-alias "@/*" \
    --use-bun \
    --skip-install \
    --yes

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
