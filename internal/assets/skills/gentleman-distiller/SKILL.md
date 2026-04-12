---
name: gentleman-distiller
description: Escanea el proyecto en busca de nuevas convenciones, skills o cambios arquitectónicos y destila reglas actualizadas para AGENTS.md.
---

# ⚗️ Gentleman Rules Distiller

## Cuándo usar esta skill
- **Trigger**: Cuando notes que las reglas en `AGENTS.md` o `CLAUDE.md` ya no reflejan la realidad del proyecto (ej: cambiaste de base de datos, agregaste un framework de testing nuevo, o creaste una nueva convención modular).
- **Proactividad**: El agente debe sugerir una "destilación" después de implementar un cambio arquitectónico importante.

## Flujo de Trabajo
1. **Escaneo**: Leer el `AGENTS.md` actual y compararlo con la estructura de archivos (`/internal`, `/modules`, `/skills`).
2. **Consultar Memoria**: Revisar las últimas observaciones de `Engram` para capturar decisiones de diseño recientes.
3. **Draft de Reglas**: Proponer un conjunto de reglas refinado que elimine lo obsoleto y refuerce lo nuevo.
4. **Actualización**: Con el permiso del usuario, ejecutar el comando `gentleman distill --apply` para consolidar los cambios.

## Principios de Destilación
- **Concisión**: Si una regla se puede explicar con un ejemplo de código, hacelo.
- **Elite**: Mantener solo lo que hace al "Gentleman Stack". No meter ruido de librerías genéricas.
- **Voseo**: Las reglas deben hablarle al agente en Rioplatense, recordándole su rol de Orquestador y Senior Architect.
