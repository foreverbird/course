import reducers from '../reducers';
import {createStore, compose, applyMiddleware} from 'redux';
import thunk from 'redux-thunk';

const middleware = [thunk];
if(process.env.NODE_ENV !== 'production'){

}

const configureStore = (preloadState={}) => {
    const store = createStore(
        reducers,
        preloadState,
        compose(
            applyMiddleware(...middleware),
            window.devToolsExtension ? window.devToolsExtension() : f => f
        )
    )
    return store;
}

export default configureStore;

