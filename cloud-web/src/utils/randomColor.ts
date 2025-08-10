const lightColors = [
  "#ADD8E6", // Light Blue
  "#90EE90", // Light Green
  "#F08080", // Light Coral
  "#FAFAD2"  // Light Goldenrod Yellow
];


const generateRandomHexColor = (infraType: string): string => {
  switch (infraType) {
    case 'compute':
      return lightColors[0]
    case 'Type2':
      return '#33FF57'; // Example color for Type2
    case 'Type3':
      return '#3357FF'; // Example color for Type3
    default:
      return '#FFFFFF'; // Default color
  }
}


export default generateRandomHexColor;
