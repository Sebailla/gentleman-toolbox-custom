# Gentle AI — Agent Skills Index

When working on this project, load the relevant skill(s) BEFORE writing any code.

## How to Use

1. Check the trigger column to find skills that match your current task
2. Load the skill by reading the SKILL.md file at the listed path
3. Follow ALL patterns and rules from the loaded skill
4. Multiple skills can apply simultaneously

## 🛡️ Protocolo de Actuación (MANDATORIO)
- **Rol del Orquestador**: El orquestador DEBE limitarse a guiar y mantener el hilo de la conversación. No debe escribir ni modificar código directamente. Toda acción técnica, incluyendo la edición de código, DEBE ser delegada a **subagentes**.
- **Cero Suposiciones**: Nunca te quedes con dudas ni infieras requerimientos, arquitectura o decisiones técnicas. Antes de tomar cualquier decisión, PREGUNTA todo lo necesario al usuario para tener el contexto completo y exacto.
- **Confirmación Constante**: Antes de delegar cambios significativos a un subagente, resume la acción y espera la confirmación del usuario.

## 🎨 Diseño UI/UX
- **Base de Diseño**: TODO lo relacionado con el diseño, UI, UX y aspectos visuales DEBE basarse estricta y únicamente en la información documentada en la carpeta `design-md`. No inventes, ni uses información externa, apégate al contenido de esa carpeta.

## Skills

| Skill | Trigger | Path |
|-------|---------|------|
| `gentle-ai-issue-creation` | When creating a GitHub issue, reporting a bug, or requesting a feature. | [`skills/issue-creation/SKILL.md`](skills/issue-creation/SKILL.md) |
| `gentle-ai-branch-pr` | When creating a pull request, opening a PR, or preparing changes for review. | [`skills/branch-pr/SKILL.md`](skills/branch-pr/SKILL.md) |
| `gentle-ai-modular-architecture` | When creating or refactoring domain modules, implementing business logic, or structuring project architecture. | [`skills/architecture/SKILL.md`](skills/architecture/SKILL.md) |
