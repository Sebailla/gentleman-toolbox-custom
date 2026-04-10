# Graphify — Guía Completa de Conceptos

## Qué es un Grafo de Conocimiento

Un grafo de conocimiento es una representación estructurada de información donde:

- **Nodos** = entidades individuales (archivos, funciones, conceptos, imágenes)
- **Conexiones (edges)** = relaciones entre entidades
- **Comunidades** = grupos de nodos que están más conectados entre sí que con el resto

El grafo se construye automáticamente a partir de tu codebase, extrayendo relaciones estructurales (imports, llamados) e inferencias semánticas (conceptos relacionados sin vínculo directo).

---

## Nodo (Node)

Un **nodo** es la unidad fundamental del grafo — representa una entidad individual.

### Qué puede ser un nodo

| Tipo | Ejemplo |
|------|---------|
| Archivo | `src/app/(dashboard)/routines/page.tsx` |
| Función | `getAuthenticatedUser()` |
| Componente | `AthleteHeader()` |
| Concepto | `Layered Architecture Pattern` |
| Documento | `plans/arquitectura-tecnica.md` |
| Imagen | `plans/boarding.png` |
| Concepto abstracto | `Role-Based Access Control` |

### Estructura de un nodo

```json
{
  "id": "getAuthenticatedUser",
  "label": "getAuthenticatedUser()",
  "file_type": "code",
  "source_file": "src/shared/lib/auth/session.ts",
  "source_location": null
}
```

### Campos

| Campo | Descripción |
|-------|-------------|
| `id` | Identificador único en el grafo |
| `label` | Nombre legible para humanos |
| `file_type` | `code`, `document`, `paper`, `image` |
| `source_file` | Archivo de origen |
| `source_location` | Ubicación específica (línea, sección) |
| `author` | Autor (si tiene frontmatter) |
| `contributor` | Contribuidor (si tiene frontmatter) |

### Degree (Grado)

El **degree** de un nodo es la cantidad de conexiones que tiene. Los nodos con degree alto son "god nodes" — las abstracciones más importantes del sistema.

```
getAuthenticatedUser() — degree: 10
getCurrentUserId() — degree: 10
```

---

## Conexión (Edge)

Una **conexión** representa una relación entre dos nodos.

### Tipos de conexiones

| Tipo | Significado | Extracción |
|------|-------------|------------|
| `calls` | A invoca B directamente (ej: función llama a función) | EXTRACTED |
| `references` | A menciona o usa B (ej: import, uso de tipo) | EXTRACTED |
| `wraps` | A envuelve a B (ej: layout wrappea pages) | EXTRACTED |
| `implements` | A implementa B (ej: acción implementa schema) | EXTRACTED |
| `shares_data_with` | A y B comparten estructura de datos | INFERRED |
| `conceptually_related_to` | A está conceptualmente relacionado con B | INFERRED |
| `semantically_similar_to` | A y B resuelven problemas similares | INFERRED |
| `rationale_for` | A explica el POR QUÉ de B | INFERRED |
| `cites` | A cita a B (papers, docs) | EXTRACTED |

### Confianza de extracción

| Tag | Score | Significado |
|-----|-------|-------------|
| `EXTRACTED` | 1.0 | La relación está explícitamente en el código (import, llamado directo) |
| `INFERRED` | 0.6 - 0.9 | Inferencia razonable basada en evidencia estructural o semántica |
| `AMBIGUOUS` | 0.1 - 0.3 | El extractor no está seguro — podría no ser real |

### Escala de confianza INFERRED

| Score | Evidencia |
|-------|-----------|
| 0.8 - 0.9 | Evidencia estructural directa (shared data, dependencia clara) |
| 0.6 - 0.7 | Inferencia razonable con algo de incertidumbre |
| 0.4 - 0.5 | Especulación débil |

### Ejemplo real del grafo

```
Onboarding Multi-Step Wizard --calls--> BiomarkersForm - Step 1
[EXTRACTED] confidence: 1.0

Athlete Workout Builder Component --rationale_for--> Cycle-Week-Day Hierarchical Structure
[INFERRED] confidence: 0.75

Landing Page --semantically_similar_to--> Fanatic Gym Documentation (llms.txt)
[INFERRED] confidence: 0.75
```

---

## Comunidad (Community)

Una **comunidad** es un grupo de nodos que están más conectados entre sí que con el resto del grafo. Es como un "barrio" — los nodos dentro comparten funcionalidad o propósito común.

### Cohesión

La **cohesión** mide qué tan fuertemente conectados están los nodos dentro de una comunidad:

| Cohesión | Significado |
|----------|------------|
| 0.9 - 1.0 | Comunidad muy cohesiva — todos los nodos se conectan entre sí |
| 0.5 - 0.8 | Buena cohesión |
| 0.2 - 0.5 | Cohesión moderada |
| 0.02 - 0.2 | Comunidad débil — hay nodos que no se conectan mucho con los demás |

### Ejemplo de comunidad

```
Comunidad 7 - "Asignación Rutinas"
Cohesión: 0.04

Nodos:
  - Assign Routine Page
  - Assign Routine To Athlete Action
  - Coach Athletes Page
  - Get Athletes Action
  - Routine Assignment Model
  - etc.
```

Esta comunidad tiene cohesión 0.04 — bastante baja. Los nodos comparten funcionalidad (asignar rutinas a atletas) pero no están todos fuertemente conectados entre sí.

### Comunidades con cohesión alta

```
Comunidad 23 - "PWA Service Workers"
Cohesión: 0.32

Comunidad 26 - "API Routes"
Cohesión: 0.4
```

Estas son comunidades pequeñas y enfocadas donde todos los nodos se conectan fuertemente.

---

## God Nodes

Los **God Nodes** son los nodos con más conexiones (degree más alto) del grafo. Representan las abstracciones más importantes del sistema — los puntos de entrada o las funciones más reutilizadas.

### Top 10 de tu codebase

```
1. getAuthenticatedUser() — 10 conexiones
2. getCurrentUserId() — 10 conexiones
3. Onboarding Multi-Step Wizard — 8 conexiones
4. Root Layout - Provider Stack — 8 conexiones
5. createNotification() — 7 conexiones
6. Dashboard Layout - Role-Based Navigation — 7 conexiones
7. createSession() — 6 conexiones
8. storeGet() — 6 conexiones
```

Estos son los nodos que más cosas conectan — si los tocás, impactan muchas partes del sistema.

---

## Conexiones Sorprendentes

Son relaciones que el grafo detecta pero que no son obvias a simple vista. Pueden ser:

- **Conceptos en diferentes archivos que resuelven el mismo problema**
- **Patrones compartidos entre partes del código que parecen independientes**
- **Dependencias implícitas no documentadas**

### Ejemplo de tu grafo

```
QueryProvider --semantically_similar_to--> Zod Schemas as Source of Truth
[INFERRED]
src/app/layout.tsx → CONTRIBUTING.md
```

Esto conecta el provider de queries con la filosofía de validación de Zod — no hay un vínculo directo en el código, pero el grafo infiere que ambos representan el mismo patrón de validación de datos.

---

## Nodos Aislados (Knowledge Gaps)

Son nodos con ≤1 conexión — código que no está siendo tocado por nada más. Pueden ser:

- **Deuda técnica** — código huérfano que debería eliminarse
- **Funcionalidad intencionalmente aislada**
- **Documentación faltante** — código que necesita más relaciones

### Ejemplo de tu grafo

```
21 nodos débilmente conectados:
  - MockAudioContext (x3)
  - ErrorBoundary Component
  - Update Routine Action
  - etc.
```

MockAudioContext aparece 3 veces y no se conecta a nada — probable candidato a eliminar o documentar.

---

## Hyperedges

Un **hyperedge** representa una relación grupal donde 3 o más nodos participan en un concepto común.

### Ejemplos de tu grafo

```
Authentication Flow
  Nodes: auth_layout, login_page, login_form, landing_page, 
         onboarding_page, get_profile_action
  [EXTRACTED 1.0]

Onboarding Multi-Step Flow
  Nodes: onboarding_page, step_indicator, biomarkers_form, 
         goals_form, injuries_form, training_sections_form
  [EXTRACTED 1.0]
```

El Authentication Flow es un hyperedge — todos estos nodos participan en el flujo de autenticación, pero no todos se llaman entre sí directamente.

---

## Interpretación del Reporte

### Resumen

```
- 1458 nodos · 2114 conexiones · 76 comunidades detectadas
- Extracción: 98% EXTRACTED · 1% INFERRED · 0% AMBIGUOUS
- INFERRED: 31 edges (avg confidence: 0.82)
```

Esto te dice qué tan confiables son las conexiones del grafo. 98% EXTRACTED significa que casi todas las relaciones están verificables en el código.

### Nodos Principales

Lista de God Nodes — las abstracciones más conectadas.

### Conexiones Sorprendentes

Las relaciones INFERRED más interesantes que el grafo detectó.

### Comunidades

Lista de todas las comunidades con:
- Nombre (si fue labelado)
- Cohesión
- Primeros nodos del grupo

### Knowledge Gaps

Lista de nodos aislados y comunidades muy pequeñas (≤2 nodos) que podrían ser ruido o deuda técnica.

### Preguntas Sugerendas

Preguntas que el grafo puede responder basándose en su estructura — especialmente útiles para auditoría de arquitectura.

---

## Cómo usar esta información

### Para entender el sistema

1. Empezá por los **God Nodes** — te dan los puntos más importantes
2. Explorá las **comunidades** relacionadas a tu área de interés
3. Seguí las **conexiones** para entender el flujo

### Para auditoría de arquitectura

1. Buscá **comunidades con baja cohesión** (0.02-0.1) — podrían estar haciendo demasiado
2. Revisá las **conexiones sorprendendentes** — pueden revelar acoplamientos no intencionales
3. Atendé los **nodos aislados** — probable deuda técnica

### Para verificar convenciones

1. Buscá si las conexiones respetan tu arquitectura (ej: servicios no tocan DB directamente)
2. Verificá que los **hyperedges** matcheen con los flujos documentados
3. Chequeá que los **conceptos** (RBAC, Zod, etc.) estén bien conectados

---

## Glosario

| Término | Definición |
|---------|------------|
| Node / Nodo | Entidad individual en el grafo |
| Edge / Conexión | Relación entre dos nodos |
| Hyperedge | Relación grupal (3+ nodos) |
| Community / Comunidad | Grupo de nodos fuertemente conectados |
| Cohesion / Cohesión | Qué tan conectados están los nodos dentro de una comunidad |
| Degree | Cantidad de conexiones de un nodo |
| God Node | Nodo con más conexiones |
| EXTRACTED | Conexión verificada en el código |
| INFERRED | Conexión razonada, no directamente verificable |
| AMBIGUOUS | Conexión incierta, requiere revisión |
| Knowledge Gap | Nodo aislado o comunidad muy pequeña |
