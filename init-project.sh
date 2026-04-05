#!/usr/bin/env bash

# ============================================================================
# Gentleman Stack Initializer
# Un script para dominar el mundo (o al menos tus proyectos)
# ============================================================================

set -e

# Configuración de colores
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
CYAN='\033[0;36m'
BOLD='\033[1m'
NC='\033[0m'

log_info() { echo -e "${CYAN}${BOLD}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}${BOLD}[OK]${NC} $1"; }
log_warn() { echo -e "${YELLOW}${BOLD}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}${BOLD}[ERROR]${NC} $1"; }

# 1. Validar nombre del proyecto
PROJECT_NAME=$1
if [[ -z "$PROJECT_NAME" ]]; then
    log_error "Falta el nombre del proyecto. Uso: ./init-project.sh <nombre>"
    exit 1
fi

# 2. Verificar estado con tu repo personal (Origin)
TOOLBOX_DIR="/Users/sebailla/Documents/Proyectos/gentleman-toolbox"
if [ -d "$TOOLBOX_DIR/.git" ]; then
    log_info "Verificando actualizaciones de tu propio Toolbox..."
    # Guardamos el dir actual para volver después del check
    pushd "$TOOLBOX_DIR" > /dev/null
    git fetch -q origin main 2>/dev/null || log_warn "No se pudo conectar al remoto (origin). Continuando con versión local..."
    
    LOCAL=$(git rev-parse HEAD)
    ORIGIN_REMOTE=$(git rev-parse origin/main 2>/dev/null || echo "$LOCAL")
    
    if [ "$LOCAL" = "$ORIGIN_REMOTE" ]; then
        log_success "Toolbox sincronizado con tu repo personal."
    else
        ORIGIN_BASE=$(git merge-base HEAD "$ORIGIN_REMOTE" 2>/dev/null || echo "$LOCAL")
        if [ "$LOCAL" = "$ORIGIN_BASE" ]; then
            log_info "Tu repo local está atrasado respecto a tu origin. Actualizando..."
            git pull -q origin main || log_warn "No se pudo hacer pull. Seguimos con tu versión local actual."
        elif [ "$ORIGIN_REMOTE" != "$ORIGIN_BASE" ]; then
            log_warn "Tenés cambios locales y remotos que divergieron."
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
log_info "Lanzando create-next-app (Next.js 16, TS, Tailwind v4, App Router)..."
# Usamos --no-install para controlar nosotros las dependencias
# Usamos --yes para evitar promts interactivos (ej: React Compiler)
bunx create-next-app@16 . \
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
    @types/react-dom @types/bcryptjs @types/jsonwebtoken \
    @commitlint/cli @commitlint/config-conventional standard-version

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

# 8. Estructura de carpetas modular (Feature-Sliced Design)
log_info "Creando estructura de carpetas modular..."
mkdir -p src/core/lib src/core/types src/core/hooks src/modules src/components/ui
mkdir -p .docs .agent/skills plans specs designs .github/workflows

# 8.5. Configurar GitHub Actions (Release Automático)
log_info "Configurando GitHub Actions para versionado semántico..."
cat > .github/workflows/release.yml <<'EOF'
name: Auto Release En Main

on:
  push:
    branches:
      - main

jobs:
  release:
    name: "🚀 Generar Release"
    runs-on: ubuntu-latest
    permissions:
      contents: write
    steps:
      - name: Checkout repo
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
          token: ${{ secrets.GITHUB_TOKEN }}

      - name: Setup Bun
        uses: oven-sh/setup-bun@v2

      - name: Run Release Script
        run: |
          git config --global user.name 'github-actions[bot]'
          git config --global user.email 'github-actions[bot]@users.noreply.github.com'
          bun run release
          git push --follow-tags origin main
EOF

# 9. Crear AGENTS.md base (Reglas para la IA)
log_info "Generando AGENTS.md base..."
cat > AGENTS.md <<EOF
# Project Rules (Gentleman Standard)

## Architecture
- Use Modular / Feature-Sliced Architecture optimized for Next.js 16 Fullstack (App Router STRICT).
- Group logic (services, types, components) by domain in \`src/modules/<feature-name>/\`.
- Keep shared global logic in \`src/core/\` and shared UI in \`src/components/ui/\`.

## Standards
- Component-Driven: Keep feature UI inside domain modules. Only generic UI (buttons, dialogs) goes in \`src/components/ui/\`.
- Zod for schema validation.
- Prisma for all DB operations.
- Tailwind CSS v4 for styling.

## 🗣️ Comunicación y Documentación
- **Idioma del Código**: Inglés (variables, funciones, clases).
- **Idioma de Comentarios y Documentación**: Español.
- **Comunicación del Agente**: Español Rioplatense (voseo, terminología técnica clara pero directa).
- **Revisiones**: El feedback del Guardian Angel debe ser siempre en español.

## 🛡️ Protocolo de Actuación (MANDATORIO)
- **Rol del Orquestador**: El orquestador DEBE limitarse a guiar y mantener el hilo de la conversación. No debe escribir ni modificar código directamente. Toda acción técnica, incluyendo la edición de código, DEBE ser delegada a **subagentes**.
- **Cero Suposiciones**: Nunca te quedes con dudas ni infieras requerimientos, arquitectura o decisiones técnicas. Antes de tomar cualquier decisión, PREGUNTA todo lo necesario al usuario para tener el contexto completo y exacto.
- **Confirmación Constante**: Antes de delegar cambios significativos a un subagente, resume la acción y espera la confirmación del usuario.

## 🎨 Diseño UI/UX
- **Base de Diseño**: TODO lo relacionado con el diseño, UI, UX y aspectos visuales DEBE basarse estricta y únicamente en la información documentada en la carpeta \`design-md\`. No inventes, ni uses información externa, apégate al contenido de esa carpeta.

## Quality & Workflow
- All new features must have unit tests (Vitest).
- Critical flows must have E2E tests (Playwright).
- Use Conventional Commits strictly.
- **Branch Naming**: When asked to create a branch, ALWAYS follow this format: `tipo/nombre-en-kebab-case`. Valid types are: `feat, fix, hotfix, chore, docs, refactor, test`. (e.g. `feat/new-dashboard`).
EOF

# 9b. Sincronizar multiverso de IA (Provider-Agnostic)
log_info "Sincronizando reglas para todos los agentes (Cursor, Claude, Gemini, Windsurf)..."
cp AGENTS.md CLAUDE.md
cp AGENTS.md GEMINI.md
cp AGENTS.md .cursorrules
cp AGENTS.md .windsurfrules

# 11b. Instalar Skill de Documentación de Usuario
log_info "Instalando skill de documentación de usuario..."
mkdir -p .agent/skills/documentar-specs-usuario
mkdir -p .agent/rules
cat > .agent/rules/agent-settings.md <<EOF
# Agent Settings & Project Preferences

## Communication
- **Format**: Always use **Spanish (Rioplatense)** for interaction with the user.
- **Documentation**: All project documentation, READMEs, and ADRs must be written in **Spanish**.

## 🛡️ Protocolo del Orquestador
- **Delegación y Liderazgo**: Como orquestador de AI, debes llevar el hilo de la conversación y planificar. NUNCA toques código. Toda modificación debe ser delegada a **subagentes**.
- **Cero Suposiciones**: Nunca te quedes con dudas ni infieras reglas de arquitectura, decisiones o qué requiere el usuario. Evalúa e interroga metódicamente antes de proceder.

## UI/UX Design Workflow
- **Standard**: La única fuente de verdad para toda tarea referida a la UI, UX o diseño visual es la carpeta \`design-md\`.
- **Action**: Siempre que debas crear o modificar una interfaz, básate obligatoria y estrictamente en la documentación y los assets presentes en \`design-md\`. Nunca inventes estilos por fuera de ese directorio.

## Specs Workflow
- **Action**: After completing a 'feat', use the 'documentar-specs-usuario' skill to document the change in /specs.

## Plannings & Design Workflow
- **Action**: Before writing code for a new feature, use the 'documentar-plan-diseno' skill to detail the architecture in /plans and UI specs in /designs.
EOF

# 11b. Instalar Skills de Documentación de Usuario y Planificación
log_info "Instalando skills de documentación (Usuario y Planificación)..."
mkdir -p .agent/skills/documentar-specs-usuario
mkdir -p .agent/skills/documentar-plan-diseno

cat > .agent/skills/documentar-specs-usuario/SKILL.md <<'EOF'
---
name: documentar-specs-usuario
description: Genera especificaciones funcionales para el usuario final en la carpeta /specs después de completar una nueva característica.
---

# 📝 Documentación de Especificaciones para el Usuario

## Cuándo usar esta skill
- **Trigger**: Después de terminar una funcionalidad nueva (\`feat\`).

## Instrucciones
1. Crear un archivo en \`specs/YYYY-MM-DD-nombre-feature.md\`.
2. Redactar en lenguaje no técnico para el usuario final.
3. Incluir: ¿Qué es?, Cómo se usa y Beneficios.
EOF

cat > .agent/skills/documentar-plan-diseno/SKILL.md <<'EOF'
---
name: documentar-plan-diseno
description: Genera y almacena documentos de planificación técnica y diseño UI/UX en las carpetas /plans y /designs antes de iniciar el código de una nueva feature.
---

# 🏗️ Documentación de Planificación y Diseño

## Cuándo usar esta skill
- **Trigger**: **ANTES** de empezar a escribir código para una funcionalidad nueva grande o compleja, o durante debates de arquitectura y UI.

## Flujo de Trabajo
1. Si es técnica (Backend, BD): Crear \`plans/YYYY-MM-DD-plan-feature.md\`.
2. Si es visual (Frontend, UX, UI): Crear \`designs/YYYY-MM-DD-diseno-feature.md\`.
3. Detallar arquitectura, dependencias y referencias a Stitch o diseño propuesto.
EOF

# 10. Configurar Husky, Commitlint y Versionado
log_info "Finalizando configuración de Husky y Commitlint..."
bunx husky init

# Añadir configuración para commitlint
cat > commitlint.config.mjs <<'EOF'
export default { extends: ['@commitlint/config-conventional'] };
EOF

# Agregar hook commit-msg para obligar conventional commits
cat > .husky/commit-msg <<'EOF'
bunx --no -- commitlint --edit $1
EOF

# Agregar validación estricta de nombres de ramas (pre-push)
cat > .husky/pre-push <<'EOF'
#!/usr/bin/env bash
LOCAL_BRANCH=$(git rev-parse --abbrev-ref HEAD)
VALID_REGEX="^(feat|fix|hotfix|chore|docs|refactor|test)\/[a-z0-9-]+$"

if [[ "$LOCAL_BRANCH" == "main" || "$LOCAL_BRANCH" == "master" || "$LOCAL_BRANCH" == "develop" ]]; then
    exit 0
fi

if [[ ! $LOCAL_BRANCH =~ $VALID_REGEX ]]; then
    echo "❌ ERROR Arquitectónico: La rama '$LOCAL_BRANCH' es un cambalache."
    echo "👉 Formato exigido: tipo/nombre-en-kebab-case (ej: feat/login-ui, fix/header-roto)"
    echo "👉 Tipos válidos: feat, fix, hotfix, chore, docs, refactor, test"
    exit 1
fi
EOF
chmod +x .husky/pre-push

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
npm pkg set scripts.release="standard-version"

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

# 13. Copiar design-md (Fuente de verdad de UI/UX)
if [ -d "$TOOLBOX_DIR/design-md" ]; then
    log_info "Copiando carpeta design-md inicial..."
    cp -r "$TOOLBOX_DIR/design-md" .
else
    log_warn "No se encontró la carpeta design-md en el toolbox, creando carpeta vacía..."
    mkdir -p design-md
fi

log_success "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
log_success "  ¡PROYECTO '$PROJECT_NAME' LISTO PARA LA GUERRA!"
log_success "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
log_info "Pasos siguientes:"
echo "  1. cd $PROJECT_NAME"
echo "  2. bun install"
echo "  3. Configura tu .env con la URL de la base de datos"
echo "  4. Empezá a laburar con el Guardian Angel cuidándote la espalda."
