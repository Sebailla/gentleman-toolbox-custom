# Skill: Gentleman Briefing (Mission Handover)

Esta skill permite empaquetar el estado mental y técnico del proyecto para que otro agente de IA (o vos mismo en el futuro) sepa exactamente dónde quedó la pelota.

## Cuándo usar esta skill
- **Trigger**: Cuando el usuario dice "prepará el traspaso", "armá el briefing" o cuando estás por terminar una sesión larga de trabajo.

## Proceso de Generación

Un buen briefing debe ser denso en información pero fácil de "digerir" por un LLM.

1.  **Estado de Guerra (Git)**: Resumí en qué rama estamos y qué hay en el stage.
2.  **Hitos Recientes (Engram)**: Leé las últimas observaciones. ¿Qué decisiones críticas tomamos? ¿Qué bugs "malditos" arreglamos?
3.  **Pendientes Inmediatos**: ¿Qué es lo siguiente que hay que picar? (Lista de TODOs técnicos).
4.  **Gotchas Arquitectónicos**: Si hay una parte del código que es "especial" o frágil, avisalo.

## Estructura del BRIEFING.md

```markdown
# 🎖️ Mission Briefing: [Nombre del Proyecto]
**Fecha/ID**: [YYYY-MM-DD-HHMM]

## 📍 Estado Actual
- **Rama**: `feat/xyz`
- **Contexto**: Estamos a mitad de la implementación de...

## 🧠 Decisiones Críticas (Last Observations)
- Se eligió X en lugar de Y por...
- Se arregló el bug Z usando...

## 🚧 Pendientes (Next Steps)
- [ ] Implementar el servicio A.
- [ ] Correr tests de integración en B.

## ⚠️ Advertencias para el siguiente Agente
- Ojo con el archivo `core/utils.ts`, tiene un problema de...
- Usar siempre el preset `full-gentleman`.
```

## Reglas de Oro

*   **Sin Relleno**: La IA no necesita que seas amable en un briefing. Sé crudo y técnico.
*   **Persistencia**: El briefing se guarda como un archivo temporal para que el siguiente agente lo tome al iniciar.
