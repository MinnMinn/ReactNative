import axios from 'axios';

const apiFoods = 'http://localhost:8080/product/';

export async function getFoods() {
    try {
        let response = await fetch(apiFoods + 'show');
        let responseJson = await response.json();
        return responseJson; //list of foods
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
    // axios.get(apiFoods + 'show')
    //   .then(res => {
    //     let responseJson = res.data;
    //     return responseJson;
    //   })
    //   .catch(error => console.log(error));
}

export async function getFoodById(id) {
    // try {
    //     let response = await fetch(apiFoods + 'show/' + id);
    //     let responseJson = await response.json();
    //     return responseJson;
    // } catch (error) {
    //     console.error(`Error is : ${error}`);
    // }
    axios.get(apiFoods + 'show/' + id)
      .then(res => {
        let responseJson = res.data;
        return responseJson;
      })
      .catch(error => console.log(error));
}

export async function deleteFoodById(id) {
    // try {
    //     let response = await fetch(apiFoods + 'delete/' + id, {
    //         method: 'DELETE'
    //     });
    //     let responseJson = await response.json();
    //     return responseJson;
    // } catch (error) {
    //     console.error(`Error is : ${error}`);
    // }
    axios.delete(apiFoods + 'delete/' + id)
      .then(res => {
        console.log(res.data);
      })
}
