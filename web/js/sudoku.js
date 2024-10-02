"use strict";

import {SudokuClient} from "./client.js";

const sudokuClient = new SudokuClient();
let selectedCell = null;
let notesEnabled = false;
let boardStartTime = null;
let nextBoardToPlay = 0;
let boardInPlay = 0;

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
    document.getElementById("sudoku-board-number").value = bdata.id;
    document.getElementById("sudoku-board-difficulty").innerText = bdata.difficulty;

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

function toggleNotes() {
    let notesBox = document.getElementById("sudoku-board-notes");

    notesEnabled = !notesEnabled;
    notesBox.checked = notesEnabled;
}

function HandleKeyDown(event) {
    if (selectedCell) {
        let cellCord = getCellCord(selectedCell);
        if (event.key >=1 && event.key <= 9) {
            if (notesEnabled) {
                ToggleOption(cellCord[0], cellCord[1], event.key);
            } else {
                SetCell(cellCord[0], cellCord[1], event.key, false);
            }
        } else if (event.key === "Backspace") {
            UnsetCell(cellCord[0], cellCord[1]);
        }
    }

    if (event.key === "Control") {
        toggleNotes();
    }
}

function formatSeconds(duration) {
    // Hours, minutes and seconds
    const hrs = ~~(duration / 3600);
    const mins = ~~((duration % 3600) / 60);
    const secs = ~~duration % 60;

    // Output like "1:01" or "4:03:59" or "123:03:59"
    let ret = "";

    if (hrs > 0) {
        ret += "" + hrs + ":" + (mins < 10 ? "0" : "");
    }

    ret += "" + mins + ":" + (secs < 10 ? "0" : "");
    ret += "" + secs;

    return ret;
}

function updateTime() {
    if(!boardStartTime) {
        return;
    }

    let timeContainer = document.getElementById("sudoku-board-time");
    let seconds = (new Date() - boardStartTime) / 1000;
    timeContainer.innerText = formatSeconds(seconds);
}

function handleBoardNumberChange() {
    nextBoardToPlay = Number(document.getElementById("sudoku-board-number").value);
    document.getElementById("sudoku-new-board-number-button").disabled = (nextBoardToPlay === boardInPlay);
}

async function handleGoButton() {
    document.getElementById("sudoku-new-board-number-button").disabled = true;
    await StartBoard(nextBoardToPlay);
}

function SetupGame() {
    let elems = document.getElementsByClassName("sudoku-cell")
    Array.from(elems).forEach(cell => {
        cell.addEventListener("click", () => HandleCellClick(cell));
    })

    document.getRootNode().addEventListener("keydown", HandleKeyDown);
    document.getElementById("sudoku-board-notes").addEventListener("change", toggleNotes);
    document.getElementById("sudoku-board-number").addEventListener("change", handleBoardNumberChange);
    document.getElementById("sudoku-new-board-number-button").addEventListener("click", handleGoButton);

    setInterval(updateTime, 1000);
}

async function StartBoard(num) {
    boardStartTime = new Date();
    updateTime();

    boardInPlay = num;
    nextBoardToPlay = num
    let bdata = null;
    if (num >= 0) {
        bdata = await sudokuClient.GetBoard(num);
    } else {
        bdata = await sudokuClient.GetRandomBoard();
    }
    FillBoard(bdata);
}

SetupGame();
await StartBoard(-1);