export const REGISTER_FAILURE = 'USER_REGISTER_FAILURE';
export const REGISTER_SUCCESS = 'USER_REGISTER_SUCCESS';
export const SIGN_OUT = 'SIGN_OUT';
export const LOGIN_FAILURE = 'LOGIN_FAILURE';
export const LOGIN_SUCCESS = 'LOGIN_SUCCESS';
export const ADD_BOOK_TO_SHOPCART = 'ADD_BOOK_TO_SHOPCART';

// 记得添加用户验证和通知
export const register = (username, password) => (dispatch, getState) => {
	if(getState().user['users'].indexOf(username) > -1) {
		dispatch({
			type: REGISTER_FAILURE
		})
	} else {
		dispatch({
			type: REGISTER_SUCCESS,
			payload: {
				username,
				password
			}
		})
	}
}
export const signOut = () => ({
	type: SIGN_OUT
})
export const login = (username, password) => (dispatch, getState) => {
  if(getState().user.users.indexOf(username) !== -1) { // 用户存在
		if(getState().user.details[username]['password'] === password) {
			dispatch({
				type: LOGIN_SUCCESS,
				payload: {
					username
				}
			})
		} else {
			dispatch({
				type: LOGIN_FAILURE,
				payload: {
					text: '密码错误'
				}
			})
		}
	} else {
		dispatch({
			type: LOGIN_FAILURE,
			payload: {
				text: '用户不存在'
			}
		})
	}
}

export const addBookToShopcart = (username, title, price, number) => ({
	type: ADD_BOOK_TO_SHOPCART,
	payload: {
		username,
		title,
		price,
		number
	}
})
