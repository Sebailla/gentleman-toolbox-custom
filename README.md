# 🧰 Gentleman Toolbox

Este es tu centro de comando para nuevos proyectos. Basado en el ecosistema **Gentle AI**, diseñado para automatizar el **Gentleman Stack**.

### 🛠️ Control de Versiones (Remotes)
Este búnker está configurado para recibir lo último de Alan pero guardar tus cambios localmente:
- **`origin`**: Tu repositorio personal (`gentleman-toolbox-custom`).
- **`upstream`**: El repositorio original de Alan (`gentle-ai`).

### 🚀 Paso 1: Inicialización
Cada vez que quieras empezar algo nuevo con todo el stack, corré el script maestro. Nota: el script solo verifica actualizaciones contra tu repo personal (`origin`) para proteger tus customizaciones:

```bash
/Users/sebailla/Documents/Proyectos/gentleman-toolbox/init-project.sh nombre-del-proyecto
```

### 🔄 Paso 1.5: ¿Cómo actualizar con novedades de Alan? (Manual)
Como tenemos un Toolbox super personalizado con el "Pibe Stack" (Stitch, idioma Rioplatense, automatización de specs), las actualizaciones de Alan (`upstream`) ahora se bajan a mano para que nada se rompa por accidente. 

Cuando quieras traerte mejoras del repo oficial:

1. **Bajate la data nueva**:
   ```bash
   cd /Users/sebailla/Documents/Proyectos/gentleman-toolbox
   git fetch upstream
   ```
2. **Fijate qué subió Alan**:
   ```bash
   git log HEAD..upstream/main --oneline
   ```
3. **Elegí cómo traer los cambios**:
   - **Forma Quirúrgica (Cherry-pick)**: ¿Te sirve solo un fix en particular? Usalo así:
     ```bash
     git cherry-pick <hash-del-commit-de-alan>
     ```
   - **Merge Todo (A lo guapo)**: Para traer todas las novedades resolviendo los conflictos a mano si chocan con tu configuración:
     ```bash
     git merge upstream/main
     ```
     *(Resolvé los conflictos, `git add .`, y `git commit`)*

## 🛠️ Paso 2: Configuración Post-Creación
Aunque el script hace casi todo por vos, para que el stack quede **TOTALMENTE COMPLETO** y funcional, seguí estos pasos dentro de la carpeta del proyecto:

### 1. Instalar dependencias bloqueadas
El script las descarga, pero siempre es bueno asegurar que el árbol esté limpio:
```bash
bun install
```

### 2. Configurar Base de Datos
Personalizá tu archivo `.env` con la URL de tu base de datos (PostgreSQL por defecto) y sincronizá Prisma:
```bash
# Edita el .env primero, luego:
bunx prisma db push
```

### 3. Verificar el Guardian Angel (GGA)
Asegurate de que el Angel esté cuidándote la espalda. Como usamos OpenCode con Minimax:
```bash
git add .
gga run
```
*Si no detecta cambios, hacé un pequeño cambio en cualquier archivo y probá de nuevo.*

### 4. Sincronizar Skills & Ecosistema (Opcional si falló el script)
Si por alguna razón el script no inyectó los skills de SDD, forzalo así:
```bash
gentle-ai install --agent opencode --preset full-gentleman
```

## 📦 El Stack Incluido:
- **Next.js 16** (Arquitectura Modular / Feature-Sliced Design estricta con App Router).
- **Prisma 7** (ORM Moderno).
- **Tailwind CSS v4** (Styling de última generación).
- **Git Workflow Pro**: Versionado Automático SemVer con GitHub Actions, Commitlint y pre-push hooks estrictos para nombres de ramas.
- **GGA** (Gentleman Guardian Angel con OpenCode/Minimax).
- **Engram**: Memoria persistente cross-session.
- **SDD**: Flujo de 9 fases para planificar antes de picar código.
- **Tests**: Vitest (Unitarios) & Playwright (E2E).

---
> [!TIP]
> **Recordatorio**: "CONCEPTS > CODE". Usá los comandos de SDD (`/sdd-new`, `/sdd-explore`) para que la IA trabaje con contexto y no como un simple autocompletado.
