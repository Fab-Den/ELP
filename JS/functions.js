let draw_pile = {
    "A": 14, "B": 4, "C": 7, "D": 5, "E": 19, "F": 2, "G": 4, "H": 2, "I": 11, "J": 1, "K": 1, "L": 6,
    "M": 5, "N": 9, "O": 8, "P": 4, "Q": 1, "R": 10, "S": 7, "T": 9, "U": 8, "V": 2, "W": 1, "X": 1, "Y": 2, "Z": 1
};


function randomLetter() {
    var keysArray = Object.keys(draw_pile);

    if (keysArray.length > 0) {
    var randomKey = keysArray[Math.floor(Math.random() * keysArray.length)];
    draw_pile[randomKey] --;

    return randomKey;
    } else {
    console.log('Error : no more letters');
    }
};

function drawLetters(turn, hands, nbLetters) {
    hands[turn] = hands[turn].concat(Array.from({ length: nbLetters }, randomLetter));   
    return hands;
};

function count(list, element){
    return list.reduce((count, list_element) => {
        return list_element === element ? count + 1 : count;
    }, 0);
}

function updateJarnacWithForLoop (turn, hands, original_word, new_word) {
    var letters = hands[(turn+1)%2];
    var difference = [];

    for (let i = 0; i <new_word.length; i++) {
        let char = new_word[i]
        if (!original_word.includes(char)) {
            difference.push(char)
        } else if (count(new_word, char) > count(original_word, char) && !difference.includes(char)) {
            difference.push(char)
        }     
    }
    return difference
};


function updateJarnac(turn, hands, original_word, new_word) {
    var letters = hands[(turn + 1) % 2];
    
    var chars = [];
    var difference = Array.from(([...new_word].flatMap(char => {
        const countNew = count(new_word, char);
        const countOriginal = count(original_word, char);
        
      // Calculate the difference in counts
        const countDifference = countNew - countOriginal;
        
        var charToAdd = (!chars.includes(char) 
            ? Array.from({ length: countDifference }, () => char) 
            : Array.from({ length: 0 }, () => char)
        );
        chars.push(char);
        return charToAdd;
    })));
   ;
    
    var chars = [];
    letters = Array.from((letters.flatMap(char => {
        const countLetters = count(letters, char);
        const countInDifference = count(difference, char);
        const nbLetters = countLetters - countInDifference;
        
        var charToAdd = (!chars.includes(char) 
            ? Array.from({ length: nbLetters }, () => char) 
            : Array.from({ length: 0 }, () => char)
        );
        chars.push(char);
        return charToAdd;
    })

    ));
    
    hands[(turn + 1) % 2] =  letters;
    return hands;
    
};

  
  
// console.log(updateJarnac(1, [["a", "n", "n", "n", "x"],["b"]], [..."sorbets"], [..."brossent"]));


function exchangeLetters(turn, hands, letters) {
    var playerLetters = hands[turn];
    
    var chars = [];

    var playerLetters = Array.from((playerLetters.flatMap( char => {
        var diff = count(playerLetters, char) - count(letters, char)


        if (!chars.includes(char)) {
            draw_pile[char] += count(letters, char)
            chars.push(char)
            return Array.from({ length: diff}, () => char)
        } else {
            chars.push(char)
            return Array.from({ length: 0 }, () => char)
        } 
    })
    ));
    hands[turn] = playerLetters;
    drawLetters(turn, hands, 3);

    return hands;
};


console.log(exchangeLetters(0, [["A", "A", "N", "R", "X", "L"],["B"]], ["A", "A", "X"]));
