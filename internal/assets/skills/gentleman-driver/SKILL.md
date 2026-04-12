# Skill: Gentleman Driver (Autonomous Loops)

Esta skill te permite operar en "Modo Conducción Autónoma". Tu objetivo es resolver bugs o fallas de tests de forma iterativa hasta lograr el éxito.

## El Ciclo de Conducción (Loop)

Cuando encuentres una falla en un test o un error de compilación:
1.  **Analizá**: Leé el output de `gentle-ai drive`. Identificá el archivo y la línea exacta del error.
2.  **Corregí**: Aplicá el fix más probable. No intentes correcciones masivas; hacé cambios atómicos.
3.  **Verificá**: Ejecutá `gentle-ai drive` nuevamente con el mismo comando de test.
4.  **Iterá**: Si vuelve a fallar, analizá el *nuevo* error (puede haber cambiado). Repetí hasta que pase.

## Reglas de Oro

*   **Límite de Bucle**: Si después de 3 intentos no lográs pasar el test, **DETENTE**. Explicá al usuario por qué estás trabado y pedí una pista. No quemes tokens en un loop infinito.
*   **Contexto Fresco**: Antes de cada intento, revisá si cambiaste algo que invalide tu suposición anterior.
*   **Rollback Manual**: Si el fix empeora las cosas (más tests fallan), revertí y probá otro camino.

## Uso del comando CLI

Ejemplo de flujo:
```bash
# Correr el driver para un paquete específico
gentle-ai drive --test="go test ./internal/auth/..." --max-loops=3
```
Respetá siempre los flags de `--test` para no correr toda la suite si solo falló un archivo.
