totals = require("fs")
  .readFileSync("./input.txt", "utf8")
  .split("\n")
  .map(Number)
  .reduce((a, e) => (e ? a.concat(a.pop() + e) : [...a, 0]), [0])
  .sort()
  .reverse();
console.log(
  totals[0],
  totals.slice(0, 3).reduce((t, n) => t + n)
);
