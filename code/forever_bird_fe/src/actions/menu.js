import fetchJsonp from 'fetch-jsonp';
export const    SELECT_OWEN_BIRD = 'SELECT_OWEN_BIRD';
export const    GET_OWEN_BIRD = 'GET_OWEN_BIRD';

export const selectOwenBird = (userName) => ({
    type : SELECT_OWEN_BIRD,
    payload : {
        text:userName
    }
})

export const getOwenBird = (userName,data) => ({
    type : GET_OWEN_BIRD,
    payload : {
        text:userName,
        birdDetail:data,
    }
})

export const getBird = (userName, pageIndex) => (dispatch,getstate) => {
    dispatch(selectOwenBird(userName));
    const count = 20;
    return fetchJsonp(`https://api.douban.com/v2/book/search?q=aa&start=${(pageIndex-1)*count}&count=${count}&fields=id,title,images,summary,rating,price`)
    .then(res => res.json())
    .then(json => dispatch(getOwenBird(userName, json, pageIndex)))
}