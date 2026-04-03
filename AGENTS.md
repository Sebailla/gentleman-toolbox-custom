# Gentle AI — Agent Skills Index

When working on this project, load the relevant skill(s) BEFORE writing any code.

## How to Use

1. Check the trigger column to find skills that match your current task
2. Load the skill by reading the SKILL.md file at the listed path
3. Follow ALL patterns and rules from the loaded skill
4. Multiple skills can apply simultaneously

## 🛡️ Protocolo de Actuación (MANDATORIO)
- **Clarificar antes de actuar**: Ante cualquier duda, ambigüedad o falta de contexto, el agente DEBE detenerse y preguntar.
- **Confirmación de Entendimiento**: Antes de realizar cualquier cambio significativo, el agente debe resumir qué entendió del requerimiento y qué plan de acción propone, esperando la confirmación del usuario.
- **Sin Suposiciones**: Nunca asumas el stack, la arquitectura o el comportamiento deseado. Validá siempre.

## Skills

| Skill | Trigger | Path |
|-------|---------|------|
| `gentle-ai-issue-creation` | When creating a GitHub issue, reporting a bug, or requesting a feature. | [`skills/issue-creation/SKILL.md`](skills/issue-creation/SKILL.md) |
| `gentle-ai-branch-pr` | When creating a pull request, opening a PR, or preparing changes for review. | [`skills/branch-pr/SKILL.md`](skills/branch-pr/SKILL.md) |
