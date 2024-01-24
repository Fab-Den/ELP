module Http_request exposing (..)

import Browser
import Html exposing (..)
import Html.Events exposing (onClick, onInput)
import Json.Decode exposing (Decoder, map2, field, string, at)
import Html.Attributes exposing (..)
import Http
import Array
import List
import Random
import Platform.Cmd as Cmd

main =
    Browser.element
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view
        }


type State
    = Failure
    | Loading
    | Success

type alias Model =
    { state : State
    , definition : List Definition
    , word_list : List String
    , current_word : String
    , random_index : Int
    , input : String
    , checked : Bool
    }

type alias Definition =
    { partOfSpeech : String
    , definitions : List String
    }

init_model = {state = Success
    , definition = []
    , word_list = ["hello", "blue", "submarine", "forest","computer"]
    , current_word = ""
    , random_index = 0
    , input = ""
    , checked = False}

init : () -> (Model, Cmd Msg)
init _ =
    (init_model
    , Http.get
          { url = "../static/words.txt"--"https://perso.liris.cnrs.fr/tristan.roussillon/GuessIt/thousand_words_things_explainer.txt"
          , expect = Http.expectString GotText
          }
    )

roll: Int -> Random.Generator Int
roll a =
  Random.int 0 a

-- UPDATE
type Msg
    = MorePlease
    | GotWordInfo (Result Http.Error (List Definition))
    | Roll
    | GetElementFromList Int
    | GotText (Result Http.Error String)
    | Change String
    | ToggleCheck

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
    case msg of
        MorePlease ->
            let
                newModel = { model | state = Loading}
            in
            (newModel, getDefinition newModel)
        GotWordInfo result ->
            case result of
                Ok definitions ->
                    ({model | definition = definitions, state = Success}, Cmd.none)

                Err _ ->
                    ({model | state = Failure}, Cmd.none)
        Roll ->
            ( model
            , Random.generate GetElementFromList (roll (List.length model.word_list))
            )        
        GetElementFromList index ->
            (
            case model.word_list |> Array.fromList |> Array.get index of
                Maybe.Just a ->
                    let
                        newModel = {model | current_word = a}
                    in
                    (newModel, getDefinition newModel)
                _ -> (model, Cmd.none)
            )
        GotText result ->
            case result of
                Ok fullText ->
                    ( {model|word_list = String.split " " fullText, state = Success} 
                    , Random.generate GetElementFromList (roll (List.length model.word_list)))

                Err _ ->
                    (model, Cmd.none)
        Change newInput -> 
            ({model | input = newInput}, Cmd.none)
        ToggleCheck ->
            ({model | checked = not(model.checked)}, Cmd.none)




-- SUBSCRIPTIONS

subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none

-- VIEW

view : Model -> Html Msg
view model =
    div []
        [ h2 [] [ text "Guess the word" ]
        , viewQuote model
        ]


viewQuote : Model -> Html Msg
viewQuote model =
    case model.state of
        Failure ->
            div []
                [ text "I could not load the word for some reason. "
                , button [ onClick MorePlease ] [ text "Try Again!" ]
                ]

        Loading ->
            text "Loading..."

        Success  -> 
            div []
            [ displayWord model
            , div [] (List.map displayDefinition model.definition)
            , checkGuess model
            , input [ placeholder "Guess the word!", value model.input, onInput Change] []
            , button [ onClick Roll ] [ text "Change word" ]
            , input [type_ "checkbox", checked model.checked, onClick ToggleCheck] [], text "Show answer"
            ]


checkGuess : Model -> Html Msg
checkGuess model =
    if model.input == model.current_word then
        div [ style "color" "green"] [ text ("You found the word \""++model.current_word++"\" !")]
    else if model.input /= "" then
        div [ style "color" "red"] [ text "Not quite there yet"]
    else
        div [] []


displayWord : Model -> Html Msg
displayWord model =
    if model.checked then
        h3 [] [ text (model.current_word)]
    else
        h3 [] []


displayDefinition : Definition -> Html Msg
displayDefinition definition =
    div []
        [ h3 [] [ text definition.partOfSpeech ]
        , ul [] (List.map displaySingleDefinition definition.definitions)
        ]

displaySingleDefinition : String -> Html Msg
displaySingleDefinition singleDefinition =
    li [] [ text singleDefinition ]


--GET DEF
getDefinition : Model -> Cmd Msg
getDefinition model =
    Http.get
        { url = "https://api.dictionaryapi.dev/api/v2/entries/en/"++model.current_word
        , expect = Http.expectJson GotWordInfo apiDecoder
        }

sentenceDecoder : Decoder String
sentenceDecoder =
    field "definition" string

definitionDecoder : Decoder Definition
definitionDecoder =
    map2 Definition
        (field "partOfSpeech" string)
        (field "definitions" (Json.Decode.list sentenceDecoder))

apiDecoder : Decoder (List Definition)
apiDecoder =
    at ["0", "meanings"](Json.Decode.list definitionDecoder)






