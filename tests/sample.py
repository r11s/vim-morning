#!/usr/bin/env python3
"""Prüfdatei für das Morning-Theme (Docstring)."""
# TODO: Kommentar mit Todo
from __future__ import annotations

import os
import re
from dataclasses import dataclass, field
from typing import Iterator

MAX_ITEMS = 16            # Zahl
RATIO = 3.14159
MASK = 0xFF
GREETING = f"Hallo {os.getenv('USER', 'welt')!r}\n"
PATTERN = re.compile(r"^\s*(\w+)\s*=\s*(.*)$")


@dataclass
class Item:
    name: str
    weight: float = 0.0
    tags: list[str] = field(default_factory=list)

    def label(self) -> str:
        return f"{self.name} ({self.weight:.1f})"


def parse(lines: Iterator[str]) -> dict[str, str]:
    result: dict[str, str] = {}
    for n, line in enumerate(lines, start=1):
        m = PATTERN.match(line)
        if not m or line.lstrip().startswith("#"):
            continue
        key, value = m.groups()
        result[key] = value if value is not None else ""
    return result


class Store:
    def __init__(self, items: list[Item] | None = None) -> None:
        self.items = items or []

    def heaviest(self) -> Item | None:
        return max(self.items, key=lambda i: i.weight, default=None)


if __name__ == "__main__":
    store = Store([Item("a", 1.5), Item("b", 2.0, ["x"])])
    try:
        print(store.heaviest().label(), True, None, 1 + 2 * 3 // 4)
    except AttributeError as exc:
        raise SystemExit(f"Fehler: {exc}") from exc
