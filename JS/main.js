const { exchangeLetters, updateJarnac,  drawLetters, updatePlay , count} = require('./functions');


const readline = require('readline');

const readline_interface = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});


let grids = [[], []]

let hands= [[], []]

let turn = 0;

let end= false

let number_jarnac = 0


function start_game(){
    // grids = [[], []]
    // hands = [[], []]
    turn = 0
    drawLetters(hands, 0, 6)
    drawLetters(hands, 1, 6)
    change_turn()
}

function end_game(){
    let points_player_1 = grids[0].reduce((partial_sum, element) => partial_sum + (element.length * element.length), 0)
    let points_player_2 = grids[1].reduce((partial_sum, element) => partial_sum + (element.length * element.length), 0)

    console.log("Player 1 got ", points_player_1)
    console.log("Player 2 got ", points_player_2)

    if (points_player_1 > points_player_2){
        console.log("Player 1 Won !")
    } else if (points_player_1 < points_player_2){
        console.log("Player 2 Won !")
    } else {
        console.log("Draw !")
    }

    readline_interface.close()
}

function step1(){
    // jarnac ou pas
    if (turn !== 0){
        step1_test().then(step2)

    } else {
        step2()
    }

}


function step2(){
    // tirer 1 lettre ou echanger

    if (turn > 1){
        step2_test().then(step3)
    } else {

        step3()
    }

}


function step3(){
    // jouer mot
    step3_chose_do().then(() => {turn += 1; change_turn()})

}

function change_turn(){

    if (grids[turn%2].length !== 8 && grids[(turn+1)%2].length !== 8){
        console.log('\033[2J')
        console.log("######## Turn of player " + ((turn%2)+1).toString() + " ########")

        step1()
    } else {
        // faire des choses pour compter le score + affichage gagnant
        end_game()
    }


}

function step3_chose_do(){
    return new Promise((resolve, reject) => {
        display_grid(turn%2)
        display_letters(turn%2)


        readline_interface.question("Do you want to input a word ? (yes/no) ", (str) => {
            if (str === "yes"){
                step3_chose_line().then(resolve)
            } else if (str === "no"){
                resolve()
            } else {
                console.log("Bad input")

                step3_chose_do().then(resolve)
            }
        })
    })
}


function step3_chose_line(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Chose a line (number / no) : ", (str) => {
            if (str === ""){
                step3_chose_do().then(resolve)
            } else if (!isNaN(Number(str)) && Number.isInteger(Number(str))){
                if (0 <= Number(str) && Number(str) < grids[turn%2].length){
                    step3_input_word(Number(str)).then(resolve)
                } else {
                    console.log("Bad line")
                    step3_chose_line().then(resolve)
                }
            } else if (str === "no"){
                if (grids[turn%2].length !== 8){
                    step3_input_word(-1).then(resolve)
                }
                else {
                    console.log("There is no more line on which we can add a word")
                    step3_chose_line().then(resolve)
                }

            } else {
                console.log("Bad input")
                step3_chose_line().then(resolve)
            }
        })
    })
}

function step3_input_word(line_number){
    return new Promise((resolve, reject) => {

        let hand_letters = hands[turn%2]
        let grid_letters = ( line_number !== -1 ? grids[turn%2][line_number] : [] )

        readline_interface.question("Chose a word : ", (str) => {

            if (str === ""){
                step3_chose_line().then(resolve)
            } else if (str.length < 3 || str.length > 9){
                console.log("Bad word length")
                step3_input_word(line_number).then(resolve)
            } else if (can_write_with([...str], hand_letters.concat(grid_letters)) && can_write_with(grid_letters, [...str]) && str.length > grid_letters.length){


                updatePlay(grids, line_number, turn, hands, [...str])

                step3_chose_do().then(resolve)
            } else {
                console.log("Bad word")
                step3_input_word(line_number).then(resolve)
            }
        })
    })
}



function step1_test(){
    return new Promise((resolve, reject) => {
        if (grids[turn%2].length < 8){
            display_grid((turn+1)%2)
            display_letters((turn+1)%2)
            step1_chose_jarnac().then(resolve)
        } else {
            resolve()
        }
    })
}

function step1_chose_jarnac(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Jarnac (yes / no) : ", (str) => {

            if (str === "yes"){


                step1_chose_line().then(resolve)

            } else if (str === "no"){
                resolve()

            } else {
                console.log("Bad input")
                step1_chose_jarnac().then(resolve)
            }
        })
    })
}



function step1_chose_line(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Chose a line (number / no) : ", (str) => {
            if (str === ""){
                step1_chose_jarnac().then(resolve)

            } else if (!isNaN(Number(str)) && Number.isInteger(Number(str))){

                if (0 <= Number(str) && Number(str) < grids[(turn+1)%2].length){
                    step1_input_word(hands[(turn+1)%2], grids[(turn+1)%2][Number(str)], Number(str)).then(resolve)
                } else {
                    console.log("Bad line")
                    step1_chose_line().then(resolve)
                }
            } else if (str === "no"){
                step1_input_word(hands[(turn+1)%2], []).then(resolve)

            } else {
                console.log("Bad input")
                step1_chose_line().then(resolve)
            }
        })
    })
}



function step1_input_word(hand_letters, grid_letters, line_number){

    return new Promise((resolve, reject) => {


        readline_interface.question("Input a word : ", (str) => {

            if (str === ""){
                step1_chose_line().then(resolve)
            } else {
                if (str.length > grid_letters.length && can_write_with(grid_letters, [...str]) && can_write_with([...str], grid_letters.concat(hand_letters)) && str.length >= 3 && str.length <= 9){
                    console.log("Mot valide")

                    updateJarnac(grids, hands, line_number, turn, [...str])

                    step1_test().then(resolve)

                } else {
                    console.log("Mot non valide")
                    step1_input_word(hand_letters, grid_letters).then(resolve)
                }
            }

        })
    })
}




function can_write_with(list1, list2){
    for (let i = 0; i < list1.length; i++){
        if (count(list1, list1[i]) > count(list2, list1[i])){
            return false
        }
    }
    return true
}



function step2_test(){
    return new Promise((resolve, reject) => {
        if (hands[turn%2].length >= 3){
            step2_chose_exchange().then(resolve)
        } else {
            drawLetters(hands, turn%2, 1)
            resolve()
        }

    })
}

function step2_chose_exchange(){
    return new Promise((resolve, reject) => {
        display_letters(turn%2)
        readline_interface.question("Exchange 3 letters instead of drawing ? (yes/no) ", (str) => {
            if (str === "yes"){
                step2_chose_letters().then(resolve)

            } else if (str === "no"){
                drawLetters(hands, turn%2, 1)
                resolve()
            }
            else {
                step2_chose_exchange().then(resolve)
            }
        })

    })
}

function step2_chose_letters(){
    return new Promise((resolve, reject) => {
        readline_interface.question("Chose 3 letters to exchange : ", (str) => {
            if (str === ""){
                step2_chose_exchange().then(resolve)

            } else if (str.length !== 3){
                console.log("You must select exactly 3 letters (no spaces)")
                step2_chose_letters().then(resolve)

            } else {
                if (can_write_with([...str], hands[turn%2])){
                    exchangeLetters(hands, turn, [...str])
                    resolve()
                } else {
                    console.log("Invalid letters")
                    step2_chose_letters().then(resolve)
                }

            }
        })
    })
}


function display_grid(index){

    let border_top = "┌" + function (){let s = ""; for (let i=0; i<8; i++){s+= "─┬"} return s}() + "─┐"
    let border_in = "├" + function (){let s = ""; for (let i=0; i<8; i++){s+= "─┼"} return s}() + "─┤"
    let border_bottom = "└" + function (){let s = ""; for (let i=0; i<8; i++){s+= "─┴"} return s}() + "─┘"

    let text = "GRID PLAYER " + (index%2+1).toString() + " : \n"
    text += border_top + "\n"
    for (let i = 0; i<8; i++){
        text += "│"
        if (i < grids[index].length){
            for (let j = 0; j < 9; j++){
                if (j<grids[index][i].length){
                    text += grids[index][i][j] + "|"
                } else {
                    text += " |"
                }

            }
        } else {
            for (let j = 0; j < 9; j++){
                text += " |"
            }
        }

        text += "\n" + (i !== 7 ? border_in : border_bottom) + "\n"

    }
    console.log(text)
}

function display_letters(index){
    let text = "LETTERS OF PLAYER " + (index%2+1).toString() + " : "

    for (let i = 0; i < hands[index].length; i++){
        text += hands[index][i] + " "
    }
    console.log(text)
}

start_game()
