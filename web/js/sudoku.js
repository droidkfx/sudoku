"use strict";

import {SudokuClient} from "./client.js";

const sudokuClient = new SudokuClient();
let selectedCell = null;
let notesEnabled = false;

function SetCell(x, y, value, locked) {
    let cell = document.getElementById("cell-" + x + "" + y);
    let cellValue = document.getElementById("cell-" + x + "" + y + "-value");
    let noteBlock = document.getElementById("cell-" + x + "" + y + "-notes");

    if (!cellValue) {
        return;
    }
    cellValue.innerText = value;
    if (locked) {
        cell.classList.add("locked");
    }

    noteBlock.classList.add("hidden");
}

function UnsetCell(x, y) {
    let cell = document.getElementById("cell-" + x + "" + y);
    let cellValue = document.getElementById("cell-" + x + "" + y + "-value");
    let noteBlock = document.getElementById("cell-" + x + "" + y + "-notes");

    cellValue.innerText = "";
    cell.classList.remove("locked");
    noteBlock.classList.remove("hidden");
}

function SetOption(x, y, option) {
    document
        .getElementById("cell-" + x + "" + y + "-note-" + option)
        .classList
        .remove("hidden");
}

function ToggleOption(x, y, option) {
    document
        .getElementById("cell-" + x + "" + y + "-note-" + option)
        .classList
        .toggle("hidden");
}

function UnsetOption(x, y, option) {
    document
        .getElementById("cell-" + x + "" + y + "-note-" + option)
        .classList
        .add("hidden");
}

function FillBoard(bdata) {
    for (let x = 0; x < 9; x++) {
        for (let y = 0; y < 9; y++) {
            let cell = bdata.board[y][x];
            if (cell === 0) {
                continue;
            }
            SetCell(x, y, cell, true);
        }
    }
}

async function HandleCellClick(cell) {
    if (selectedCell) {
        selectedCell.classList.remove("selected");
    }
    if (selectedCell === cell) {
        selectedCell = null;
        return;
    }
    cell.classList.add("selected");
    selectedCell = cell;
}

function getCellCord(cell) {
    let id = cell.id;
    return [parseInt(id[5]), parseInt(id[6])];
}

function HandleKeyDown(event) {
    if (!selectedCell) {
        return;
    }
    let cellCord = getCellCord(selectedCell);
    if (event.key >=1 && event.key <= 9) {
        if (notesEnabled) {
            ToggleOption(cellCord[0], cellCord[1], event.key);
        } else {
            SetCell(cellCord[0], cellCord[1], event.key, false);
        }
    } else if (event.key === "Backspace") {
        UnsetCell(cellCord[0], cellCord[1]);
    } else if (event.key === "Control") {
        notesEnabled = !notesEnabled;
    }
}

async function SetupBoard() {
    let elems = document.getElementsByClassName("sudoku-cell")
    Array.from(elems).forEach(cell => {
        cell.addEventListener("click", () => HandleCellClick(cell));
    })

    document.getRootNode().addEventListener("keydown", HandleKeyDown);

    FillBoard(await sudokuClient.GetBoard(0));
}

await SetupBoard();