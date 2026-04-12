# 🎭 El Gentleman Stack - Manual de Operaciones 2026

Este manual define el flujo de trabajo sagrado para construir software de alta calidad usando el ecosistema **Gentle AI**, **OpenCode** y el **Gentleman Stack**. No es solo código, es **disciplina arquitectónica**.

---

## 🛠️ 1. El Comienzo: Inicialización del Búnker

Para cada nuevo proyecto, no pierdas tiempo configurando carpetas. El script maestro ya lo hace por vos con **Next.js 15, Prisma 7, Tailwind v4 y SDD Skills**.

```bash
/Users/sebailla/Documents/Proyectos/gentleman-toolbox/init-project.sh nombre-de-tu-app
```

**Qué incluye el búnker:**
- **Atomic Design ready**: Carpetas configuradas en `src/components`.
- **Inteligencia inyectada**: 12 skills de SDD copiados a `.agent/skills/`.
- **Guardian Angel**: Configurado con **OpenCode:minimax-m2.5**.
- **Memoria (Engram)**: El cerebro del stack, encargado de recordar decisiones y contexto.

---

## 🧠 2. Memoria Persistente (Engram)

**Engram** es el sistema que permite a Antigravity (y a cualquier agente del stack) recordar qué hicimos en sesiones pasadas, qué decisiones arquitectónicas tomamos y cómo resolvimos bugs específicos.

### Comandos CLI Útiles
Podés interactuar con la memoria directamente desde tu terminal:

| Comando | Descripción |
| :--- | :--- |
| `engram tui` | **(RECOMENDADO)** Abre la interfaz visual para navegar memorias. |
| `engram search "texto"` | Busca en toda la base de conocimientos por palabras clave. |
| `engram save "Título" "Mensaje"` | Guarda una nota mental manual (usar `--project` para filtrar). |
| `engram context` | Muestra un resumen del contexto reciente que la IA tiene cargado. |
| `engram stats` | Estado de salud de la base de datos de memoria. |

### ¿Cómo funciona la IA?
Cuando lanzás un comando `/sdd-*`, la IA utiliza el protocolo **MCP (Model Context Protocol)** para consultar `engram` automáticamente. Al finalizar una tarea con `/sdd-archive`, el resumen se guarda en Engram para futuras referencias.

---

## 🏗️ 3. El Flujo SDD (Spec-Driven Development)

**REGLA DE ORO**: Nunca escribas código sin antes planificar. Usá estos comandos en tu chat de **OpenCode**.

| Fase | Comando | Qué hace |
| :--- | :--- | :--- |
| **Exploración** | `/sdd-explore "tarea"` | Analiza el código actual y propone soluciones. |
| **Propuesta** | `/sdd-propose "tarea"` | Crea un `proposal.md` con la solución técnica. |
| **Especificación** | `/sdd-spec "tarea"` | Define los specs técnicos y tests necesarios. |
| **Diseño** | `/sdd-design "tarea"` | Planifica la estructura de archivos y arquitectura. |
| **Tareas** | `/sdd-tasks "tarea"` | Crea la lista de pasos atómicos para la IA. |
| **Fast-Forward** | `/sdd-ff "tarea"` | **(RECOMENDADO)** Corre Explorar -> Tareas en un solo paso. |
| **Aplicación** | `/sdd-apply "tarea"` | La IA empieza a escribir el código y los tests. |
| **Verificación** | `/sdd-verify "tarea"` | Comprueba que los tests pasen y el código sea sólido. |
| **Archivado** | `/sdd-archive "tarea"` | Limpia los planes y guarda el conocimiento en **Engram**. |

---

## 🌿 3.5. Gestión de Ramas y Versiones (GitFlow Simplificado)

Para que el búnker no sea un caos, seguimos una jerarquía de ramas estricta:

1.  **`main` (Santuario)**: Es el reflejo exacto de lo que está en producción. **PROHIBIDO** pushear cambios directos a esta rama. Solo se actualiza vía Pull Request desde `develop`. Al fusionar en `main`, se dispara el versionado automático (`vX.X.X`).
2.  **`develop` (Motor)**: Es la rama principal de trabajo. Es el punto de partida y de llegada de todas las tareas.
3.  **`tipo/nombre-tarea` (Silos)**: Cada tarea se trabaja en su propia rama saliendo de `develop`.
    - Formato: `feat/login-view`, `fix/broken-header`, `refactor/api-calls`.
    - **Regla de Oro**: La IA es proactiva. Si le pedís una tarea, va a crear la rama automáticamente sin preguntarte.

### Versionado Automático
Usamos **Conventional Commits** para que el sistema sepa cuándo subir la versión:
- `feat:` -> Sube versión **MINOR** (ej. 1.0.0 -> 1.1.0).
- `fix:` -> Sube versión **PATCH** (ej. 1.1.0 -> 1.1.1).
- `feat/fix!:` (con el signo de exclamación) -> Sube versión **MAJOR** (ej. 1.1.1 -> 2.0.0).

Cada vez que el script de inicialización crea un búnker, realiza un "Ritual de Día Cero": crea el primer commit en `main`, le clava el tag `v1.0.0` y te deja automáticamente en `develop` para empezar a laburar.

## 🏗️ 3.6. Arquitectura Modular (Modular Vertical Slicing)

Olvidate de la vieja arquitectura por componentes sueltos. El proyecto se instaura bajo un modelo **Modular nativo para Next.js 16 (App Router)** basado en **Vertical Slicing**:

### Estructura de un Módulo (`src/modules/[module-name]/`)
Cada funcionalidad independiente es un módulo con sus propias capas:
- **`services/`**: Lógica de negocio pura y consultas a la DB. Prohibido usar APIs de Next.js (cookies, revalidate) acá.
- **`actions.ts`**: El controlador. Valida inputs con **Zod**, llama a servicios y maneja la revalidación de caché. Siempre `'use server'`.
- **`components/`**: UI específica del dominio.
- **`types.ts`**: Contrato del dominio (Interfaces y Schemas).
- **`index.ts`**: La API pública del módulo. **IMPORTANTE**: Otros módulos solo pueden importar de acá.

### Reglas de Dependencia
- `src/core/*`: Lógica, tipos y hooks compartidos globales.
- `src/components/ui/`: Solo componentes visuales genéricos (Botones, inputs).
- **Unidireccionalidad**: Los módulos no se cruzan. Si necesitás algo de otro módulo, usá su `index.ts`.

---

## 🧰 4. Catálogo de Skills (Toolbox)

Los **Skills** son archivos de inteligencia modular que Antigravity carga dinámicamente según la tarea que le pidas. Aquí tenés el inventario de lo que tenés instalado en tu búnker:

### A. Core SDD (El Motor de Trabajo)
Estas capacidades permiten el flujo de **Spec-Driven Development**.
- **`sdd-init` / `sdd-explore`**: Inicialización del contexto y escaneo de código existente.
- **`sdd-propose` / `sdd-spec`**: Generación de propuestas técnicas y especificaciones de tests.
- **`sdd-design` / `sdd-tasks`**: Diseño de arquitectura de archivos y desglose en tareas atómicas.
- **`sdd-apply` / `sdd-verify`**: Implementación de código y validación automatizada de resultados.
- **`sdd-archive`**: Limpieza de planos y guardado de conocimiento en **Engram**.

### B. Infraestructura y Fábrica
- **`skill-creator`**: Skill maestro diseñado para crear *otros* skills siguiendo el estándar.
- **`skill-registry`**: El bibliotecario. Escanea todos tus skills (globales y locales) y prepara el mapa de herramientas para la IA.

### C. Especialidades Técnicas
- **`go-testing`**: Experto en el ecosistema Go. Genera mocks, corre tests y verifica la cobertura de código.

### D. Operaciones y Calidad
- **`branch-pr`**: Automatiza la creación de ramas y Pull Requests siguiendo convenciones.
- **`issue-creation`**: Transforma bugs o requerimientos en Issues de GitHub bien documentados.

### E. Autonomía y Estrategia (La Evolución)
Estas son las nuevas capacidades integradas para un búnker inteligente.

#### 🕹️ Comandos CLI (Explícitos)
Correlos en tu terminal o pedile a la IA que los ejecute.

| Comando | Descripción | Cuándo usarlo |
| :--- | :--- | :--- |
| `gentle-ai distill` | Filtra el contexto del proyecto. | Al inicio de sesión. |
| `gentle-ai doctor --fix` | Diagnostica y repara el búnker. | Si algo falla o faltan reglas. |
| `gentle-ai drive --test="X"` | Inicia loop autónomo de corrección. | Si un test no pasa. |
| `gentle-ai learn` | Extrae instintos de la sesión. | Al terminar tu día. |
| `gentle-ai sentinel install` | Activa el Centinela (Pre-commit). | En cada nuevo repo. |
| `gentle-ai status --score` | Calcula el AHI (Salud Arquitectónica). | Para medir deuda técnica. |
| `gentle-ai briefing` | Genera reporte de traspaso. | Para cambiar de agente/sesión. |
| `gentle-ai plan --feature="X"` | Genera scaffolding modular. | Para empezar un módulo nuevo. |
| `gentle-ai judge <target>` | Inicia el Juicio Adversario. | Antes de un merge crítico. |
| `gentle-ai blueprint` | Genera mapa de dependencias. | Para auditar el acoplamiento. |
| `gentle-ai refactor --instincts` | Aplica tus gustos al código viejo. | Para evitar la podredumbre técnica. |
| `gentle-ai pr-fix <url>` | Corrige comentarios de PR. | Para automatizar el feedback de GitHub. |
| `gentle-ai console` | Lanza el dashboard interactivo. | Para gestión visual del búnker. |
| `gentle-ai sync` | Sincroniza tácticas globales. | Para mantener el búnker actualizado. |

#### 🧠 Skills de Contexto (Implícitos)
La IA los carga automáticamente según la necesidad. No necesitás invocarlos, están ahí para protegerte.
- **`go-reviewer` / `typescript-reviewer`**: Opinión experta durante el `/sdd-verify`.
- **`gentleman-coach`**: Tu mentor senior. Te frena si hacés "cambalache".
- **`gentleman-storyteller`**: Redacta tus PRs con tono profesional y contexto real.
- **`gentleman-driver`**: Lógica de "pensamiento" para el loop de arreglos.
- **`judgment-day`**: Protocolo de doble revisión ciega.
- **`gentleman-blueprint`**: Traductor de grafos a visión de arquitecto.
- **`instinct-refactor`**: El motor de evolución del proyecto.
- **`pr-fix`**: El brazo ejecutor que limpia tus Pull Requests.
- **`gentleman-console`**: El cerebro visual del sistema.
- **`fleet-sync`**: El enlace de comunicación con el Cuartel General.

### G. Arquitecto Supremo (Techo de Cristal)
Este es el estado final de la evolución. Tu búnker está conectado al mundo.

- **`pr-fix`**: No pierdas tiempo corrigiendo estilos o sugerencias menores en GitHub. Dejá que el búnker lea los comentarios de tus compañeros y los arregle por vos.
- **`console`**: Si te abruman los comandos, tirá un `console`. Tenés todo el estado de tu arquitectura, tus instintos y tus acciones rápidas en un dashboard "pipi cucu".
- **`sync`**: La inteligencia artificial evoluciona cada día. Corré `sync` para asegurarte de que tu búnker local esté usando las últimas reglas y patrones descubiertos por la comunidad Gentleman.

### F. Arsenal de Elite (Nivel Pipi Cucu)
Estas son herramientas de alta gama para proyectos que no pueden fallar.

- **`judge`**: Si una feature es muy compleja, pedile un juicio. No te conformes con una opinión; exigí consenso.
- **`blueprint`**: Tiralo una vez por semana. Si las flechas se cruzan mucho, tenés un problema de diseño.
- **`refactor --instincts`**: Los proyectos mueren porque dejan de parecerse a lo que el autor quería. Este comando mantiene viva tu esencia en cada archivo.

## 😇 5. El Guardian Angel (GGA)

El Angel es tu mentor. No te deja hacer commits con código basura.

```bash
# Paso 1: Agregá tus cambios al stage
git add .

# Paso 2: Pedí la bendición del Angel
gga run
```

### ⚠️ El "Secreto" del Guardian Angel (GGA)
**RECUERDA**: El Angel sólo revisa archivos que están en el **Stage** de Git. Si no tiraste `git add`, el Angel no ve nada.
```bash
git add .
gga run
```

**Si falla (STATUS: FAILED):** Arreglá lo que te pide. No intentes engañarlo, él sabe lo que es el buen código.
**Si pasa (STATUS: PASSED):** Ya podés hacer el commit tranquilo.

## 🩹 5.5. ¿Qué hacer si el Angel tira un FAILED?

Si el Guardian Angel te da un veredicto de **FAILED**, usá al agente para redimir el código. No lo hagas a mano si podés automatizarlo.

### 🧠 Estrategia de Repregunta (Prompt Maestro)
Copiá el reporte de errores y pegalo en el chat de OpenCode con este comando:

> **Prompt:** *"El Guardian Angel me tiró un FAILED: [PEGAR REPORTE]. Arreglalo de raíz usando `/sdd-ff 'Fix GGA: [Resumen del error]'`. Aseguráte de cumplir con todas las reglas de AGENTS.md, separar tipos en `types.ts` y usar Zod si corresponde. ¡Quiero el PASS ya!"*

### 🔥 Por qué usar SDD para arreglar fallos:
1.  **Traza de Tareas**: La IA desglosa cada violación del Angel en una tarea específica.
2.  **Validación**: Te obliga a generar los tests que el Angel te está reclamando.
3.  **Persistencia**: La solución queda documentada en los planes, no es un "parche" al aire.

---

## 6. Solución de Problemas (Troubleshooting)

Aquí tenés los problemas que resolvimos "a los hachazos" hoy. Si te pasa de nuevo, tirá de manual:

### 6.1. Error de versión de Tcl (have 9.0.0, need 8.5)
Si al usar `gga` te tira un conflicto de versión de Tcl, corré esto:
```bash
brew unlink tcl-tk && brew install tcl-tk@8 && brew link --force tcl-tk@8
```

### 6.2. Model Not Found (OpenCode)
Si OpenCode dice que no encuentra el modelo, verificá el ID exacto en el `.gga`.
**Formatos aceptados:**
- `PROVIDER="opencode"` (Usa el modelo por defecto de tu CLI).
- `PROVIDER="opencode:minimax/minimax-m2.5"` (Formato `vendor/modelo`).
- `PROVIDER="opencode:minimax-m2.5"` (ID de modelo plano).

### 6.3. Los Skills no aparecen en el IDE
Si el agente no te responde a comandos de `/sdd-...`, es que los skills no llegaron a la carpeta del proyecto.
```bash
cp -r ~/.config/opencode/skills/* .agent/skills/
```

---

## 7. Tabla de Providers (Archivo .gga)
Si decidís cambiar de cerebro, aquí tenés los IDs para el campo `PROVIDER`:

| Servicio | String en .gga | Notas |
| :--- | :--- | :--- |
| **OpenCode** | `opencode` | Recomendado con Minimax/DeepSeek local. |
| **Claude** | `claude` | Solo si tenés saldo en el Angel (GGA). |
| **Gemini** | `gemini` | Usa la API Key configurada globalmente. |
| **Ollama** | `ollama:llama3.2` | Para modelos 100% locales en tu Mac. |
| **LMStudio** | `lmstudio` | Ideal para jugar con modelos GGUF. |

---

## 8. 🛠️ Personalización (Agregar Skills)

Podés expandir la inteligencia del stack agregando tus propios "chips" de conocimiento.

### ¿Cómo crear un Skill?
1. **Definir el objetivo**: Creá una carpeta en `~/.config/opencode/skills/` (para que sea global) o en `.agent/skills/` (solo para el proyecto).
2. **Crear `SKILL.md`**: Usar el frontmatter obligatorio (name, description, trigger).
3. **Reglas y Ejemplos**: Escribí las reglas que la IA **debe** seguir y ejemplos de código real.

### Comandos de Expansión
- Para crear uno rápidamente, pedíle a la IA: `"Crea un skill para [Tecnología] siguiendo el estándar skill-creator"`.
- Los skills globales en `~/.config/opencode/skills/` se inyectarán automáticamente en cada `init-project.sh`.

---

## 9. Generación de Nuevos Skills (Fábrica de Inteligencia)

Si necesitás que el agente aprenda un patrón nuevo o una tecnología específica, usá el **`skill-creator`**.

### Cómo pedir un Skill:
En tu chat de OpenCode, simplemente pedilo:
> *"Usá el `skill-creator` para generarme un nuevo skill en `.agent/skills/nombre-del-skill` para [Tecnología/Patrón]."*

### Qué incluye un Skill:
- **SKILL.md**: El cerebro con reglas, ejemplos y comandos.
- **Triggers**: Instrucciones para que la IA sepa cuándo cargar este conocimiento.
- **Registro**: Recordale al agente que lo agregue a la tabla de `AGENTS.md`.

---

## 10. Mantenimiento y Comandos Útiles

Si sentís que el agente está "perdido" o se olvidó de los skills:

### Re-Inyectar Inteligencia
```bash
# Sincroniza los skills globales del sistema al proyecto
cp -r ~/.config/opencode/skills/* .agent/skills/
```

### Actualizar el Ecosistema
```bash
# Actualiza las herramientas globales (gentle-ai, engram, gga)
gentle-ai install --agent opencode --preset full-gentleman
```

### Gestión de MCP (Model Context Protocol)
Los MCPs permiten que la IA use herramientas externas (Notion, GitHub, etc.).
- **Comando**: `opencode mcp add` (Sigue las instrucciones en pantalla).
- **Lista**: `opencode mcp ls` (Verifica conexiones).
- **Configuración Manual**: Editá el archivo `~/.config/opencode/opencode.json` en la sección `"mcp"`.

### Base de Datos (Prisma)
```bash
# Sincroniza cambios en el schema.prisma con tu DB local
bunx prisma db push

# Abre el visor de datos para ver qué hay en las tablas
bunx prisma studio
```

---

## 📜 12. Optimización para OpenCode

Este Toolbox está diseñado para brillar con **OpenCode**, ya sea que uses modelos de Anthropic (Claude) o modelos optimizados como **Minimax-m2.5** o **DeepSeek**.

### Reglas para OpenCode
1.  **Contexto es Rey**: Usá siempre `gentle-ai distill` antes de empezar. Los modelos locales/específicos agradecen que no les tires basura.
2.  **Instintos Locales**: Los instintos en `.gentleman/instincts.md` son la clave para que tu IA no se olvide de tus preferencias de tipado o estilos.
3.  **Harness Residencial**: El `sentinel` es obligatorio. Como a veces saltamos de herramienta en herramienta, el Centinela asegura que el estándar se mantenga sin importar el agente.

---

## 📜 13. Filosofía y Protocolo del Gentleman
- **CONCEPTS > CODE**: Entendé por qué usás Atomic Design antes de crear un átomo.
- **SOLID FOUNDATIONS**: La arquitectura es la que aguanta el peso del proyecto a largo plazo.
- **AI IS A TOOL**: Vos sos Tony Stark. Vos decidís la arquitectura, la IA es el brazo que ejecuta.
- **EL ROL DEL ORQUESTADOR**: El agente principal es un coordinador. No edita código directo, ni infiere contextos o requerimientos. Pregunta absolutamente todo antes de actuar y delega todo el trabajo pesado a subagentes.
- **DISEÑO DICTADO, NO INVENTADO**: La única fuente de verdad para todo lo estético, UI y UX, es la información de la carpeta `design-md`. El diseño se clona automáticamente al inicializar el proyecto, proveyendo al agente de los tokens y directivas fijadas.

---
