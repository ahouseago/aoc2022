s = {
  "A X": 4,
  "A Y": 8,
  "A Z": 3,
  "B X": 1,
  "B Y": 5,
  "B Z": 9,
  "C X": 7,
  "C Y": 2,
  "C Z": 6,
};
o = {
  "A X": 3,
  "A Y": 4,
  "A Z": 8,
  "B X": 1,
  "B Y": 5,
  "B Z": 9,
  "C X": 2,
  "C Y": 6,
  "C Z": 7,
};
console.log(
  require("fs")
    .readFileSync("./input.txt", "utf8")
    .split("\n")
    .reduce((a, l) => (l ? [a[0] + s[l], a[1] + o[l]] : a), [0, 0])
);
