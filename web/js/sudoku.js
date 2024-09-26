"use strict";

import {SudokuClient} from "./client.js";

const sudokuClient = new SudokuClient();

function SetCell(x, y, value, locked) {
    let cell = document.getElementById("cell-" + x + "" + y);
    let cellValue = document.getElementById("cell-" + x + "" + y + "-value");
    let noteBlock = document.getElementById("cell-" + x + "" + y + "-notes");

    if (!cellValue) {
        console.error("Cell not found: " + x + "" + y);
        console.error(cell- + x + "" + y + "-value");
        return;
    }
    cellValue.innerText = value;
    if (locked) {
        cell.classList.add("locked");
    }

    noteBlock.classList.add("hidden");
}

function SetOption(x, y, option) {
    document
        .getElementById("cell-" + x + "" + y + "-note-" + option)
        .classList
        .remove("hidden");
}

function UnsetOption(x, y, option) {
    document
        .getElementById("cell-" + x + "" + y + "-note-" + option)
        .classList
        .add("hidden");
}

async function FillBoard(number) {
    let bdata = await sudokuClient.GetBoard(number);
    for (let x = 0; x < 9; x++) {
        for (let y = 0; y < 9; y++) {
            let cell = bdata.board[y][x];
            SetCell(x, y, cell, true);
        }
    }
}

await FillBoard(0);