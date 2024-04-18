# NRW-Lehrer Durchschnittliche Wochenstunden pro Halbjahr berechnen
Dieses Programm berechnet die durchschnittliche Anzahl von Wochenstunden pro Halbjahr für Lehrer in Nordrhein-Westfalen anhand eines Excel-Exports aus Schild-NRW.

## Voraussetzungen
- (Golang)[https://golang.org/] installiert

## Installation
1. Dieses Repository klonen
2. In das Verzeichnis des Repositories wechseln
3. `go build` ausführen
4. Das Programm mit `./stunden-berechner <date> <input_file>` ausführen, wobei `<date>` das Datum des Halbjahresendes im Format `dd.mm.` und `<input_file>` der Pfad zur Excel-Datei ist.