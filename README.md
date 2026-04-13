# 🧰 Gentleman Toolbox

Este es tu centro de comando para nuevos proyectos. Basado en el ecosistema **Gentle AI**, diseñado para automatizar el **Gentleman Stack**.

### 🛠️ Control de Versiones (Remotes)
Este búnker está configurado para recibir lo último de Alan pero guardar tus cambios localmente:
- **`origin`**: Tu repositorio personal (`gentleman-toolbox-custom`).
- **`upstream`**: El repositorio original de Alan (`gentle-ai`).

### 🚀 Paso 1: Instalación de SAI

Para tener el control total de tu búnker y protegerte de actualizaciones que puedan romper tus personalizaciones, usamos **SAI (Seba AI)** como punto de entrada inteligente.

1. **Compilar e Instalar**:
   ```bash
   cd /Users/sebailla/Documents/Proyectos/gentleman-toolbox
   chmod +x install-sai.sh
   ./install-sai.sh
   ```
2. **Crear el acceso global** (Recomendado):
   ```bash
   sudo ln -sf "$(pwd)/bin/sai" /usr/local/bin/sai
   ```

Una vez instalado, usá siempre `sai` en lugar de `gentle-ai`.

### 🏗️ Paso 2: Inicialización de Proyectos
Cada vez que quieras empezar algo nuevo, corré el inicializador desde **cualquier carpeta** de tu sistema:

```bash
init-project nombre-del-proyecto
```

> `init-project` se instala globál junto con `sai` cuando corremos el `install-sai.sh`. No hay que recordar rutas absolutas.

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
Si por alguna razón el script no inyectó los skills de SDD, forzalo usando el overlay:
```bash
sai install --agent opencode --preset full-gentleman
```

---
Para más detalles sobre cada comando y la arquitectura del stack, consulta el [SAI_STACK_MANUAL.md](file:///Users/sebailla/Documents/Proyectos/gentleman-toolbox/SAI_STACK_MANUAL.md).

## 📦 El Stack Incluido:
- **Next.js 16** (Arquitectura Modular / Feature-Sliced Design estricta con App Router).
- **Prisma 7** (ORM Moderno).
- **Tailwind CSS v4** (Styling de última generación).
- **Git Workflow Pro**: Versionado Automático SemVer con GitHub Actions, Commitlint y pre-push hooks estrictos para nombres de ramas.
- **GGA** (Gentleman Guardian Angel con OpenCode/Minimax).
- **Engram**: Memoria persistente cross-session.
- **SDD**: Flujo de 9 fases para planificar antes de picar código.
- **Orquestador Estricto**: La IA asume el rol exclusivo de coordinadora. No toca código, no asume contexto y delega las tareas en subagentes especializados.
- **Única Fuente de Verdad UI/UX**: Todas las interfaces derivan estrictamente de la carpeta `design-md`, que se propaga y clona automáticamente a cada nuevo proyecto.
- **Tests**: Vitest (Unitarios) & Playwright (E2E).

---
> [!TIP]
> **Recordatorio**: "CONCEPTS > CODE". Usá los comandos de SDD (`/sdd-new`, `/sdd-explore`) para que la IA trabaje con contexto y no como un simple autocompletado.
