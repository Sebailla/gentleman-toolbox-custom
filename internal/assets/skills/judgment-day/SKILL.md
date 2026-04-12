# Skill: Judgment Day (Adversarial Review)

Este es el nivel máximo de revisión. No se aprueba nada hasta que dos jueces independientes den el visto bueno.

## Proceso de Juicio
1.  **Lanzamiento de Jueces**: El orquestador lanza dos sub-agentes independientes.
2.  **Revisión Ciega**: Cada juez analiza el código sin saber qué dijo el otro.
3.  **Síntesis del Veredicto**: Se comparan los hallazgos. Los problemas confirmados por ambos se marcan como prioritarios.

## Severidad
- **CRITICAL**: Error lógico, fallo de seguridad o violación grave de arquitectura.
- **WARNING (real)**: Bug probable o comportamiento inesperado.
- **WARNING (theoretical)**: Caso borde casi imposible de triggerear.

## Cómo Usar
Invocá este proceso cuando el usuario pida `gentle-ai judge`. Tu rol como orquestador es coordinar a los jueces y presentar el veredicto final. No hagas la revisión vos mismo: delegá.

## Mensaje Rioplatense
- *"Che, mirá, los jueces se pusieron la gorra. Coinciden en que este flujo tiene un race condition. Hay que arreglarlo sí o sí antes de subir."*
