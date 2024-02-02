// VARIABLES
var draw_pile = {
    "A": 14, "B": 4, "C": 7, "D": 5, "E": 19, "F": 2, "G": 4, "H": 2, "I": 11, "J": 1, "K": 1, "L": 6,
    "M": 5, "N": 9, "O": 8, "P": 4, "Q": 1, "R": 10, "S": 7, "T": 9, "U": 8, "V": 2, "W": 1, "X": 1, "Y": 2, "Z": 1
};


// FUNCTIONS
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

function drawLetters(hands, player_index, nbLetters) {
    hands[player_index] = hands[player_index].concat(Array.from({ length: nbLetters }, randomLetter));
};

function remove_letters_from_hand(hands, player_index, letters_to_remove){
    let chars = []

    console.log("HANDDDD : ", hands)

    hands[player_index] = Array.from((hands[player_index].flatMap(char => {

            let countInHand = count(hands[player_index], char);
            let countInToRemove = count(letters_to_remove, char);

            // nombre de lettres restantes après avoir retiré celles utilisées pour le mot
            let nbLetters = countInHand - countInToRemove;

            let charToAdd = (!chars.includes(char)
                    ? Array.from({ length: nbLetters }, () => char)
                    : Array.from({ length: 0 }, () => char)
            );

            chars.push(char);
            return charToAdd;
        })
    ));
    console.log("HANDDDD : ", hands)
}

function count(list, element){
    return list.reduce((count, list_element) => {
        return list_element === element ? count + 1 : count;
    }, 0);
}

function updateJarnacWithForLoop (turn, hands, original_word, new_word) {
    let letters = hands[(turn+1)%2];
    let difference = [];

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


function difference_between_and(long, short){
    let chars = [];
    return Array.from((long.flatMap(char => {

        let countNew = count(long, char);
        let countOriginal = count(short, char);

        // Calculate the difference in counts
        let countDifference = countNew - countOriginal;

        let charToAdd = (!chars.includes(char)
                ? Array.from({ length: countDifference }, () => char)
                : Array.from({ length: 0 }, () => char)
        );
        chars.push(char);
        return charToAdd;
    })));
}



function updateJarnac(grids, hands, line_number, turn, new_word) {
        console.log("Mot originel : ", grids[(turn+1)%2][line_number])
        let difference = difference_between_and(new_word, grids[(turn+1)%2][line_number])
        console.log("DIfférences de lettres : ", difference)
        console.log("Mains avant d'enlever les lettres manquantes :", hands)
        remove_letters_from_hand(hands, (turn + 1) % 2, difference)
        console.log("Mains après : " ,hands)

        if (line !== -1){
            grids[(turn + 1) % 2].splice(line, 1);
        }

        grids[turn % 2].push([...new_word]);

};


function exchangeLetters(hands, turn, letters) {
    let chars = [];

    hands[turn%2] = Array.from((hands[turn%2].flatMap( char => {
        let diff = count(hands[turn%2], char) - count(letters, char)

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

    drawLetters(hands, turn%2, 3);

};


// fonction pour mettre à jour toutes les variables de jeu après un placement de mot
// fonctionne meme pour une nouvelle ligne tant que le numéro de ligne est bon
// renvoie grids et hands mis à jour (meme fonctionnement que updateJarnac)
function updatePlay(grids, line, turn, hands, new_word) {


    let original_word = (grids[turn][line] !== undefined) ? grids[turn][line] : [];
    let letters = hands[turn];

    let difference = difference_between_and(new_word, original_word)

    remove_letters_from_hand(hands, turn%2, difference)

    if (line !== -1){
        grids[turn%2][line] = [...new_word];
    } else {
        grids[turn%2].push([...new_word])
    }
};

// TESTS


let grids = [[], [["B", "O", "N", "J", "O", "U"],["T", "E", "S", "T"]]]
let line = 0
let turn = 0
let hands = [[], ["R", "R", "G", "H"]]
let original_word = ["B", "O", "N", "J", "O", "U"]
let new_word = [..."BONJOUR"]



/*
y = updatePlay(
    grids = [[], []], 
    line = 0, turn = 1,
    hands = [[], [..."AJOUR"]],
    new_word = [..."JOUR"]);

console.log("grids :", y[0], "hands :", y[1]);
*/
// console.log(exchangeLetters(0, [["A", "A", "N", "R", "X", "L"],["B"]], ["A", "X", "L"]));

// let hands = [["A", "B", "C", "D", "E"], []]

// remove_letters_from_hand(hands, 1, ["R"])
// console.log(hands)
let original = ["B", "O", "N", "J", "O", "U"]
let word = ["B", "O", "N", "J", "O", "U", "N"]
console.log(difference_between_and(word, original))
console.log(original)
console.log(word)

module.exports = {exchangeLetters, updateJarnac, drawLetters, updatePlay };