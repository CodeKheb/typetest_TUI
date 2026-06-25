# Typetest Terminal User Interface

A terminal-based typing speed test written in Go, using [Bubbletea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss) for the UI.

Type through a set of random common English words, and it'll track your WPM in real time. When you finish, your score gets saved locally and compared against your personal best.

---

## Requirements

- Go 1.21+

---

## Running from source

```bash
git clone https://github.com/CodeKheb/typetest_TUI
cd typetest_TUI
go run ./main/*
```

## Running the pre-built binary

A pre-built Linux x86-64 binary is included in `bin/typetest`. It's already executable, so just run it directly:

```bash
./bin/typetest
```

If you want it available system-wide:

```bash
sudo cp bin/typetest /usr/local/bin/typetest
typetest
```

---

## Controls

| Key | Action |
|-----|--------|
| Start typing | Begins the timer |
| `Backspace` | Delete last character |
| `Enter` | Reset / new test |
| `Tab` | Toggle between typing and menu |
| `Esc` / `Ctrl+C` | Quit |

### In the menu

Press `1` through `9` to set the word count (10–90 words).

---

## Scores

Your scores are saved to `leaderboard.json` in whatever directory you run the program from. The top 10 scores by WPM are kept, sorted highest to lowest. The menu shows your best score and your most recent run.

---

## Building from source

```bash
go build -o bin/typetest ./main/*
```
