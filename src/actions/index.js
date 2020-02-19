import * as types from '../constants/actionTypes';
import getFoods from '../sagas/Api';

export const listAll = (foods) => {
    return {
        type: types.LIST_ALL,
        foods
    };
}

export function getFoodById(food) {
    return { 
        type: types.FOOD_DETAIL,
        food
    };
}

// export function fetchError() {
//     return { type: 'FETCH_ERROR' };
// }

// export function fetchDataThunk() {
//     return dispatch => {
//         dispatch(fetchSuccess(foods));
//         getFoods()
//         .then(res => fetchSuccess(res))
//     };
// }