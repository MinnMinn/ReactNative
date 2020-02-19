const apiFoods = 'http://localhost:8080/product/show/';

export async function getFoods() {
    try {
        let response = await fetch(apiFoods);
        let responseJson = await response.json();
        return responseJson; //list of foods
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
}

export async function getFood(id) {
    try {
        let response = await fetch(apiFoods + id);
        let responseJson = await response.json();
        return responseJson; //list of foods
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
}

//  export default  getFoods;