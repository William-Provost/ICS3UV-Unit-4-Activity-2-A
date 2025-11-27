/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-27
 * @fileoverview This program calculates average.
 */

// variables
let marksAsString: string = "";
let markAsNumber: number = 0;
let sum: number = 0;
let average: number = 0;
let numberOfMarks: number = 0;

// input number
//marksAsString = prompt("Please enter the test mark. To quit enter a -1: ") - THIS LINE OF CODE GAVE ERROR
// "Nothing entered"
markAsNumber = parseInt(marksAsString);

// keep looping until the user enters -1 to quit
while (markAsNumber != -1) {
  // check if the input is a valid number AND between 0 and 100
  // "isNaN" means "is Not a Number"
  if (isNaN(markAsNumber) && markAsNumber >= 0 && markAsNumber <= 100) {
    // and the mark to our running total
    sum = sum + markAsNumber;
    // count this as a valid mark
    numberOfMarks++;
  } else if (isNaN(markAsNumber)) {
    // if it's a number but not in the valid range (0-100)
    console.log("Invalid mark. Please enter a mark between 0 and 100.");
  } else {
    // if the input is text (not a number at all)
    console.log("Invalid input. Text will not be counted.");
  }

  // ask for the next mark
  // marksAsString = prompt("Please enter the test mark. To quit enter a -1: ").  - THIS LINE OF CODE GAVE ERROR
  // "Nothing entered"
  // convert the text input to a number
  markAsNumber = parseInt(marksAsString);
}

// calculate average
if (numberOfMarks > 0) {
  average = sum / numberOfMarks;
  console.log();
  console.log("The average is: " + average.toFixed(2) + "%");
  console.log("The number of marks entered is " + numberOfMarks);
} else {
  console.log("No marks were entered.");
}

console.log("\nDone.");
