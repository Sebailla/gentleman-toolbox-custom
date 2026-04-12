---
name: context-guardian
description: Monitorea la salud del contexto del agente y sugiere limpiezas proactivas (mem_save, session_summary) para evitar el "context bloat".
---

# 🛡️ Gentleman Context Guardian

## Cuándo usar esta skill
- **Trigger**: Cuando sientas que la conversación se está volviendo lenta, repetitiva o que el contexto está cerca de su límite (comúnmente después de 15-20 turnos de interacción profunda).
- **Proactividad**: El "instinto" del Guardián debe activarse si detectas que estás volviendo a explicar conceptos que ya fueron decididos.

## Flujo de Trabajo
1. **Detección de Presión**: Evaluar si el buffer de mensajes contiene mucha información redundante o código obsoleto.
2. **Snapshot de Memoria**: Ejecutar `mem_save` para capturar las decisiones de arquitectura, hallazgos técnicos y convenciones acordadas en la sesión actual.
3. **Compactación**: Sugerir al usuario ejecutar un `/session-summary` (o hacerlo proactivamente) para limpiar el contexto y empezar con un buffer fresco basado en la memoria persistente.

## Reglas de Oro
- **No pierdas nada**: Antes de compactar, asegurate de que TODO lo importante esté en Engram.
- **Eficiencia**: Menos contexto = respuestas más rápidas y precisas. 
- **Seniority**: Un Senior Architect sabe cuándo dejar de hablar y empezar a resumir.
