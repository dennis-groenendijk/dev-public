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

type Color = typeof Colors[number];

function colorCode(color: Color): number {
  return Colors.indexOf(color);
}

export function decodedResistorValue([color1, color2, color3]: Color[]): string {
  let value = (colorCode(color1) * 10 + colorCode(color2)) * Math.pow(10, colorCode(color3));
  return value < 1000 ? value.toString() + " ohms" : (value / 1000).toString() + " kiloohms";
}
