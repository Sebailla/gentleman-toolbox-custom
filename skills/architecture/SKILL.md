---
name: gentle-ai-modular-architecture
description: >
  Enforces the Modular Vertical Slicing architecture pattern for Next.js projects.
  Trigger: When creating or refactoring domain modules, implementing business logic, or structuring project architecture.
license: Apache-2.0
metadata:
  author: gentleman-programming
  version: "1.0"
---

# Gentle AI — Modular Vertical Slicing Skill

## When to Use

Load this skill whenever you are:
- Creating a new module in `src/modules/`.
- Refactoring existing logic into services or actions.
- Defining domain types and schemas.
- Structuring routing in `src/app/`.

## Core Principle: Modular Vertical Slicing

The project is organized in **Módulos de Dominio** dentro de `src/modules/`. Cada módulo es una unidad funcional independiente que sigue un patrón de capas.

### Modular Folder Structure
Cada nuevo módulo debe seguir estrictamente este esquema:

```plaintext
src/modules/[module-name]/
├── components/   # UI Components (Client/Server)
├── services/     # Business Logic & DB Queries (Server Only)
├── actions.ts    # Server Actions (Orchestration & Validation)
├── types.ts      # Domain Interfaces & Zod Schemas
└── index.ts      # Public API (Export only what's necessary)
```

## Layered Responsibilities (Modern MVC)

### A. The Service Layer (services/)
**Responsabilidad**: Lógica de negocio pura, acceso a datos (Supabase/Prisma), integraciones externas.
**Regla**: No debe usar `revalidatePath`, `cookies()`, o manejar `FormData`. Debe retornar datos limpios o lanzar errores de negocio.

### B. The Action Layer (actions.ts)
**Responsabilidad**: Actuar como controlador. Validar inputs con Zod, llamar a servicios, manejar el manejo de errores para la UI y revalidar el caché de Next.js.
**Regla**: Siempre usar `'use server'`.

### C. The Presentation Layer (components/)
**Responsabilidad**: Renderizar UI.
**Regla**: Los componentes no deben realizar consultas directas a la base de datos si la lógica es compleja; deben delegar en el service (si es Server Component) o en la action (si es Client Component).

## Strict Dependency Rules

1.  **Unidireccionalidad**: Los módulos pueden importar de `shared` o `lib`, pero nunca un módulo de dominio debe importar directamente archivos internos de otro módulo de dominio.
2.  **Public API**: Si el módulo A necesita algo del módulo B, debe importarlo exclusivamente desde `src/modules/B/index.ts`.
3.  **App Layer**: La carpeta `src/app/` solo debe contener archivos de ruteo (`page.tsx`, `layout.tsx`) que actúen como "glue code" llamando a componentes de los módulos.

## Coding Instructions for the AI

- **Orden de implementación**: Primero define el esquema en `types.ts`, luego el servicio en `services/`, la acción en `actions.ts` y finalmente el componente.
- **Refactorización**: Mueve cualquier lógica de base de datos fuera de los componentes hacia la capa de servicios.
- **Nomenclatura**: Los servicios deben agruparse en un objeto o clase (ej. `ProductService`) para mantener el namespace limpio.
