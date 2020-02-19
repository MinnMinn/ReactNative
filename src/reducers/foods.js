import * as types from '../constants/actionTypes';

var initalState = [];

const foodReducer = (state = initalState, action) => {
    switch (action.type) {
        case types.LIST_ALL:
            state = [...action.foods];
            return [...state];
        // case FETCH_FAILED:
        //     return [];
        // case ADD_MOVIE:
        //     return [
        //         ...state,
        //         action.newMovie
        //     ];
        default:
            return [...state]; //state does not change
    }
}

export default foodReducer;