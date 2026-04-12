# Skill: Instinct Refactor (Continuous Evolution)

Esta skill permite que el proyecto evolucione orgánicamente. No permitas que el código viejo se pudra.

## Cuándo usar esta skill
- **Trigger**: Cuando el usuario ejecuta `gentle-ai refactor --instincts`.
- **Contexto**: Tenés que haber leído previamente el archivo `.gentleman/instincts.md`.

## Proceso de Refactorización
1.  **Carga de Instintos**: Leé las reglas grabadas por el usuario en sesiones previas (Naming, Types, Patterns).
2.  **Identificación de Candidatos**: Buscá archivos que no sigan estas reglas.
3.  **Aplicación Quirúrgica**: Realizá los cambios sin romper la funcionalidad. Usá tests si están disponibles (`gentle-ai drive`).

## Reglas de Oro
- **No rompas nada**: Si no estás seguro de un cambio, pedí aclaración.
- **Minimalismo**: Aplicá los instintos pero no hagas refactors masivos de lógica de negocio a menos que sea necesario.
- **Tono**: Explicá el cambio: *"Che, vi que en este archivo estábamos usando ANY, así que lo tipé siguiendo tu instinto de 'Strict Types' que grabamos ayer. Quedó pipi cucu."*

## Auditoría
Al terminar, generá un mini-reporte de cuántos archivos fueron "purificados" por el instinto.
