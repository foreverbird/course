import {SELECT_OWEN_BIRD, GET_OWEN_BIRD} from '../actions/menu';

const selectOwenBird = (state={birdDetail:{}}, action) => {
    switch(action.type){
        case SELECT_OWEN_BIRD:
            return action.playload.text;
        case GET_OWEN_BIRD:
            return {
            ...state,
            birdDetail:action.playload.birdDetail}
        default:
            return state;
    }
}

export default selectOwenBird