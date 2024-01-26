function randomLetter() {
    var keysArray = Object.keys(draw_pile);

    if (keysArray.length > 0) {
    var randomKey = keysArray[Math.floor(Math.random() * keysArray.length)];
    draw_pile[randomKey] --;

    return randomKey
    } else {
    console.log('Error : no more letters');
    }
}

let letter = randomLetter()
console.log(letter)
console.log(draw_pile[letter])