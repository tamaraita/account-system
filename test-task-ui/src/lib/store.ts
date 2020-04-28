import {
    Action,
    configureStore,
    getDefaultMiddleware,
    ThunkAction,
} from '@reduxjs/toolkit';

import TransactionsState, { transactionsReducer } from './ducks/transactions';


const store = configureStore({
    reducer: transactionsReducer,
    middleware: [
        ...getDefaultMiddleware<TransactionsState>(),
    ],
});

export type AppDispatch = typeof store.dispatch;
export type AppThunk = ThunkAction<void, TransactionsState, null, Action<string>>;

export default store;
