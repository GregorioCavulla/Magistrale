import os
import re
from pathlib import Path

# 📁 Cartella principale degli appunti
BASE_DIR = Path("/path/alla/tua/cartella")  # <-- QUI devi mettere il tuo path!

# 🔍 Regex per trovare i titoli Markdown
HEADER_REGEX = re.compile(r'^(#{1,4})\s+(.*)', re.MULTILINE)

# 🗂️ Prende tutti i file markdown
md_files = list(BASE_DIR.rglob('*.md'))

# 📚 Mappa di tutti i titoli trovati { titolo_normalizzato : nome_file }
title_to_file = {}

# 1. Prima passata: indicizza tutti i titoli
for md_file in md_files:
    content = md_file.read_text(encoding='utf-8')
    for match in HEADER_REGEX.finditer(content):
        header_level, header_text = match.groups()
        header_key = header_text.strip().lower()
        title_to_file[header_key] = md_file.stem  # Nome file senza estensione

# 2. Seconda passata: sostituisce i titoli con i link
for md_file in md_files:
    content = md_file.read_text(encoding='utf-8')

    def replace_header(match):
        header_level, header_text = match.groups()
        header_key = header_text.strip().lower()

        # Non cambiare il titolo principale del file (presumiamo sia il primo # titolo)
        if header_level == '#' and match.start() < 10:
            return match.group(0)

        # Se troviamo un altro file con questo titolo
        if header_key in title_to_file and title_to_file[header_key] != md_file.stem:
            linked_text = f'[[{title_to_file[header_key]}]]'
            return f'{header_level} {linked_text}'
        else:
            return match.group(0)

    # Applica la sostituzione
    new_content = HEADER_REGEX.sub(replace_header, content)

    # Scrive il file modificato
    md_file.write_text(new_content, encoding='utf-8')

print("✅ Collegamenti automatici creati senza toccare i titoli principali!")
