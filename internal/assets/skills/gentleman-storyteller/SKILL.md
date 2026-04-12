# Skill: Gentleman Storyteller (PR Creator)

Esta skill te permite redactar Pull Requests (PRs) que no solo describen el código, sino que cuentan la historia del cambio con el tono y estándar del Gentleman.

## Proceso de Relato

Cuando el usuario te pida "armar la PR" o "preparar el commit":

1.  **Analizar el Diff**: Usá `git diff --cached` para ver qué se va a subir exactamente.
2.  **Consultar Memoria**: Leé Engram (`mem_context`) para entender el "por qué" detrás del cambio (decisiones, correcciones del usuario, blockers superados).
3.  **Redactar**: Seguí la estructura oficial de Gentleman Toolbox.

## Estructura de la PR (Output)

### 🚀 Impacto
*Breve resumen de qué mejora esto para el usuario o el sistema.*

### 🛠️ Cambios Técnicos
*   **Módulo X**: Implementación de tal cosa.
*   **Refactor**: Mejora en tal servicio para evitar tal problema.
*   *Uso de la arquitectura Modular Vertical Slicing.*

### 🧪 Verificación Realizada
*   ✅ Tests unitarios ejecutados.
*   ✅ Verificación manual del flujo X.
*   ✅ Resultado del `gentle-ai doctor`.

## Reglas de Oro

*   **Sin Humo**: No uses palabras vacías como "mejoras generales" o "ajustes menores". Sé específico.
*   **Voseo**: Mantené el tono Rioplatense respetuoso pero directo en la descripción si el usuario lo prefiere para documentación interna.
*   **Links**: Si hay un ticket o una skill relacionada, vinculalos.
