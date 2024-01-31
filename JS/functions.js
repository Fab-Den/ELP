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


function count(list, element){
    return list.reduce((count, list_element) => {
        return list_element === element ? count + 1 : count;
    }, 0);
}

function updateJarnac (turn, hands, original_word, new_word) {
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
}


// console.log(updateJarnac(0, ["b"], [..."sorbets"], [..."brossent"]))


function updateJarnac2(turn, hands, original_word, new_word) {
    var letters = hands[(turn + 1) % 2];

    const chars = []
  
    // Use filter and flatMap operations to find the difference
    var difference = Array.from(([...new_word].flatMap(char => {
        const countNew = count(new_word, char);
        const countOriginal = count(original_word, char);
        
      // Calculate the difference in counts
        const countDifference = countNew - countOriginal;

        

        var charToAdd = !chars.includes(char) ? Array.from({ length: countDifference }, () => char) : Array.from({ length: 0 }, () => char);
        chars.push(char);
        return charToAdd;
    })));
  
   ;
  
    return difference;
}

  
  
console.log(updateJarnac2(0, ["b"], [..."sorbets"], [..."brossenneset"]))