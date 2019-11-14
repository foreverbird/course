import { REGISTER_FAILURE, REGISTER_SUCCESS, SIGN_OUT, LOGIN_SUCCESS, LOGIN_FAILURE, ADD_BOOK_TO_SHOPCART } from '../actions';

const user = (state = {
	currentUser: '',
	users:[],
	details: {},
	shopcart: {},
	isRegisterSuccess: false,
	infoOfLogin: '',
	isLoginSuccess: false
}, action) => {
	let shopcart = {};
	switch (action.type) {
		case REGISTER_FAILURE:
		  return {
				...state,
				isRegisterSuccess: false
			}
		case REGISTER_SUCCESS:
		    let users = [...state.users, action.payload.username];
			let details = {...state.details, [action.payload.username]: { username: action.payload.username, password: action.payload.password } }
			shopcart = { ...state.shopcart }
			return {
				...state,
				users,
				details,
				shopcart,
				currentUser: action.payload.username,
				isRegisterSuccess: true
			}
		case SIGN_OUT:
			return {
				...state,
				currentUser: ''
			}
		case LOGIN_FAILURE:
			let info = action.payload.text;
			return {
				...state,
				infoOfLogin: info,
				isLoginSuccess: false
			}
		case LOGIN_SUCCESS:
			return {
				...state,
				infoOfLogin: '',
				isLoginSuccess: true,
				currentUser: action.payload.username,
			}
		default:
			return state;
	}
}

export default user;
