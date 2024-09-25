"use strict";

export class SudokuClient {
    constructor() {
        // noinspection JSIgnoredPromiseFromCall
        this.Health();
    }

    async Health() {
        this.live = await (fetch("/health").then(response => response.ok));
    }
}