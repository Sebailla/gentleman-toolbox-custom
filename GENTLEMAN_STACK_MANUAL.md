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
- **Memoria (Engram)**: Listo para recordar tus decisiones.

---

## 🏗️ 2. El Flujo SDD (Spec-Driven Development)

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

## 😇 3. El Guardian Angel (GGA)

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

## 🩹 3.5. ¿Qué hacer si el Angel tira un FAILED?

Si el Guardian Angel te da un veredicto de **FAILED**, usá al agente para redimir el código. No lo hagas a mano si podés automatizarlo.

### 🧠 Estrategia de Repregunta (Prompt Maestro)
Copiá el reporte de errores y pegalo en el chat de OpenCode con este comando:

> **Prompt:** *"El Guardian Angel me tiró un FAILED: [PEGAR REPORTE]. Arreglalo de raíz usando `/sdd-ff 'Fix GGA: [Resumen del error]'`. Aseguráte de cumplir con todas las reglas de AGENTS.md, separar tipos en `types.ts` y usar Zod si corresponde. ¡Quiero el PASS ya!"*

### 🔥 Por qué usar SDD para arreglar fallos:
1.  **Traza de Tareas**: La IA desglosa cada violación del Angel en una tarea específica.
2.  **Validación**: Te obliga a generar los tests que el Angel te está reclamando.
3.  **Persistencia**: La solución queda documentada en los planes, no es un "parche" al aire.

---

## 4. Solución de Problemas (Troubleshooting)

Aquí tenés los problemas que resolvimos "a los hachazos" hoy. Si te pasa de nuevo, tirá de manual:

### 4.1. Error de versión de Tcl (have 9.0.0, need 8.5)
Si al usar `gga` te tira un conflicto de versión de Tcl, corré esto:
```bash
brew unlink tcl-tk && brew install tcl-tk@8 && brew link --force tcl-tk@8
```

### 4.2. Model Not Found (OpenCode)
Si OpenCode dice que no encuentra el modelo, verificá el ID exacto en el `.gga`.
**Formatos aceptados:**
- `PROVIDER="opencode"` (Usa el modelo por defecto de tu CLI).
- `PROVIDER="opencode:minimax/minimax-m2.5"` (Formato `vendor/modelo`).
- `PROVIDER="opencode:minimax-m2.5"` (ID de modelo plano).

### 4.3. Los Skills no aparecen en el IDE
Si el agente no te responde a comandos de `/sdd-...`, es que los skills no llegaron a la carpeta del proyecto.
```bash
cp -r ~/.config/opencode/skills/* .agent/skills/
```

---

## 5. Tabla de Providers (Archivo .gga)
Si decidís cambiar de cerebro, aquí tenés los IDs para el campo `PROVIDER`:

| Servicio | String en .gga | Notas |
| :--- | :--- | :--- |
| **OpenCode** | `opencode` | Recomendado con Minimax/DeepSeek local. |
| **Claude** | `claude` | Solo si tenés saldo en el Angel (GGA). |
| **Gemini** | `gemini` | Usa la API Key configurada globalmente. |
| **Ollama** | `ollama:llama3.2` | Para modelos 100% locales en tu Mac. |
| **LMStudio** | `lmstudio` | Ideal para jugar con modelos GGUF. |

---

## 6. Generación de Nuevos Skills (Fábrica de Inteligencia)

Si necesitás que el agente aprenda un patrón nuevo o una tecnología específica, usá el **`skill-creator`**.

### Cómo pedir un Skill:
En tu chat de OpenCode, simplemente pedilo:
> *"Usá el `skill-creator` para generarme un nuevo skill en `.agent/skills/nombre-del-skill` para [Tecnología/Patrón]."*

### Qué incluye un Skill:
- **SKILL.md**: El cerebro con reglas, ejemplos y comandos.
- **Triggers**: Instrucciones para que la IA sepa cuándo cargar este conocimiento.
- **Registro**: Recordale al agente que lo agregue a la tabla de `AGENTS.md`.

---

## 7. Mantenimiento y Comandos Útiles

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

### Base de Datos (Prisma)
```bash
# Sincroniza cambios en el schema.prisma con tu DB local
bunx prisma db push

# Abre el visor de datos para ver qué hay en las tablas
bunx prisma studio
```

---

## 📜 8. Filosofía del Gentleman
- **CONCEPTS > CODE**: Entendé por qué usás Atomic Design antes de crear un átomo.
- **SOLID FOUNDATIONS**: La arquitectura es la que aguanta el peso del proyecto a largo plazo.
- **AI IS A TOOL**: Vos sos Tony Stark. Vos decidís la arquitectura, la IA es el brazo que ejecuta.

> "No programamos para que funcione, programamos para que sea mantenible por humanos (o IAs inteligentes)."

---
