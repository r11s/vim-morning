# Morning (Vim Port)

A port of Vim’s classic **“morning.vim”** colorscheme for Visual Studio Code: a light grey background with black text and a small, calm palette.

------

## 🎨 About the Theme

`morning.vim` was written by **Bram Moolenaar**, the author of Vim, and ships with Vim. It is a light colorscheme with a grey background (`#e4e4e4`), blue comments and dark red, bold statements. This port keeps the colors of the GUI version (`guifg`/`guibg`) and maps Vim’s highlight groups to TextMate scopes and VS Code UI colors.

### ✳️ Highlights

| Group         | Color (hex)            | Description                  |
| ------------- | ---------------------- | ---------------------------- |
| **Normal**    | `#000000` on `#e4e4e4` | Black on light grey          |
| **Comment**   | `#0000ff`              | Blue                         |
| **Constant**  | `#ff00ff`              | Magenta (strings, numbers)   |
| **Statement** | `#a52a2a`              | Brown-red, bold              |
| **PreProc**   | `#6a0dad`              | Purple                       |
| **Type**      | `#2e8b57`              | Sea green, bold              |
| **Special**   | `#6a5acd`              | Slate blue                   |
| **Function**  | `#008787`              | Teal (links to `Identifier`) |

------

## 🎭 Variants

| Theme                          | Description                                 |
| ------------------------------ | ------------------------------------------- |
| **Morning (Vim Port)**         | The port, with bold keywords and types      |
| **Morning (Vim Port, no bold)**| Same colors, without bold text              |

The sidebar, activity bar and panel use slightly darker greys (`#d6d6d6`, `#cccccc`), because Morning itself only knows one background. Title bar and status bar share Vim’s `TitleBar` look (`#cccccc` with black text); tabs, list selection and widgets follow `TabLine`, `PmenuSel` and `Pmenu`.

Not 1:1:
- Vim gives `Constant` a slightly lighter background (`#eeeeee`). VS Code cannot color the background of single tokens in the editor.
- Vim’s `Search` is light text on purple. VS Code cannot switch the text color of a match, so matches are highlighted in yellow.
- The terminal uses the palette from the colorscheme (`g:terminal_ansi_colors`) as it is. It is unusual: ANSI “black” is the light background color.

------

## 🧩 Installation

Search for **Morning (Vim Port)** by **r11s** in the Extensions view, or install a local `.vsix` via **Extensions: Install from VSIX…**. Then select the theme under **Preferences: Color Theme**.

------

## ⚙️ Recommended Settings

```jsonc
{
  "editor.semanticHighlighting.enabled": true,
  "editor.bracketPairColorization.enabled": false,
  "editor.guides.bracketPairs": false,
  "editor.renderWhitespace": "none",
  "editor.cursorBlinking": "solid"
}
```

------

## ⚖️ License

Released under the [Vim License](https://github.com/vim/vim/blob/master/runtime/doc/uganda.txt).

This theme is a **port** of **“morning.vim”** (original author and maintainer: **Bram Moolenaar** <Bram@vim.org>; current version in the Vim colorschemes, https://github.com/vim/colorschemes, Last Change 2026 Oct 02). It is not the original. All modifications and adaptations are provided under the same terms as Vim itself. See `LICENSE`.

------

## 🛠 Maintainers

- **Ported by:** Rainer Strauss ([r11s](https://github.com/r11s))
- **Original morning.vim author:** Bram Moolenaar
- **Vim Project:** https://www.vim.org
