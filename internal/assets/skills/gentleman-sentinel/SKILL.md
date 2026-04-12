# Skill: Gentleman Sentinel (Git Hooks)

Esta skill regula el comportamiento del toolbox como un centinela activo del repositorio. Su objetivo es evitar que código que rompa las reglas o que degrade la salud del proyecto llegue a ser commiteado.

## El Centinela en Acción

El Centinela se manifiesta principalmente como un hook de `pre-commit`. 

1.  **Validación Pre-Commit**:
    *   Cada vez que se intenta un `git commit`, el Centinela ejecuta `gentle-ai doctor`.
    *   Si el proyecto no tiene `AGENTS.md` o si las réplicas están desactualizadas, el commit debería ser observado.
2.  **Mantenimiento de Ganchos**:
    *   Si el usuario te pide "activar el centinela", usá el comando `gentle-ai sentinel install`.
    *   Esta skill sabe distinguir entre un hook nativo de Git (`.git/hooks/pre-commit`) y Husky (`.husky/pre-commit`).

## Reglas para el Agente

*   **No forzar**: Si un commit falla por el centinela, no sugieras `--no-verify` a menos que sea una emergencia justificada. Primero intentá arreglar el problema con `gentle-ai doctor --fix`.
*   **Sincronización**: Después de actualizar reglas con `distill`, recordá que el centinela va a validar que esas reglas se propaguen a las réplicas.

## Uso del comando CLI

```bash
# Instalar o actualizar el centinela
gentle-ai sentinel install
```
