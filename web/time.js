const systemTimeSpan = document.getElementById("system-time");
const cwYearSpan = document.getElementById("cw-year");

function updateTime() {
    systemTimeSpan.textContent = new Date().toLocaleTimeString();
}

function updateCwYear() {
    const date = new Date();
    const year = date.getFullYear();
    const month = date.getMonth().toString().padStart(2, "0");
    cwYearSpan.textContent = `${month}/${year}`;
}

updateTime();
updateCwYear();
setInterval(updateTime, 1000);