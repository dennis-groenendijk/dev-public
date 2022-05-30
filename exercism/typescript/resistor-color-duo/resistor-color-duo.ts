const Colors = [
    "black",
    "brown",
    "red",
    "orange",
    "yellow",
    "green",
    "blue",
    "violet",
    "grey",
    "white",
];

export function decodedValue([color1, color2]: Array<string>): number {
  return Number(`${Colors.indexOf(color1)}${Colors.indexOf(color2)}`)
}
