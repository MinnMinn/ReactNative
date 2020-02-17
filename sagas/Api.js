const urlGetMovies = 'http://localhost:8080/user/show';

async function getMoviesFromApi() {
    try {
        let response = await fetch(urlGetMovies);
        let responseJson = await response.json();
        console.log(responseJson)
        return responseJson.data; //list of foods
    } catch (error) {
        console.error(`Error is : ${error}`);
    }
}

export default getMoviesFromApi;