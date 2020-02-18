const apiGetAllFoods = 'http://localhost:8080/product/show';

async function getFoods() {
    try {
        let response = await fetch(apiGetAllFoods);
        let responseJson = await response.json();
        return responseJson; //list of foods
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
}

export default getFoods;