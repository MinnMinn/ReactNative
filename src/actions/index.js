import * as types from '../constants/actionTypes';

export const listFoods = (foods) => {
    return {
        type: types.LIST_ALL,
        foods
    };
}

export function getFood(food) {
    return { 
        type: types.FOOD_DETAIL,
        food
    };
}

export function deleteFood(food) {
    return { 
        type: types.FOOD_DETAIL,
        food
    };
}