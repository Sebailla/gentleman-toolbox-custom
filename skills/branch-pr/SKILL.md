---
name: gentle-ai-branch-pr
description: >
  Standard PR and Branching workflow for Gentleman Búnkers (GitFlow simplificado).
  Trigger: When creating a pull request, opening a PR, or preparing changes for review.
license: Apache-2.0
metadata:
  author: gentleman-programming
  version: "3.0"
---

# Gentle AI — Branch & PR Skill (Gentleman Standard)

## When to Use

Load this skill whenever you need to:
- Create a branch for a new task (feature, fix, refactor).
- Prepare changes for review and merge.
- Sync branches or release to production.

## Critical Rules

1.  **Regla de Oro**: Proactividad absoluta. El Agente crea la rama de tarea automáticamente (`tipo/nombre`) al detectar un nuevo requerimiento, sin esperar a que el usuario lo pida.
2.  **Strict Branching Model**:
    - `main`: Production, **INALTEABLE**. Solo lectura vía PR desde `develop`.
    - `develop`: Integración y desarrollo. Base para todas las tareas.
    - `tipo/nombre-en-kebab-case`: Ramas de corta vida sacadas de `develop`.
2.  **Conventional Commits**: Todos los commits deben seguir el estándar para permitir el versionado automático.
3.  **No AI Attribution**: Nunca agregues trailers como "Co-authored-by: AI" a los commits.
4.  **No Force Push a ramas protegidas**: Prohibido en `main` y `develop`.

## Workflow

### 1. Preparación del terreno
Antes de empezar cualquier tarea, asegurate de tener lo último de lo último.
```bash
git checkout develop && git pull origin develop
```

### 2. Creación de la rama de tarea
Saca una rama de `develop` con el formato correcto.
```bash
git checkout -b <tipo>/<nombre-tarea>
# Ejemplos: feat/login-ui, fix/broken-header, docs/readme-update
```

### 3. Desarrollo e Implementación
Trabajá siguiendo los principios del búnker (Modular Vertical Slicing, Clean Code).

### 4. Commits Conscientes
Usá siempre el formato de commits convencionales.
```bash
git add .
git commit -m "feat(auth): add google provider to login"
```

### 5. Integración (PR a Develop)
Una vez terminada la tarea y validados los tests, abrí un PR contra `develop`.
```bash
gh pr create --base develop --title "feat: <descripcion>" --body "Closes #<issue>"
```

### 6. Release (PR a Main)
Cada cierto tiempo, o cuando una feature grande está lista, se fusiona `develop` en `main` para disparar el versionado automático (vX.X.X).

## Branch Naming Pattern

```plaintext
^(feat|fix|chore|docs|style|refactor|perf|test|build|ci|revert)\/[a-z0-9-]+$
```

| Tipo | Propósito |
|------|-----------|
| `feat/` | Nueva funcionalidad. |
| `fix/` | Corrección de un bug. |
| `docs/` | Cambios en documentación. |
| `refactor/` | Cambio de código que no altera funcionalidad. |
| `chore/` | Mantenimiento, dependencias, herramientas. |
| `test/` | Adición o corrección de tests. |

## Conventional Commits

```plaintext
<type>(<scope>)!: <description>
```

- **feat**: Una nueva característica.
- **fix**: Una corrección de error.
- **chore**: Tareas de mantenimiento.
- **!**: Indica un *Breaking Change*.
