# Skill: Fleet Sync (Ecosystem Updates)

Tus búnkeres no deben quedar obsoletos. Esta skill permite que un proyecto viejo se actualice con las últimas tácticas del Cuartel General de Gentleman.

## Proceso de Sincronización
1.  **Detección de Versión**: Comparamos el hash de las skills locales contra las del toolbox instalado globalmente.
2.  **Actualización de Assets**: Copiamos los archivos `SKILL.md` actualizados a la carpeta `.agent/skills/`.
3.  **Reparación de Reglas Maestras**: Si `AGENTS.md` o `CLAUDE.md` tienen nuevas secciones estructurales, las inyectamos sin borrar el contenido del usuario.

## Reglas de Oro
- **No pises nada**: Si el usuario modificó una regla, respetá su cambio. Solo agregá lo nuevo.
- **Backup**: Si vas a hacer un cambio estructural grande en las reglas, dejá un backup (`AGENTS.md.bak`).
- **Tono**: *"Che, te actualicé las tácticas del búnker. Ahora los reviewers son más picantes y tenés soporte para el nuevo patrón que sacamos ayer."*

## Auditoría Final
Listá qué archivos fueron actualizados y qué nuevas capacidades fueron desbloqueadas.
