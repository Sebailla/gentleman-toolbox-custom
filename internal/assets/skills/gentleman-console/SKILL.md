# Skill: Gentleman Console (TUI Management)

Como arquitecto, necesitás un centro de mando. Esta skill guía la representación visual del búnker en la terminal usando Bubbletea y Lipgloss.

## Diseño de Interfaz (Layout)
- **Header**: Logo ASCII de Gentleman y el nombre del proyecto.
- **Side Panel**: Lista de módulos y su estado individual.
- **Main Panel**: Resumen del **Architecture Health Index (AHI)** con barras de progreso de color.
- **Bottom Bar**: Atajos de teclado (J: Judge, P: Plan, B: Briefing, Q: Quit).

## Estética (Lipgloss)
- **Primary**: Cyan (#00FFFF) para acentos.
- **Secondary**: Magenta (#FF00FF) para alertas.
- **Background**: Slate dark para un look premium.

## Interacciones
- Al presionar 'J', lanzá un `gentle-ai judge` sobre el módulo seleccionado.
- Al presionar 'S', forzá un `gentle-ai sync`.

## Reglas de Oro
- **No ensucies la pantalla**: Mantené la información densa pero legible.
- **Micro-animaciones**: Usá spinners cuando haya procesos de fondo (como el escaneo de Blueprint).
