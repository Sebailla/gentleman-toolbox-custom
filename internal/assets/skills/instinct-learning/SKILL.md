# Skill: Instinct Learning (Aprendizaje Continuo)

Esta skill te permite evolucionar en base a la interacción con el usuario. Tu objetivo es detectar patrones implícitos y convertirlos en guías explícitas (Instintos).

## Extracción de Instintos

Al final de una tarea importante o cuando el usuario use el comando `gentle-ai learn`:
1.  **Recuperar Memoria**: Buscá en Engram (`mem_context`) las últimas correcciones que te hizo el usuario ("no hagas esto", "preferí aquello").
2.  **Identificar Patrones**:
    *   ¿Corrigió mi estilo de código?
    *   ¿Me pidió un patrón arquitectónico específico varias veces?
    *   ¿Hay un error recurrente que siempre cometo?
3.  **Consolidar**: Creá una entrada corta y accionable.
    *   **Malo**: "Al usuario no le gusta el desorden".
    *   **Bueno**: "Instinto: Usar siempre punteros para receptores de struct en el paquete /internal/store para consistencia".

## Almacenamiento

Los instintos se guardan localmente en `.gentleman/instincts.md`. Este archivo es prioridad de lectura para el Orquestador al iniciar cada sesión.

## Protocolo de Mejora

1. Ejecutá `gentle-ai learn`.
2. Analizá el feedback recibido en la sesión actual.
3. Actualizá el archivo de instintos local.
4. Notificá al usuario: "He aprendido X instinto basado en nuestra interacción. Lo tendré en cuenta de ahora en más".
